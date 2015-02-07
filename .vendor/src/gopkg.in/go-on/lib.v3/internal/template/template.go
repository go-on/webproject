package template

import (
	"bytes"
	"fmt"
	"gopkg.in/go-on/lib.v3/internal/template/placeholder"
	"io"
	"net/http"

	"gopkg.in/go-on/lib.v3/internal/replacer"
)

type (
	stringer interface {
		String() string
	}

	writerto interface {
		WriteTo(io.Writer) (int64, error)
	}
)

type Template struct {
	*Buffer
	repl *replacer.Template
}

func New(name string) *Template {
	t := &Template{}
	t.Buffer = newBuffer(name)
	return t
}

func (t *Template) New() *Buffer {
	return newBuffer(t.Buffer.name)
}

func (t *Template) WriteSetter(p placeholder.Setter) (err error) {
	// fmt.Printf("writing: %#v\n", replacer.Placeholder(p.Name()).String())
	_, err = t.Buffer.WriteString(replacer.Placeholder(p.Name()).String())
	return
}

func (t *Template) MustWriteSetter(s placeholder.Setter) {
	err := t.WriteSetter(s)
	if err != nil {
		panic(err.Error())
	}
}

func (t *Template) Add(data ...interface{}) (err error) {
	for _, d := range data {
		switch v := d.(type) {
		case placeholder.Setter:
			err = t.WriteSetter(v)
		case []byte:
			_, err = t.Buffer.Write(v)
		case string:
			_, err = t.Buffer.WriteString(v)
		case byte:

			err = t.Buffer.WriteByte(v)
		case rune:
			_, err = t.Buffer.WriteRune(v)
		case writerto:
			_, err = v.WriteTo(t.Buffer)
		case stringer:
			_, err = t.Buffer.WriteString(v.String())
		default:
			_, err = t.Buffer.WriteString(fmt.Sprintf("%v", v))
		}
		if err != nil {
			return
		}
	}
	return
}

func (t *Template) String() string {
	return t.Replace().String()
}

// add data to the template
func (t *Template) MustAdd(data ...interface{}) *Template {
	for _, d := range data {
		switch v := d.(type) {
		// case placeholder.Placeholder:
		// t.MustWriteString(v.Name())
		case placeholder.Setter:
			t.MustWriteSetter(v)
		case []byte:
			t.MustWrite(v)
		case string:
			t.MustWriteString(v)
		case byte:
			t.MustWriteByte(v)
		case rune:
			t.MustWriteRune(v)
		case writerto:
			_, err := v.WriteTo(t.Buffer)
			if err != nil {
				panic(err.Error())
			}
		case stringer:
			t.MustWriteString(v.String())
		case *Template:
			t.MustWriteSetter(t)
		default:
			panic(fmt.Sprintf("unknown type: %T", v))
			// t.MustWriteString(fmt.Sprintf("%v", v))
		}
	}
	return t
}

func (t *Template) MustWrite(b []byte) {
	_, err := t.Buffer.Write(b)
	if err != nil {
		panic(err.Error())
	}
}

func (t *Template) MustWriteString(s string) {
	_, err := t.Buffer.WriteString(s)
	if err != nil {
		panic(err.Error())
	}
}

func (t *Template) MustWriteByte(b byte) {
	err := t.Buffer.WriteByte(b)
	if err != nil {
		panic(err.Error())
	}
}

func (t *Template) MustWriteRune(r rune) {
	_, err := t.Buffer.WriteRune(r)
	if err != nil {
		panic(err.Error())
	}
}

func (t *Template) MustWriteTo(w io.Writer) {
	_, err := t.Buffer.WriteTo(w)
	if err != nil {
		panic(err.Error())
	}
}

func (t *Template) Parse() *Template {
	//t.repl.ParseBytes(t.Buffer.Bytes())
	t.repl = replacer.NewTemplateBytes(t.Buffer.Bytes())
	return t
}

func mixedSetters(mixed ...interface{}) (ss []placeholder.Setter) {
	for _, m := range mixed {
		switch v := m.(type) {
		case View:
			ss = append(ss, v.Set(v)...)
		case *View:
			ss = append(ss, v.Set(v)...)
		case placeholder.Setter:
			ss = append(ss, v)
		case []placeholder.Setter:
			ss = append(ss, v...)
		default:
			panic(fmt.Sprintf("unsupported type: %T, supported are: View, *View, Setter and []Setter", v))
		}
	}
	return
}

func (r *Template) ReplaceMixed(mixed ...interface{}) (bf *Buffer) {
	ss := mixedSetters(mixed...)
	return r.Replace(ss...)
}

func (r *Template) ReplaceMixedTo(b *bytes.Buffer, mixed ...interface{}) (bf *Buffer) {
	ss := mixedSetters(mixed...)
	return r.ReplaceTo(b, ss...)
}

func (r *Template) ReplaceTo(b *bytes.Buffer, setters ...placeholder.Setter) (bf *Buffer) {
	m := map[replacer.Placeholder]string{}
	for _, s := range setters {
		m[replacer.Placeholder(s.Name())] = s.SetString()
	}

	// fmt.Printf("replacemap: %v\n", m)
	b.Write(r.repl.ReplaceStrings(m))

	bf = &Buffer{Buffer: b, name: r.Buffer.Name()}
	return
}

type templateServer struct {
	t *Template
}

func mkHandler(ph placeholder.PlaceholderHandler) http.Handler {
	return http.HandlerFunc(
		func(wr http.ResponseWriter, rq *http.Request) {
			ph.ServeHTTP(wr, rq)
		})
}

// TODO check how it could work for subtemplates and multi replacements
// and how it could integrate with context and router and integrate it with html lib
func (t *Template) Handle(phs ...placeholder.PlaceholderHandler) http.Handler {
	m := make(map[string]http.Handler, len(phs))
	for _, ph := range phs {
		m[ph.Name()] = mkHandler(ph)
	}
	return t.repl.NewHandler(m)
}

func (t *Template) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Replace().ServeHTTP(w, r)
}

func (r *Template) Replace(setters ...placeholder.Setter) (bf *Buffer) {
	b := bytes.Buffer{}
	return r.ReplaceTo(&b, setters...)
}
