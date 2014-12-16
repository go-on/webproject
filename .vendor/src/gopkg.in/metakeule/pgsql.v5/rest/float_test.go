package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type FloatTest struct {
	Id        *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Float     *fat.Field `type:"float"       db:"float"           rest:"CRU"`
	FloatNull *fat.Field `type:"float"       db:"floatnull NULL"  rest:" RU" `
}

var FLOAT_TEST = fat.Proto(&FloatTest{}).(*FloatTest)
var CRUDFloatTest *CRUD

func init() {
	registry.MustRegisterTable("floattest", FLOAT_TEST)

	db.Exec("DROP TABLE floattest")

	floatTestTable := registry.TableOf(FLOAT_TEST)
	_, e := db.Exec(floatTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table floattest: \nError: %s\nSql: %s\n", e.Error(),
			floatTestTable.Create()))
	}

	CRUDFloatTest = NewCRUD(registry, FLOAT_TEST)
}

func TestFloatCreate(t *testing.T) {
	id, err := CRUDFloatTest.Create(db, b(`
	{
		"Float": 2.5
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create FloatTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDFloatTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created floattest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Float"]) != "2.5" {
		t.Errorf("floattest Float is not 2.5, but %#v", x["Float"])
	}

	if x["FloatNull"] != nil {
		t.Errorf("floattest FloatNull is %#v, but should be nil", x["FloatNull"])
	}

}

func TestFloatUpdate(t *testing.T) {
	id, _ := CRUDFloatTest.Create(db, b(`
	{
		"Float": 2.2
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDFloatTest.Update(db, id, b(`
	{
		"Float": 6,
		"FloatNull": 2.5
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update floattest with id %s: %s", id, err)
		return
	}

	x, err = CRUDFloatTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created floattest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Float"]) != "6" {
		t.Errorf("floattest Float is not 6, but %#v", x["Float"])
	}

	if jsonify(x["FloatNull"]) != "2.5" {
		t.Errorf("floattest FloatNull is not 2.5, but %#v", x["FloatNull"])
	}
}
