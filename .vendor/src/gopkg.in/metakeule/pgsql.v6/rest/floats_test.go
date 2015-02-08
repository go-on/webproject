package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v6"

	"testing"
)

type FloatsTest struct {
	Id         *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Floats     *fat.Field `type:"[]float"     db:"floats"          rest:"CRU"`
	FloatsNull *fat.Field `type:"[]float"     db:"floatsnull NULL" rest:" RU"`
}

var FLOATS_TEST = fat.Proto(&FloatsTest{}).(*FloatsTest)
var CRUDFloatsTest *CRUD

func init() {
	registry.MustRegisterTable("floatstest", FLOATS_TEST)

	DB.Exec("DROP TABLE floatstest")

	floatsTestTable := registry.TableOf(FLOATS_TEST)
	_, e := DB.Exec(floatsTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table floatstest: \nError: %s\nSql: %s\n", e.Error(),
			floatsTestTable.Create()))
	}

	CRUDFloatsTest = NewCRUD(registry, FLOATS_TEST)
}

func TestFloatsCreate(t *testing.T) {
	id, err := CRUDFloatsTest.Create(DB, b(`
	{
		"Floats": [2.5,6]
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create FloatsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDFloatsTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created floatstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Floats"]) != "[2.5,6]" {
		t.Errorf("floatstest Floats is not [2.5,6], but %#v", x["Floats"])
	}

	if x["FloatsNull"] != nil {
		t.Errorf("floatstest FloatsNull is %#v, but should be nil", x["FloatsNull"])
	}

}

func TestFloatsEmpty(t *testing.T) {
	id, err := CRUDFloatsTest.Create(DB, b(`
	{
		"Floats": []
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create FloatsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDFloatsTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created floatstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Floats"]) != "[]" {
		t.Errorf("floatstest Floats is not [], but %#v", x["Floats"])
	}

	if x["FloatsNull"] != nil {
		t.Errorf("floatstest FloatsNull is %#v, but should be nil", x["FloatsNull"])
	}

}

func TestFloatsUpdate(t *testing.T) {
	id, _ := CRUDFloatsTest.Create(DB, b(`
	{
		"Floats": [2.2,2.5]
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDFloatsTest.Update(DB, id, b(`
	{
		"Floats": [2.5,6],
		"FloatsNull": [2.5,2.2]
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update floatstest with id %s: %s", id, err)
		return
	}

	x, err = CRUDFloatsTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created floatstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Floats"]) != "[2.5,6]" {
		t.Errorf("floatstest Floats is not [2.5,6], but %#v", x["Floats"])
	}

	if jsonify(x["FloatsNull"]) != "[2.5,2.2]" {
		t.Errorf("floatstest FloatsNull is not [2.5,2.2], but %#v", x["FloatsNull"])
	}
}
