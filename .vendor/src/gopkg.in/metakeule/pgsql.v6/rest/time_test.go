package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v6"

	"testing"
)

type TimeTest struct {
	Id       *fat.Field `type:"string uuid"      db:"id UUIDGEN PKEY" rest:" R "`
	Time     *fat.Field `type:"time timestamp"   db:"time"            rest:"CRU"`
	TimeNull *fat.Field `type:"time timestamptz" db:"timenull NULL"   rest:" RU"`
}

var TIME_TEST = fat.Proto(&TimeTest{}).(*TimeTest)
var CRUDTimeTest *CRUD

func init() {
	registry.MustRegisterTable("timetest", TIME_TEST)

	DB.Exec("DROP TABLE timetest")

	timeTestTable := registry.TableOf(TIME_TEST)
	_, e := DB.Exec(timeTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table timetest: \nError: %s\nSql: %s\n", e.Error(),
			timeTestTable.Create()))
	}

	CRUDTimeTest = NewCRUD(registry, TIME_TEST)
}

func TestTimeCreate(t *testing.T) {
	id, err := CRUDTimeTest.Create(DB, b(`
	{
		"Time": "2001-02-13T23:04:45Z"
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create TimeTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDTimeTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created timetest with id %s: %s", id, err)
		return
	}

	if x["Time"] != "2001-02-13T23:04:45Z" {
		t.Errorf("timetest Time is not 2001-02-13T23:04:45Z, but %#v", x["Time"])
	}

	if x["TimeNull"] != nil {
		t.Errorf("timetest TimeNull is %#v, but should be nil", x["TimeNull"])
	}

}

func TestTimeUpdate(t *testing.T) {
	id, _ := CRUDTimeTest.Create(DB, b(`
	{
		"Time": "2001-02-13T23:04:45Z"
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDTimeTest.Update(DB, id, b(`
	{
		"Time": "2011-12-13T23:04:45Z",
		"TimeNull": "2011-12-13T23:04:45+03:00"
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update timetest with id %s: %s", id, err)
		return
	}

	x, err = CRUDTimeTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created timetest with id %s: %s", id, err)
		return
	}

	if x["Time"] != "2011-12-13T23:04:45Z" {
		t.Errorf("timetest Time is not \"2011-12-13T23:04:45Z\", but %#v", x["Time"])
	}

	if x["TimeNull"] != "2011-12-13T20:04:45Z" {
		t.Errorf("timetest TimeNull is not \"2011-12-13T20:04:45Z\", but %#v", x["TimeNull"])
	}
}
