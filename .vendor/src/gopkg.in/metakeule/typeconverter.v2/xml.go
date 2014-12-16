package typeconverter

import (
	xl "encoding/xml"
	_ "fmt"
	"strings"
	"time"
)

type Xmler interface {
	Xml() string
}

func Xml(s string) XmlType {
	var str string
	if s != "" {
		err := xl.NewDecoder(strings.NewReader(s)).Decode(&str)
		if err != nil {
			panic("'" + s + "' is  not a valid xml: " + err.Error())
		}
	}
	return XmlType(s)
}

type XmlType string

func (ø XmlType) String() string { return string(ø) }
func (ø XmlType) Xml() string    { return string(ø) }

func (ø XmlType) Int() int {
	var i float64
	err := xl.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to int")
	}
	return Float(i).Int()
}

func (ø XmlType) Float() float64 {
	var i float64
	err := xl.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to float")
	}
	return i
}

func (ø XmlType) Time() time.Time {
	var i time.Time
	err := xl.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to time")
	}
	return i
}

func (ø XmlType) Bool() bool {
	var i bool
	err := xl.Unmarshal([]byte(ø), &i)
	if err != nil {
		panic("can't convert " + ø + " to bool")
	}
	return i
}

type arrayXmlHelper struct {
	Int     []int
	Float64 []float64
	String  []string
	Time    []time.Time
}

func (ø XmlType) Array() []interface{} {
	var err error
	ia := []interface{}{}
	ah := arrayXmlHelper{}
	err = xl.Unmarshal([]byte("<arrayXmlHelper>"+ø+"</arrayXmlHelper>"), &ah)
	if err != nil {
		panic("can't convert " + ø + " to array")
	}
	for _, v := range ah.Int {
		ia = append(ia, v)
	}
	for _, v := range ah.Float64 {
		ia = append(ia, v)
	}
	for _, v := range ah.String {
		ia = append(ia, v)
	}
	for _, v := range ah.Time {
		ia = append(ia, v)
	}
	return ia
}
