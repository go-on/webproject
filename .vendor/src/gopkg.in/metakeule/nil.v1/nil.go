package nil

import (
	"fmt"
	"reflect"
)

func IsNilable(obj interface {
	Kind() reflect.Kind
}) bool {
	switch obj.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	default:
		return false
	}
}

// returns true, if the inner value of i is nil
// otherwise false
func Is(i interface{}) bool {
	return i == nil || (IsNilable(reflect.ValueOf(i)) && reflect.ValueOf(i).IsNil())
}

func IsNilValue(v reflect.Value) bool {
	return IsNilable(v) && v.IsNil()
}

// panics if the inner value of i is nil
func MustNot(i interface{}) {
	if IsNil(i) {
		panic(fmt.Sprintf("%#v (%T) must not be nil", i, i))
	}
}

// convenience function to be used instead of MustNot
// if the package is included with a dot
func Nø(i interface{}) { MustNot(i) }

// convenience function to be used instead of Is
// if the package is included with a dot
func IsNil(i interface{}) bool { return Is(i) }
