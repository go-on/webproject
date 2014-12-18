package main

import (
	"fmt"
)

import (
	"database/sql"
	"gopkg.in/go-on/pq.v2"
	"log"
	"os"
)

var DB *sql.DB

func Connect(driver string, str string) *sql.DB {
	cs, err := pq.ParseURL(str)
	if err != nil {
		panic(err.Error())
	}
	db, err := sql.Open(driver, cs)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func init() {
	DB = Connect("postgres", os.Getenv("PG_URL"))
}

func main() {
	r := DB.QueryRow("Select 3,4")
	var a, b int
	err := r.Scan(&a, &b)

	if err != nil {
		log.Println("Error ", err)

	}

	fmt.Println(a, b)
}
