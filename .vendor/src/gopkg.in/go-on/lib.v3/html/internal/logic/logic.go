package logic

import (
	. "gopkg.in/go-on/lib.v3/html/internal/element"
)

func IF_(condition func() bool, ifData interface{}) interface{} {
	if condition() {
		return ifData
	}
	return nil
}

func IF_ELSE_(condition func() bool, ifData, elseData interface{}) interface{} {
	if condition() {
		return ifData
	}
	return elseData
}

func TIMES_(no int, fn func(n int) interface{}) *Element {
	elems := []interface{}{}
	for i := 0; i < no; i++ {
		elems = append(elems, fn(i))
	}
	return Elements(elems...)
}
