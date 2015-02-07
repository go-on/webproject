package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"
	"testing"
)

type BoolsTest struct {
	Id        *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Bools     *fat.Field `type:"[]bool"      db:"bools"           rest:"CRU"`
	BoolsNull *fat.Field `type:"[]bool"      db:"boolsnull NULL"  rest:" RU"`
}

var BOOLS_TEST = fat.Proto(&BoolsTest{}).(*BoolsTest)
var CRUDBoolsTest *CRUD

func init() {
	registry.MustRegisterTable("boolstest", BOOLS_TEST)

	db.Exec("DROP TABLE boolstest")

	boolsTestTable := registry.TableOf(BOOLS_TEST)
	_, e := db.Exec(boolsTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table boolstest: \nError: %s\nSql: %s\n", e.Error(),
			boolsTestTable.Create()))
	}

	CRUDBoolsTest = NewCRUD(registry, BOOLS_TEST)
}

func TestBoolsCreate(t *testing.T) {
	id, err := CRUDBoolsTest.Create(db, b(`
	{
		"Bools": [true,false]
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create BoolsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDBoolsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created boolstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Bools"]) != "[true,false]" {
		t.Errorf("boolstest Bools is not [true, false], but %#v", jsonify(x["Bools"]))
	}

	if x["BoolsNull"] != nil {
		t.Errorf("boolstest BoolsNull is %#v, but should be nil", x["BoolsNull"])
	}

}

func TestBoolsEmpty(t *testing.T) {
	id, err := CRUDBoolsTest.Create(db, b(`
	{
		"Bools": []
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create BoolsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDBoolsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created boolstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Bools"]) != "[]" {
		t.Errorf("boolstest Bools is not [], but %#v", jsonify(x["Bools"]))
	}

	if x["BoolsNull"] != nil {
		t.Errorf("boolstest BoolsNull is %#v, but should be nil", x["BoolsNull"])
	}

}

func TestBoolsUpdate(t *testing.T) {
	id, _ := CRUDBoolsTest.Create(db, b(`
	{
		"Bools": [true,false]
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDBoolsTest.Update(db, id, b(`
	{
		"Bools": [false,true],
		"BoolsNull": [true,true]
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update boolstest with id %s: %s", id, err)
		return
	}

	x, err = CRUDBoolsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created boolstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Bools"]) != "[false,true]" {
		t.Errorf("boolstest Bools is not [false,true], but %#v", x["Bools"])
	}

	if jsonify(x["BoolsNull"]) != "[true,true]" {
		t.Errorf("boolstest BoolsNull is not [true,true], but %#v", x["BoolsNull"])
	}
}
