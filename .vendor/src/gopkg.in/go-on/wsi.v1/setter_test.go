package wsi

import (
	"fmt"
	"testing"
)

func TestSetter(t *testing.T) {
	var i int
	var s string
	var f float64
	var b bool

	tests := []struct {
		input       interface{}
		setter      Setter
		expectedErr bool
		expected    string
	}{
		// correct types
		{&i, SetInt(4), false, "4"},
		{&i, SetInt(40), false, "40"},
		{&b, SetBool(false), false, "false"},
		{&b, SetBool(true), false, "true"},
		{&s, SetString("hi"), false, "hi"},
		{&f, SetFloat64(3.4), false, "3.4"},

		// false types
		{&b, SetInt(40), true, "40"},
		{&i, SetBool(true), true, "true"},
		{&f, SetString("hi"), true, "hi"},
		{&i, SetFloat64(3.4), true, "3.4"},
	}

	for _, test := range tests {
		err := test.setter.Set(test.input)
		// if got, want := err, test.err; got != want {
		if test.expectedErr && err == nil {
			t.Errorf("%#v.Set(%T) = nil, want error", test.setter, test.input)
		}

		if !test.expectedErr && err != nil {
			t.Errorf("%#v.Set(%T) = %T, want nil", test.setter, test.input, err)
		}

		if err == nil {
			if got, want := toStr(test.input), test.expected; got != want {
				t.Errorf("var x %T; %#v.Set(x); x = %s, want %s", test.input, test.setter, got, want)
			}
		}
	}

}

func toStr(i interface{}) string {

	switch t := i.(type) {
	case *int:
		return fmt.Sprintf("%v", *t)
	case *string:
		return fmt.Sprintf("%v", *t)
	case *bool:
		return fmt.Sprintf("%v", *t)
	case *float64:
		return fmt.Sprintf("%v", *t)
	default:
		return fmt.Sprintf("%v", t)
	}

}
