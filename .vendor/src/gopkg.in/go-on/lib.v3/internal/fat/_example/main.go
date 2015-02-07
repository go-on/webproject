package main

import (
	"fmt"
	"time"

	"github.com/go-on/fat"
	"github.com/go-on/template"
	"github.com/go-on/template/templfat"
	"gopkg.in/metakeule/fmtdate.v1"
)

type (
	Person struct {
		FirstName        *fat.Field `type:"string" enum:"Peter|Paul|Mary"`
		LastName         *fat.Field `type:"string"`
		Age              *fat.Field `type:"int" default:"32"`
		FieldsOfInterest *fat.Field `type:"[]string"`
		Points           *fat.Field `type:"[]float"`
		Votes            *fat.Field `type:"map[string]int"`
		Birthday         *fat.Field `type:"time"`
		Datings          *fat.Field `type:"[]time"`
		Meetings         *fat.Field `type:"map.time"`
	}
)

var (
	// prototype of a Person, used to get field infos, like placeholders
	// and to create new
	PERSON  = fat.Proto(&Person{}).(*Person)
	details = template.New("details").MustAdd("\n--------------\nDETAILS",
		"\nThe first name: ", templfat.Placeholder(PERSON.FirstName),
		"\nThe last name: ", templfat.Placeholder(PERSON.LastName),
		"\nThe age: ", templfat.Placeholder(PERSON.Age),
		"\nThe fields of interest: ", templfat.Placeholder(PERSON.FieldsOfInterest),
		"\nThe points: ", templfat.Placeholder(PERSON.Points),
		"\nThe votes: ", templfat.Placeholder(PERSON.Votes),
		"\nThe birthday: ", templfat.Placeholder(PERSON.Birthday),
		"\nThe datings: ", templfat.Placeholder(PERSON.Datings),
		"\nThe meetings: ", templfat.Placeholder(PERSON.Meetings),
		"\n-------------\n\n").Parse()
)

// use fat.New to create a new Person and have the field informations
// available via the prototype
func NewPerson() *Person { return fat.New(PERSON, &Person{}).(*Person) }

func main() {
	now := time.Now()
	bday, _ := fmtdate.Parse("DD.MM.YYYY", "02.01.1952")

	peter := NewPerson()
	peter.FirstName.MustSet("Peter")
	peter.Points.MustScanString(`[2,3,4.5]`)
	peter.LastName.MustScanString("Pan")
	peter.Votes.MustSet(fat.MapType("int", "Mary", 3, "Paul", 2))
	peter.Birthday.Set(bday)
	peter.Datings.Set(fat.Times(now, now.Add(2*time.Hour)))
	peter.Meetings.Set(fat.MapType("time", "Paul", now.Add(5*time.Hour)))
	peter.FieldsOfInterest.MustSet(fat.Strings("cooking", "swimming"))

	paul := NewPerson()
	paul.FirstName.MustSet("Paul")
	paul.LastName.MustSet("Panzer")
	paul.Age.MustSet(42)
	paul.Age.MustScanString("53")
	paul.Points.MustSet(fat.Floats(1.0, 2.3))
	paul.Votes.MustScanString(`{"Peter": 45}`)
	paul.Meetings.Set(fat.Map(map[string]time.Time{
		"Peter": now.Add(5 * time.Hour),
		"Mary":  now.Add(7 * time.Hour),
	}))

	fmt.Printf("%s: %s is set? %v\n", peter.FirstName.Name(), peter.FirstName, peter.FirstName.IsSet)
	fmt.Printf("%s: %s is set? %v\n", peter.Age.Name(), peter.Age, peter.Age.IsSet)

	fmt.Println(
		details.Replace(templfat.Setters(peter)...),
	)

	fmt.Printf("%s: %s is  set? %v\n", paul.FirstName.Name(), paul.FirstName, paul.FirstName.IsSet)
	fmt.Printf("%s: %s is set? %v\n", paul.Age.Name(), paul.Age, paul.Age.IsSet)
	fmt.Println(
		details.Replace(templfat.Setters(paul)...),
	)
}
