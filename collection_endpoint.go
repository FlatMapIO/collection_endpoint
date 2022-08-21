package collection_endpoint

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CollectionEndpoint struct {
	db    *sql.DB
	table string

	spec         QuerySpec
	errorHandler func(error, http.ResponseWriter)
}

func NewCollectionEndpoint(db *sql.DB, table string, spec QuerySpec) (CollectionEndpoint, error) {
	if db == nil {
		return CollectionEndpoint{}, errors.New("db is nil")
	}
	if table == "" {
		return CollectionEndpoint{}, errors.New("table is empty")
	}

	var ret = CollectionEndpoint{db, table, spec, defaultErrorHandler}
	if err := spec.complete(); err != nil {
		return ret, err
	}
	return ret, nil
}
func (c *CollectionEndpoint) SetErrorHandler(h func(err error, w http.ResponseWriter)) {
	c.errorHandler = h
}

func (c (CollectionEndpoint)) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	info, err := NewQueryInfo(&c.spec, query)
	if err != nil {
		c.errorHandler(err, w)
		return
	}

	if datasql, totalsql, err := info.GenerateSQL(c.table); err != nil {
		c.errorHandler(err, w)
		return
	} else {
		dataRows, err := c.db.Query(datasql)
		if err != nil {
			c.errorHandler(err, w)
			return
		}
		totalRows, err := c.db.Query(totalsql)
		if err != nil {
			c.errorHandler(err, w)
			return
		}
		var (
			total int
			data  json.RawMessage
		)
		for totalRows.Next() {
			if err := totalRows.Scan(&total); err != nil {
				c.errorHandler(err, w)
				return
			}
		}
		for dataRows.Next() {
			if err := dataRows.Scan(&data); err != nil {
				c.errorHandler(err, w)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"total": %d, "data": %s}`, total, data)))
	}
}

func defaultErrorHandler(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"errors": "%s"}`, err.Error())))
}