package meta

import (
	"fmt"
	"reflect"
)

// stolen from https://ahmetalpbalkan.com/blog/golang-take-slices-of-any-type-as-input-parameter/
func SliceArg(arg interface{}) (out []interface{}) {
	slice := takeArg(arg, reflect.Slice)
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value) {
	val = reflect.ValueOf(arg)
	if val.Kind() != kind {
		panic(fmt.Sprintf("%T is not of kind %s, but %s", arg, kind.String(), val.Kind().String()))
	}
	return
}
