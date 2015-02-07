package benchmark

import (
	"bytes"
	"text/template"
)

type t struct {
	*template.Template
}

func NewTemplate() *t {
	return &t{}
}

const tname = "t-1"

func (ø *t) Parse(s string) (err error) {
	ø.Template, err = template.New(tname).Parse(s)
	return
}

func (ø *t) Replace(data map[string]string, buf *bytes.Buffer) error {
	return ø.Template.Execute(buf, data)
}
