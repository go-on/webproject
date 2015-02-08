package pgsql

import (
	"database/sql"
	"fmt"
	"gopkg.in/metakeule/typeconverter.v2"
	"strings"
	"testing"
)

var _ = typeconverter.String

type FakeDB struct {
	TransactionStarted bool
	LastQuery          string
	LastQueryParams    []interface{}
}

func (ø *FakeDB) LastInsertId() (id int64, err error) {
	id = 1
	return
}

func (ø *FakeDB) RowsAffected() (id int64, err error) {
	id = 1
	return
}

func (ø *FakeDB) Close() error {
	return nil
}

func (ø *FakeDB) Begin() (tx *sql.Tx, ſ error) {
	return
}

func (ø *FakeDB) Exec(query string, args ...interface{}) (s sql.Result, err error) {
	ø.LastQuery = query
	ø.LastQueryParams = args
	s = ø
	return
}

func (ø *FakeDB) Query(query string, args ...interface{}) (s *sql.Rows, err error) {
	ø.LastQuery = query
	ø.LastQueryParams = args
	s = &sql.Rows{}
	return
}

func (ø *FakeDB) QueryRow(query string, args ...interface{}) (s *sql.Row) {
	ø.LastQuery = query
	ø.LastQueryParams = args
	s = &sql.Row{}
	return
}

func (ø *FakeDB) Prepare(query string) (s *sql.Stmt, err error) {
	ø.LastQuery = query
	ø.LastQueryParams = []interface{}{}
	s = &sql.Stmt{}
	return
}

var Fake = &FakeDB{}

func lastQuery() string {
	return Fake.LastQuery
}

type hasLastname string

func (ø hasLastname) ValidateRow(vals map[*Field]interface{}) error {
	if vals[LASTNAME] == nil {
		return nil
	}
	if ln := vals[LASTNAME].(typeconverter.Stringer).String(); ln != string(ø) {
		return fmt.Errorf("does not have lastname %s but %s", ø, ln)
	}
	return nil
}

type hasAge int

func (ø hasAge) ValidateRow(vals map[*Field]interface{}) error {
	if vals[AGE] == nil {
		return nil
	}
	if ln := vals[AGE].(typeconverter.Inter).Int(); ln != int(ø) {
		return fmt.Errorf("does not have AGE %v but %v", ø, ln)
	}
	return nil
}

var ID = NewField("Id", IntType, PrimaryKey|Serial)
var FIRSTNAME = NewField("FirstName", VarChar(123), NullAllowed)
var LASTNAME = NewField("LastName", VarChar(25))
var AGE = NewField("Age", IntType, NullAllowed)
var VITA = NewField("Vita", TextType, NullAllowed, Selection("a", "b"))
var PERSON = NewTable("person", ID, FIRSTNAME, LASTNAME, AGE, VITA)

func init() {
	ln := hasLastname("Duck")
	ag := hasAge(22)
	schema := NewSchema("test", PERSON)
	PERSON.Schema = schema
	PERSON.AddValidator(OrRowValidator{&ln, &ag})
}

func q() {
	fmt.Println(Fake.LastQuery)
}

func has(contained string) bool {
	return strings.Contains(lastQuery(), contained)
}

func NewPerson() *Row { return NewRow(Fake, PERSON) }

/*
func TestRowInsert(t *testing.T) {
	p := NewPerson()
	p.Set(FIRSTNAME, "Donald", LASTNAME, "Duck")
	p.Save()
	if !has("INSERT") {
		q()
		err(t, "insert statement should contain INSERT", lastQuery(), "INSERT")
	}
}
*/

/*
func TestRowUpdate(t *testing.T) {
	p := NewPerson()
	p.Set(ID, 2, FIRSTNAME, "Donald", LASTNAME, "Duck")
	p.Save()
	if !has("UPDATE") {
		q()
		err(t, "update statement should contain UPDATE", lastQuery(), "UPDATE")
	}
}
*/

func TestRowDelete(t *testing.T) {
	p := NewPerson()
	p.Set(ID, 2)
	p.Delete()
	if !has("DELETE") {
		q()
		err(t, "delete statement should contain DELETE", lastQuery(), "DELETE")
	}
}

func TestRowSelect(t *testing.T) {
	p := NewPerson()
	p.Select(FIRSTNAME, LASTNAME, Limit(12))
	if !has("SELECT") {
		q()
		err(t, "select statement should contain SELECT", lastQuery(), "SELECT")
	}
}

/*
func TestRowValidation(t *testing.T) {
	p := NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückD", FIRSTNAME, "12", VITA, "a", AGE, 22)
	errs := p.ValidateAll()
	if len(errs) > 0 {
		err(t, "should have no validation errors", errs, nil)
	}

	p = NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückD", FIRSTNAME, "12", VITA, "a")
	errs = p.ValidateAll()
	if len(errs) > 0 {
		err(t, "should have no validation errors (no age)", errs, nil)
	}

	p = NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückD")
	errs = p.ValidateAll()
	if len(errs) > 0 {
		err(t, "should have no validation errors (just lastname)", errs, nil)
	}

	p = NewPerson()
	p.Set(FIRSTNAME, "12", VITA, "a", AGE, 22)
	errs = p.ValidateAll()
	if len(errs) != 1 {
		err(t, "should have one validation error", errs, `Validation Error in "person"."LastName": nil (null) is not allowed`)
	}

	p = NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückDü", FIRSTNAME, "12", VITA, "a", AGE, 22)
	errs = p.ValidateAll()
	if len(errs) != 1 {
		err(t, "should have one validation error", errs, `Validation Error in "person"."LastName": "DückDückDückDückDückDückDü" with length 26 is too long for a varchar(25)`)
	}

	p = NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückD", FIRSTNAME, "12", VITA, "a", AGE, 21)
	errs = p.ValidateAll()
	if len(errs) != 1 {
		err(t, "should have one validation error", errs, `Validation Error in "person": does not have AGE 22 but 21`)
	}

	p = NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückD", FIRSTNAME, "12", VITA, "hu")
	errs = p.ValidateAll()
	if len(errs) != 1 {
		err(t, "should have one validation error", errs, `Validation Error in "person"."Vita": "hu" is not in the selection: pgsql.Selection{"a", "b"}`)
	}

	p = NewPerson()
	p.Set(LASTNAME, "DückDückDückDückDückDückD", FIRSTNAME, 12)
	e := p.Save()
	if e == nil {
		q()
		err(t, "should have one validation error", e, `error when setting field "person"."FirstName" to value 12: value 12 type int is incompatible with type varchar(123)`)
	}
}
*/
