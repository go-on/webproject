package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type TimesTest struct {
	Id        *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Times     *fat.Field `type:"[]time"      db:"times"           rest:"CRU"`
	TimesNull *fat.Field `type:"[]time"      db:"timesnull NULL"  rest:" RU"`
}

var TIMES_TEST = fat.Proto(&TimesTest{}).(*TimesTest)
var CRUDTimesTest *CRUD

func init() {
	registry.MustRegisterTable("timestest", TIMES_TEST)

	db.Exec("DROP TABLE timestest")

	timesTestTable := registry.TableOf(TIMES_TEST)
	_, e := db.Exec(timesTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table timestest: \nError: %s\nSql: %s\n", e.Error(),
			timesTestTable.Create()))
	}

	CRUDTimesTest = NewCRUD(registry, TIMES_TEST)
}

func TestTimesCreate(t *testing.T) {
	id, err := CRUDTimesTest.Create(db, b(`
	{
		"Times": ["2001-02-13T23:04:45Z","2011-02-13T23:04:45Z"]
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create TimesTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDTimesTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created timestest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Times"]) != `["2001-02-13T23:04:45Z","2011-02-13T23:04:45Z"]` {
		t.Errorf("timestest Times is not [\"2001-02-13T23:04:45Z\",\"2011-02-13T23:04:45Z\"], but %#v", jsonify(x["Times"]))
	}

	if x["TimesNull"] != nil {
		t.Errorf("timestest TimesNull is %#v, but should be nil", x["TimesNull"])
	}

}

func TestTimesEmpty(t *testing.T) {
	id, err := CRUDTimesTest.Create(db, b(`
	{
		"Times": []
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create TimesTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDTimesTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created timestest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Times"]) != `[]` {
		t.Errorf("timestest Times is not [], but %#v", jsonify(x["Times"]))
	}

	if x["TimesNull"] != nil {
		t.Errorf("timestest TimesNull is %#v, but should be nil", x["TimesNull"])
	}

}

func TestTimesUpdate(t *testing.T) {
	id, _ := CRUDTimesTest.Create(db, b(`
	{
		"Times": ["2001-02-13T23:04:45Z","2011-02-13T23:04:45Z"]
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDTimesTest.Update(db, id, b(`
	{
		"Times": ["2001-04-13T23:04:45Z","2011-12-13T23:04:45Z"],
		"TimesNull": ["2006-02-13T23:04:45Z","2011-02-13T23:04:45Z"]
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update timestest with id %s: %s", id, err)
		return
	}

	x, err = CRUDTimesTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created timestest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Times"]) != `["2001-04-13T23:04:45Z","2011-12-13T23:04:45Z"]` {
		t.Errorf("timestest Times is not [\"2001-04-13T23:04:45Z\",\"2011-12-13T23:04:45Z\"], but %#v", jsonify(x["Times"]))
	}

	if jsonify(x["TimesNull"]) != `["2006-02-13T23:04:45Z","2011-02-13T23:04:45Z"]` {
		t.Errorf("timestest TimesNull is not [\"2006-02-13T23:04:45Z\",\"2011-02-13T23:04:45Z\"], but %#v", jsonify(x["TimesNull"]))
	}
}
