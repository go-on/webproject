package pgsql

import (
	"fmt"
)

type Trigger struct {
	Name      string
	Condition Sqler
	Table     *Table
	Body      Sqler
}

func (ø *Trigger) Sql() SqlType {
	return Sql(ø.Name)
}

func (ø *Trigger) Create() SqlType {
	s := `CREATE TRIGGER %s %s ON %s %s`
	return Sql(fmt.Sprintf(s, ø.Name, ø.Condition.Sql().String(), ø.Table.Sql().String(), ø.Body.Sql().String()))
}
