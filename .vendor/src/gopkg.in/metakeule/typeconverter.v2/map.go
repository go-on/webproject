package typeconverter

import (
	js "encoding/json"
	"fmt"
)

type Mapper interface {
	Map() map[string]interface{}
}

func Map(m map[string]interface{}) MapType { return MapType(m) }

type MapType map[string]interface{}

func (ø MapType) Map() map[string]interface{} { return map[string]interface{}(ø) }
func (ø MapType) String() string              { return ø.Json() }

func (ø MapType) Json() string {
	b, err := js.Marshal(ø)
	if err != nil {
		panic("can't convert " + fmt.Sprintf("%v", ø) + " to json")
	}
	return string(b)
}
