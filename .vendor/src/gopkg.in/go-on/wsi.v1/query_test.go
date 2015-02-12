package wsi

import (
	"database/sql"
	"gopkg.in/go-on/builtin.v1"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gopkg.in/go-on/pq.v2"
	"gopkg.in/metakeule/dbwrap.v2"
)

var fake, DB = dbwrap.NewFake()
var realDB *sql.DB

func searchPersonFromDB(limit, offset int, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if limit == 0 {
		limit = 30
	}
	return DBQuery(realDB, `SELECT 2 AS "Id", 'hiho' AS "Name", null AS "Notes" ORDER BY "Id" LIMIT $1 OFFSET $2`, limit, offset)
}

func searchPeronsIdsNames(limit, offset int, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if limit == 0 {
		limit = 20
	}
	return DBQuery(DB, `SELECT "Id","Name" FROM person ORDER BY "Name" LIMIT $1 OFFSET $2`, limit, offset)
}

func searchPersonIds(limit, offset int, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if limit == 0 {
		limit = 10
	}
	return DBQuery(DB, `SELECT "Id" FROM person ORDER BY "Id" LIMIT $1 OFFSET $2`, limit, offset)
}
func searchPersonErr(limit, offset int, w http.ResponseWriter, r *http.Request) (Scanner, error) {
	if limit == 0 {
		limit = 30
	}
	return DBQuery(realDB, `SELECT 2 AS "Id, 'hiho' AS "Name", null AS "Notes" ORDER BY "Id" LIMIT $1 OFFSET $2`, limit, offset)
}

type person struct {
	Id    int
	Name  string
	Age   int              `json:",omitempty" sql:",omitempty"`
	Notes builtin.Stringer `json:",omitempty"` // optional
	err   error
}

func newPersonMapper() interface{} {
	return &person{}
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errHandler := func(rr *http.Request, err error) { p.err = err }
	var fn func(limit, offset int, w http.ResponseWriter, r *http.Request) (Scanner, error)
	switch r.URL.Path {
	case "/a":
		fn = searchPeronsIdsNames
	case "/b":
		fn = searchPersonIds
	case "/real":
		fn = searchPersonFromDB
	case "/err":
		fn = searchPersonErr
	}

	Ressource{newPersonMapper, errHandler}.ServeQuery(fn, w, r)
}

func init() {
	fake.SetNumInputs(2)
}

func TestRealDB(t *testing.T) {
	pg_url := os.Getenv("PG_URL")
	if pg_url == "" {
		t.SkipNow()
	}
	u, err := pq.ParseURL(pg_url)
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
		t.Errorf("got err: %s", p.err.Error())
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

func TestRealDBErr(t *testing.T) {
	pg_url := os.Getenv("PG_URL")
	if pg_url == "" {
		t.SkipNow()
	}
	u, err := pq.ParseURL(pg_url)
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
	req, _ := http.NewRequest("GET", "/err", nil)

	p.ServeHTTP(rec, req)

	if p.err == nil {
		t.Errorf("expected err, got nil")
	}

}

func TestQueryRun(t *testing.T) {
	p := &person{}

	queryA := `SELECT "Id","Name" FROM person ORDER BY "Name" LIMIT $1 OFFSET $2`
	queryB := `SELECT "Id" FROM person ORDER BY "Id" LIMIT $1 OFFSET $2`

	tests := []struct {
		url    string
		limit  int64
		offset int64
		query  string
	}{
		{"/a", 20, 0, queryA},
		{"/a?limit=-1&offset=-20", 20, 0, queryA},
		{"/a", 20, 0, queryA},
		{"/a?limit=12", 12, 0, queryA},
		{"/a?limit=0&offset=2", 20, 2, queryA},
		{"/a?offset=4", 20, 4, queryA},
		{"/b?offset=5", 10, 5, queryB},
	}

	for _, test := range tests {

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.url, nil)

		p.ServeHTTP(rec, req)

		lastQuery, lastVals := fake.LastQuery()
		if want, got := test.query, lastQuery; want != got {
			t.Errorf("%s => query = %#v, want: %#v", test.url, got, want)
		}

		if want, got := test.limit, lastVals[0].(int64); want != got {
			t.Errorf("%s => limit = %v, want: %v", test.url, got, want)
		}

		if want, got := test.offset, lastVals[1].(int64); want != got {
			t.Errorf("%s => offset = %v, want: %v", test.url, got, want)
		}

	}
}
