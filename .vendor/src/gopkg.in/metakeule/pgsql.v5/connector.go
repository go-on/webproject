package pgsql

import (
	"database/sql"
	"fmt"
	"sync"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DBComplete interface {
	DB
	Close() (ſ error)
	Begin() (tx *sql.Tx, ſ error)
}

/*
func (ø *Row) setSearchPath() {
	if !ø.isTransaction() {
		if ø.Table.Schema != nil {
			schemaName := ø.Table.Schema.Name
			sql := `SET search_path = "` + schemaName + `"`
			if ø.Debug {
				fmt.Println(sql)
			}
			_, _ = ø.DB.Exec(sql)
		}
	}
}
*/
/*
type SchemaDb struct {
	Name string
	DB DB
}

func (ø *SchemaDb) setSearchPath(q) string {
	return `SET search_path = "` + ø.Name + `" ;` + q
}

func (* SchemaDb) Exec(query string, args ...interface{}) (sql.Result, error) {

}

	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
*/

var dbLock = make(chan int, 1)

func Open(driverName, dataSourceName string) (ø *DbWrapper, ſ error) {
	ø = &DbWrapper{dataSourceName: dataSourceName, driverName: driverName}
	ø.RWMutex = &sync.RWMutex{}
	dbLock <- 1
	return
}

// wraps a *sql.DB in order to do x queries at a time, to prevent
//    pq: too many connections for role "user"
// errors
// to be used with github.com/metakeule/pq (only one global connection)
type DbWrapper struct {
	*sync.RWMutex
	dataSourceName string
	driverName     string
	db             *sql.DB
	Debug          bool
}

func (ø *DbWrapper) Close() (ſ error) {
	ø.Lock()
	defer ø.Unlock()
	ſ = ø.db.Close()
	ø.db = nil
	if ø.Debug {
		fmt.Println("disconnect from db")
	}
	return
}

func (ø *DbWrapper) connect() (ſ error) {
	ø.Lock()
	defer ø.Unlock()
	if ø.db == nil {
		ø.db, ſ = sql.Open(ø.driverName, ø.dataSourceName)
		//ø.db.SetMaxIdleConns(-1) //for go 1.1
		if ø.Debug {
			fmt.Println("connect to db")
		}
		if ſ != nil {
			fmt.Printf("can't connect to DB: %#v\n", ſ.Error())
		}
	}
	return
}

func (ø *DbWrapper) Begin() (tx *sql.Tx, ſ error) {
	<-dbLock
	defer func() {
		dbLock <- 1
	}()
	tx, ſ = ø.db.Begin()
	return
}

func (ø *DbWrapper) Exec(query string, args ...interface{}) (res sql.Result, ſ error) {
	<-dbLock
	defer func() {
		dbLock <- 1
	}()
	ſ = ø.connect()

	if ſ != nil {
		return
	}

	if ø.Debug {
		fmt.Printf("Exec: \n------\n%s\n-----\nwith args: %#v\n", query, args)
	}
	st, ſ := ø.db.Prepare(query)
	if ſ != nil {
		return
	}
	res, ſ = st.Exec(args...)
	return
}

func (ø *DbWrapper) Prepare(query string) (st *sql.Stmt, ſ error) {
	panic("use Exec, Query or QueryRow directly (they use Prepare internally)")
	return
}

func (ø *DbWrapper) Query(query string, args ...interface{}) (rows *sql.Rows, ſ error) {
	<-dbLock
	defer func() {
		dbLock <- 1
	}()
	ſ = ø.connect()

	if ſ != nil {
		return
	}

	if ø.Debug {
		fmt.Printf("Query: \n------\n%s\n-----\nwith args: %#v\n", query, args)
	}
	st, ſ := ø.db.Prepare(query)
	if ſ != nil {
		return
	}
	rows, ſ = st.Query(args...)
	return
}

func (ø *DbWrapper) QueryRow(query string, args ...interface{}) (r *sql.Row) {
	<-dbLock
	defer func() {
		dbLock <- 1
	}()
	ſ := ø.connect()
	if ſ != nil {
		panic(ſ.Error())
	}
	if ø.Debug {
		fmt.Printf("QueryRow: \n------\n%s\n-----\nwith args: %#v\n", query, args)
	}
	st, ſ := ø.db.Prepare(query)
	if ſ != nil {
		panic(ſ.Error())
	}
	r = st.QueryRow(args...)
	return
}
