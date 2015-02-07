package fat

import (
	"fmt"
)

type Type interface {
	Typ() string
	Get() interface{}  // must be typecasted
	String() string    // represent value as string
	Scan(string) error // sets the value based on a string
	Set(interface{}) error
}

// maps type to a generator
var newTypeFuncs = map[string]func() Type{
	"int":    func() Type { return Int(0) },
	"float":  func() Type { return Float(float64(0.0)) },
	"bool":   func() Type { return Bool(false) },
	"string": func() Type { return String("") },
	"time":   func() Type { return Time(zeroTime) },

	"[string]int":    func() Type { return newMap("int") },
	"[string]float":  func() Type { return newMap("float") },
	"[string]bool":   func() Type { return newMap("bool") },
	"[string]string": func() Type { return newMap("string") },
	"[string]time":   func() Type { return newMap("time") },

	"[]int":    func() Type { return newSlice("int") },
	"[]float":  func() Type { return newSlice("float") },
	"[]bool":   func() Type { return newSlice("bool") },
	"[]string": func() Type { return newSlice("string") },
	"[]time":   func() Type { return newSlice("time") },
}

func newType(typ string) Type {
	f, ok := newTypeFuncs[typ]
	if !ok {
		panic(fmt.Sprintf("can't create value of type: %s: unknown type", typ))
	}
	return f()
}

func newMap(typ string) Type   { return &map_{typ: typ, Map: map[string]Type{}} }
func newSlice(typ string) Type { return &slice{typ: typ, Slice: []Type{}} }
