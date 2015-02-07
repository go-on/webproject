package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type StringsTest struct {
	Id          *fat.Field `type:"string uuid" db:"id UUIDGEN PKEY"  rest:" R "`
	Strings     *fat.Field `type:"[]string"    db:"strings"          rest:"CRU"`
	StringsNull *fat.Field `type:"[]string"    db:"stringsnull NULL" rest:" RU"`
}

var STRINGS_TEST = fat.Proto(&StringsTest{}).(*StringsTest)
var CRUDStringsTest *CRUD

func init() {
	registry.MustRegisterTable("stringstest", STRINGS_TEST)

	db.Exec("DROP TABLE stringstest")

	stringsTestTable := registry.TableOf(STRINGS_TEST)
	_, e := db.Exec(stringsTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table stringstest: \nError: %s\nSql: %s\n", e.Error(),
			stringsTestTable.Create()))
	}

	CRUDStringsTest = NewCRUD(registry, STRINGS_TEST)
}

func TestStringsCreate(t *testing.T) {
	id, err := CRUDStringsTest.Create(db, b(`
	{
		"Strings": ["a u","b"]
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create StringsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDStringsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created stringstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Strings"]) != `["a u","b"]` {
		t.Errorf("stringstest Strings is not [\"a u\",\"b\"], but %#v", jsonify(x["Strings"]))
	}

	if x["StringsNull"] != nil {
		t.Errorf("stringstest StringsNull is %#v, but should be nil", x["StringsNull"])
	}

}

func TestStringsEmpty(t *testing.T) {
	id, err := CRUDStringsTest.Create(db, b(`
	{
		"Strings": []
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create StringsTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDStringsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created stringstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Strings"]) != `[]` {
		t.Errorf("stringstest Strings is not [], but %#v", jsonify(x["Strings"]))
	}

	if x["StringsNull"] != nil {
		t.Errorf("stringstest StringsNull is %#v, but should be nil", x["StringsNull"])
	}

}

func TestStringsUpdate(t *testing.T) {
	id, _ := CRUDStringsTest.Create(db, b(`
	{
		"Strings": ["c","b"]
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDStringsTest.Update(db, id, b(`
	{
		"Strings": ["d","g"],
		"StringsNull": ["a ","x"]
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update stringstest with id %s: %s", id, err)
		return
	}

	x, err = CRUDStringsTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created stringstest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Strings"]) != `["d","g"]` {
		t.Errorf("stringstest Strings is not [\"d\",\"g\"], but %#v", jsonify(x["Strings"]))
	}

	if jsonify(x["StringsNull"]) != `["a","x"]` {
		t.Errorf("stringstest StringsNull is not [\"a\",\"x\"], but %#v", jsonify(x["StringsNull"]))
	}
}
