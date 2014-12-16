package main

import (
	"fmt"
	"gopkg.in/metakeule/nil.v1"
)

func main() {
	var x interface{}
	fmt.Println("is nil ", nil.Is(x))
	nil.MustNot(x) // panics
}
