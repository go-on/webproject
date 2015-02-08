package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"

	"github.com/go-on/fat"
	"gopkg.in/go-on/router.v2"
	"gopkg.in/metakeule/pgsql.v6"
	"gopkg.in/metakeule/pgsql.v6/pgsqlfat"
	"gopkg.in/metakeule/pgsql.v6/rest"
)

/*
type matching should be performed the following way:

fat.Field looks if type tag matches
  []string, []int, []time, []float
or
  map.string map.int map.time map.float
or
  string, int, time, float

pgsql looks if type tag matches
  []int32, []int16, []int8, []int, []float, []string, []date, []timetz, []time
or
  xml, json,
or
  uuid, varchar(x), char(x), text, int32, int16, int8, int, float, string, date, timetz, time

goh4 looks, if type tag matches
  class, html, text, string (-> text)

fat.Field respects sprintf tag when String() and scanf tag when ScanString()

furthermore, fat.Field looks for nil to get info, if nil is allowed (Get() returns then
nil, if value is not set and there is no default (we should add a IsNil() method))

pgsql looks also for nil matches to find out, if the field may have null values

*/
/*
type Person_ struct {
	Id         *fat.Field `type:"string uuid"   db:"id UUIDGEN PKEY"                rest:"CDL"  template:"customers"`
	FirstName  *fat.Field `type:"string"        db:"firstname"       index:"name.1" rest:"CRUL"                    enum:"Paul|Benny|Nadja" `
	LastName   *fat.Field `type:"string"        db:"lastname"        index:"name.0" rest:"CRU"  template:"customers"`
	Age        *fat.Field `type:"int"           db:"age"                            rest:"RU"                      default:"32" `
	Interests  *fat.Field `type:"[]string"      db:"interests NULL"                             template:"customers"`
	Vita       *fat.Field `type:"string text"   db:"vita NULL"                                            `
	Class      *fat.Field `type:"string class"  db:"class"                                      template:"customers" sprintf:"db-person-%s" scanf:"db-person-%s"`
	Company    *fat.Field `type:"string uuid"   db:"company fkey.company.id"                         `
	CustomerNo *fat.Field `type:"int"           db:"customer_no unique serial"      rest:"RU"                      default:"32" `
}
*/

type Person struct {
	Id               *fat.Field `type:"string uuid"        db:"id UUIDGEN PKEY"         rest:" R DL"`
	FirstName        *fat.Field `type:"string varchar(66)" db:"firstname"               rest:"CRU L" enum:"Paul|Benny|Nadja" `
	LastName         *fat.Field `type:"string varchar(80)" db:"lastname"                rest:"CRU  "`
	Age              *fat.Field `type:"int"                db:"age"                     rest:" RU"   default:"32"`
	FieldsOfInterest *fat.Field `type:"[]string"           db:"field_of_interest NULL"`
	Vita             *fat.Field `type:"string text"        db:"vita NULL"`
}

type Corporation struct {
	Id        *fat.Field `type:"string uuid"        db:"id UUIDGEN PKEY" rest:" R DL"`
	Name      *fat.Field `type:"string varchar(66)" db:"name"            rest:"CRU L"`
	FoundedAt *fat.Field `type:"time date"          db:"founded_at NULL" rest:"CRU L"`
}

/*
CREATE TABLE company
(
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name character varying(66) NOT NULL,
  age integer,
  ratings integer[],
  tags character varying[],
  updated_at timestamp without time zone NOT NULL,
  CONSTRAINT company_pkey PRIMARY KEY (id)
)
*/

/*
type Company struct {
	Id      *fat.Field `type:"string"       db:"id"                pgsql.type:"uuid"        pgsql.flags:"uuidgen,pkey" rest:"d,r,l"`
	Name    *fat.Field `type:"string"       db:"firstname"         pgsql.type:"varchar(66)"                            rest:"c,u,r,l" fat.enum:"Paul|Benny|Nadja" `
	Age     *fat.Field `type:"int"          db:"age"                                                                   rest:"r,u"     fat.default:"32"`
	Ratings *fat.Field `type:"[]string" db:"field_of_interest"                          pgsql.flags:"NULL"`
	Tags *fat.Field `type:"[]string" db:"field_of_interest"                          pgsql.flags:"NULL"`
	Ratings *fat.Field `type:"[]string" db:"field_of_interest"                          pgsql.flags:"NULL"`
}
*/

var PERSON = fat.Proto(&Person{}).(*Person)
var CORPORATION = fat.Proto(&Corporation{}).(*Corporation)
var registry = pgsqlfat.NewRegistries()

func init() {
	CORPORATION.Name.Validator = fat.Validaters(fat.StringMustNotBeEmpty)

	registry.MustRegisterTable("person", PERSON)
	registry.MustRegisterTable("corporation", CORPORATION)

}

