package rest

import (
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/fat"

	// . "gopkg.in/metakeule/pgsql.v5"

	"testing"
)

type StringTest struct {
	Id         *fat.Field `type:"string uuid"         db:"id UUIDGEN PKEY" rest:" R "`
	String     *fat.Field `type:"string varchar(122)" db:"string"          rest:"CRU"`
	StringNull *fat.Field `type:"string text"         db:"stringnull NULL" rest:" RU"`
}

var STRING_TEST = fat.Proto(&StringTest{}).(*StringTest)
var CRUDStringTest *CRUD

func init() {
	registry.MustRegisterTable("stringtest", STRING_TEST)

	db.Exec("DROP TABLE stringtest")

	stringTestTable := registry.TableOf(STRING_TEST)
	_, e := db.Exec(stringTestTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table stringtest: \nError: %s\nSql: %s\n", e.Error(),
			stringTestTable.Create()))
	}

	CRUDStringTest = NewCRUD(registry, STRING_TEST)
}

func TestStringCreate(t *testing.T) {
	id, err := CRUDStringTest.Create(db, b(`
	{
		"String": "hello"
	}
 	`), false, "")

	if err != nil {
		t.Errorf("can't create StringTest: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var x map[string]interface{}

	x, err = CRUDStringTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created stringtest with id %s: %s", id, err)
		return
	}

	if x["String"] != "hello" {
		t.Errorf("stringtest String is not \"hello\", but %#v", x["String"])
	}

	if x["StringNull"] != nil {
		t.Errorf("stringtest StringNull is %#v, but should be nil", x["StringNull"])
	}

}

func TestStringUpdate(t *testing.T) {
	id, _ := CRUDStringTest.Create(db, b(`
	{
		"String": "hi"
	}
	`), false, "")

	var x map[string]interface{}
	err := CRUDStringTest.Update(db, id, b(`
	{
		"String": "hello",
		"StringNull": "world"
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update stringtest with id %s: %s", id, err)
		return
	}

	x, err = CRUDStringTest.Read(db, id)

	if err != nil {
		t.Errorf("can't get created stringtest with id %s: %s", id, err)
		return
	}

	if x["String"] != "hello" {
		t.Errorf("stringtest String is not \"hello\", but %#v", x["String"])
	}

	if x["StringNull"] != "world" {
		t.Errorf("stringtest StringNull is not \"world\", but %#v", x["StringNull"])
	}
}
