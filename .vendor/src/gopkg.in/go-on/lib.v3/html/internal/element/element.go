package element

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/go-on/wrap.v2"

	"gopkg.in/go-on/builtin.v1"
	"gopkg.in/go-on/lib.v3/internal/template"
	"gopkg.in/go-on/lib.v3/types"
	"gopkg.in/go-on/lib.v3/types/placeholder"

	// "github.com/go-on/replacer"

	"net/http"
	"strings"
)

type flag int

const (
	_                              = iota
	hasDefaults               flag = 1 << iota // element has default flags
	IdForbidden                                // element should not have an id attribute
	ClassForbidden                             // element should not have a class attribute
	SelfClosing                                // element is selfclosing and contains no content
	Inline                                     // element is an inline element (only for visible elements)
	FormField                                  // element is a field of a form
	Invisible                                  // element doesn't render anything visible
	WithoutEscaping                            // element does not escape inner Text
	WithoutDecoration                          // element just prints the InnerHtml
	JavascriptSpecialEscaping                  // content contains javascript and needs special escaping or </
)

var flagNames = map[flag]string{
	hasDefaults:               "hasDefaults",
	IdForbidden:               "IdForbidden",
	ClassForbidden:            "ClassForbidden",
	SelfClosing:               "SelfClosing",
	Inline:                    "Inline",
	FormField:                 "FormField",
	Invisible:                 "Invisible",
	WithoutEscaping:           "WithoutEscaping",
	WithoutDecoration:         "WithoutDecoration",
	JavascriptSpecialEscaping: "JavascriptSpecialEscaping",
}

func (ø flag) String() string {
	return flagNames[ø]
}

type handlerFuncElement func(http.ResponseWriter, *http.Request)

func (h handlerFuncElement) ServeHTTP(w http.ResponseWriter, r *http.Request) { h(w, r) }
func (h handlerFuncElement) String() (s string)                               { return }

// var _ ElementLike = NewElement("")

// empty list of elements and others
func Elements(objects ...interface{}) (t *Element) {
	// important: tag as empty string is handled in a special way (in IsParentAllowed for example)
	t = NewElement("", WithoutDecoration)
	t.Add(objects...)
	return
}

// an Elementer might be parent of an Element
// by implementing a type that fulfills this interface
// you might peek into the execution.
// when String() method is called, the html of the
// tree is built and when SetParent() it is embedded in another Elementer
// it could be combined with the Pather interface that allows you to modify specific
// css selectors for any children Elements

// ElementLike behaves like an HTML Element.
type ElementLike interface {
	// Tag returns the HTML tag
	Tag() string

	// HTML returns the HTML string
	HTML() string

	// String returns the same as HTML
	String() string

	// Add adds different objects to the element
	Add(objects ...interface{})
}

func jsSpecialEscape(in string) string {
	return strings.Replace(in, `</`, `<\/`, -1)
}

// the base of what becomes a tag when printed
type Element struct {
	tag        string
	flags      flag
	Id         types.Id
	Descr      types.Descr
	Parent     ElementLike
	Attributes []types.Attribute
	Classes    []types.Class
	Children   []builtin.Stringer
	Styles     []types.Style
	//PlaceholderHandler []template.PlaceholderHandler
}

// contruct a new element with some flags.
//
// the tag constructors A(), Body(),... use these method, see tags.go file for examples
//
// use it for your own tags
//
// the following flags are supported
//
// 	IdForbidden                        // element should not have an id attribute
// 	ClassForbidden                     // element should not have a class attribute
// 	SelfClosing                        // element is selfclosing and contains no content
// 	Inline                             // element is an inline element (only for visible elements)
// 	Field                              // element is a field of a form
// 	Invisible                          // element doesn't render anything visible
// 	WithoutEscaping                    // element does not escape inner Text
// 	WithoutDecoration                  // element just prints the InnerHtml
//
// see Add() and Set() methods for how to modify the Element
func NewElement(tag string, flags ...flag) (ø *Element) {
	ø = &Element{
		tag:   tag,
		flags: hasDefaults,
	}

	for _, flag := range flags {
		ø.flags = ø.flags | flag
	}
	return
}

func (ø *Element) HTML() string { return ø.String() }
func (ø *Element) Tag() string  { return ø.tag }

