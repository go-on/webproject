package main

import (
	"fmt"
	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/html/internal/match"
	. "gopkg.in/go-on/lib.v2/types"
	// . "gopkg.in/go-on/lib.v2/html/tag"
)

func main() {
	content := DIV(
		Id("content"), // sets the id attribute
		A(
			Class("button"), // sets the class attribute
			Attribute{"href", "#"},
			"something",                                     // sets the inner text
			Style{"color", "red"},                           // sets the style attribute
			STRONG(Id("sth"), HTMLString("not <escaped>!")), // sets some inner html.Element
		),
	)

	children := content.Children
	// el := children[0].(*Element).
	inner := InnerHtml(children[0].(*Element)) // everything inside
	buttons := match.All(content, match.New(Class("button")))
	sth := match.Any(content, match.New(Id("sth")))

	fmt.Printf(`
children[0].Classes(): %#v
inner: %#v
buttons[0].Tag(): %#v
sth: %#v
`, children[0].(*Element), inner, buttons[0].Tag(), sth.String())

	fmt.Println(content.String())
}
