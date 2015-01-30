# wsi
integrates web requests with database/sql (golang)

[ ![Codeship Status for go-on/wsi](https://codeship.io/projects/88d33190-89c8-0132-2266-4676ffdbdc37/status)](https://codeship.io/projects/59797) [![GoDoc](https://godoc.org/github.com/go-on/wsi?status.png)](http://godoc.org/github.com/go-on/wsi)

## Example

```go

import (
    "gopkg.in/go-on/wsi.v1"
    "net/http"
)

type Person struct {
    Id   int
    Name string
    Age  int `json:",omitempty"`
}

// maps the columns to the fields of a new Person; must be a pointer method
func (p *Person) MapColumns(colToField map[string]interface{}) {
    colToField["id"] = &p.Id
    colToField["name"] = &p.Name
    colToField["age"] = &p.Age
}

var newPerson wsi.Ressource = func() wsi.ColumnsMapper { return &Person{} }

func findPersons(opts wsi.QueryOptions, w http.ResponseWriter, r *http.Request) (wsi.Scanner, error) {
    sc, err := wsi.DBQuery(
        DB, 
        "SELECT id,name from person ORDER BY $1 LIMIT $2 OFFSET $3", 
        strings.Join(opts.OrderBy, ","), 
        opts.Limit, 
        opts.Offset,
    )
    // handle error writes on your own
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`Database error. Contact the administrator.`))
    }
    return sc, err
}


func main() {
    // you might set an error handler as well, see the api docs
    servePersons := newPerson.Query(findPersons)

    http.Handle("/person/", servePersons)
    // will serve: [{"Id":12,"Name":"Adrian"},{"Id":24,"Name":"George"},...]

    http.ListenAndServe(":8080",nil)    
}

```

You may define your own `wsi.Encoder` if you want to deliver something other than json.