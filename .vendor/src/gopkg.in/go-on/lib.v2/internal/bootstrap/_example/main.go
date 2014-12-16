package main

import (
	"gopkg.in/go-on/lib.v2/html/tag"

	. "github.com/go-on/bootstrap/bs3"
)

func main() {
	tag.DIV(
		Success,
		BtnGroupXs,
		//Animated,
		"Success",

		tag.BUTTON(
			Btn, BtnDefault, BtnDanger,

			"Help!",
		),
	).Print()
}
