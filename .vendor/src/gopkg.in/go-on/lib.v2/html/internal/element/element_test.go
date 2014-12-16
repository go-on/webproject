package element

import (
	"bytes"
	"fmt"
	types "gopkg.in/go-on/lib.v2/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStrings(t *testing.T) {
	ahref := NewElement("a")
	ahref.Add(
		types.Attribute{"href", "#"},
		"buy <now!>",
		types.Class("external"),
		types.Class("button"),
		types.Id("buy"),
	)

	script := NewElement("script", JavascriptSpecialEscaping)
	script.Add("hi</ho")

	withoutDeco := NewElement("a", WithoutDecoration)
	withoutDeco.Add(types.Descr("wrapper"), ahref)

	tests := map[string]string{
		NewElement("a").String():               "<a></a>",
		NewElement("br", SelfClosing).String(): "<br />",
		ahref.String():                         `<a id="buy" class="external button" href="#">buy &lt;now!&gt;</a>`,
		ahref.HTML():                           ahref.String(),
		ahref.Tag():                            "a",
		withoutDeco.String():                   "<!-- Begin: wrapper -->" + ahref.String() + "<!-- End: wrapper -->",
		WithoutDecoration.String():             "WithoutDecoration",
		script.String():                        `<script>hi<\/ho</script>`,
	}

	for result, should := range tests {
		if result != should {
			t.Errorf("expected: %#v, got %#v", should, result)
		}
	}
}

func TestAdding(t *testing.T) {
	var bf bytes.Buffer
	bf.WriteString("written in <buffer>")

	tests := map[string]interface{}{
		`<a class="external"></a>`:                                    types.Class("external"),
		`<a id="buy"></a>`:                                            types.Id("buy"),
		`<a>buy &lt;now!&gt;</a>`:                                     "buy <now!>",
		`<a><p>Hi</p></a>`:                                            types.HTMLString("<p>Hi</p>"),
		`<a>&lt;p&gt;Ho&lt;/p&gt;</a>`:                                types.Text("<p>Ho</p>"),
		`<a><!-- no comment --></a>`:                                  types.Comment("no comment"),
		`<!-- Begin: just a test --><a></a><!-- End: just a test -->`: types.Descr("just a test"),
		`<a style="color:red;"></a>`:                                  types.Style{"color", "red"},
		`<a><span></span></a>`:                                        NewElement("span"),
		`<a>4</a>`:                                                    4,
		`<a>written in &lt;buffer&gt;</a>`:                            &bf,
		`<a href="#" target="_blank"></a>`:                            []types.Attribute{{"href", "#"}, {"target", "_blank"}},
		`<a class="external button"></a>`:                             []types.Class{"external", "button"},
		`<a style="color:red;display:block;"></a>`:                    []types.Style{{"color", "red"}, {"display", "block"}},
		`<a><span></span>text<strong></strong></a>`:                   Elements(NewElement("span"), "text", NewElement("strong")),
		`<a></a>`: []types.Attribute{{"id", "nix"}, {"class", "ignore"}, {"style", "nothing"}},
	}

	for should, adder := range tests {
		element := NewElement("a")
		element.Add(adder)
		result := element.String()
		if result != should {
			t.Errorf("expected: %#v, got %#v", should, result)
		}
	}
}

func index(rw http.ResponseWriter, req *http.Request) { rw.Write([]byte("index")) }

func code302(rw http.ResponseWriter, req *http.Request) { rw.Write([]byte("302")); rw.WriteHeader(302) }

type code404 struct{}

func (code404) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("404"))
	rw.WriteHeader(404)
}

func code500(rw http.ResponseWriter, req *http.Request) { rw.Write([]byte("500")); rw.WriteHeader(500) }

type article string

func (a article) HTML() string {
	return fmt.Sprintf(`<article>%s</article>`, types.EscapeHTML(string(a)))
}

func (a article) Tag() string                { return "article" }
func (a article) String() string             { return a.HTML() }
func (a article) Add(objects ...interface{}) {}

func serve(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("served"))
}

func TestHandler(t *testing.T) {
	div := NewElement("div")
	div.Add(
		index,
		types.Descr("here the index"),
	)

	doc := NewElement("doc", WithoutDecoration)
	doc.Add(
		http.HandlerFunc(index),
		types.Descr("here the index in doc"),
	)

	doc302 := NewElement("doc", WithoutDecoration)
	doc302.Add(code302)

	doc404 := NewElement("doc", WithoutDecoration)
	doc404.Add(code404{})

	doc500 := NewElement("doc", WithoutDecoration)
	doc500.Add(code500)

	serv1 := NewElement("body")
	serv1.Add(serve)

	tests := map[string]http.Handler{
		`<a></a>`:  NewElement("a"),
		`<br />hi`: Elements(NewElement("br", SelfClosing), "hi"),
		`<!-- Begin: here the index --><div>index</div><!-- End: here the index -->`:    div,
		`<!-- Begin: here the index in doc -->index<!-- End: here the index in doc -->`: doc,
		``:                                doc302,
		`404`:                             doc404,
		`Server Error`:                    doc500,
		`<article>w&lt;how&gt;</article>`: Elements(article("w<how>")),
		`<body>served</body>`:             serv1,
	}
	for should, handler := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		handler.ServeHTTP(rec, req)
		result := rec.Body.String()

		if result != should {
			t.Errorf("expected: %#v, got %#v", should, result)
		}
	}
}
