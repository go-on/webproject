package compiler

import (
	"bytes"
	"fmt"
	"gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/internal/replacer"
	"net/http"
)

func mkPhHandler(e *element.Element, key string, buf *bytes.Buffer) (phdl []*placeholderHandler) {
	phdl = []*placeholderHandler{}
	element.Pre(e, buf)

	// fmt.Printf("checking %s\n", e.Tag())

	if len(e.Children) == 0 || element.Is(e, element.SelfClosing) {
		// fmt.Printf("no children\n")
		element.Post(e, buf)
		return
	}

	// fmt.Printf("no children %d\n", len(e.Children))
	for i, in := range e.Children {
		// fmt.Printf("children %T\n", in)
		switch ch := in.(type) {
		case *element.Element:
			phdl = append(
				phdl,
				mkPhHandler(ch, fmt.Sprintf("%s/%d", key, i), buf)...,
			)
		case http.Handler:
			// fmt.Printf("is http.Handler: %T\n", ch)
			pha := replacer.Placeholder(fmt.Sprintf("%s-%d", key, i))
			phdl = append(phdl, &placeholderHandler{pha, ch})
			buf.WriteString(pha.String())
		default:
			buf.WriteString(in.String())
		}
	}

	element.Post(e, buf)
	return
}

type placeholderHandler struct {
	replacer.Placeholder
	http.Handler
}
