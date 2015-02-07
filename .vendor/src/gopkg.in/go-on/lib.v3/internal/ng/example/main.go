package main

import (
	"fmt"
	// "gopkg.in/metakeule/goh4.v5"
	. "gopkg.in/go-on/lib.v3/html"
	. "gopkg.in/go-on/lib.v3/types"
	// . "gopkg.in/go-on/lib.v3/html/tag"
	// . "gopkg.in/metakeule/goh4.v5/tag/short"
	"gopkg.in/go-on/lib.v3/internal/ng"
	// "strings"
)

func main() {
	fmt.Printf("%T: %#v\n", ng.Show("currentSection"), ng.Show("currentSection").String())
	fmt.Println(
		DIV(Class("col-xs-4"),
			ng.Show("currentSection"),
			DIV(
				Class("row"),
				"huhu",
			),
		).String(),
	)
}
