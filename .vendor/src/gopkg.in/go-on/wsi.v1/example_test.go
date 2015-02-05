// +build go1.1

package wsi_test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-on/wsi"
)

type Person struct {
	Id   int
	Name string
	Age  int `json:",omitempty"`
}

// maps the given column to a pointer of a fields of the Person
// must be a pointer method
func (p *Person) Map(column string) (fieldPtr interface{}) {
	switch column {
	case "id":
		return &p.Id
	case "name":
		return &p.Name
	case "age":
		return &p.Age
	default:
		panic("unknown column " + column)
	}
}

// newPerson is a function that creates a new person.
// we need this as wsi.Ressource to generate the http.Handlers
var newPerson wsi.RessourceFunc = func() wsi.Mapper { return &Person{} }

// findPersonsFake fakes our query, for a realistic query, see findPersons
// if any error happens, it must write to the response writer and return an error
func findPersonsFake(opts wsi.QueryOptions, w http.ResponseWriter, r *http.Request) (wsi.Scanner, error) {
	return wsi.NewTestQuery([]string{"id", "name"}, testData...), nil
}

// creates a http.Handler based on findPersonsFake that writes the resulting persons as json
// we are using the fake query here to avoid the need for a database, you may replace findPersonsFake
// with findPersons if you have a real database connection
var findHandler = newPerson.Query(findPersonsFake).SetErrorCallback(printErr)

var DB *sql.DB

// findPersons defines the search sql.
// it must handle edge case, like limit = 0 or max limits, however limit and offset will never be < 0
func findPersons(opts wsi.QueryOptions, w http.ResponseWriter, r *http.Request) (wsi.Scanner, error) {
	if len(opts.OrderBy) == 0 {
		opts.OrderBy = append(opts.OrderBy, "id asc")
	}

	// handle max limit
	limit := opts.Limit
	if limit == 0 || limit > 30 {
		limit = 30
	}

	return wsi.DBQuery(
		DB,
		"SELECT id,name from person ORDER BY $1 LIMIT $2 OFFSET $3",
		strings.Join(opts.OrderBy, ","),
		limit,
		opts.Offset,
	)
}

// createPerson creates a person based on the values of the given ColumnsMapper
// and writes to the given responsewriter
// we need to return an error here, even if we handle the response writing, so that the general
// error handler may be called
func createPerson(m wsi.Mapper, w http.ResponseWriter, r *http.Request) error {
	// we fake a created response here
	res := map[string]interface{}{"Id": 400, "Name": m.Map("name")}
	w.WriteHeader(http.StatusCreated)
	wsi.ServeJSON(res, w)
	return nil
}

// creates a http.Handler based on createPerson that load persons as json
var addHandler = newPerson.Exec(createPerson).SetErrorCallback(printErr)

func Example() {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/person/", nil)
	findHandler.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())

	fmt.Println("-----")

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/person", strings.NewReader(`{"Name":"Peter"}`))
	addHandler.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())

	// Output:
	// [{"Id":12,"Name":"Adrian"}
	// ,{"Id":24,"Name":"George"}
	// ]
	// -----
	// {"Id":400,"Name":"Peter"}
	//
}

var testData = []map[string]wsi.Setter{
	map[string]wsi.Setter{"id": wsi.SetInt(12), "name": wsi.SetString("Adrian")},
	map[string]wsi.Setter{"id": wsi.SetInt(24), "name": wsi.SetString("George")},
}

// an example error handler
func printErr(r *http.Request, err error) {
	fmt.Printf("Error in route GET %s: %T %s\n", r.URL.Path, err, err.Error())
}