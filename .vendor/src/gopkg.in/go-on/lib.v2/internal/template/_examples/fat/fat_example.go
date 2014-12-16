package main

import (
	"bytes"
	"fmt"

	"github.com/go-on/fat"
	"github.com/go-on/template"
	f "github.com/go-on/template/templfat"
)

type Person struct {
	FirstName *fat.Field `type:"string"`
	LastName  *fat.Field `type:"string"`
	Age       *fat.Field `type:"int"`
	Vita      *fat.Field `type:"string"`
}

var PERSON = fat.Proto(&Person{}).(*Person)

func NewPerson() (p *Person) { return fat.New(PERSON, &Person{}).(*Person) }

var detailsView = template.New("person-details").MustAdd(
	"------------\n",
	`Name:`, f.Placeholder(PERSON.LastName), ", ", f.Placeholder(PERSON.FirstName), "\n",
	`Age:`, f.Placeholder(PERSON.Age), "\n",
	`Vita:`, f.Placeholder(PERSON.Vita), "\n",
).Parse()

func PersonDetails(ps ...*Person) (bf *bytes.Buffer) {
	switch len(ps) {
	case 0:
		return
	case 1:
		return detailsView.Replace(f.Setters(ps[0])...).Buffer
	default:
		bf = &bytes.Buffer{}
		for _, pp := range ps {
			detailsView.ReplaceTo(bf, f.Setters(pp)...)
		}
		return
	}
}

func main() {
	pete := NewPerson()
	pete.FirstName.Set("Pete")
	pete.LastName.Set("Norman")
	pete.Age.Set(55)
	pete.Vita.Set(`I was never born.`)

	paul := NewPerson()
	paul.FirstName.Set("Paul")
	paul.LastName.Set("Simon")
	paul.Age.Set(65)
	paul.Vita.Set(`I was born.`)

	fmt.Println(PersonDetails(pete))
	fmt.Println(PersonDetails(paul, pete))
}
