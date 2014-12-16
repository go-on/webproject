package main

import (
	"fmt"
)

type A struct {
	B string
}

func printer(a ...A) {
	fmt.Printf("%#v\n", a)
}

func change(as []A) {
	for i := 0; i < 5; i++ {
		a := as[i]
		a.B = "hu"
		as[i] = a
	}
}

func main() {
	as := make([]A, 10)
	//as := [12]A{}
	change(as)
	fmt.Println(as[0:5])

	printer(as[0:5]...)
}
