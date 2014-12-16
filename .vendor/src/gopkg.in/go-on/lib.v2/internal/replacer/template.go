package replacer

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

var (
	delimiterStartString = "\u2423\u2023" // ␣‣
	delimiterEndString   = "\u220e\u2423" // ∎␣
	delimiterStart       = []byte(delimiterStartString)
	delimiterEnd         = []byte(delimiterEndString)
	delimiterStartLen    = len(delimiterStart)
	delimiterEndLen      = len(delimiterEnd)
)

type (
	Place struct {
		Pos         int
		Placeholder Placeholder
	}

	Places []Place

	Template struct {
		Original []byte
		Places   Places
		Length   int
	}
)

// fullfill sort.Interface.
func (p Places) Len() int           { return len(p) }
func (p Places) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Places) Less(i, j int) bool { return p[i].Pos < p[j].Pos }

func NewTemplateBytes(b []byte) *Template {
	r := &Template{}
	r.parseBytes(b)
	return r
}

func NewTemplateString(s string) *Template {
	r := &Template{}
	r.parseString(s)
	return r
}

func (r *Template) Placeholders() (p []Placeholder) {
	uniq := map[Placeholder]bool{}
	p = []Placeholder{}
	for _, place := range r.Places {
		if !uniq[place.Placeholder] {
			p = append(p, place.Placeholder)
			uniq[place.Placeholder] = true
		}
	}
	return
}

func (r *Template) ReplaceStrings(replacements map[Placeholder]string) (res []byte) {
	var last int
	res = make([]byte, 0, r.Length)
	//res = []byte{}
	for _, place := range r.Places {
		res = append(res, r.Original[last:place.Pos]...)
		res = append(res, replacements[place.Placeholder]...)
		last = place.Pos
	}
	res = append(res, r.Original[last:r.Length]...)
	return
}

func (r *Template) ReplaceBytes(replacements map[Placeholder][]byte) (res []byte) {
	var last int
	//res = []byte{}
	res = make([]byte, 0, r.Length)
	for _, place := range r.Places {
		res = append(res, r.Original[last:place.Pos]...)
		res = append(res, replacements[place.Placeholder]...)
		last = place.Pos
	}
	res = append(res, r.Original[last:r.Length]...)
	return
}

func (t *Template) NewSetter() *Setter {
	return newSetter(t)
}

func (r *Template) parseBytes(in []byte) {
	lenIn := len(in)
	r.Places = []Place{}
	pos := 0
	for i := 0; i < lenIn; i++ {
		foundstart := bytes.Index(in[i:], delimiterStart)
		if -1 < foundstart {
			start := foundstart + i
			copy(in[pos:pos+start-i], in[i:start])
			pos = pos + start - i
			startPlaceH := start + delimiterStartLen
			foundend := bytes.Index(in[startPlaceH:], delimiterEnd)
			if -1 < foundend {
				end := foundend + start + delimiterStartLen
				r.Places = append(r.Places, Place{pos, Placeholder(string(in[startPlaceH:end]))})
				i = end + delimiterEndLen - 1
			} else {
				panic(fmt.Sprintf("starting delimiter at position %d has no ending delimiter", foundstart))
			}
		} else {
			copy(in[pos:lenIn-i+pos], in[i:])
			pos = pos + lenIn - i
			break
		}
	}
	r.Original = in
	r.Length = pos
	sort.Sort(r.Places)
}

func (r *Template) parseString(in string) {
	lenIn := len(in)
	r.Original = make([]byte, lenIn)
	r.Places = []Place{}
	pos := 0
	for i := 0; i < lenIn; i++ {
		foundstart := strings.Index(in[i:], delimiterStartString)
		if -1 < foundstart {
			start := foundstart + i
			copy(r.Original[pos:pos+start-i], []byte(in[i:start]))
			pos = pos + start - i
			startPlaceH := start + delimiterStartLen
			foundend := strings.Index(in[startPlaceH:], delimiterEndString)
			if -1 < foundend {
				end := foundend + start + delimiterStartLen
				r.Places = append(r.Places, Place{pos, Placeholder(string(in[startPlaceH:end]))})
				i = end + delimiterEndLen - 1
			} else {
				panic(fmt.Sprintf("starting delimiter at position %d has no ending delimiter", foundstart))
			}
		} else {
			copy(r.Original[pos:lenIn-i+pos], []byte(in[i:]))
			pos = pos + lenIn - i
			break
		}
	}
	r.Length = pos
	sort.Sort(r.Places)
}

func (t *Template) NewHandler(m map[string]http.Handler) http.Handler {
	return &templateHandler{m, t}
}

type templateHandler struct {
	replacements map[string]http.Handler
	Template     *Template
}

func (ts *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var last int
	for _, place := range ts.Template.Places {
		w.Write(ts.Template.Original[last:place.Pos])
		handler, exists := ts.replacements[place.Placeholder.Name()]
		if exists {
			handler.ServeHTTP(w, req)
		}
		last = place.Pos
	}
	w.Write(ts.Template.Original[last:ts.Template.Length])
	return
}
