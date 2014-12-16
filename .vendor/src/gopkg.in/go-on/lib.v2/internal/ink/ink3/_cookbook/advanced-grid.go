package main

import (
	"fmt"

	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/html/internal/element"
	. "gopkg.in/go-on/lib.v2/types"
)

var (
	_ = E_nbsp
	_ = A
	_ = Element{}
)

var elements = Elements(

	NewDocType("<!DOCTYPE html>",

		HTML(
			Attrs_("lang", "en"),

			HEAD(

				META(
					Attrs_("charset", "utf-8"),
				),

				META(
					Attrs_("http-equiv", "X-UA-Compatible", "content", "IE=edge,chrome=1"),
				),

				TITLE(

					"Advanced Grid",
				),

				META(
					Attrs_("name", "description", "content", ""),
				),

				META(
					Attrs_("name", "author", "content", "ink, cookbook, recipes"),
				),

				META(
					Attrs_("content", "True", "name", "HandheldFriendly"),
				),

				META(
					Attrs_("content", "320", "name", "MobileOptimized"),
				),

				META(
					Attrs_("name", "viewport", "content", "width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"),
				),
				Comment(" Place favicon.ico and apple-touch-icon(s) here  "),

				LINK(
					Attrs_("href", "../img/favicon.ico", "rel", "shortcut icon"),
				),

				LINK(
					Attrs_("rel", "apple-touch-icon-precomposed", "href", "../img/touch-icon.57.png"),
				),

				LINK(
					Attrs_("rel", "apple-touch-icon-precomposed", "sizes", "72x72", "href", "../img/touch-icon.72.png"),
				),

				LINK(
					Attrs_("rel", "apple-touch-icon-precomposed", "sizes", "114x114", "href", "../img/touch-icon.114.png"),
				),

				LINK(
					Attrs_("rel", "apple-touch-startup-image", "href", "../img/splash.320x460.png", "media", "screen and (min-device-width: 200px) and (max-device-width: 320px) and (orientation:portrait)"),
				),

				LINK(
					Attrs_("rel", "apple-touch-startup-image", "href", "../img/splash.768x1004.png", "media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:portrait)"),
				),

				LINK(
					Attrs_("rel", "apple-touch-startup-image", "href", "../img/splash.1024x748.png", "media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:landscape)"),
				),
				Comment(" load inks CSS "),

				LINK(
					Attrs_("rel", "stylesheet", "type", "text/css", "href", "../css/ink-flex.min.css"),
				),

				LINK(
					Attrs_("href", "../css/font-awesome.min.css", "rel", "stylesheet", "type", "text/css"),
				),
				Comment(" load inks CSS for IE8 "),
				Comment("[if lt IE 9 ]>\n            <link rel=\"stylesheet\" href=\"../css/ink-ie.min.css\" type=\"text/css\" media=\"screen\" title=\"no title\" charset=\"utf-8\">\n        <![endif]"),
				Comment(" test browser flexbox support and load legacy grid if unsupported "),

				SCRIPT(
					Attrs_("type", "text/javascript", "src", "../js/modernizr.js"),
				),

				SCRIPT(
					Attrs_("type", "text/javascript"),
					"Modernizr.load({\n              test: Modernizr.flexbox,\n              nope : '../css/ink-legacy.min.css'\n            });",
				),
				Comment(" load inks javascript files "),

				SCRIPT(
					Attrs_("type", "text/javascript", "src", "../js/holder.js"),
				),

				SCRIPT(
					Attrs_("type", "text/javascript", "src", "../js/ink-all.min.js"),
				),

				SCRIPT(
					Attrs_("type", "text/javascript", "src", "../js/autoload.js"),
				),

				STYLE(
					Attrs_("type", "text/css"),
					"body {\n                background: #ededed;\n            }\n\n            header {\n                border-bottom: 1px solid #cecece;\n            }\n\n            footer {\n                background: #ccc;\n            }",
				),
			),

			BODY(

				DIV(
					Class("ink-grid"),
					Comment("[if lte IE 9 ]>\n            <div class=\"ink-alert basic\" role=\"alert\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p>\n                    <strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n                </p>\n            </div>\n            <![endif]"),
					Comment(" Add your site or application content here "),

					HEADER(
						Class("clearfix vertical-padding"),

						DIV(
							Class("logo xlarge-push-left large-push-left"),

							A(
								Attrs_("href", "#"),

								IMG(
									Attrs_("src", "holder.js/150x90/auto/ink", "alt", ""),
								),
							),
						),

						NAV(
							Class("ink-navigation xlarge-push-right large-push-right half-top-space"),

							UL(
								Class("menu horizontal black"),

								LI(
									Class("active"),

									A(
										Attrs_("href", "#"),
										"item",
									),
								),

								LI(

									A(
										Attrs_("href", "#"),
										"item",
									),
								),

								LI(

									A(
										Attrs_("href", "#"),
										"item",
									),
								),
							),
						),
					),

					DIV(
						Class("column-group vertical-space"),

						DIV(
							Class("all-100"),

							IMG(
								Attrs_("src", "holder.js/1650x300/ink/auto"),
							),
						),
					),

					H1(

						"heading",
					),

					DIV(
						Class("column-group gutters"),

						DIV(
							Class("xlarge-50 large-50 all-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),

						DIV(
							Class("xlarge-25 large-25 all-50"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),

						DIV(
							Class("xlarge-25 large-25 all-50"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),
					),

					DIV(
						Class("column-group gutters"),

						DIV(
							Class("xlarge-50 large-50 all-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),

						DIV(
							Class("xlarge-25 large-25 all-50"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),

						DIV(
							Class("xlarge-25 large-25 all-50"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),
					),

					H2(

						"heading",
					),

					DIV(
						Class("column-group gutters"),

						DIV(
							Class("all-33 small-100 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),

						DIV(
							Class("all-33 small-100 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),

						DIV(
							Class("all-33 small-100 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/600x480/ink/auto"),
							),

							P(
								Class("quarter-top-space"),
								"\"Sit down,\" Edwin counselled soothingly. \"Granser's all right. He's just  gettin' to the Scarlet Death, ain't you, Granser? He's just goin' to  tell us about it right now. Sit down, Hare-Lip. Go ahead, Granser.\"",
							),
						),
					),

					DIV(
						Class("column-group gutters"),

						DIV(
							Class("xlarge-20 large-20 medium-25 small-50 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),

						DIV(
							Class("xlarge-20 large-20 medium-25 small-50 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),

						DIV(
							Class("xlarge-20 large-20 medium-25 small-50 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),

						DIV(
							Class("xlarge-20 large-20 medium-25 small-50 tiny-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),

						DIV(
							Class("xlarge-20 large-20 hide-small hide-medium tiny-100"),

							IMG(
								Attrs_("src", "holder.js/960x640/ink/auto"),
							),
						),
					),
				),

				FOOTER(
					Class("clearfix"),

					DIV(
						Class("ink-grid"),

						UL(
							Class("unstyled inline half-vertical-space"),

							LI(
								Class("active"),

								A(
									Attrs_("href", "#"),
									"Welcome",
								),
							),

							LI(

								A(
									Attrs_("href", "#"),
									"Portfolio",
								),
							),

							LI(

								A(
									Attrs_("href", "#"),
									"About",
								),
							),
						),

						P(
							Class("note"),
							"copyright © photoguy",
						),
					),
				),
			),
		),
	),
)

func main() {
	fmt.Println(elements)
}
