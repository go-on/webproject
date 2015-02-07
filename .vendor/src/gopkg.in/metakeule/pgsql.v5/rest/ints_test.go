package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type IntsTest struct {
	Id       *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Ints     *fat.Field `type:"[]int"       db:"ints"            rest:"CRU"`
	IntsNull *fat.Field `type:"[]int"       db:"intsnull NULL"   rest:" RU"`
}

var INTS_TEST = fat.Proto(&IntsTest{}).(*IntsTest)
var CRUDIntsTest *CRUD

func init() {
	registry.MustRegisterTable("intstest", INTS_TEST)

	db.Exec("DROP TABLE intstest")

	intsTestTable := registry.TableOf(INTS_TEST)
	_, e := db.Exec(intsTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table intstest: \nError: %s\nSql: %s\n", e.Error(),
			intsTestTable.Create()))
	}

	CRUDIntsTest = NewCRUD(registry, INTS_TEST)
}

func TestIntsCreate(t *testing.T) {
	id, err := CRUDIntsTest.Create(db, b(`
	{
		"Ints": [2,3]
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create IntsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDIntsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created intstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Ints"]) != "[2,3]" {
		t.Errorf("intstest Ints is not [2,3], but %#v", x["Ints"])
	}

	if x["IntsNull"] != nil {
		t.Errorf("intstest IntsNull is %#v, but should be nil", x["IntsNull"])
	}

}

func TestIntsEmpty(t *testing.T) {
	id, err := CRUDIntsTest.Create(db, b(`
	{
		"Ints": []
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create IntsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDIntsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created intstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Ints"]) != "[]" {
		t.Errorf("intstest Ints is not [], but %#v", x["Ints"])
	}

	if x["IntsNull"] != nil {
		t.Errorf("intstest IntsNull is %#v, but should be nil", x["IntsNull"])
	}
}

func TestIntsUpdate(t *testing.T) {
	id, _ := CRUDIntsTest.Create(db, b(`
	{
		"Ints": [1,2]
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDIntsTest.Update(db, id, b(`
	{
		"Ints": [4,5],
		"IntsNull": [3,5]
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update intstest with id %s: %s", id, err)
		return
	}

	x, err = CRUDIntsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created intstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Ints"]) != "[4,5]" {
		t.Errorf("intstest Ints is not [4,5], but %#v", x["Ints"])
	}

	if jsonify(x["IntsNull"]) != "[3,5]" {
		t.Errorf("intstest IntsNull is not [3,5], but %#v", x["IntsNull"])
	}
}
