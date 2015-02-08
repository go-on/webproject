package rest

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"os"

	. "gopkg.in/metakeule/pgsql.v6"
	"gopkg.in/metakeule/pgsql.v6/pgsqlfat"
	"gopkg.in/go-on/lib.v3/internal/fat"
	"gopkg.in/metakeule/dbwrap.v2"

	"strings"
	// "net/url"
	//"fmt"
	"testing"

	"gopkg.in/go-on/pq.v2"
)

func configureDB() string {
	dbconnectString := "postgres://docker:docker@127.0.0.1:5432/pgsqltest?sslmode=disable"
	if dbconn := os.Getenv("PG_TEST"); dbconn != "" {
		dbconnectString = dbconn
	}
	// fmt.Printf("PG_TEST is %#v\n", os.Getenv("PG_TEST"))
	// fmt.Printf("dbconnectString is set to %#v\n", dbconnectString)
	return dbconnectString
}

type testdrv struct {
	Query string
}

func (td testdrv) Open(connectString string) (driver.Conn, error) {
	return pq.Open(connectString)
}

var (
	//	db                *sql.DB
	wrapperDriverName = "dbtest"
	testdb            = testdrv{}
	dbconnectString   = configureDB()
)

func connect(driver string, str string) *sql.DB {
	cs, err := pq.ParseURL(str)
	if err != nil {
		panic(err.Error())
	}
	d, err := sql.Open(driver, cs)
	if err != nil {
		panic(err.Error())
	}
	return d
}

/*
types to check (everything also with NULL)

not null

id/uuid
id/serial

int
string/varchar
string/text
bool
float
time

[]int
[]string
[]bool
[]float
[]time


map[string]int
map[string]string
map[string]bool
map[string]float
map[string]time


null

int
string/varchar
string/text
bool
float
time

[]int
[]string
[]bool
[]float
[]time


map[string]int
map[string]string
map[string]bool
map[string]float
map[string]time

*/

type Company struct {
	Id        *fat.Field `type:"string uuid"        db:"id UUIDGEN PKEY" rest:" R DL"`
	Name      *fat.Field `type:"string varchar(66)" db:"name"            rest:"CRU L"`
	Age       *fat.Field `type:"int"                db:"age NULL"        rest:"CRU L"`
	UpdatedAt *fat.Field `type:"time timestamp"     db:"updated_at NULL"      rest:"CRU L"`
}

var COMPANY = fat.Proto(&Company{}).(*Company)
var CRUDCompany *CRUD
var _ = strings.Contains

func makeDB() *sql.DB {
	dbWrap := dbwrap.New(wrapperDriverName, testdb)

	dbWrap.HandlePrepare = func(conn driver.Conn, query string) (driver.Stmt, error) {
		testdb.Query = query
		/*		if strings.Contains(query, "Update") {
				fmt.Printf("-- Prepare --\n%s\n", query)
			}*/
		return conn.Prepare(query)
	}

	dbWrap.HandleExec = func(conn driver.Execer, query string, args []driver.Value) (driver.Result, error) {
		testdb.Query = query
		// fmt.Printf("-- Exec --\n%s\n", query)
		/*	if strings.Contains(query, "Update") {
			fmt.Printf("-- Exec --\n%s\n", query)
		}*/
		return conn.Exec(query, args)
	}

	return connect(wrapperDriverName, dbconnectString)
}

var DB = makeDB()
var _, _ = DB.Exec(`CREATE EXTENSION "uuid-ossp"`)

func init() {

	//db := makeDB()
	registry.MustRegisterTable("company", COMPANY)

	DB.Exec("DROP TABLE company")

	companyTable := registry.TableOf(COMPANY)
	_, e := DB.Exec(companyTable.Create().String())
	if e != nil {
		panic(fmt.Sprintf("Can't create table company: \nError: %s\nSql: %s\n", e.Error(), companyTable.Create()))
	}

	CRUDCompany = NewCRUD(registry, COMPANY)

}

/*
func parseQuery(q string) url.Values {
	vals, err := url.ParseQuery(q)
	if err != nil {
		panic(fmt.Sprintf("error in request url: %s", err))
	}
	return vals
}
*/

func b(in string) []byte { return []byte(in) }

