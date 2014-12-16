package typeconverter

import (
	xl "encoding/xml"
	"fmt"
)

type Booler interface {
	Bool() bool
}

type BoolType bool

func Bool(b bool) BoolType { return BoolType(b) }

func (ø BoolType) String() string { return fmt.Sprintf("%v", ø.Bool()) }
func (ø BoolType) Json() string   { return ø.String() }
func (ø BoolType) Bool() bool     { return bool(ø) }

func (ø BoolType) Xml() string {
	b, err := xl.Marshal(bool(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
