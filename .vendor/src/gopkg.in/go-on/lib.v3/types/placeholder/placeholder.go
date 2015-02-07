package placeholder

import (
	"fmt"
	"reflect"

	"gopkg.in/go-on/lib.v3/internal/replacer"
	"gopkg.in/go-on/lib.v3/internal/template"
	"gopkg.in/go-on/lib.v3/types"

	ph "gopkg.in/go-on/lib.v3/internal/template/placeholder"
)

var (
	commentType   = reflect.TypeOf(types.Comment("")).Name()
	descrType     = reflect.TypeOf(types.Descr("")).Name()
	classType     = reflect.TypeOf(types.Class("")).Name()
	idType        = reflect.TypeOf(types.Id("")).Name()
	htmlType      = reflect.TypeOf(types.HTMLString("")).Name()
	textType      = reflect.TypeOf(types.Text("")).Name()
	attributeType = reflect.TypeOf(types.Attribute{}).Name()
	tagType       = reflect.TypeOf(types.Tag("")).Name()
	styleType     = reflect.TypeOf(types.Style{}).Name()
)

// TODO: extract the template.Setter interface in a extra subpackage and just include that subpackage
type Placeholder interface {
	ph.Setter
	Set(val interface{}) ph.Setter
	Setf(format string, val ...interface{}) ph.Setter
	String() string
	Type() interface{}
	// Handle(http.Handler) template.PlaceholderHandler
	// HandleFunc(func(http.ResponseWriter, *http.Request)) template.PlaceholderHandler
}

func New(thing interface{}) Placeholder {

	switch ø := thing.(type) {
	case types.Comment:
		return newTPh(template.NewPlaceholder(commentType+"."+string(ø)), ø)
	case types.Descr:
		return newTPh(template.NewPlaceholder(descrType+"."+string(ø)), ø)
	case types.Class:
		return newTPh(template.NewPlaceholder(classType+"."+string(ø)), ø)
	case types.Id:
		return newTPh(template.NewPlaceholder(idType+"."+string(ø)), ø)
	case types.HTMLString:
		return newTPh(template.NewPlaceholder(htmlType+"."+string(ø)), ø)
	case types.Text:
		t := template.NewPlaceholder(textType+"."+string(ø), handleStrings(types.EscapeHTML))
		return newTPh(t, ø)
	case types.Attribute:
		fn := func(in string) string {
			return (types.Attribute{ø.Key, in}).String()
		}
		t := template.NewPlaceholder(attributeType+"."+ø.Value, handleStrings(fn))
		return newTPh(t, ø)
	case types.Tag:
		return newTPh(template.NewPlaceholder(tagType+"."+string(ø)), ø)
	case types.Style:
		fn := func(in string) string {
			return (types.Style{ø.Property, in}).String()
		}
		t := template.NewPlaceholder(styleType+"."+ø.Value, handleStrings(fn))
		return newTPh(t, ø)
	case string:
		t := template.NewPlaceholder(textType+"."+ø, handleStrings(types.EscapeHTML))
		return newTPh(t, types.Text(ø))
	default:
		str := fmt.Sprintf("%v", ø)
		t := template.NewPlaceholder(textType+"."+str, handleStrings(types.EscapeHTML))
		return newTPh(t, types.Text(str))
	}

}

func newTPh(ph ph.Placeholder, i interface{}) typedPlaceholder {
	return typedPlaceholder{ph, i}
}

type typedPlaceholder struct {
	ph.Placeholder
	typ interface{}
}

func (ø typedPlaceholder) String() string {
	return replacer.Placeholder(ø.Name()).String()
}

func (ø typedPlaceholder) Type() interface{} {
	return ø.typ
}

func handleStrings(trafo func(string) string) func(interface{}) string {
	return func(in interface{}) (out string) {
		if in == nil {
			return ""
		}
		var s string
		switch v := in.(type) {
		case fmt.Stringer:
			s = v.String()
		case string:
			s = v
		default:
			s = fmt.Sprintf("%v", v)
		}
		return trafo(s)
	}
}
