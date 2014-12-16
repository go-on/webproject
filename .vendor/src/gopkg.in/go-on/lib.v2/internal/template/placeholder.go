package template

import (
	// "bytes"
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/template/placeholder"
	"io"
	// "net/http"
	// "reflect"
	// "strings"

	// "github.com/go-on/replacer"
)

type TemplatePlaceholder struct {
	name, Value string
	Escaper     []func(interface{}) string
}

func NewPlaceholder(name string, escaper ...func(interface{}) string) placeholder.Placeholder {
	return TemplatePlaceholder{name: name, Escaper: escaper}
}

func (p TemplatePlaceholder) Name() string { return p.name }

func (p TemplatePlaceholder) Set(val interface{}) placeholder.Setter {
	var value string
	if len(p.Escaper) == 0 {
		// panic("no escaper")
		value = fmt.Sprintf("%v", val)
	}

	for i, esc := range p.Escaper {
		if i == 0 {
			value = esc(val)
			continue
		}
		value = esc(value)
	}
	// fmt.Printf("value for %#v is: %#v\n", p.name, value)
	return TemplatePlaceholder{name: p.name, Value: value, Escaper: p.Escaper}
}

func (p TemplatePlaceholder) Setf(format string, vals ...interface{}) placeholder.Setter {
	return p.Set(fmt.Sprintf(format, vals...))
}

func (p TemplatePlaceholder) SetString() string {
	return p.Value
}

func (p TemplatePlaceholder) WriteTo(w io.Writer) (int64, error) {
	// fmt.Printf("writing %s\n", p.Value)
	i, err := w.Write([]byte(p.Value))
	return int64(i), err
}
