package main

import (
	"database/sql"
	//	"database/sql/driver"
	"fmt"
	//"github.com/mattn/go-sqlite3"
	"gopkg.in/metakeule/dbwrap.v2"
	"gopkg.in/metakeule/pgsql.v6"
)

type areaTable struct {
	TABLE    *pgsql.Table `name:"koelnart-area_focus" unique:"area#work"`
	Area     *pgsql.Field `name:"area"     type:"varchar(255)" flag:"pkey" enum:"foto-own,foto-other"`
	Work     *pgsql.Field `name:"work"     type:"int"`
	Position *pgsql.Field `name:"position" type:"int"          flag:"pkey"`
}

var AreaTable = &areaTable{}

type Area struct {
	Work     int    `db.select:"allareas"`
	Area     string `db.select:"allareas"`
	Position int    `db.select:"allareas"`
}

/*
type dB struct {
	*dbwrap.Wrapper
}

func (d *dB) Begin() (tx driver.Tx, e error) { return }
func (d *dB) Close() error                   { return nil }
func (d *dB) NumInput() int                  { return 0 }

func (d *dB) Prepare(s string) (stm driver.Stmt, e error) {
	fmt.Printf("\n-- --->8--- PREPARE --->8---\n%s\n-- ---8<--- PREPARE ---8<---\n\n", s)
	stm = d
	return
}

func (d *dB) Exec(vals []driver.Value) (res driver.Result, err error) { return }
func (d *dB) Query([]driver.Value) (res driver.Rows, err error)       { return }

var dbdrv = &dB{dbwrap.New("d", &sqlite3.SQLiteDriver{})}

*/
var db *sql.DB
var fake *dbwrap.Fake

func init() {
	/*
		dbdrv.HandleOpen = func(s string, conn driver.Conn) (driver.Conn, error) { return dbdrv, nil }
		dbdrv.HandlePrepare = func(d driver.Conn, s string) (driver.Stmt, error) { return dbdrv, nil }

		sql.Register("debug", dbdrv)
	*/
	fake, db = dbwrap.NewFake()

	td := pgsql.MustTableDefinition(AreaTable)
	// set foreign key
	//td.Fields["work"].Add(work.Id, OnDeleteCascade)
	// set default value
	td.Fields["area"].Add(pgsql.Sql("'default area'"))
	td.MustFinish()
}

func AllAreas(dB *sql.DB, result interface{}) (num int, err error) {
	return pgsql.NewRow(dB, AreaTable.TABLE).SelectByStructs(result, "all")
}

func main() {
	fmt.Println(AreaTable.TABLE.Create().String())
	db.Exec("hu")
	fmt.Println(fake.LastQuery)
	db.Query("ho")
	fmt.Println(fake.LastQuery)
	db.Prepare("he")
	fmt.Println(fake.LastQuery)
	db.QueryRow("hui")
	fmt.Println(fake.LastQuery)
	/*
	   areas := make([]Area, 21)
	   num, err := AllAreas(db, areas)
	   if err == nil {
	       areas = areas[0:num]
	       fmt.Printf("areas: %#v\n", areas)
	   }
	*/

	db.Query("select * from a")
	fmt.Println(fake.LastQuery)
	/*
		fmt.Println(AreaTable.Area.Name)
		fmt.Println(AreaTable.Area.Type)
		fmt.Println(AreaTable.Area.Is(pgsql.PrimaryKey))
		fmt.Println(AreaTable.Area.Selection)
		fmt.Println(AreaTable.TABLE.QueryField("Position").Name)
	*/
}
