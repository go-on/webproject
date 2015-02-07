package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/go-on/lib.v3/html/internal/cssextract"
	"gopkg.in/go-on/lib.v3/html/internal/htmlconv"
)

func transform(classOrId string) string {
	//return classOrId
	var buf bytes.Buffer

	nextUpper := true

	for _, r := range classOrId {
		if r == '-' || r == '_' {
			buf.WriteString("_")
			// nextUpper = true
			continue
		}

		if nextUpper {
			buf.WriteString(strings.ToUpper(string(r)))
			nextUpper = false
			continue
		}

		buf.WriteString(string(r))
	}

	return buf.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing parameter: cmd, allow: convert")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("missing parameter: file (file to convert/extract)")
		os.Exit(1)
	}

	cmd := os.Args[1]
	file := os.Args[2]

	switch cmd {
	case "extract-classes":
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		parser := cssextract.Parse(string(b))

		var buf bytes.Buffer
		buf.WriteString(`
			package class

			import (
         "gopkg.in/go-on/lib.v3/types"
			)

		  var (

			`)

		for _, class := range parser.Classes {
			fmt.Fprintf(&buf, "%s = types.Class(%#v)\n", transform(class), class)
		}

		buf.WriteString(`
      )
		`)

		fmt.Println(buf.String())

	case "extract-ids":
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		parser := cssextract.Parse(string(b))

		var buf bytes.Buffer
		buf.WriteString(`
			package id

			import (
         "gopkg.in/go-on/lib.v3/types"
			)

		  var (

			`)

		for _, id := range parser.Ids {
			fmt.Fprintf(&buf, "%s = types.Id(%#v)\n", transform(id), id)
		}

		buf.WriteString(`
      )
		`)

		fmt.Println(buf.String())

	case "convert":
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		parser := htmlconv.Parser{StripPrefixes: true, TrimSpace: true}

		fmt.Printf(`
	package main

	import (
		 "fmt"
		. "gopkg.in/go-on/lib.v3/types"
		. "gopkg.in/go-on/lib.v3/html"
		. "gopkg.in/go-on/lib.v3/html/internal/element"	   
	)

	var (
    _ = E_nbsp
    _ = A
    _ = Element{}
	)

	var elements = %s

	func main() {
		fmt.Println(elements)
	}
		`,
			parser.Parse(string(b)),
		)
	}

}
