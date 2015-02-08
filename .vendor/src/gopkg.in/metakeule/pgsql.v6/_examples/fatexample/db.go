package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"gopkg.in/go-on/pq.v2"
	"gopkg.in/metakeule/dbwrap.v2"
	"os"
)

var (
	DB                *sql.DB
	wrapperDriverName = "dbwrapper1"
)

type pqdrv bool

func (ø pqdrv) Open(connectString string) (driver.Conn, error) { return pq.Open(connectString) }

func Connect(driver string, str string) *sql.DB {
	cs, err := pq.ParseURL(str)
	if err != nil {
		panic(err.Error())
	}
	db, err := sql.Open(driver, cs)
	//db, err := sql.Open("xyz", cs)
	_ = cs
	//db, err := sql.Open(driver, "xyz")
	if err != nil {
		fmt.Printf("Error: %#v %T\n", err, err)
		// panic(err.Error())
	}

	//r, e := db.Query("select 1")
	_, e := db.Exec("select 1")
	if e != nil {
		fmt.Printf("exec failed: %#v %s\n", e, e)
	}

	//_ = r

	/*
		r := db.QueryRow("select 1")
		if r == nil {
			fmt.Printf("queriing single row failed: %#v %s\n", r, r)
		}
		var i int
		e := r.Scan(&i)
		if e != nil {
			fmt.Printf("scanning single row failed: %#v %s\n", e, e)
		}
	*/
	return db
}

func init() {

	DBWrap := dbwrap.New(wrapperDriverName, pqdrv(true))

	DBWrap.HandlePrepare = func(conn driver.Conn, query string) (driver.Stmt, error) {
		fmt.Println("\n-- PREPARE", query)
		return conn.Prepare(query)
	}

	DBWrap.HandleExec = func(conn driver.Execer, query string, args []driver.Value) (driver.Result, error) {
		fmt.Println("\n-- EXEC", query)
		return conn.Exec(query, args)
	}

	DBWrap.HandleOpen = func(name string, conn driver.Conn) (driver.Conn, error) {
		conn.(driver.Execer).Exec(`SET search_path = "public"`, []driver.Value{})
		return conn, nil
	}

	fmt.Println(os.Getenv("PG_URL"))
	pg_url := "postgres://docker:docker@172.17.0.2:5432/pgsqltest?schema=public"
	DB = Connect(wrapperDriverName, pg_url)

	/*
		r, e := DB.Query("select 'name' from company where id = '9d12b0e6-773f-432a-b31a-a77a87dbd7d'")
		if e != nil {
			fmt.Printf("can't select from company\n")

		}

		if r != nil && r.Err() != nil {
			fmt.Printf("Err %T\n", r.Err())
		}

		for r.Next() {
			var n string
			er := r.Scan(&n)

			if er != nil {
				fmt.Printf("can't scan frmo company\n")
			}

			fmt.Printf("name is: %#v\n", n)
		}
	*/
	/*
		r := DB.QueryRow("select 'name' from company where name = 'x'")
		if r == nil {
			fmt.Printf("no row for company\n")
			return
		}

		var n string
		er := r.Scan(&n)

		if er != nil {
			fmt.Printf("can't scan from company\n")
		}
	*/

}
