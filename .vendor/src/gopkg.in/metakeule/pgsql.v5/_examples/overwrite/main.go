package main

import (
	"fmt"
	. "gopkg.in/metakeule/pgsql.v5/examples/overwrite/person"
)

type Beatle Person

// return me converted to my Prototype
func (ø *Beatle) Prototype() *Person { n := Person(*ø); return &n }

// use methods from another type that is based on the same prototype (Person)
func (ø *Beatle) AsSinger() *Singer { s := Singer(*ø); return &s }

func (ø *Beatle) Hi() {
	fmt.Print(ø.Name + " says: \"")
	ø.Person.Hi() // call the overwritten method, that has access to own properties
	fmt.Println("\"")
}

func NewBeatle(name string) (ø *Beatle) {
	// reference to Person instance is important, otherwise no Inherit method
	ø = &Beatle{Person: &Person{}, Name: name}
	ø.Inherit(ø) // important!!
	return
}

func main() {
	f := NewBeatle("John")
	f.Hi()
	f.GoodBye()
	f.Name = "Paul"
	f.Hi()
	f.AsSinger().Sing()
	f.GoodBye()
}
