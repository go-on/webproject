package main

import (
	"fmt"
	. "github.com/go-on/css"
	. "github.com/go-on/css/c"
	. "github.com/go-on/css/color"
	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/html/tag"
)

var bestClass = Class("best")
var mainId = Id("main")
var newsClass = Class("news")

func main() {
	css := Css(DIV(mainId, bestClass),
		BackgroundColor(GREEN),

		Css(SECTION(newsClass),
			BorderBottom(Px(2), GREEN, Dotted_),

			Css(ARTICLE(),
				FontSize(Px(12)),

				Css(H2(),
					FontSize(Px(25)),

					Css(P(),
						MarginBottom(Px(10)),
					),
				),
			),
		),
	)

	fmt.Println(css)
}
