package person

import (
	"fmt"
)

type person interface {
	Prototype() *Person // should be implemented by all "subclasses", convert into a Person
}

type Person struct {
	*Person // recursiv inheritence to build a pattern
	Name    string
	object  person // keep a reference to the object, so we may Sync to it
}

func (ø *Person) Inherit(o person) { ø.object = o }

// Sync based on conversion based on Prototype()
func (ø *Person) Sync() {
	n := ø.object.Prototype()
	n.object = ø.object
	*(ø) = *n
}

// every method that accesses the objects properties has to call Sync first
// if the methods are called via the Prototype method, no Sync would be required
func (ø *Person) Hi()      { ø.Sync(); fmt.Print("Hi from " + ø.Name) }
func (ø *Person) GoodBye() { ø.Sync(); fmt.Print("Goodbye from " + ø.Name + "\n") }

type Singer Person

func (ø *Singer) Sing() { ø.Sync(); fmt.Println(ø.Name + " sings") }
