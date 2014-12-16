package main

import (
	"bytes"
	"fmt"
	"html"
	"os"

	"github.com/go-on/template"
	"github.com/go-on/template/placeholder"
)

func Html(name string) (t placeholder.Placeholder) {
	return template.NewPlaceholder(name)
}

func Text(name string) (t placeholder.Placeholder) {
	return template.NewPlaceholder(
		name,
		func(in interface{}) (out string) { return html.EscapeString(in.(string)) },
	)
}

var (
	person   = Text("person")
	greeting = Html("greeting")
	T        = template.New("t").MustAdd("<h1>Hi, ", person, "</h1>", greeting).Parse()
)

func main() {
	fmt.Println(
		T.Replace(
			person.Set("S<o>meone"),
			greeting.Set("<div>Hi</div>"),
		),
	)

	hw := T.Replace(person.Set("W<o>rld"), greeting.Set("<div>Hello</div>"))
	// hw.WriteTo(os.Stdout)
	fmt.Printf("%s\n", hw.Bytes())
	fmt.Fprintln(os.Stdout, "---")
	// hw.WriteTo(os.Stdout)
	fmt.Printf("%s\n", hw.Bytes())

	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {

		T.ReplaceTo(&buffer,
			person.Setf("Bugs <Bunny> %v", i+1),
			greeting.Set("<p>How are you?</p>\n"))
	}
	fmt.Println(buffer.String())
}
