package template

import (
	// "bytes"
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/template/placeholder"
	// "io"
	// "net/http"
	"reflect"
	"strings"

	// "github.com/go-on/replacer"
)

type Escaper map[string]func(interface{}) string

type View struct {
	_type         string
	_tag          string
	_placeholders map[string]placeholder.Placeholder
}

func (esc Escaper) View(stru interface{}, tag string) *View {
	s := &View{_type: structName(stru), _tag: tag}
	s.scanPlaceholders(stru, esc)
	return s
}

func structName(stru interface{}) string {
	return strings.Replace(fmt.Sprintf("%T", stru), "*", "", 1)
}

func (str *View) Tag() string  { return str._tag }
func (str *View) Type() string { return str._type }

func (str *View) Placeholder(field string) placeholder.Placeholder {
	p, ok := str._placeholders[field]
	if !ok {
		panic(fmt.Sprintf("no placeholder for field %s in struct %s (tag: %s)", field, str._type, str._tag))
	}
	return p
}

func (str *View) HasPlaceholder(field string) bool {
	_, ok := str._placeholders[field]
	return ok
}

func (str *View) Set(stru interface{}) (ss []placeholder.Setter) {
	if structName(stru) != str._type {
		panic(fmt.Sprintf("wrong type: %T, needed %s or *%s", stru, str._type, str._type))
	}
	for field, ph := range str._placeholders {
		f := _Field(stru, field)
		// we need to handle the nil pointers differently,
		// since they may be handled via interfaces
		// and then they are not nil
		if f.Kind() == reflect.Ptr && f.IsNil() {
			ss = append(ss, ph.Set(nil))
			continue
		}
		ss = append(ss, ph.Set(f.Interface()))
	}
	return
}

func (str *View) scanPlaceholders(stru interface{}, escaper Escaper) {
	str._placeholders = map[string]placeholder.Placeholder{}
	_EachRaw(stru,
		func(field reflect.StructField, v reflect.Value) {
			phName := fieldName(stru, field.Name, str._tag)
			// ph := NewPlaceholder(phName)
			ph := TemplatePlaceholder{name: phName}
			if t := field.Tag.Get(str._tag); t != "" {
				if t != "-" { // "-" signals ignorance
					for _, escaperKey := range strings.Split(t, ",") {
						escFunc, ok := escaper[escaperKey]
						if !ok {
							panic("unknown escaper " + escaperKey + " needed by " + phName)
						}
						ph.Escaper = append(ph.Escaper, escFunc)
					}
					str._placeholders[field.Name] = ph
				}
				return
			}

			escFunc, ok := escaper[""]
			if !ok {
				panic(`missing empty escaper (key: "") needed by ` + phName)
			}
			ph.Escaper = append(ph.Escaper, escFunc)
			str._placeholders[field.Name] = ph
		})
	return
}

/*
// returns a fieldname as used in placeholders of structs
func FieldName(stru interface{}, field string) string {
	f := meta.Struct.Field(stru, field)
	r := fieldName(stru, field)
	if f.Interface() == nil {
		panic("field does not exist: " + r)
	}
	return r
}
*/

func fieldName(stru interface{}, field string, tag string) string {
	return strings.Replace(fmt.Sprintf("%T.%s#%s", stru, field, tag), "*", "", 1)
}
