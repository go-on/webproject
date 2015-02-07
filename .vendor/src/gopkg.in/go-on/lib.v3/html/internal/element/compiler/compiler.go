package compiler

import (
	"bytes"
	"net/http"
	"reflect"
	"strings"

	ht "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	"gopkg.in/go-on/lib.v3/internal/replacer"
	"gopkg.in/go-on/wrap.v2"
)

type templateDispatcher struct {
	*replacer.Template
	*Dispatcher
	replacements map[string]http.Handler
}

func (tw *templateDispatcher) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var last int
	var placeholder string
	var escaper string
	state := tw.Dispatcher.new()
	state.MethodByName("Prepare").Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(req)})

	for _, place := range tw.Template.Places {
		placeholder = place.Placeholder.Name()
		w.Write(tw.Original[last:place.Pos])
		last = place.Pos
		if handler, exists := tw.replacements[placeholder]; exists {
			handler.ServeHTTP(w, req)
			continue
		}

		if i := strings.Index(placeholder, ":"); i > -1 {
			escaper = placeholder[:i]
			placeholder = placeholder[i+1:]
			m, has := tw.methods[placeholder]
			if !has {
				continue
			}
			switch escaper {
			case "text":
				m.Call([]reflect.Value{
					state,
					reflect.ValueOf(&wrap.EscapeHTML{w}),
					reflect.ValueOf(req),
				})
				continue
			}
		}
		if m, has := tw.methods[placeholder]; has {
			m.Call([]reflect.Value{state, reflect.ValueOf(w), reflect.ValueOf(req)})
		}
	}
	w.Write(tw.Original[last:tw.Length])
}

type templateHandler struct {
	*replacer.Template
	replacements map[string]http.Handler
}

func (tw *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var last int
	var placeholder string

	for _, place := range tw.Template.Places {
		placeholder = place.Placeholder.Name()
		w.Write(tw.Original[last:place.Pos])
		last = place.Pos
		if handler, exists := tw.replacements[placeholder]; exists {
			handler.ServeHTTP(w, req)
		}
	}
	w.Write(tw.Original[last:tw.Length])
}

func DocHandler(doc *ht.DocType) http.Handler {
	var buffer bytes.Buffer
	h := mkPhHandler(doc.Element, "document", &buffer)
	// fmt.Printf("length dochandlers: %d\n", len(h))
	r := replacer.NewTemplateString(doc.DocType + "\n" + buffer.String())
	return handler(r, h...)
}

func ElementHandler(name string, e *element.Element) http.Handler {
	var buffer bytes.Buffer
	h := mkPhHandler(e, name, &buffer)
	r := replacer.NewTemplateBytes(buffer.Bytes())
	return handler(r, h...)
}

func handler(r *replacer.Template, phs ...*placeholderHandler) http.Handler {
	m := make(map[string]http.Handler, len(phs))
	for _, ph := range phs {
		m[ph.Name()] = ph.Handler
	}
	return &templateHandler{r, m}
}
