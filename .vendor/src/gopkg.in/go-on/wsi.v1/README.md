# wsi
integrates web requests with database/sql (golang)

[ ![Codeship Status for go-on/wsi](https://codeship.io/projects/88d33190-89c8-0132-2266-4676ffdbdc37/status)](https://codeship.io/projects/59797) [![GoDoc](https://godoc.org/github.com/go-on/wsi?status.png)](http://godoc.org/github.com/go-on/wsi)

## Example

```go

import (
    "database/sql"
    "gopkg.in/go-on/wsi.v1"
    "net/http"
    "log"
)

type Person struct {
    ID   int `sql:"id"`
    Name string `sql:"name"`
    Age  int `json:",omitempty" sql:"-"`
}

func newPerson() interface{} { 
    return &Person{} 
}

func logErr(r *http.Request, err error) {
    log.Printf("Error in route GET %s: %T %s\n", r.URL.Path, err, err.Error())
}

var db *sql.DB // TODO setup the db connection and create person table

func findPersons(limit, offset int, w http.ResponseWriter, r *http.Request) (wsi.Scanner, error) {
    return wsi.DBQuery(
        db, 
        `SELECT id,name from person ORDER BY name ASC LIMIT $1 OFFSET $2`, 
        limit, 
        offset,
    )
}

func main() {
    http.Handle("/person/", wsi.Ressource{newPerson,logErr}.Query(findPersons))
    // will serve: [{"ID":12,"Name":"Adrian"},{"ID":24,"Name":"George"},...]

    http.ListenAndServe(":8080",nil)    
}

```

You may define your own `wsi.Encoder` if you want to deliver something other than json.