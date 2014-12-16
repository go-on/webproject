package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type MapTest struct {
	Id      *fat.Field `type:"int"            db:"id SERIAL PKEY" rest:" R "`
	Map     *fat.Field `type:"[string]string" db:"map"            rest:"CRU"`
	MapNull *fat.Field `type:"[string]int"    db:"mapnull NULL"   rest:" RU"`
}

var MAP_TEST = fat.Proto(&MapTest{}).(*MapTest)
var CRUDMapTest *CRUD

func init() {
	registry.MustRegisterTable("maptest", MAP_TEST)

	db.Exec("DROP TABLE maptest")

	mapTestTable := registry.TableOf(MAP_TEST)
	_, e := db.Exec(mapTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table maptest: \nError: %s\nSql: %s\n", e.Error(),
			mapTestTable.Create()))
	}

	CRUDMapTest = NewCRUD(registry, MAP_TEST)
}

func TestMapCreate(t *testing.T) {
	id, err := CRUDMapTest.Create(db, b(`
	{
		"Map": {"a":"b"}
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create MapTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDMapTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created maptest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Map"]) != `{"a":"b"}` {
		t.Errorf("maptest Map is not {\"a\":\"b\"}, but %#v", x["Map"])
	}

	if x["MapNull"] != nil {
		t.Errorf("maptest MapNull is %#v, but should be nil", x["MapNull"])
	}

}

func TestMapUpdate(t *testing.T) {
	id, _ := CRUDMapTest.Create(db, b(`
	{
		"Map": {"b":"c"}
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDMapTest.Update(db, id, b(`
	{
		"Map": {"d":"e"},
		"MapNull": {"f":5}
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update maptest with id %s: %s", id, err)
		return
	}

	x, err = CRUDMapTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created maptest with id %s: %s", id, err)
		return
	}

	if jsonify(x["Map"]) != `{"d":"e"}` {
		t.Errorf("maptest Map is not {\"d\":\"e\"}, but %#v", x["Map"])
	}

	if jsonify(x["MapNull"]) != `{"f":5}` {
		t.Errorf("maptest MapNull is not {\"f\":5}, but %#v", x["MapNull"])
	}
}
