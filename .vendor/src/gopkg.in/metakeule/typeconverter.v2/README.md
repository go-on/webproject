typeconverter
=============

[![Build Status](https://secure.travis-ci.org/metakeule/typeconverter.png)](http://travis-ci.org/metakeule/typeconverter)

conversion between the following go types (it is possible to add your own):

 - int
 - bool
 - float32
 - float64
 - *time.Time
 - string/Stringer
 - json
 - xml
 - []interface{}
 - map[string]interface{}

Example
-------

For the main types you may simply use Convert()

```go
package main

import (
	"fmt"
	conv "github.com/metakeule/typeconverter"
	"time"
)

func main() {
	var s string

	// convert time to string
	t1, _ := time.Parse(time.RFC3339, "2011-01-26T18:53:18+01:00")
	conv.Convert(t1, &s)
	fmt.Println(s) // 2011-01-26T18:53:18+01:00

	// convert string back to time
	var t2 time.Time
	conv.Convert(s, &t2)
	fmt.Println(t2.Format(time.RFC3339)) // 2011-01-26T18:53:18+01:00
}
```

But you may build your own converters upon it, add own types and conversion functions and overwrite the default conversions.

Documentation
-------------

Look into the examples directory and see the Documentation at http://godoc.org/github.com/metakeule/typeconverter

typeconverter is based on https://github.com/metakeule/dispatch, so you might want to look there too.