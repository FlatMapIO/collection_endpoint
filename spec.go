package collection_endpoint

import (
	"database/sql"
	"errors"
)

type OP string

const (
	OP_EQ       OP = "eq"
	OP_NE       OP = "ne"
	OP_GT       OP = "gt"
	OP_GE       OP = "ge"
	OP_LT       OP = "lt"
	OP_LE       OP = "le"
	OP_IN       OP = "in"
	OP_NIN      OP = "nin"
	OP_LIKE     OP = "like"
	OP_NLIKE    OP = "nlike"
	OP_BETWEEN  OP = "between"
	OP_NBETWEEN OP = "nbetween"
)

type Type int

type FilterSpec struct {
	Required  bool
	Operators []OP
	opSet     keySet
}

type DefaultQuery map[string]string
type QuerySpec struct {
	Database *sql.DB
	Table    string
	Default  DefaultQuery
	Fields   []string
	Sorts    []string

	Limit  LimitSpec
	Filter map[string]FilterSpec

	completed bool
	sortsSet  keySet
	fieldsSet keySet
}

func (q *QuerySpec) complete() error {
	if q.completed {
		return nil
	}
	if q.Database == nil {
		return errors.New("Databas is required")
	}
	if q.Table == "" {
		return errors.New("table is required")
	}
	q.sortsSet = newKeySet(q.Sorts)
	q.fieldsSet = newKeySet(q.Fields)
	for _, f := range q.Filter {
		f.opSet = newKeySet(f.Operators)
	}
	q.completed = true
	return nil
}

type LimitSpec struct {
	Min uint
	Max uint
}

var LimitByMin1Max100 = LimitSpec{
	Min: 1,
	Max: 100,
}