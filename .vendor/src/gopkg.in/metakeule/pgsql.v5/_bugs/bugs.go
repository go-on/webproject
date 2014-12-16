package main

import (
	"database/sql"
	"fmt"
	//"github.com/lib/pq"
	"github.com/metakeule/pq"
	"net/http"
	"os"
	"strings"
)

var DB *DbWrapper
var dbLock = make(chan int, 1)
var httpLock = make(chan int, 1)

//var nSql = 15
var nSql = 1
var reqNum = 0

func Connect(url string) (db *DbWrapper) {
	p, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}

	db, err = Open("postgres", p)
	db.Debug = true
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
		err = res.Scan(&ii)
		if err != nil {
			fmt.Printf("Error (Scan): %s\n", err.Error())
			continue
		}
		if reqNum*100+i != ii {
			fmt.Printf("Error: %v != %v\n", reqNum*100+i, ii)
		}
	}
	fmt.Printf("Num Errors: %v\nFailed: \n%s\n\n", numErrs, strings.Join(failed, "\n"))
	fmt.Fprintf(w, "Num Errors: %v\nFailed: \n%s\n\n", numErrs, strings.Join(failed, "\n"))
	httpLock <- 1
}

func Open(driverName, dataSourceName string) (ø *DbWrapper, ſ error) {
	ø = &DbWrapper{dataSourceName: dataSourceName, driverName: driverName}
	dbLock <- 1
	return
}

// wraps a *sql.DB in order to do x queries at a time, to prevent
//    pq: too many connections for role "user"
// errors
type DbWrapper struct {
	dataSourceName string
	driverName     string
	db             *sql.DB
	Debug          bool
}

func (ø *DbWrapper) Close() (err error) {
	err = ø.db.Close()
	if ø.Debug {
		fmt.Println("disconnect from db")
	}
	dbLock <- 1
	return
}

func (ø *DbWrapper) connect() (err error) {
	<-dbLock
	ø.db, err = sql.Open(ø.driverName, ø.dataSourceName)
	// you may try this with go 1.1 beta2
	//ø.db.SetMaxIdleConns(-1)
	if ø.Debug {
		fmt.Println("connect to db")
	}
	if err != nil {
		fmt.Printf("can't connect to DB: %s\n", err.Error())
	}
	return
}

func (ø *DbWrapper) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	err = ø.connect()

	if err != nil {
		dbLock <- 1
		return
	}

	if ø.Debug {
		fmt.Printf("Query: \n------\n%s\n-----\nwith args: %#v\n", query, args)
	}
	st, err := ø.db.Prepare(query)
	defer func() {
		if st != nil {
			st.Close()
		}
		ø.Close()
	}()
	if err != nil {
		return
	}
	rows, err = st.Query(args...)
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
	DB = Connect(os.Getenv("PG_URL"))
	httpLock <- 1
	http.HandleFunc("/", handler)
	fmt.Println("serving on localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error (Serving): %s\n", err)
	}
}
