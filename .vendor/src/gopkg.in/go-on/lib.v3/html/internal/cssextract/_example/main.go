package main

import (
	"fmt"

	"gopkg.in/go-on/lib.v3/html/cssextract"
)

var styles = `
  table.a { 
			display: none; 
			font-weight:bold; 
			background-image: url('image.jpg');
		}`

func main() {
	p := cssextract.Parse(styles)

	fmt.Printf("%#v\n", p)
}
