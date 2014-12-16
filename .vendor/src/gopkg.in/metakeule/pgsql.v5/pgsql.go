package pgsql

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Database struct {
	Name    string
	Schemas []*Schema
}

func ToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
