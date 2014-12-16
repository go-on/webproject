package typeconverter

import (
	js "encoding/json"
	"strings"
	"time"
)

type Jsoner interface {
	Json() string
}

func Json(s string) JsonType {
	if strings.Contains(s, "{") {
		m := map[string]interface{}{}
		err := js.Unmarshal([]byte(s), &m)
		if err != nil {
			panic("'" + s + "' is  not a valid json: " + err.Error())
		}
	}
	return JsonType(s)
}

type JsonType string

func (ø JsonType) Int() int {
	var i float64
	err := js.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to int")
	}
	return Float(i).Int()
}

func (ø JsonType) Float() float64 {
	var i float64
	err := js.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to float")
	}
	return i
}

func (ø JsonType) Time() time.Time {
	var i time.Time
	err := js.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to time")
	}
	return i
}

func (ø JsonType) String() string { return string(ø) }
func (ø JsonType) Json() string   { return string(ø) }

func (ø JsonType) Bool() bool {
	var i bool
	err := js.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to bool")
	}
	return i
}

func (ø JsonType) Array() []interface{} {
	ia := []interface{}{}
	err := js.Unmarshal([]byte(ø), &ia)
	if err != nil {
		panic("can't convert " + ø + " to array")
	}
	return ia
}

func (ø JsonType) Map() map[string]interface{} {
	ia := map[string]interface{}{}
	err := js.Unmarshal([]byte(ø), &ia)
	if err != nil {
		panic("can't convert " + ø + " to array")
	}
	return ia
}
