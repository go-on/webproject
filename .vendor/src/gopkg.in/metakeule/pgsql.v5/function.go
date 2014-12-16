package pgsql

import (
	"fmt"
	"strings"
)

type Function struct {
	Name     string
	Input    []Type
	Output   Type
	Body     Sqler
	Language Sqler
}

func (ø *Function) Sql() SqlType {
	return Sql(ø.Name)
}

func (ø *Function) Call(p ...Sqler) SqlType {
	params := []string{}
	for _, sq := range p {
		params = append(params, sq.Sql().String())
	}
	return Sql(fmt.Sprintf("%s(%s)", ø.Name, strings.Join(params, ", ")))
}

func (ø *Function) Create() SqlType {
	s := `
CREATE OR REPLACE FUNCTION %s(%s)
  RETURNS %s AS
$$
%s
$$
  LANGUAGE %s`

	language := "sql"
	if ø.Language != nil {
		language = ø.Language.Sql().String()
	}

	input := []string{}
	for i, in := range ø.Input {
		input = append(input, fmt.Sprintf("p%v %s", i, in.String()))
	}

	return Sql(fmt.Sprintf(s, ø.Name, strings.Join(input, ","), ø.Output.String(), ø.Body.Sql().String(), language))
}
