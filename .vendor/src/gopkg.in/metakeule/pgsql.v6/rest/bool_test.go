package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v6"
	"testing"
)

type BoolTest struct {
	Id       *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY" rest:" R "`
	Bool     *fat.Field `type:"bool"        db:"bool"            rest:"CRU"`
	BoolNull *fat.Field `type:"bool"        db:"boolnull NULL"   rest:" RU"`
}

var BOOL_TEST = fat.Proto(&BoolTest{}).(*BoolTest)
var CRUDBoolTest *CRUD

func init() {
	registry.MustRegisterTable("booltest", BOOL_TEST)

	DB.Exec("DROP TABLE booltest")

	boolTestTable := registry.TableOf(BOOL_TEST)
	_, e := DB.Exec(boolTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table booltest: \nError: %s\nSql: %s\n", e.Error(),
			boolTestTable.Create()))
	}

	CRUDBoolTest = NewCRUD(registry, BOOL_TEST)
}

func TestBoolCreate(t *testing.T) {
	id, err := CRUDBoolTest.Create(DB, b(`
	{
		"Bool": true
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create BoolTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDBoolTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created booltest with id %s: %s", id, err)
		return
	}

	if x["Bool"].(bool) != true {
		t.Errorf("booltest Bool is not true, but %#v", x["Bool"])
	}

	if x["BoolNull"] != nil {
		t.Errorf("booltest BoolNull is %#v, but should be nil", x["BoolNull"])
	}

}

func TestBoolUpdate(t *testing.T) {
	id, _ := CRUDBoolTest.Create(DB, b(`
	{
		"Bool": true
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDBoolTest.Update(DB, id, b(`
	{
		"Bool": false,
		"BoolNull": true
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update booltest with id %s: %s", id, err)
		return
	}

	x, err = CRUDBoolTest.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created booltest with id %s: %s", id, err)
		return
	}

	if x["Bool"].(bool) != false {
		t.Errorf("booltest Bool is not false, but %#v", x["Bool"])
	}

	if x["BoolNull"].(bool) != true {
		t.Errorf("booltest BoolNull is not true, but %#v", x["BoolNull"].(bool))
	}
}
