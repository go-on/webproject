package prettify

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-on/wrap.v2"

	"gopkg.in/go-on/wrap-contrib.v2/wraps"
)

func TestPrettify(t *testing.T) {
	corpus := map[string]string{
		"<a> hi \t\nho</a>\n<b>hü ho</b> <div>yeah</div>":                     "<a> hi \t ho</a>\n<b>hü ho</b>\n<div>yeah</div>",
		"<div>\n\t\t \t \t<div class=\"inner\"><p>here we go</p></div></div>": "<div>\n\t<div class=\"inner\">\n\t\t<p>here we go</p>\n\t</div>\n</div>",
		"<div>does <b>not</b> change</div>":                                   "<div>does <b>not</b> change</div>",
	}

	for in, out := range corpus {
		res := Prettify(in)
		// println(res + "\n\n")
		if res != out {
			t.Errorf("expecting: %#v, got: %#v", out, res)
		}
	}

}

func TestWrap(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	wrap.New(
		Wrap,
		wraps.HTMLString("<div><div><p>text</p></div></div>"),
	).ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("wrong status code,should be 200, is: %d", rec.Code)
	}

	expected := "<div>\n\t<div>\n\t\t<p>text</p>\n\t</div>\n</div>"
	got := rec.Body.String()

	if expected != got {
		t.Errorf("expecting: %#v, got: %#v", expected, got)
	}
}
