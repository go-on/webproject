package htmlfat

import (
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/replacer"
	"gopkg.in/go-on/lib.v2/internal/template/placeholder"
	"reflect"
	"strings"
	"sync"

	ph "gopkg.in/go-on/lib.v2/html/internal/placeholder"

	//html "gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/internal/fat"
	"gopkg.in/go-on/lib.v2/internal/meta"
	// html "gopkg.in/go-on/lib.v2/internal/shared"
	"gopkg.in/go-on/lib.v2/internal/template"
)

/*
   html support for fat fields
*/

type registry_ struct {
	*sync.RWMutex
	escapeRegistry map[string]string
}

var (
	registry = &registry_{&sync.RWMutex{}, map[string]string{}}

// maps struct-types to maps of field to escaper
)

func findEscaper(tag string) (escaperKey string) {
	for k := range ph.Escaper {
		if k == "" {
			continue
		}
		if k == "text" {
			continue
		}
		if strings.Contains(tag, k) {
			return k
		}
	}
	return "text"
}

// register a struct in the registry
// should be called at initialization time
func Register(østruct interface{}) {
	registry.Lock()
	defer registry.Unlock()
	fn := func(field *meta.Field) {
		f, ok := field.Value.Interface().(*fat.Field)
		if ok {
			if f.StructType() == "" {
				panic(fmt.Sprintf("struct %s has no prototype (not initialized with fat.Proto)", reflect.TypeOf(østruct).String()))
			}
			// println(f.Path())
			registry.escapeRegistry[f.Path()] = findEscaper(field.Type.Tag.Get("type"))
		}
	}
	stru, err := meta.StructByValue(reflect.ValueOf(østruct))

	if err != nil {
		panic(err.Error())
	}
	stru.Each(fn)
}

/*
func Escape(øfield *fat.Field, val interface{}) html.HTMLString {
	return html.HTMLString(Placeholder(øfield).Set(val).SetString())
}

func Escapef(øfield *fat.Field, format string, vals ...interface{}) html.HTMLString {
	return html.HTMLString(Placeholder(øfield).Setf(format, vals...).SetString())
}
*/

func Placeholder(øfield *fat.Field) string {
	return replacer.Placeholder(øfield.Path()).String()
}

func Setter(øfield *fat.Field) placeholder.Setter {
	registry.RLock()
	defer registry.RUnlock()
	ty, ok := registry.escapeRegistry[øfield.Path()]
	if !ok {
		panic(fmt.Sprintf("struct of field %s is not registered with html/fat", øfield.Path()))
	}
	_ = replacer.P
	ph := template.NewPlaceholder(øfield.Path(), ph.Escaper[ty])
	return ph.Set(øfield.String())
}

func Setters(østruct interface{}) (s []placeholder.Setter) {
	fn := func(field *meta.Field) {
		if f, ok := field.Value.Interface().(*fat.Field); ok {
			s = append(s, Setter(f))
		}
	}

	stru, err := meta.StructByValue(reflect.ValueOf(østruct))
	if err != nil {
		panic(err.Error())
	}

	stru.Each(fn)
	return
}

// TODO setters for each field of struct

/*
var Escaper = template.Escaper{
    "text":     handleStrings(html.EscapeString, true),
    "":         handleStrings(html.EscapeString, true),
    "html":     handleStrings(idem, true),
    "px":       units("%vpx"),
    "%":        units("%v%%"),
    "em":       units("%vem"),
    "pt":       units("%vpt"),
    "urlparam": handleStrings(url.QueryEscape, false),
}

type view struct {
    *template.View
}

type placeholder struct {
    template.Placeholder
}

func (p placeholder) String() string {
    return "@@" + p.Name() + "@@"
}

func (v *view) Placeholder(field string) placeholder {
    return placeholder{v.View.Placeholder(field)}
}

func View(stru interface{}, tag string) *view {
    return &view{Escaper.View(stru, tag)}
}
*/
