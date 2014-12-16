package main

import (
	"fmt"
	conv "gopkg.in/metakeule/typeconverter.v2"
)

type MyString string

func (ø MyString) Int() int {
	return 42
}

func MyConverter() *conv.BasicConverter {
	var s string
	c := conv.New()

	// if the input is of the type string
	// transform it to MyString and call the output
	// dispatcher
	c.Input.SetHandler(s,
		func(in interface{}, out interface{}) (err error) {
			c.Output.Dispatch(out, MyString(in.(string)))
			return
		})

	// if input should be transformed to string
	// change the output and add " was the answer" to normal string conversion

	c.Output.SetHandler(&s,
		func(out interface{}, in interface{}) (err error) {
			*out.(*string) = in.(conv.Stringer).String() + " was the answer"
			return
		})
	return c
}

func main() {
	c := MyConverter()

	var i int
	c.Convert("the answer?", &i)
	fmt.Printf("the answer is: %#v\n", i) // the answer is: 42

	var s string
	c.Convert(42, &s)
	fmt.Println(s) // 42 was the answer

	var f float64
	err := c.Convert("won't work", &f)
	fmt.Println(err) // interface conversion: main.MyString is not typeconverter.Floater: missing method Float
}
