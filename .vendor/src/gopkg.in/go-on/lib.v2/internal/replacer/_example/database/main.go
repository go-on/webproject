package main

import (
	"database/sql"
	"fmt"
	"github.com/go-on/replacer"
	"net/http"
	"sync"

	"github.com/lib/pq"
)

var (
	db     *sql.DB
	Name   = replacer.Placeholder("name")
	Wheels = replacer.Placeholder("wheels")
)

type Vehicles struct {
	*sync.Mutex
	template *replacer.Template
}

func (v *Vehicles) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("select name, wheels from vehicle order by random(), name limit 80 ")

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		w.WriteHeader(500)
		return
	}

	for rows.Next() {
		var name, wheels []byte
		err = rows.Scan(&name, &wheels)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			w.WriteHeader(500)
			return
		}

		v.Lock()
		s := v.template.NewSetter()
		s.SetBytes(Name, name)
		s.SetBytes(Wheels, wheels)
		w.Write(s.Bytes())
		v.Unlock()
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	conn, err := pq.ParseURL(`postgres://docker:docker@172.17.0.2:5432/pgsqltest?schema=public`)
	panicOnErr(err)

	db, err = sql.Open("postgres", conn)
	panicOnErr(err)
	// db.SetMaxOpenConns(10)

	/*
		for i := 0; i < 10000; i++ {
			_, err := db.Exec(fmt.Sprintf("insert into vehicle (name, wheels) values('car-%d', %d)", i, i))
			if err != nil {
				fmt.Println(err)
			}
		}
	*/
}

func main() {
	t := replacer.NewTemplateString(
		"Name: " + Name.String() +
			" Wheels: " + Wheels.String() +
			"\n------------\n",
	)

	http.Handle("/", &Vehicles{&sync.Mutex{}, t})
	http.ListenAndServe(":7978", nil)
}
