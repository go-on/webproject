package meta

import (
	"reflect"
)

func FinalType(i interface{}) reflect.Type {
	ty := reflect.TypeOf(i)
	if ty.Kind() == reflect.Ptr {
		return reflect.TypeOf(i).Elem()
	}
	return ty
}

func FinalValue(i interface{}) reflect.Value {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		return reflect.ValueOf(i).Elem()
	}
	return val
}