// adds new inner content or properties based on Stringer objects and returns an error if changes could not be applied
//
// the following types are handled in a special way:
//
//  - Descr: sets a description of the eleent, that with be rendered to a comment
//  - Style: set a single style
//  - Styles: sets multiple styles
//  - Attr: set a single attribute   // do not set id or class via Attr(s) directly, use Id() and Class() instead
//  - Attrs: sets multiple attribute
//  - Class: adds a class
//  - Id: sets the id
//  - *Css: applies the css, see ApplyCss()
//
// the following types are added to  the inner content:
//
// 	- Text: ís escaped if the WithoutEscaping flag isn't set
// 	- Html: is never escaped
//
// If the Stringer can be casted to an Elementer (as Element can), it is added to the inner content as well
// otherwise it is handled like Text(), that means any type implementing Stringer can be added as (escaped) text
//func (ø *Element) Add(objects ...interface{}) ElementLike {
func (ø *Element) Add(objects ...interface{}) {
	for _, o := range objects {
		if o == nil {
			continue
		}
		switch v := o.(type) {
		/*
			case template.PlaceholderHandler:
				ø.Add(replacer.Placeholder(v.Name()).String())
				// ø.placeholderHandler = append(ø.placeholderHandler, v)
				continue
			case template.Placeholder:
				ø.Add(replacer.Placeholder(v.Name()).String())
				continue
			case *template.Template:
				ø.Add(replacer.Placeholder(v.Name()))
				continue
		*/
		case placeholder.Placeholder:
			addPlaceholder(ø, v)
			continue
		case string:
			addText(ø, v)
		case types.Text:
			addText(ø, string(v))
		case types.HTMLString:
			addChild(ø, v)
		case types.Comment:
			addChild(ø, v)
		case types.Descr:
			ø.Descr = v
		case types.Attribute:
			ø.Attributes = append(ø.Attributes, v)
		case []types.Attribute:
			ø.Attributes = append(ø.Attributes, v...)
		case []types.Class:
			addClasses(ø, v)
		case types.Class:
			addClass(ø, v)
		case types.Id:
			ø.Id = v
		case types.Style:
			addStyle(ø, v)
		case []types.Style:
			addStyles(ø, v)
		case *Element:
			addChild(ø, v)
		case ElementLike:
			addChild(ø, v)
		case http.HandlerFunc:
			addChild(ø, handlerFuncElement(v))
		case func(http.ResponseWriter, *http.Request):
			addChild(ø, handlerFuncElement(v))
		case http.Handler:
			addChild(ø, handlerFuncElement(v.ServeHTTP))
		case types.HTMLer:
			addChild(ø, &htmlerstring{v})
		case builtin.Stringer:
			addText(ø, v.String())
		default:
			addText(ø, fmt.Sprintf("%v", v))
		}
	}
}

// WriteTo does implement io.WriterTo interface
func (ø *Element) WriteTo(w io.Writer) (num int64, err error) {
	var n int64
	var nn int
	nn, err = Pre(ø, w)
	if err != nil {
		return
	}
	num += int64(nn)

	if !Is(ø, SelfClosing) && len(ø.Children) > 0 {
		n, err = InnerHtmlBf(ø, w)
		if err != nil {
			return
		}
		num += n
	}
	nn, err = Post(ø, w)
	if err != nil {
		return
	}
	num += int64(nn)
	return
}

// returns the html with inner content (and the own tags if WithoutDecoration is not set)
func (ø *Element) String() (res string) {
	var buf bytes.Buffer
	ø.WriteTo(&buf)
	return buf.String()
}

func (ø *Element) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !Is(ø, SelfClosing) && len(ø.Children) > 0 {
		inner := serveInner(ø, w, r)
		switch inner.Code {
		case 302, 301:
			inner.FlushHeaders()
			inner.FlushCode()
			// stop
			return
		}
		if inner.HasChanged() {
			if len(inner.Header()) > 0 {
				inner.FlushHeaders()
			}
			if inner.Code != 0 {
				inner.FlushCode()
			}
			Pre(ø, w)
			inner.Buffer.WriteTo(w)
			Post(ø, w)
			return
		}
	}
	Pre(ø, w)
	Post(ø, w)
}

func serveInner(ø *Element, w http.ResponseWriter, r *http.Request) (outer *wrap.Buffer) {
	outer = wrap.NewBuffer(w)
	for _, in := range ø.Children {

		switch ch := in.(type) {
		//case *Element:
		case http.Handler:
			buf := wrap.NewBuffer(outer)
			ch.ServeHTTP(buf, r)

			switch buf.Code {
			case 302, 301:
				buf.FlushHeaders()
				buf.FlushCode()
				return
			case 404:
				buf.FlushAll()
			case 500:
				fmt.Fprint(outer, "Server Error")
			default:
				if buf.IsOk() {
					buf.FlushAll()
				}
			}
		case io.WriterTo:
			ch.WriteTo(outer)
		default:
			fmt.Fprint(outer, in.String())
		}

	}
	return
}

func (e *Element) Template(name string) *template.Template {
	return template.New(name).
		MustAdd(e.HTML()).
		Parse()
}

func (ø *Element) Selector() string {
	var s = ø.tag
	// sele := []string{ø.tag.Selector()}
	if string(ø.Id) != "" {
		s += ø.Id.Selector()
	}
	for _, c := range ø.Classes {
		s += c.Selector()
	}
	return s
}
