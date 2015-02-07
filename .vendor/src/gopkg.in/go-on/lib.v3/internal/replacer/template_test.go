package replacer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var template = []byte{}
var Expected = ""
var Map = map[Placeholder][]byte{}

func Prepare() {
	Map = map[Placeholder][]byte{}
	orig := []string{}
	exp := []string{}
	for i := 0; i < 5; i++ {
		orig = append(orig, fmt.Sprintf(`a string with `+Placeholder("replacement%v").String(), i))
		exp = append(exp, fmt.Sprintf("a string with repl%v", i))
		Map[Placeholder(fmt.Sprintf("replacement%v", i))] = []byte(fmt.Sprintf("repl%v", i))
	}
	Expected = strings.Join(exp, "")
	template = []byte(strings.Join(orig, ""))
}

func TestReplaceMulti(t *testing.T) {
	Prepare()
	repl := NewTemplateBytes(template)
	res := string(repl.ReplaceBytes(Map))
	if res != Expected {
		t.Errorf("unexpected result: %#v, expected: %#v", res, Expected)
	}
}

func TestReplaceSyntaxError(t *testing.T) {
	repl := NewTemplateString("before " + Placeholder("one").String() + Placeholder("two").String() + Placeholder("one").String() + " after")

	m := map[Placeholder][]byte{
		Placeholder("one"): []byte("1"),
		Placeholder("two"): []byte("2"),
	}

	expected := "before 121 after"
	//	var buffer bytes.Buffer
	res := string(repl.ReplaceBytes(m))
	if res != expected {
		t.Errorf("unexpected result: %#v, expected: %#v", res, expected)
	}
}

func TestReplaceSyntaxErrorString(t *testing.T) {
	repl := NewTemplateString(
		"before " +
			Placeholder("one").String() +
			Placeholder("two").String() +
			Placeholder("one").String() +
			" after",
	)

	m := map[Placeholder]string{
		Placeholder("one"): "1",
		Placeholder("two"): "2",
	}
	expected := "before 121 after"
	res := string(repl.ReplaceStrings(m))
	// var buffer bytes.Buffer
	if res != expected {
		t.Errorf("unexpected result: %#v, expected: %#v", res, expected)
	}
}

func TestReplaceNoPlaceholders(t *testing.T) {
	repl := NewTemplateBytes([]byte("before after"))

	m := map[Placeholder]string{
		Placeholder("one"): "1",
		Placeholder("two"): "2",
	}

	expected := "before after"
	// var buffer bytes.Buffer
	res := string(repl.ReplaceStrings(m))
	if res != expected {
		t.Errorf("unexpected result: %#v, expected: %#v", res, expected)
	}
}

func TestReplaceNoReplacements(t *testing.T) {
	repl := NewTemplateBytes([]byte("before " + Placeholder("one").String() + Placeholder("two").String() + Placeholder("one").String() + " after"))

	m := map[Placeholder][]byte{}

	expected := "before  after"
	// var buffer bytes.Buffer
	res := string(repl.ReplaceBytes(m))
	if res != expected {
		t.Errorf("unexpected result: %#v, expected: %#v", res, expected)
	}
}

func TestServer(t *testing.T) {
	one := func(wr http.ResponseWriter, req *http.Request) {
		fmt.Fprint(wr, req.URL.Path+"1")
	}

	two := func(wr http.ResponseWriter, req *http.Request) {
		fmt.Fprint(wr, req.URL.Path+"2")
	}

	server := NewTemplateString(
		"before " +
			Placeholder("one").String() +
			Placeholder("two").String() +
			" middle " +
			Placeholder("one").String() +
			" after",
	).NewHandler(MapHandlers("one", http.HandlerFunc(one), "two", http.HandlerFunc(two)))

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hi", nil)

	server.ServeHTTP(rec, req)

	res := rec.Body.String()
	expected := "before /hi1/hi2 middle /hi1 after"
	if res != expected {
		t.Errorf("wrong  body, expected: %#v, got: %#v", expected, res)
	}
}

func TestServerSetter(t *testing.T) {
	_one_ := Placeholder("one")
	_two_ := Placeholder("two")
	_three_ := Placeholder("three")

	setter := NewTemplateString(
		"before " +
			_one_.String() + _two_.String() +
			" middle " +
			_one_.String() + _three_.String() +
			" after",
	).NewSetter()

	setter.SetString(_one_, "1")
	setter.SetString(_two_, "2")

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hi", nil)

	setter.ServeHTTP(rec, req)

	res := rec.Body.String()
	expected := "before 12 middle 1 after"

	if res != expected {
		t.Errorf("wrong  body, expected: %#v, got: %#v", expected, res)
	}
}
