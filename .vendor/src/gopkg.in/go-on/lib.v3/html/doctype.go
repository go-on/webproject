package html

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	. "gopkg.in/go-on/lib.v3/html/internal/element"
	"gopkg.in/go-on/lib.v3/internal/template"
	"gopkg.in/go-on/lib.v3/types"
	"gopkg.in/go-on/wrap.v2"
)

// pseudo element for placeholder
func doc(objects ...interface{}) (t *Element) {
	t = NewElement("doc", WithoutDecoration)
	t.Add(objects...)
	return
}

type DocType struct {
	*Element
	DocType string
}

func NewDocType(doctype string, objects ...interface{}) *DocType {
	e := doc(objects...)
	dc := &DocType{
		Element: e,
		DocType: doctype,
	}
	return dc
}

func docTypeXml(doctypeString string, objects ...interface{}) (d *DocType) {
	objects = append(objects, types.Attribute{"xmlns", "http://www.w3.org/1999/xhtml"})
	return NewDocType(doctypeString, objects...)
}

func (ø *DocType) String() string {
	var buf bytes.Buffer
	ø.WriteTo(&buf)
	return buf.String()
}

func (ø *DocType) WriteTo(w io.Writer) (num int64, err error) {
	var n int64
	var nn int
	nn, err = fmt.Fprint(w, ø.DocType+"\n")

	if err != nil {
		return
	}
	num += int64(nn)

	n, err = ø.Element.WriteTo(w)
	num += n
	return
}

func (ø *DocType) HTML() string {
	return ø.String()
}

func (ø *DocType) Tag() string {
	return "-doctype-"
}

var _ ElementLike = &DocType{}

func (ø *DocType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	inner := wrap.NewBuffer(w)
	ø.Element.ServeHTTP(inner, r)
	switch inner.Code {
	case 302, 301:
		inner.FlushHeaders()
		inner.FlushCode()
		// stop
		return
	}
	fmt.Fprint(w, ø.DocType+"\n")
	inner.FlushAll()
}

func (ø *DocType) Template() *template.Template {
	return template.New("document").
		MustAdd(ø.HTML()).
		Parse()
}

func HTML4_01Strict(objects ...interface{}) (t *DocType) {
	return NewDocType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
   "http://www.w3.org/TR/html4/strict.dtd">`, objects...)
}

func HTML4_01Transitional(objects ...interface{}) (t *DocType) {
	return NewDocType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/loose.dtd">`, objects...)
}

func HTML4_01Frameset(objects ...interface{}) (t *DocType) {
	return NewDocType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN"
   "http://www.w3.org/TR/html4/frameset.dtd">`, objects...)
}

func XHTML1_0Strict(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`, objects...)
}

func XHTML1_0Transitional(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`, objects...)
}

func XHTML1_0Frameset(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`, objects...)
}

func XHTML1_1(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
   "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`, objects...)
}

func XHTML1_1Basic(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN"
    "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`, objects...)
}

func HTML5(objects ...interface{}) (t *DocType) {
	return NewDocType(`<!DOCTYPE HTML>`, objects...)
}

func MathML2_0(objects ...interface{}) (t *DocType) {
	return NewDocType(`<!DOCTYPE math PUBLIC "-//W3C//DTD MathML 2.0//EN"
  "http://www.w3.org/Math/DTD/mathml2/mathml2.dtd">`, objects...)
}

func MathML1_01(objects ...interface{}) (t *DocType) {
	return NewDocType(`<!DOCTYPE math SYSTEM
  "http://www.w3.org/Math/DTD/mathml1/mathml.dtd">`, objects...)
}
