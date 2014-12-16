package compiler

import (
	"fmt"
	. "gopkg.in/go-on/lib.v2/html"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func stripWhiteSpace(in string) string {
	return strings.Replace(strings.Replace(strings.Replace(in, "\n", "", -1), "\t", "", -1), " ", "", -1)
}

func body(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("one"))
}

func TestDocCompiler(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	DocHandler(
		HTML5(
			HTML(
				BODY(
					body,
					"-two-",
					SPAN(),
				),
			),
		),
	).ServeHTTP(rec, req)

	expected := `<!DOCTYPE HTML><html><body>one-two-<span></span></body></html>`

	expected = stripWhiteSpace(expected)

	got := stripWhiteSpace(rec.Body.String())

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}

func TestElementCompiler(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	ElementHandler("body",
		BODY(
			body,
			"-two-",
			SPAN(),
		),
	).ServeHTTP(rec, req)

	expected := `<body>one-two-<span></span></body>`

	expected = stripWhiteSpace(expected)

	got := stripWhiteSpace(rec.Body.String())

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}

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

func TestDocDispatcher(t *testing.T) {

	dispatcher := NewDispatcher(&layout{})
	titleHTML := dispatcher.HTML("Title")
	titleText := dispatcher.Text("Title")
	body := dispatcher.Text("Body")

	doc := HTML5(
		HTML(
			BODY(H1(titleText), H2(titleHTML), P(body)),
		),
	)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?title=h<i>h</i>o&body=<hu>ho", nil)

	dispatcher.DocHandler(doc).ServeHTTP(rec, req)

	expected := `
	<!DOCTYPE HTML>
	<html>
		<body>
			<h1>title is &#34;h&lt;i&gt;h&lt;/i&gt;o&#34;</h1>
			<h2>title is "h<i>h</i>o"</h2>
			<p>body is &#34;&lt;hu&gt;ho&#34;</p>
		</body>
	</html>
		`

	expected = stripWhiteSpace(expected)

	got := stripWhiteSpace(rec.Body.String())

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}

func TestElementDispatcher(t *testing.T) {

	dispatcher := NewDispatcher(&layout{})
	titleHTML := dispatcher.HTML("Title")
	titleText := dispatcher.Text("Title")
	body := dispatcher.Text("Body")

	elem := BODY(H1(titleText), H2(titleHTML), P(body))

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?title=h<i>h</i>o&body=<hu>ho", nil)

	dispatcher.ElementHandler("body", elem).ServeHTTP(rec, req)

	expected := `
		<body>
			<h1>title is &#34;h&lt;i&gt;h&lt;/i&gt;o&#34;</h1>
			<h2>title is "h<i>h</i>o"</h2>
			<p>body is &#34;&lt;hu&gt;ho&#34;</p>
		</body>
		`

	expected = stripWhiteSpace(expected)

	got := stripWhiteSpace(rec.Body.String())

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}
