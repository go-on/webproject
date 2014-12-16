package html

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoctypes(t *testing.T) {
	inner := HTML(BODY("hi"))

	tests := map[string]string{
		HTML5(inner).Tag():                         `-doctype-`,
		XHTML1_1Basic().HTML():                     "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML Basic 1.1//EN\"\n    \"http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd\">\n",
		HTML5(inner).Template().Replace().String(): "<!DOCTYPE HTML>\n<html><body>hi</body></html>",
		XHTML1_0Transitional(inner).String():       "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\"\n   \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n<html><body>hi</body></html>",
		MathML2_0().String():                       "<!DOCTYPE math PUBLIC \"-//W3C//DTD MathML 2.0//EN\"\n  \"http://www.w3.org/Math/DTD/mathml2/mathml2.dtd\">\n",
		HTML4_01Frameset().String():                "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Frameset//EN\"\n   \"http://www.w3.org/TR/html4/frameset.dtd\">\n",
		HTML4_01Strict().String():                  "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\"\n   \"http://www.w3.org/TR/html4/strict.dtd\">\n",
		HTML4_01Transitional().String():            "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\"\n   \"http://www.w3.org/TR/html4/loose.dtd\">\n",
		XHTML1_0Frameset().String():                "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Frameset//EN\"\n   \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd\">\n",
		XHTML1_0Strict().String():                  "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\"\n   \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\">\n",
		XHTML1_1().String():                        "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\"\n   \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">\n",
		MathML1_01().String():                      "<!DOCTYPE math SYSTEM\n  \"http://www.w3.org/Math/DTD/mathml1/mathml.dtd\">\n",
	}

	for got, expected := range tests {
		if got != expected {
			t.Errorf("got: %#v expected: %#v", got, expected)
		}
	}

}

func TestDoctypeServe(t *testing.T) {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		code := req.URL.Query().Get("code")
		switch code {
		case "302":
			rw.Header().Set("Location", "/302")
			rw.WriteHeader(302)
		case "301":
			rw.Header().Set("Location", "/301")
			rw.WriteHeader(301)
		default:
			rw.Header().Set("Location", "/default")
			rw.Write([]byte("ok"))
		}
	}

	doc := HTML5(HTML(BODY("hi", handler)))

	tests := map[string]int{
		"/?code=302": 302,
		"/?code=301": 301,
		"/":          200,
	}

	for url, expected := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)

		doc.ServeHTTP(rec, req)

		if rec.Code != expected {
			t.Errorf("got: %d expected: %d", rec.Code, expected)
		}
	}

}
