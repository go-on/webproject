// Copyright 2013 Marc René Arns. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package typeconverter provides an flexible way to do conversion between types.

Out of the box the following go types are supported (it is possible to add your own and overwrite the conversions):

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

For the main types you may simply use Convert()

	package main

	import (
		"fmt"
		conv "gopkg.in/metakeule/typeconverter.v2"
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

But you may build your own converters upon it, add own types and conversion functions and overwrite the default conversions.

Look into the examples directory.
*/
package typeconverter
