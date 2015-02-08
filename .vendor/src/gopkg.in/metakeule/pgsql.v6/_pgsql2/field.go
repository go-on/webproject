package pgsql2

type Field struct {
	Name       string
	flags      flag
	Default    *Value
	Type       Type
	Table      *Table
	ForeignKey IsField
}

type Value struct {
	Field IsField
	Value interface{}
}

func (ø *value) Validate() error {
	return ø.Field.Validate(ø.Value)
}

func Value_(field IsField, val interface{}) (ø *Value, err error) {
	ø = &Value{field, val}
	err = ø.Validate()
	if err != nil {
		ø = nil
	}
	return
}

/*
	every field of a table must be a type (struct),
	that inherits from the Field struct and by that fullfills the
	IsField interface
*/

var fName = &Field{Name: "firstname"}

type FirstName string

func FirstName_(val string) (ø *FirstName) {
	return &FirstName{fName, val}
}

var aGe = &Field{Name: "age"}

type Age struct {
	*Field
	Value int
}

func Age_(val int) (ø *Age) {
	return &Age{fName, val}
}

// here the table definition
type Table struct {
}

type Person struct {
	*Table
	FirstName *FirstName
	Age       *Age
}
