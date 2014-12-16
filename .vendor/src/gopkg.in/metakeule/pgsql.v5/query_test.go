package pgsql

import (
	"fmt"
	"strings"
	"testing"
)

var RoleId = NewField("Id", IntType, PrimaryKey|Serial)
var RoleName = NewField("Name", VarChar(123))
var ROLE_TABLE = NewTable("Role", RoleId, RoleName)

var Id = NewField("Id", IntType, PrimaryKey|Serial)
var FirstName = NewField("FirstName", VarChar(123), NullAllowed)
var LastName = NewField("LastName", VarChar(125))
var Age = NewField("Age", IntType)

var Role = NewField("Role", IntType, ROLE_TABLE.Field("Id"))
var Vita = NewField("Vita", TextType, NullAllowed, Selection("a", "b"))
var TABLE = NewTable("person", Id, FirstName, LastName, Age, Vita, Role)
var _ = fmt.Println

func normalize(in string) (out string) {
	out = strings.Replace(in, "\t", " ", -1)
	out = strings.Replace(out, "\n", " ", -1)
	out = strings.Replace(out, "  ", " ", -1)
	out = strings.Replace(out, "  ", " ", -1)
	out = strings.ToLower(out)
	return
}

func hasSql(s Sqler, subs string) bool {
	return strings.Contains(normalize(s.Sql().String()), subs)
}

func TestInsert(t *testing.T) {
	i := Insert(TABLE, Set(FirstName, "donald"))
	res := `insert into "person" ("firstname") values ($userinput$donald$userinput$::varchar(123))`
	if !hasSql(i, res) {
		err(t, "sql should contain insert statement", normalize(i.Sql().String()), res)
	}
}

func TestInsertMap(t *testing.T) {
	m := map[*Field]interface{}{FirstName: "donald"}
	i := InsertMap(TABLE, m)
	res := `insert into "person" ("firstname") values ($userinput$donald$userinput$::varchar(123))`
	if !hasSql(i, res) {
		err(t, "sql should contain insert (map) statement", normalize(i.Sql().String()), res)
	}
}

func TestDelete(t *testing.T) {
	i := Delete(TABLE, Where(Or(Equals(FirstName, "donald"), Equals(LastName, "duck"))))
	res := `delete  from "person" where ("person"."firstname" = $userinput$donald$userinput$::text) or ("person"."lastname" = $userinput$duck$userinput$::text) `
	if !hasSql(i, res) {
		err(t, "sql should contain delete statement", normalize(i.Sql().String()), res)
	}
}

func TestUpdate(t *testing.T) {
	i := Update(TABLE, Set(FirstName, "daisy"), Where(In(FirstName, "donald", "dagobert")))
	res := `update "person" set "firstname" = $userinput$daisy$userinput$::varchar(123) where "person"."firstname" in($userinput$donald$userinput$::text, $userinput$dagobert$userinput$::text) `
	if !hasSql(i, res) {
		err(t, "sql should contain update statement", normalize(i.Sql().String()), res)
	}
}

func TestSelect(t *testing.T) {
	i := Select(TABLE, FirstName, As(LastName, "Name", TextType), Where(In(FirstName, "donald", "dagobert")))
	res := `select "person"."firstname", "person"."lastname" as "name" from "person" where "person"."firstname" in($userinput$donald$userinput$::text, $userinput$dagobert$userinput$::text)`
	if !hasSql(i, res) {
		err(t, "sql should contain select statement", normalize(i.Sql().String()), res)
	}
}

func TestSelectLimit(t *testing.T) {
	i := Select(TABLE, FirstName, Limit(23))
	res := `select "person"."firstname" from "person" limit 23`
	if !hasSql(i, res) {
		err(t, "sql should contain select limit statement", normalize(i.Sql().String()), res)
	}
}

func TestSelectOrderBy(t *testing.T) {
	i := Select(TABLE, FirstName, LastName, OrderBy(LastName, ASC, FirstName, DESC))
	res := `select "person"."firstname", "person"."lastname" from "person" order by "person"."lastname" asc, "person"."firstname" desc`
	if !hasSql(i, res) {
		err(t, "sql should contain select limit statement", normalize(i.Sql().String()), res)
	}
}

func TestSelectGroupBy(t *testing.T) {
	i := Select(TABLE, As(Call("count", LastName), "no", IntType), Age, GroupBy(Age))
	res := `select "person"."age", count("person"."lastname") as "no" from "person" group by "person"."age"`
	if !hasSql(i, res) {
		err(t, "sql should contain select limit statement", normalize(i.Sql().String()), res)
	}
}

func TestSelectLeftJoin(t *testing.T) {
	i := Select(TABLE, FirstName, LastName, As(Sql(`"r"."Name"`), "Name", TextType), LeftJoin(Role, RoleId, "r"))
	res := `select "person"."firstname", "person"."lastname", "r"."name" as "name" from "person" left join "role" "r" on ("person"."role" = "r"."id")`
	if !hasSql(i, res) {
		err(t, "sql should contain select limit statement", normalize(i.Sql().String()), res)
	}
}
