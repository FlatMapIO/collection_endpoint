package collection_endpoint

import (
	"fmt"
	"net/http"
)

type CollectionEndpoint struct {
	spec         QuerySpec
	errorHandler func(error, http.ResponseWriter)
}

func NewCollectionEndpoint(spec QuerySpec) (CollectionEndpoint, error) {
	var ret = CollectionEndpoint{spec, defaultErrorHandler}
	if err := spec.complete(); err != nil {
		return ret, err
	}
	return ret, nil
}
func (c *CollectionEndpoint) SetErrorHandler(h func(err error, w http.ResponseWriter)) {
	c.errorHandler = h
}

func (c *CollectionEndpoint) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// merge
		// query := r.URL.Query()

	}
}
func defaultErrorHandler(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"errors": "%s"}`, err.Error())))
}