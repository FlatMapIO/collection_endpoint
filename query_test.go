package collection_endpoint

import (
	"net/url"
	"testing"
)

func TestBuildQueryInfo(t *testing.T) {
	info, err := NewQueryInfo(&QuerySpec{
		Table: "sample",
		Default: DefaultQuery{
			"limit": "10",
		},
		Fields: []string{
			"id", "name", "created_at",
		},
		Sorts: []string{
			"name", "created_at",
		},
		Limit:  LimitByMin1Max100,
		Filter: map[string]FilterSpec{},
	}, url.Values{
		"fields":              []string{"id,name,created_at"},
		"sort":                []string{"+id,-name,created_at"},
		"created_at[between]": []string{"('2020-01-01', '2020-01-31')"},
		"id[eq]":              []string{"1"},
	})
	if err != nil {
		t.Fatal(err)
	}
	data, total, err := info.BuildSQL()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("\n%s", data)
	t.Logf("\n%s", total)
}