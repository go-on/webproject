package main

import (
	"fmt"

	"github.com/go-on/replacer"
)

func main() {
	template := replacer.Placeholder("name").String() + " - " + replacer.Placeholder("animal").String()

	res := replacer.NewTemplateString(template).ReplaceStrings(
		replacer.MapStrings(
			"animal", "Duck",
			"name", "Donald",
		),
	)

	// the returned replacement is a []byte
	fmt.Printf("%s\n", res)
}
