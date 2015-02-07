package template

import (
	"fmt"
	"reflect"
)

// copied from github.com/metakeule/meta

func _Kind(i interface{}) reflect.Kind {
	return reflect.TypeOf(i).Kind()
}

func _Is(k reflect.Kind, i interface{}) bool {
	//fmt.Println(ø.Kind(i))
	return _Kind(i) == k
}

func _FinalValue(i interface{}) reflect.Value {
	if _Is(reflect.Ptr, i) {
		return reflect.ValueOf(i).Elem()
	}
	return reflect.ValueOf(i)
}

func _FinalType(i interface{}) reflect.Type {
	if _Is(reflect.Ptr, i) {
		return reflect.TypeOf(i).Elem()
	}
	return reflect.TypeOf(i)
}

func _Field(s interface{}, field string) (f reflect.Value) {
	fv := _FinalValue(s)
	f = fv.FieldByName(field)
	return
}

func _Inspect(i interface{}) (s string) {
	if reflect.TypeOf(i).Kind().String() == "float64" || reflect.TypeOf(i).Kind().String() == "float32" {
		s = fmt.Sprintf("%f (%s)", i, reflect.TypeOf(i))
	} else {
		s = fmt.Sprintf("%#v (%s)", i, reflect.TypeOf(i))
	}
	return
}

func _Panicf(s string, i ...interface{}) {
	panic(fmt.Sprintf(s, i...))
}

func _EachRaw(s interface{}, fn func(field reflect.StructField, val reflect.Value)) {
	if s == nil {
		return
	}
	ft := _FinalType(s)
	fv := _FinalValue(s)
	if !(fv.Type().Kind() == reflect.Struct) {
		_Panicf("%s is not a struct / pointer to a struct", _Inspect(s))
	}

	elem := ft.NumField()
	for i := 0; i < elem; i++ {
		fn(ft.Field(i), fv.Field(i))
	}
	return
}
