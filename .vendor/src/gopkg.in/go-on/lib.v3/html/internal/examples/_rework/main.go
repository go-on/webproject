package main

import (
	"fmt"
	. "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element/compiler"

	// "gopkg.in/go-on/wrap.v2"
	"net/http"
)

type layout struct {
	title string
	body  string
}

func (m *layout) Title(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "title is %#v", m.title)
}

func (m *layout) Body(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "body is %#v", m.body)
}

func (m *layout) Prepare(rw http.ResponseWriter, req *http.Request) {
	m.title = req.URL.Query().Get("title")
	m.body = req.URL.Query().Get("body")
}

func main() {
	dispatcher := compiler.NewDispatcher(&layout{})
	titleHTML := dispatcher.HTML("Title")
	titleText := dispatcher.Text("Title")
	body := dispatcher.Text("Body")

	doc := HTML5(
		BODY(H1(titleText), H2(titleHTML), P(body)),
	)

	err := http.ListenAndServe(":8182", dispatcher.DocHandler(doc))

	if err != nil {
		panic(err)
	}

}
