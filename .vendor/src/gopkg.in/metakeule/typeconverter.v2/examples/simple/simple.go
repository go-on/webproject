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
