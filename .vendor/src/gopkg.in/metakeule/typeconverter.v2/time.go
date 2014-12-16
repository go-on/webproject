package typeconverter

import (
	js "encoding/json"
	xl "encoding/xml"
	"time"
)

type Timer interface {
	Time() time.Time
}

type TimeType time.Time

func Time(t time.Time) TimeType { return TimeType(t) }

func (ø TimeType) String() string  { return time.Time(ø).Format(time.RFC3339) }
func (ø TimeType) Time() time.Time { return time.Time(ø) }
func (ø TimeType) Int() int        { return int(ø.Time().Unix()) }

func (ø TimeType) Json() string {
	b, err := js.Marshal(time.Time(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to json")
	}
	return string(b)
}

func (ø TimeType) Xml() string {
	b, err := xl.Marshal(time.Time(ø))
	if err != nil {
		panic("can't convert " + ø.String() + " to xml")
	}
	return string(b)
}
