package typeconverter

import (
	js "encoding/json"
	xl "encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type Stringer interface {
	String() string
}
type StringType string

func String(s string) StringType { return StringType(s) }

func (ø StringType) String() string { return string(ø) }

func (ø StringType) Int() int {
	ii, err := strconv.ParseInt(ø.String(), 10, 32)
	if err != nil {
		var f float64
		f, err = strconv.ParseFloat(ø.String(), 64)
		if err != nil {
			panic("can't convert " + ø.String() + "to int")
		}
		return Float(f).Int()
	}
	return int(ii)
}

func (ø StringType) Float() float64 {
	f, err := strconv.ParseFloat(ø.String(), 64)
	if err != nil {
		panic("can't convert " + ø.String() + "to float64")
	}
	return f
}

func (ø StringType) Time() time.Time {
	tt, err := time.Parse(time.RFC3339, ø.String())
	if err != nil {
		panic("can't convert " + ø.String() + "to time")
	}
	return tt
}

func (ø StringType) Bool() bool {
	b, err := strconv.ParseBool(ø.String())
	if err != nil {
		panic("can't convert " + ø.String() + "to bool")
	}
	return b
}

func (ø StringType) Json() string {
	b, err := js.Marshal(ø.String())
	if err != nil {
		panic("can't convert " + fmt.Sprintf("%v", ø.String()) + " to json")
	}
	return string(b)
}

func (ø StringType) Xml() string {
	b, err := xl.Marshal(string(ø))
	if err != nil {
		panic("can't convert " + string(ø) + " to xml")
	}
	return string(b)
}

// checks, if something is a string or Stringer
func isString(i interface{}) bool {
	if _, ok := i.(string); ok {
		return true
	}

	if _, ok := i.(fmt.Stringer); ok {
		return true
	}

	return false
}
