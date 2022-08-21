package collection_query

import (
	"database/sql"
	ce "github.com/FlatMapIO/collection_endpoint"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExample(t *testing.T) {
	db, err := sql.Open("mysql", "sa:sa@tcp(localhost:3306)/test")
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := ce.NewCollectionEndpoint(ce.QuerySpec{
		Database: db,
		Table:    "userdata",
		Default: ce.DefaultQuery{
			"limit":               "10",
			"created_at[between]": `('2020-01-01', '2020-01-31')`,
		},
		Fields: []string{
			"id", "name", "created_at",
		},
		Sorts: []string{
			"id", "name", "created_at",
		},
		Limit: ce.LimitByMin1Max100,
	})
	if err != nil {
		t.Fatal(err)
	}
	router := chi.NewRouter()
	router.Get("/users", endpoint.Handler())

	server := httptest.NewServer(router)
	defer server.Close()

	cases := map[string]string{
		"limit":            "?fields=id,name,order&limit=10",
		"filter datetime":  "?fields=id,name,order&created_at[inr]=[2020-01-01 12-33,2020-01-02]",
		"filter id equals": "?fields=id,name,order&,id[eq]=1|id[eq]=2",
	}
	for name, query := range cases {
		run(t, name, server.URL+"/users", query)
	}
}

func run(t *testing.T, label, url, query string) {
	t.Helper()
	res, err := http.Get(url + query)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal(res.Status)
	}
	data, err := io.ReadAll(res.Body)
	t.Logf("%s", data)

}