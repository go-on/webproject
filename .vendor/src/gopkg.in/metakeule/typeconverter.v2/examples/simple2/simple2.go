package main

import (
	"fmt"
	conv "gopkg.in/metakeule/typeconverter.v2"
)

func main() {
	var s string

	// convert int to string
	conv.Convert("huho", &s)
	fmt.Println(s) // 2011-01-26T18:53:18+01:00
}
