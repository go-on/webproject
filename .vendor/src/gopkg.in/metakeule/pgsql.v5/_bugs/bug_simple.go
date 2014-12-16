package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	//"github.com/metakeule/pq"
	"net/http"
	"os"
	"strings"
)

var DB *sql.DB

var nSql = 15
var maxDBCons = 30
var maxHttpConnects = 20
var maxIdleDBCons = maxDBCons - maxHttpConnects
var httpLock = make(chan int, maxHttpConnects)

//var nSql = 1
var reqNum = 0

func Connect(url string) (db *sql.DB) {
	p, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}

	db, err = sql.Open("postgres", p)
	//db.SetMaxIdleConns(-1)
	db.SetMaxIdleConns(maxIdleDBCons)
	if err != nil {
		panic(err.Error())
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	<-httpLock
	reqNum++
	numErrs := 0
	failed := []string{}
	for i := 0; i < nSql; i++ {
		res, err := DB.Query(fmt.Sprintf("SELECT '%v'::int", reqNum*100+i))
		if err != nil {
			numErrs++
			s := strings.ToUpper(fmt.Sprintf("Error (Query): %s", err.Error()))
			failed = append(failed, s)
			fmt.Println(s)
			continue
		}
		var ii int
		res.Next()
		//for res.Next() {
		err = res.Scan(&ii)
		if err != nil {
			fmt.Printf("Error (Scan): %s\n", err.Error())
			continue
		}
		res.Close()
		//}
	}
	// fmt.Printf("Num Errors: %v\nFailed: \n%s\n\n", numErrs, strings.Join(failed, "\n"))
	// fmt.Fprintf(w, "Num Errors: %v\nFailed: \n%s\n\n", numErrs, strings.Join(failed, "\n"))
	httpLock <- 1
}

func Open(driverName, dataSourceName string) (ø *sql.DB, ſ error) {
	ø, ſ = sql.Open(driverName, dataSourceName)
	return
}

/*
   prepare the postgres server with

       ALTER Role username CONNECTION LIMIT 10;

   run webserver with

       export DB_URL=postgres://user:password@localhost:5432/database ; go run main.go

   run ab with

       ab -n 2000 -c 200 http://localhost:8080/

*/
func main() {
	for i := 0; i < maxHttpConnects; i++ {
		httpLock <- 1
	}
	DB = Connect(os.Getenv("PG_URL"))
	http.HandleFunc("/", handler)
	fmt.Println("serving on localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error (Serving): %s\n", err)
	}
}
