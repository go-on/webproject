package person

import (
	. "gopkg.in/metakeule/pgsql.v5"
)

var Id = NewField("Id", IntType, PrimaryKey|Serial)
var FirstName = NewField("FirstName", VarChar(123), NullAllowed)
var LastName = NewField("LastName", VarChar(125))
var Age = NewField("Age", IntType)
var Tags = NewField("Tags", StringsType, NullAllowed)
var Vita = NewField("Vita", TextType, NullAllowed, Selection("a", "b"))
var Person = NewTable("Person", Id, FirstName, LastName, Age, Vita, Tags)

func New(db DB) *Row { return NewRow(db, Person) }
