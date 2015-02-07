package element

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/go-on/builtin.v1"
	"gopkg.in/go-on/lib.v3/types"
	"gopkg.in/go-on/lib.v3/types/placeholder"
)

type htmlerstring struct {
	types.HTMLer
}

func (h *htmlerstring) String() string {
	return h.HTMLer.HTML()
}

// prepare the id attribute for output
func AttrsString(ø *Element) (res string) {
	var buffer bytes.Buffer
	if !Is(ø, IdForbidden) && ø.Id != "" {
		buffer.WriteString(" " + types.Attribute{"id", string(ø.Id)}.String())
	}
	if !Is(ø, ClassForbidden) && len(ø.Classes) > 0 {
		buffer.WriteString(" " + types.Attribute{"class", classAttrString(ø.Classes)}.String())
	}
	if !Is(ø, Invisible) && len(ø.Styles) > 0 {
		buffer.WriteString(" " + types.Attribute{"style", css(ø.Styles)}.String())
	}

	for _, v := range ø.Attributes {
		// ignore id and class and style attributes, since they should be set via Id and Classes properties or Add
		if v.Key == "id" || v.Key == "class" || v.Key == "style" {
			continue
		}
		buffer.WriteString(" " + v.String())
	}
	return buffer.String()
}

func classAttrString(classes []types.Class) (s string) {
	var buffer bytes.Buffer
	for _, cl := range classes {
		buffer.WriteString(" " + string(cl))
	}
	return buffer.String()[1:]
}

func css(fds []types.Style) (s string) {
	var buffer bytes.Buffer
	for _, v := range fds {
		buffer.WriteString(v.String())
	}
	return buffer.String()
}

// adds css properties to the style attribute, same keys are overwritten
func addStyles(ø *Element, v []types.Style) {
	ø.Styles = append(ø.Styles, v...)
}

func addStyle(ø *Element, v types.Style) {
	ø.Styles = append(ø.Styles, v)
}

func addText(ø *Element, text string) {
	s := types.Text(text)

	if !Is(ø, WithoutEscaping) {
		s = types.Text(types.EscapeHTML(text))
	}

	if Is(ø, JavascriptSpecialEscaping) {
		s = types.Text(jsSpecialEscape(text))
	}
	ø.Children = append(ø.Children, s)
}

func addChild(ø *Element, child builtin.Stringer) {
	ø.Children = append(ø.Children, child)
}

func addClass(ø *Element, class types.Class) {
	ø.Classes = append(ø.Classes, class)
}

func addClasses(ø *Element, classes []types.Class) {
	ø.Classes = append(ø.Classes, classes...)
}

func addPlaceholder(ø *Element, v placeholder.Placeholder) {
	switch tp := v.Type().(type) {
	case types.Descr:
		ø.Add(types.Descr(v.String()))
	case types.Id:
		ø.Add(types.Id(v.String()))
	case types.Class:
		ø.Add(types.Class(v.String()))
	case types.HTMLString:
		ø.Add(types.HTMLString(v.String()))
	case types.Text:
		ø.Add(types.Text(v.String()))
	case types.Attribute:
		ø.Add(types.Attribute{tp.Key, v.String()})
	case types.Tag:
		ø.tag = v.String()
	case types.Style:
		ø.Add(types.Style{tp.Property, v.String()})
	default:
		ø.Add(types.Text(v.String()))
	}
}

func InnerHtmlBf(ø *Element, w io.Writer) (num int64, err error) {
	var n int64
	var nn int
	for _, in := range ø.Children {
		switch ch := in.(type) {
		case *Element:
			n, err = ch.WriteTo(w)
		case io.WriterTo:
			n, err = ch.WriteTo(w)
		default:
			nn, err = fmt.Fprint(w, in.String())
			n = int64(nn)
		}
		if err != nil {
			return
		}
		num += n
	}
	return
}

func InnerHtml(ø *Element) (res string) {
	var buffer bytes.Buffer
	InnerHtmlBf(ø, &buffer)
	return buffer.String()
}

/*
func InnerHtml(ø *Element) (res string) {
	var buffer bytes.Buffer
	for _, in := range ø.Children {
		buffer.WriteString(in.String())
	}
	return buffer.String()
}
*/

// Is checks if a given flag is set, e.g.
//
// 	Is(Inline)
//
// checks for the Inline flag
func Is(ø *Element, f flag) bool { return ø.flags&f != 0 }

// Pre writes properties of an element to w before the inner elements are written to w
func Pre(ø *Element, w io.Writer) (num int, err error) {
	var n int
	if ø.Descr != "" {
		n, err = fmt.Fprintf(w, "<!-- Begin: %s -->", ø.Descr)
		if err != nil {
			return
		}
		num += n
	}
	if Is(ø, WithoutDecoration) {
		return
	}
	if Is(ø, SelfClosing) {
		n, err = fmt.Fprintf(w, "<%s%s />", ø.Tag(), AttrsString(ø))
		if err != nil {
			return
		}
		num += n
		return
	}
	n, err = fmt.Fprintf(w, "<%s%s>", ø.Tag(), AttrsString(ø))
	num += n
	return
}

// Post writes properties of an element to w after the inner elements are written to w
func Post(ø *Element, w io.Writer) (num int, err error) {
	var n int
	if Is(ø, WithoutDecoration) || Is(ø, SelfClosing) {
		if ø.Descr != "" {
			n, err = fmt.Fprintf(w, "<!-- End: %s -->", ø.Descr)
			if err != nil {
				return
			}
			num += n
		}
		return
	}
	n, err = fmt.Fprintf(w, "</%s>", ø.Tag())
	if err != nil {
		return
	}
	num += n
	if ø.Descr != "" {
		n, err = fmt.Fprintf(w, "<!-- End: %s -->", ø.Descr)
		if err != nil {
			return
		}
		num += n
	}
	return
}