func TestCRUDCreate(t *testing.T) {
	//id, err := CRUDCompany.Create(db, parseQuery(`Name=testcreate&Age=42&UpdatedAt=2013-12-12 02:10:02`))
	id, err := CRUDCompany.Create(DB, b(`
	{
		"Name": "testcreate",
		"Age": 42,
		"UpdatedAt": "2013-12-12T02:10:02Z"
  }
 	`), false, "")

	if err != nil {
		t.Errorf("can't create company: %s", err)
		return
	}

	if id == "" {
		t.Errorf("got empty id")
		return
	}

	var comp map[string]interface{}

	comp, err = CRUDCompany.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created company with id %s: %s", id, err)
		return
	}

	/*
		c, ok := comp.(*Company)

		if !ok {
			t.Errorf("result is no *Company, but %T", comp)
			return
		}
	*/
	if comp["Name"] != "testcreate" {
		t.Errorf("company name is not testcreate, but %#v", comp["Name"])
	}

	if comp["Age"].(int64) != 42 {
		t.Errorf("company Age is not 42, but %#v", comp["Age"])
	}

	if comp["UpdatedAt"] != "2013-12-12T02:10:02Z" {
		t.Errorf("company updatedat is not 2013-12-12T2:10:02Z, but %#v", comp["UpdatedAt"])
	}
}

func TestCRUDUpdate(t *testing.T) {
	//id, _ := CRUDCompany.Create(DB, parseQuery("Name=testupdate&Age=42&UpdatedAt=2013-12-12 02:10:02"))
	id, _ := CRUDCompany.Create(DB, b(`
	{
		"Name": "testupdate",
		"Age": 42,
		"UpdatedAt": "2013-12-12T02:10:02Z"
	}
	`), false, "")

	var comp map[string]interface{}
	//	fmt.Printf("uuid: %#v\n", id)

	//	err := CRUDCompany.Update(DB, id, parseQuery("Name=testupdatechanged&Age=43&Ratings=[0,6,7]&Tags=[\"a\",\"b\"]&UpdatedAt=2014-01-01 00:00:02"))

	//err := CRUDCompany.Update(DB, id, parseQuery("Name=testupdatechanged&Age=43&Ratings=[0,6,7]&Tags=[\"a\",\"b\"]"))
	/*
		err := CRUDCompany.Update(DB, id, b(`
		{
			"Name": "testupdatechanged",
			"Age": 43,
			"Ratings" : [0,6,7],
			"Tags": ["a","b"]
		}
		`))
	*/

	err := CRUDCompany.Update(DB, id, b(`
	{
		"Name": "testupdatechanged",
		"Age": 43
	}
	`), false, "")

	if err != nil {
		t.Errorf("can't update company with id %s: %s", id, err)
		return
	}

	comp, err = CRUDCompany.Read(DB, id)

	if err != nil {
		t.Errorf("can't get created company with id %s: %s", id, err)
		return
	}

	/*
		c, ok := comp.(*Company)

		if !ok {
			t.Errorf("result is no *Company, but %T", comp)
			return
		}
	*/

	if comp["Name"] != "testupdatechanged" {
		t.Errorf("company name is not testupdatechanged, but %#v", comp["Name"])
	}

	if comp["Age"].(int64) != 43 {
		t.Errorf("company age is not 43, but %#v", comp["Age"])
	}

	/*
		if c.UpdatedAt.String() != "2014-01-01 00:00:02" {
			t.Errorf("company UpdatedAt is not 2014-01-01 0:00:02, but %#v", c.UpdatedAt.String())
		}
	*/
}

func TestCRUDDelete(t *testing.T) {
	//id, _ := CRUDCompany.Create(DB, parseQuery("Name=testdelete&Age=42&UpdatedAt=2013-12-12 02:10:02"))
	id, _ := CRUDCompany.Create(DB, b(`
	{
		"Name": "testdelete",
		"Age": 42,
		"UpdatedAt": "2013-12-12T02:10:02Z"
	}
	`), false, "")
	err := CRUDCompany.Delete(DB, id)
	if err != nil {
		t.Errorf("can't delete company with id %s: %s", id, err)
		return
	}

	_, err = CRUDCompany.Read(DB, id)

	if err == nil {
		t.Errorf("can get deleted company with id %s, but should not", id)
		return
	}
}

