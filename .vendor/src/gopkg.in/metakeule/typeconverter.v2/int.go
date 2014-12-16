package typeconverter

import (
	xl "encoding/xml"
	"fmt"
	"time"
)

type Inter interface {
	Int() int
}

type IntType int

func Int(b int) IntType { return IntType(b) }

func (ø IntType) String() string  { return fmt.Sprintf("%v", ø.Int()) }
func (ø IntType) Int() int        { return int(ø) }
func (ø IntType) Float() float64  { return float64(ø) }
func (ø IntType) Json() string    { return ø.String() }
func (ø IntType) Time() time.Time { return time.Unix(int64(ø), 0) }

func (ø IntType) Xml() string {
	b, err := xl.Marshal(int(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}

type IntType64 int64

func Int64(b int64) IntType64 { return IntType64(b) }

func (ø IntType64) String() string  { return fmt.Sprintf("%v", ø.Int()) }
func (ø IntType64) Int() int        { return int(ø) }
func (ø IntType64) Float() float64  { return float64(ø) }
func (ø IntType64) Json() string    { return ø.String() }
func (ø IntType64) Time() time.Time { return time.Unix(int64(ø), 0) }

func (ø IntType64) Xml() string {
	b, err := xl.Marshal(int(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
