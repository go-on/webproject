package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type IntTest struct {
	Id      *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Int     *fat.Field `type:"int"         db:"int"             rest:"CRU"`
	IntNull *fat.Field `type:"int"         db:"intnull NULL"    rest:" RU"`
}

var INT_TEST = fat.Proto(&IntTest{}).(*IntTest)
var CRUDIntTest *CRUD

func init() {
	registry.MustRegisterTable("inttest", INT_TEST)

	db.Exec("DROP TABLE inttest")

	intTestTable := registry.TableOf(INT_TEST)
	_, e := db.Exec(intTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table inttest: \nError: %s\nSql: %s\n", e.Error(),
			intTestTable.Create()))
	}

	CRUDIntTest = NewCRUD(registry, INT_TEST)
}

func TestIntCreate(t *testing.T) {
	id, err := CRUDIntTest.Create(db, b(`
	{
		"Int": 2
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create IntTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDIntTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created inttest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Int"]) != "2" {
		t.Errorf("inttest Int is not 2, but %#v", x["Int"])
	}

	if x["IntNull"] != nil {
		t.Errorf("inttest IntNull is %#v, but should be nil", x["IntNull"])
	}

}

func TestIntUpdate(t *testing.T) {
	id, _ := CRUDIntTest.Create(db, b(`
	{
		"Int": 2
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDIntTest.Update(db, id, b(`
	{
		"Int": 3,
		"IntNull": 4
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update inttest with id %s: %s", id, err)
		return
	}

	x, err = CRUDIntTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created inttest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Int"]) != "3" {
		t.Errorf("inttest Int is not 3, but %#v", x["Int"])
	}

	if jsonify(x["IntNull"]) != "4" {
		t.Errorf("inttest IntNull is not 4, but %#v", x["IntNull"])
	}
}