func TestCRUDList(t *testing.T) {
	DB.Exec("delete from company")
	//	id1, _ := CRUDCompany.Create(DB, parseQuery("Name=testlist1&Age=42&UpdatedAt=2013-12-12 02:10:02"))
	id1, err := CRUDCompany.Create(DB, b(`
	{
		"Name": "testlist1",
		"Age": 42,
		"UpdatedAt": "2013-12-12T02:10:02Z"
	}
	`), false, "")

	if err != nil {
		panic(err.Error())
	}
	//	id2, _ := CRUDCompany.Create(DB, parseQuery("Name=testlist2&Age=43&UpdatedAt=2013-01-30 02:10:02"))
	//id2, _ := CRUDCompany.Create(DB, parseQuery("Name=testlist2&Age=43"))
	id2, err2 := CRUDCompany.Create(DB, b(`
	{
		"Name": "testlist2",
		"Age": 43
	}
	`), false, "")

	if err2 != nil {
		panic(err2.Error())
	}

	//CRUDCompany.Update(db, id1, parseQuery("Name=testlist1&Age=42&Ratings=[0,6,7]&Tags=[\"a\",\"b\"]&UpdatedAt=2014-01-03 02:10:02"))
	// CRUDCompany.Update(db, id1, b(`
	// {
	// 	"Name": "testlist1",
	// 	"Age": 42,
	// 	"Ratings": [0,6,7],
	// 	"Tags": ["a","b"],
	// 	"UpdatedAt": "2014-01-03 02:10:02"
	// }
	// `))

	CRUDCompany.Update(DB, id1, b(`
	{
		"Name": "testlist1",
		"Age": 42,
		"UpdatedAt": "2014-01-03T02:10:02Z"
	}
	`), false, "")

	//CRUDCompany.Update(db, id2, parseQuery("Name=testlist2&Age=43&Ratings=[6,7,8]"))
	// registry.Field

	var c *Company = &Company{}
	// ty := reflect.TypeOf(c)
	// ty.
	tyPath := pgsqlfat.TypeString(c) // ty.PkgPath() + ty.String()
	// println(tyPath, "vs", "*github.com/metakeule/pgsql/rest.Company")
	// "*github.com/metakeule/pgsql/rest.Company"
	companyNameField := registry.Field(tyPath, "Name")

	if companyNameField == nil {
		panic("can't find field for COMPANY.Name")
	}

	total, comps, err := CRUDCompany.List(DB, 10, ASC, companyNameField, 0)
	// comps, err := CRUDCompany.List(db, 10, OrderBy(companyNameField, ASC))

	if err != nil {
		t.Errorf("can't list created company with id1 %s and id2 %s: %s", id1, id2, err)
		return
	}

	c1 := comps[0] //.(*Company)
	c2 := comps[1] //.(*Company)

	if len(comps) != 2 {
		t.Errorf("results are not 2 companies, but %d", len(comps))
	}

	if total != 2 {
		t.Errorf("total results are not 2 companies, but %d", total)
	}

	/*
		if !ok1 || !ok2 {
			t.Errorf("results are no *Company, but %T and %T", comps[0], comps[1])
			return
		}
	*/

	if c1["Name"] != "testlist1" {
		t.Errorf("company 1 name is not testlist1, but %#v", c1["Name"])
	}

	if c2["Name"] != "testlist2" {
		t.Errorf("company 2 name is not testlist2, but %#v", c2["Name"])
	}

	if c1["Age"].(int64) != 42 {
		t.Errorf("company 1 age is not 42, but %#v", c1["Age"])
	}

	if c2["Age"].(int64) != 43 {
		t.Errorf("company 2 age is not 43, but %#v", c2["Age"])
	}

	/*
		if c1.Ratings.String() != "[0,6,7]" {
			t.Errorf("company 1 Ratings is not [0,6,7], but %#v", c1.Ratings.String())
		}

		if c1.Tags.String() != `["a","b"]` {
			t.Errorf("company 1 Tags is not [\"a\",\"b\"], but %#v", c1.Tags.String())
		}
	*/
	if c1["UpdatedAt"] != `2014-01-03T02:10:02Z` {
		t.Errorf("company 1 UpdatedAt is not 2014-01-03T02:10:02Z, but %#v", c1["UpdatedAt"])
	}

	// fmt.Printf("updatedat is set: %v\n", c2.UpdatedAt.IsSet)

	// fmt.Println(c2.UpdatedAt.String())
}

func jsonify(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}
