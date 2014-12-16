package main

import (
	"os"

	. "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/html/internal/element"
	. "gopkg.in/go-on/lib.v2/types"
	"gopkg.in/go-on/lib.v2/types/placeholder"
)

var (
	name    = placeholder.New(Text("name"))
	content = DIV(name, element.NewElement("p"), "hello world").Template("content")
	layout  = SECTION(content).Template("layout")
)

func main() {

	all := content.New()
	content.ReplaceTo(all.Buffer, name.Set("<heino>"))
	content.ReplaceTo(all.Buffer, name.Set("<hannelore>"))

	layout.Replace(all).WriteTo(os.Stdout)
}
