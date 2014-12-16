package main

import (
	"fmt"
	conv "gopkg.in/metakeule/typeconverter.v2"
)

type Special string

// fullfills Inter interface
func (ø Special) Int() int {
	return 42
}

// conversion method
// You might simply let the conversion panic, since Convert() catches any errors and returns them.
// But if you return your own error it will be passed through
func toSpecial(out interface{}, in interface{}) (err error) {
	*out.(*Special) = Special(in.(conv.Stringer).String())
	return
}

func main() {
	c := conv.New()

	sp := Special("")
	c.Output.SetHandler(&sp, toSpecial) // register output handler for the pointer (important) type

	var i int
	c.Convert(Special("hello"), &i)
	fmt.Printf("Special to int: %v\n", i) // Special to int: 42

	var s Special
	c.Convert(42, &s)
	fmt.Printf("42 to Special: %v\n", s) // 42 to Special: 42

	var str string
	err := c.Convert(Special("won't work"), &str)
	fmt.Println(err) // interface conversion: main.Special is not typeconverter.Stringer: missing method String
}