func main() {
	RESTRouter := router.NewETagged()
	REST := rest.NewREST(DB, registry, RESTRouter)

	REST.Mount(
		PERSON, "person",
		rest.CREATE|rest.READ|rest.UPDATE|rest.LIST,
		nil)

	REST.Mount(
		CORPORATION, "corporation",
		rest.CREATE|rest.READ|rest.UPDATE|rest.DELETE|rest.LIST,
		rest.MaxLimit(10).
			SetSortFields(
			CORPORATION.Id,
			CORPORATION.Name,
			CORPORATION.FoundedAt,
		),
	)

	/*
		RESTPerson := rest.NewCRUD(PERSON).Mount(DB, RESTRouter, "person", rest.MaxLimit(10))
		RESTPerson.ReadRoute()
		RESTPerson.ListRoute()
		RESTPerson.UpdateRoute()
		RESTPerson.CreateRoute()

		RESTCorporation := rest.NewCRUD(CORPORATION).Mount(DB, RESTRouter, "corporation",
			rest.MaxLimit(10).SetSortFields(CORPORATION.Id, CORPORATION.Name, CORPORATION.FoundedAt))
		RESTCorporation.ReadRoute()
		RESTCorporation.DeleteRoute()
		RESTCorporation.ListRoute()
		RESTCorporation.UpdateRoute()
		RESTCorporation.CreateRoute()

	*/
	router.MustMount("/api/v1", RESTRouter)

	//rack.New(h, ...)

	http.ListenAndServe(":8080", nil)
}

func initCorporation() {
	corporationTable := registry.TableOf(CORPORATION)

	_, e := DB.Exec(corporationTable.Create().String())
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	r := pgsql.NewRow(DB, corporationTable)
	r.Debug = true
	e = r.Set(registry.FieldOf(CORPORATION.Name), "Know")
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	e = r.Save()

	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

}

func insertCompany() {
	corporationTable := registry.TableOf(CORPORATION)

	r := pgsql.NewRow(DB, corporationTable)
	r.Debug = true
	e := r.Set(registry.FieldOf(CORPORATION.Name), "Stridor")
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	e = r.Save()

	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	fmt.Printf("inserted corporation with id %s\n", r.GetString(registry.FieldOf(CORPORATION.Id)))
}

func insertCompany2() {
	/*
			CURRVAL('sequence-name')

		users_id_seq

		pg_get_serial_sequence()

		SELECT currval(pg_get_serial_sequence('users', 'id'));
	*/
	/*
		tx, e1 := DB.Begin()

		if e1 != nil {
			fmt.Printf("can't start transaction: %s\n", e1.Error())
			return
		}
	*/
	//	_, err := tx.Exec("insert into corporation (name) values ('test1')")
	r := DB.QueryRow("insert into corporation (name) values ('test1') RETURNING(id)")
	var id string
	err := r.Scan(&id)

	if err != nil {
		fmt.Printf("Error while inserting: %s\n", err.Error())
		return
	}

	fmt.Printf("new uuid: %s\n", id)

	/*
		r := tx.QueryRow("SELECT currval(pg_get_serial_sequence('corporation', 'id'))")
		var id string
		e2 := r.Scan(&id)

		if e2 != nil {
			fmt.Printf("can't scan id: %s\n", e2.Error())
			return
		}
	*/
}

func main2a() {
	rt := router.New()

	prest := rest.NewREST(DB, registry, rt).Mount(PERSON, "person", rest.READ, nil)
	pget := prest[rest.READ]

	router.MustMount("/api/v1", rt)

	fmt.Println(pget.MustURL("person_id", "9d12b0e6-773f-432a-b31a-a77a87dbd7d1"))

	http.ListenAndServe(":8080", nil)
}

//func NewPerson() *Person { return fat.New(PERSON, &Person{}).(*Person) }

func main2() {
	_ = fmt.Println

	//personTable := pgsql.TableOf(PERSON)
	/* fmt.Println(personTable.Create()) */
	/*
		_, e := DB.Exec(personTable.Create().String())
		if e != nil {
			fmt.Printf("Error: %s\n", e.Error())
			return
		}

		r := pgsql.NewRow(DB, personTable)
		r.Debug = true
		e = r.Set(
			rest.FieldOf(PERSON.FirstName), "Benny",
			rest.FieldOf(PERSON.LastName), "Bergheim",
			rest.FieldOf(PERSON.Age), 41)
		if e != nil {
			fmt.Printf("Error: %s\n", e.Error())
			return
		}

		e = r.Save()

		if e != nil {
			fmt.Printf("Error: %s\n", e.Error())
			return
		}
	*/
	//mt.Printf("%v", personTable.PrimaryKey)

	rt := router.New()

	prest := rest.NewCRUD(registry, PERSON).Mount(DB, rt, "person", nil)
	pget := prest.ReadRoute()

	router.MustMount("/api/v1", rt)

	fmt.Println(pget.MustURL("person_id", "9d12b0e6-773f-432a-b31a-a77a87dbd7d1"))

	http.ListenAndServe(":8080", nil)

	/*
		var p = NewPerson()
		e := rest.Get(DB, "9d12b0e6-773f-432a-b31a-a77a87dbd7d1", p)

		if e != nil {
			fmt.Printf("Error: %s\n", e.Error())
			return
		}

		bt, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		fmt.Printf("%s\n", bt)
	*/
}
