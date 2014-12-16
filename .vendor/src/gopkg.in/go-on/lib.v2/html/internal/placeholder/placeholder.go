package placeholder

import (
	"fmt"
	. "gopkg.in/go-on/builtin.v1"
	"gopkg.in/go-on/lib.v2/internal/replacer"

	// . "gopkg.in/go-on/lib.v2/internal/shared"

	// ph2 "gopkg.in/go-on/lib.v2/internal/shared/placeholder"
	ph "gopkg.in/go-on/lib.v2/internal/template/placeholder"
	"html"
	// "net/http"
	"net/url"
	"os"
	"path/filepath"
	// "reflect"
	"runtime"
	// "strings"
	// "time"

	"gopkg.in/go-on/lib.v2/internal/template"
)

var Escaper = template.Escaper{
	"text":     handleStrings(html.EscapeString, true),
	"":         handleStrings(html.EscapeString, true),
	"html":     handleStrings(idem, true),
	"comment":  handleStrings(idem, true),
	"px":       units("%vpx"),
	"%":        units("%v%%"),
	"em":       units("%vem"),
	"pt":       units("%vpt"),
	"urlparam": handleStrings(url.QueryEscape, false),
}

/*
type view struct {
	*template.View
}

type placeholder struct {
	ph.Placeholder
}

func (p placeholder) String() string {
	return replacer.Placeholder(p.Name()).String()
}

func (v *view) Placeholder(field string) placeholder {
	return placeholder{v.View.Placeholder(field)}
}

func View(stru interface{}, tag string) *view {
	return &view{Escaper.View(stru, tag)}
}
*/

//Placeholder("Link")

func units(format string) func(interface{}) string {
	return func(in interface{}) (out string) {
		switch v := in.(type) {
		case int, int8, int16, int32, int64, float32, float64:
			return fmt.Sprintf(format, v)
		default:
			panic("unsupported type: " + fmt.Sprintf("%v (%T)", v, v))
		}
		return
	}
}

/*
// takes different types and outputs a string
func Str(in interface{}) string {
	switch v := in.(type) {
	// case *template.Placeholder:
	// return replacer.Placeholder(*v).Name()).String()
	case ph.Placeholder:
		return replacer.Placeholder(v.Name()).String()
	case Stringer:
		return v.String()
	case string:
		return v
	}
	panic("unsupported type: " + fmt.Sprintf("%v (%T)", in, in))
}
*/

func idem(in string) (out string) { return in }

// is  used by FillStruct, see github.com/metakeule/template
/*
var Transformer = map[string]func(interface{}) string{
	"text": handleStrings(html.EscapeString, true),
	"html": handleStrings(idem, false),
}
*/

/*
// shortcut for template.FillStruct with transformer
func FillStruct(ptrToStruct interface{}) map[string]string {
	return template.FillStruct("goh4", Transformer, ptrToStruct)
}
*/

func handleStrings(trafo func(string) string, allowAll bool) func(interface{}) string {
	return func(in interface{}) (out string) {
		if in == nil {
			return ""
		}
		var s string
		switch v := in.(type) {
		case Stringer:
			s = v.String()
		case string:
			s = v
		default:
			if allowAll {
				s = fmt.Sprintf("%v", v)
			} else {
				panic("unsupported type: " + fmt.Sprintf("%v (%T)", v, v))
			}
		}
		return trafo(s)
	}
}

type typedPlaceholder struct {
	ph.Placeholder
	typ interface{}
}

/*
func Handle(handler http.Handler) ph.PlaceholderHandler {
	return ph.NewPlaceholderHandler(ph2.New(HTMLString(fmt.Sprintf("handle %v", time.Now().UnixNano()))), handler)
}

func HandleFunc(fn func(http.ResponseWriter, *http.Request)) ph.PlaceholderHandler {
	return ph.NewPlaceholderHandler(ph2.New(HTMLString(fmt.Sprintf("handlefunc %v", time.Now().UnixNano()))), http.HandlerFunc(fn))
}
*/

/*
func (tph typedPlaceholder) Handle(h http.Handler) ph.PlaceholderHandler {
	return template.NewPlaceholderHandler(tph.Placeholder, h)
}

func (tph typedPlaceholder) HandleFunc(fn func(http.ResponseWriter, *http.Request)) ph.PlaceholderHandler {
	return template.NewPlaceholderHandler(tph.Placeholder, http.HandlerFunc(fn))
}
*/

func (ø typedPlaceholder) String() string {
	return replacer.Placeholder(ø.Name()).String()
}

func (ø typedPlaceholder) Type() interface{} {
	return ø.typ
}

func newTPh(ph ph.Placeholder, i interface{}) typedPlaceholder {
	return typedPlaceholder{ph, i}
}

func stripGoPath(path string) {
	gopath := filepath.SplitList(os.Getenv("GOPATH"))[0]
	// gopath := strings.Split(os.Getenv("GOPATH"), ":")[0]
	if gopath == "" {
		panic("GOPATH not set")
	}
}

func caller(skip int) string {
	_, file, num, ok := runtime.Caller(skip)
	if !ok {
		panic("can't get caller")
	}
	return fmt.Sprintf("%s:%v", file, num)
}
