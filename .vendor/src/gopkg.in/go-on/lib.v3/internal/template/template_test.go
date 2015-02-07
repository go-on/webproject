package template

import (
	"bytes"
	// "fmt"
	// "net/http"
	// "net/http/httptest"
	"testing"
)

var (
	firstname = NewPlaceholder("firstname")
	lastname  = NewPlaceholder("lastname")
	templ     = New("t").MustAdd("Hello, ", firstname, " ", lastname, "!\n").Parse()
	expected  = "Hello, Donald Duck!\nHello, Mickey Mouse!\n"
)

func TestTemplate(t *testing.T) {
	var b bytes.Buffer
	templ.ReplaceTo(&b, firstname.Set("Donald"), lastname.Set("Duck"))
	templ.ReplaceTo(&b, firstname.Set("Mickey"), lastname.Set("Mouse"))

	if r := b.String(); r != expected {
		t.Errorf("Error in setting: expected\n\t%#v\ngot\n\t%#v\n", expected, r)
	}
}

/*
func TestTemplateServer(t *testing.T) {
	firstnameH := func(wr http.ResponseWriter, req *http.Request) {
		// fmt.Println("Donald")
		fmt.Fprint(wr, "Donald")
	}

	lastnameH := func(wr http.ResponseWriter, req *http.Request) {
		// fmt.Println("Duck")
		fmt.Fprint(wr, "Duck")
	}

	_ = lastnameH
	server := templ.Handle(
		NewPlaceholderHandler(lastname, http.HandlerFunc(lastnameH)),
		NewPlaceholderHandler(firstname, http.HandlerFunc(firstnameH)),
	)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	server.ServeHTTP(rec, req)

	body := rec.Body.String()
	exp := "Hello, Donald Duck!\n"
	if body != exp {
		t.Errorf("Error in setting: expected\n\t%#v\ngot\n\t%#v\n", exp, body)
	}
}
*/
