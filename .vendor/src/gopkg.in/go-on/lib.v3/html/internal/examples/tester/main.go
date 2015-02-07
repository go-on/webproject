package main

import (
	"fmt"

	. "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	// "gopkg.in/go-on/lib.v3/html/tag"
)

func main() {
	fmt.Println(element.Elements(LI("a"), LI("b")).String())
}
