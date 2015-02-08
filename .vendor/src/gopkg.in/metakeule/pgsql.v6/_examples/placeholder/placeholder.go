package main

import (
	"database/sql"
	"fmt"
	. "gopkg.in/metakeule/pgsql.v6"
	"gopkg.in/metakeule/pgsql.v6/examples/person"
)

var (
	Age__             = person.Age.Placeholder()
	LastName__        = person.LastName.Placeholder()
	Tag__             = person.Tags.Placeholder()
	SearchFirstName__ = SearchEnd("searchFirstName").Placeholder()
)

func update(row *Row) {
	q := MustCompile(row.UpdateQuery())
	fmt.Println(
		q.New().Replace(
			LastName__.Set(`$userinput$hih%so`),
			Age__.Set(34),
		),
	)
}

func insert(row *Row) {
	q := MustCompile(row.InsertQuery())
	fmt.Println(
		q.New().Replace(
			LastName__.Set(`bunny`),
			Age__.Set(134),
		),
	)
}

func remove() {
	// row.DeleteQuery is not support, since it is not easy to
	// pass an id as placeholder and you can do this with normal Delete query
	q := MustCompile(
		Delete(
			person.Person,
			Where(
				And(
					Equals(person.Age, Age__),
					Equals(person.LastName, LastName__),
					Like(person.FirstName, SearchFirstName__),
				),
			),
		),
	)
	fmt.Println(
		q.New().Replace(
			Age__.Set(42),
			// uncomment placeholder to get an error
			// of a missing placeholder
			LastName__.Set("Poetschki"),
			Tag__.Set("funny"),
			SearchFirstName__.Set("Nadja"),
		),
	)
}

func selecT() {
	q := MustCompile(
		Select(
			person.Person, person.FirstName, person.LastName, person.Vita,
			Where(
				And(
					Equals(person.Age, Age__),
					Equals(person.LastName, LastName__),
					Like(person.FirstName, SearchFirstName__),
				),
			),
		),
	)

	fmt.Println(
		q.New().Replace(
			Age__.Set(32),
			LastName__.Set("$userinput$Arns"),
			Tag__.Set("fun"),
			SearchFirstName__.Set("Benny"),
		),
	)
}

func main() {
	row := person.New(&sql.DB{})
	row.SetId("4")
	err := row.Set(person.LastName, LastName__, person.Age, Age__)
	if err != nil {
		fmt.Println(err.Error())
	}

	insert(row)
	update(row)
	remove()
	selecT()
}
