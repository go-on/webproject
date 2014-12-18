package main

import (
	"database/sql"
	"fmt"
	"gopkg.in/go-on/pq.v2"
	. "gopkg.in/metakeule/pgsql.v5"
	"gopkg.in/metakeule/pgsql.v5/sharedquery"
	"os"
	"time"
)

var (
	numParallelQueryRequests = 12
	numParallelQueries       = 12
	collectingTime           = time.Second * 3
	timeout                  = time.Second * 4
	Db                       *sql.DB
	Qm                       *sharedquery.QueryManager
)

var TABLE = NewTable("koelnart-news")

var Id = TABLE.NewField("id", UuidType, PrimaryKey, UuidGenerate)
var Name = TABLE.NewField("name", VarChar(255))
var Description = TABLE.NewField("description", HtmlType)
var Sidebar = TABLE.NewField("sidebar", HtmlType)
var Position = TABLE.NewField("position", IntType)

func Connect(str string) *sql.DB {
	cs, ſ := pq.ParseURL(str)
	if ſ != nil {
		panic(ſ)
	}

	db, ſ := sql.Open("postgres", cs)
	if ſ != nil {
		panic(ſ)
	}
	//db.SetMaxIdleConns(0)
	// db.Debug = true
	return db
}

func init() {
	Db = Connect(os.Getenv("PG_URL"))
	Qm = sharedquery.New(Db, numParallelQueryRequests, numParallelQueries, collectingTime, timeout)

}

func f1() {
	defer func() {
		finished <- true
	}()
	// 
	sel1 := Select(TABLE, Id, Name, Description, Where(Equals(Name, "timeslices")))
	res1, err := Qm.Query(sel1.(*SelectQuery))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res1.HasErrors() {
		fmt.Println(res1.Errors)
	} else {
		for {
			r, next := res1.Next()
			if next == false {
				break
			}
			if r == nil {
				continue
			}
			fmt.Println("---")
			// fmt.Printf("r: %v\n", r.AsStrings())
			for k, v := range r.AsStrings() {
				fmt.Printf("%s: %s\n", k, v)
			}
		}
	}
}

func f2() {
	defer func() {
		finished <- true
	}()

	sel1 := Select(TABLE, Id, Name, Where(Equals(Name, "eight")))
	res1, err := Qm.Query(sel1.(*SelectQuery))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res1.HasErrors() {
		fmt.Println(res1.Errors)
	} else {
		for {
			r, next := res1.Next()
			if next == false {
				break
			}
			if r == nil {
				continue
			}
			//fmt.Printf("r: %v\n", r.AsStrings())
			// fmt.Printf("r: %v\n", r)
			fmt.Println("---")
			for k, v := range r.AsStrings() {
				fmt.Printf("%s: %s\n", k, v)
			}

		}
	}
}

var finished = make(chan bool, 2)

func main() {

	go Qm.Run()

	go f1()
	go f2()

	<-finished
	<-finished
}
