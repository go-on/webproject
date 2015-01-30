package wsi

import (
	"database/sql"
	"gopkg.in/go-on/pq.v2"
	"gopkg.in/metakeule/dbwrap.v2"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var fake, db = dbwrap.NewFake()
var realDB *sql.DB

func searchPersonFromDB(opts QueryOptions, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if len(opts.OrderBy) == 0 {
		opts.OrderBy = append(opts.OrderBy, "id ASC")
	}
	limit := opts.Limit
	if limit == 0 {
		limit = 30
	}
	return DBQuery(realDB, `SELECT 2 AS "id", 'hiho' AS "name" ORDER BY $1 LIMIT $2 OFFSET $3`, strings.Join(opts.OrderBy, ","), limit, opts.Offset)
}

func searchPeronsIdsNames(opts QueryOptions, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if len(opts.OrderBy) == 0 {
		opts.OrderBy = append(opts.OrderBy, "id ASC")
	}
	limit := opts.Limit
	if limit == 0 {
		limit = 20
	}
	return DBQuery(db, "SELECT id,name FROM person ORDER BY $1 LIMIT $2 OFFSET $3", strings.Join(opts.OrderBy, ","), limit, opts.Offset)
}

func searchPersonIds(opts QueryOptions, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if len(opts.OrderBy) == 0 {
		opts.OrderBy = append(opts.OrderBy, "id ASC")
	}
	limit := opts.Limit
	if limit == 0 {
		limit = 10
	}
	return DBQuery(db, "SELECT id FROM person ORDER BY $1 LIMIT $2 OFFSET $3", strings.Join(opts.OrderBy, ","), limit, opts.Offset)
}

func (p *person) MapColumns(colToField map[string]interface{}) {
	colToField["id"] = &p.Id
	colToField["name"] = &p.Name
}

type person struct {
	Id   int
	Name string
	Age  int `json:",omitempty"`
	err  error
}

func newPersonMapper() ColumnsMapper {
	return &person{}
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errHandler := func(rr *http.Request, err error) { p.err = err }
	var fn func(QueryOptions, http.ResponseWriter, *http.Request) (Scanner, error)
	switch r.URL.Path {
	case "/a":
		fn = searchPeronsIdsNames
	case "/b":
		fn = searchPersonIds
	case "/real":
		fn = searchPersonFromDB
	}
	RessourceFunc(newPersonMapper).Query(fn).SetErrorCallback(errHandler).ServeHTTP(w, r)
}

func init() {
	fake.SetNumInputs(3)
}

func TestRealDB(t *testing.T) {
	pg_url := os.Getenv("PG_URL")
	if pg_url == "" {
		t.SkipNow()
	}
	u, err := pq.ParseURL(pg_url + "?sslmode=disable")
	if err != nil {
		panic(err)
	}
	realDB, err = sql.Open("postgres", u)
	if err != nil {
		panic(err)
	}
	defer realDB.Close()
	p := &person{}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/real", nil)

	p.ServeHTTP(rec, req)

	if p.err != nil {
		t.Errorf("got err: %s", err.Error())
	}

	got := rec.Body.String()
	expected := `[{"Id":2,"Name":"hiho"}
]`

	if got != expected {
		t.Errorf("Body: %#v, expected: %#v", got, expected)
	}

	got = rec.Header().Get("Content-Type")
	expected = `application/json; charset=utf-8`

	if got != expected {
		t.Errorf("Content-Type: %#v, expected: %#v", got, expected)
	}

}

func TestQueryRun(t *testing.T) {
	p := &person{}

	queryA := "SELECT id,name FROM person ORDER BY $1 LIMIT $2 OFFSET $3"
	queryB := "SELECT id FROM person ORDER BY $1 LIMIT $2 OFFSET $3"

	tests := []struct {
		url     string
		limit   int64
		offset  int64
		orderBy string
		query   string
	}{
		{"/a", 20, 0, "id ASC", queryA},
		{"/a?sort=-name&limit=-1&offset=-20", 20, 0, "name DESC", queryA},
		{"/a?sort=name&sort=id", 20, 0, "name ASC,id ASC", queryA},
		{"/a?sort=-name&sort=id", 20, 0, "name DESC,id ASC", queryA},
		{"/a?sort=-name&sort=id&limit=12", 12, 0, "name DESC,id ASC", queryA},
		{"/a?sort=-name&sort=id&limit=0&offset=2", 20, 2, "name DESC,id ASC", queryA},
		{"/a?offset=4", 20, 4, "id ASC", queryA},
		{"/b?offset=5", 10, 5, "id ASC", queryB},
	}

	for _, test := range tests {

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.url, nil)

		p.ServeHTTP(rec, req)

		lastQuery, lastVals := fake.LastQuery()
		if want, got := test.query, lastQuery; want != got {
			t.Errorf("%s => query = %#v, want: %#v", test.url, got, want)
		}

		if want, got := test.orderBy, lastVals[0]; want != got {
			t.Errorf("%s => orderBy = %#v, want: %#v, got: %#v", test.url, got, want)
		}

		if want, got := test.limit, lastVals[1].(int64); want != got {
			t.Errorf("%s => limit = %v, want: %v", test.url, got, want)
		}

		if want, got := test.offset, lastVals[2].(int64); want != got {
			t.Errorf("%s => offset = %v, want: %v", test.url, got, want)
		}

	}
}
