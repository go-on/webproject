package compiler

import (
	"bytes"
	"fmt"
	ht "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/internal/replacer"
	"go/ast"
	"net/http"
	"reflect"
	"strings"
)

type Dispatcher struct {
	methods map[string]reflect.Value
	typ     reflect.Type
}

var q http.HandlerFunc
var handlerfunctype = reflect.TypeOf(q)

func isHandlerFunc(val reflect.Value, method string) bool {
	if val.MethodByName(method).Type().ConvertibleTo(handlerfunctype) {
		return true
	}
	return false
}

type Preparer interface {
	Prepare(http.ResponseWriter, *http.Request)
}

func NewDispatcher(h Preparer) *Dispatcher {
	d := &Dispatcher{methods: map[string]reflect.Value{}}
	val := reflect.ValueOf(h)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("handler %T must be pointer type", h))
	}
	d.typ = val.Elem().Type()
	for i := 0; i < val.NumMethod(); i++ {
		m := val.Type().Method(i)
		if ast.IsExported(m.Name) && m.Name != "Prepare" && isHandlerFunc(val, m.Name) {
			d.methods[m.Name] = m.Func
		}
	}
	return d
}

func (d *Dispatcher) HTML(s string) replacer.Placeholder {
	return d.method(s)
}

func (d *Dispatcher) Text(s string) replacer.Placeholder {
	return d.method("text:" + s)
}

func (d *Dispatcher) DocHandler(doc *ht.DocType) http.Handler {
	var buffer bytes.Buffer
	h := mkPhHandler(doc.Element, "document", &buffer)
	r := replacer.NewTemplateString(doc.DocType + "\n" + buffer.String())
	return d.handler(r, h...)
}

func (d *Dispatcher) ElementHandler(name string, e *element.Element) http.Handler {
	var buffer bytes.Buffer
	h := mkPhHandler(e, name, &buffer)
	r := replacer.NewTemplateBytes(buffer.Bytes())
	return d.handler(r, h...)
}

func (d *Dispatcher) handler(r *replacer.Template, phs ...*placeholderHandler) http.Handler {
	m := make(map[string]http.Handler, len(phs))
	for _, ph := range phs {
		m[ph.Name()] = ph.Handler
	}
	return &templateDispatcher{r, d, m}
}

func (d *Dispatcher) checkMethod(method string) {
	if !ast.IsExported(method) {
		panic(`method "` + method + `" is not exported`)
	}
	_, has := d.methods[method]
	if !has {
		panic(`type "*` + d.typ.String() + `" has no method "` + method + `" that is a http.HandlerFunc`)
	}
}

func (d *Dispatcher) method(s string) replacer.Placeholder {
	if i := strings.Index(s, ":"); i > -1 {
		d.checkMethod(s[i+1:])
	} else {
		d.checkMethod(s)
	}

	return replacer.Placeholder(s)
}

func (d *Dispatcher) new() reflect.Value {
	return reflect.New(d.typ)
}
