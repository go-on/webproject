package typeconverter

import (
	js "encoding/json"
	xl "encoding/xml"
)

type Arrayer interface {
	Array() []interface{}
}

func Array(a []interface{}) ArrayType { return ArrayType(a) }

type ArrayType []interface{}

func (ø ArrayType) Array() (a []interface{}) {
	a = ø
	return
}
func (ø ArrayType) String() string { return ø.Json() }

func (ø ArrayType) Json() string {
	b, err := js.Marshal(ø)
	if err != nil {
		panic("can't convert " + ø.String() + " to json")
	}
	return string(b)
}

func (ø ArrayType) Xml() string {
	b, err := xl.Marshal(ø)
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
