package main

import (
	"fmt"

	. "gopkg.in/go-on/lib.v3/html"
	. "gopkg.in/go-on/lib.v3/html/internal/element"
	. "gopkg.in/go-on/lib.v3/types"
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

					"4 Boxes",
				),

				META(
					Attrs_("name", "description", "content", ""),
				),

				META(
					Attrs_("name", "author", "content", "ink, cookbook, recipes"),
				),

				META(
					Attrs_("name", "HandheldFriendly", "content", "True"),
				),

				META(
					Attrs_("name", "MobileOptimized", "content", "320"),
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
					Attrs_("type", "text/css", "href", "../css/ink-flex.min.css", "rel", "stylesheet"),
				),

				LINK(
					Attrs_("rel", "stylesheet", "type", "text/css", "href", "../css/font-awesome.min.css"),
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
					Attrs_("src", "../js/autoload.js", "type", "text/javascript"),
				),

				STYLE(
					Attrs_("type", "text/css"),
					"body {\n                background: #ededed;\n            }\n\n            header h1 small:before  {\n                content: \"|\";\n                margin: 0 0.5em;\n                font-size: 1.6em;\n            }\n\n            footer {\n                background: #ccc;\n            }",
				),
			),

			BODY(

				DIV(
					Class("ink-grid"),
					Comment("[if lte IE 9 ]>\n            <div class=\"ink-alert basic\" role=\"alert\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p><strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n                </p>\n            </div>\n            <![endif]"),
					Comment(" Add your site or application content here "),

					HEADER(
						Class("vertical-space"),

						H1(

							"site name",

							SMALL(

								"smaller text",
							),
						),

						NAV(
							Class("ink-navigation"),

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
						Class("button-toolbar"),

						DIV(
							Class("button-group"),

							BUTTON(
								Class("ink-button"), Attrs_("disabled", ""),
								"left",
							),

							BUTTON(
								Class("ink-button"),
								"middle",
							),

							BUTTON(
								Class("ink-button"),
								"right",
							),
						),

						DIV(
							Class("button-group"),

							BUTTON(
								Class("ink-button"),
								"left",
							),

							BUTTON(
								Class("ink-button"), Attrs_("disabled", ""),
								"middle",
							),

							BUTTON(
								Class("ink-button"),
								"right",
							),
						),

						DIV(
							Class("button-group"),

							BUTTON(
								Class("ink-button"),
								"left",
							),

							BUTTON(
								Class("ink-button"),
								"middle",
							),

							BUTTON(
								Class("ink-button"), Attrs_("disabled", ""),
								"right",
							),
						),

						DIV(
							Class("button-group"),

							BUTTON(
								Class("ink-button green"), Attrs_("disabled", ""),
								"left",
							),

							BUTTON(
								Class("ink-button green"),
								"middle",
							),

							BUTTON(
								Class("ink-button green"),
								"right",
							),
						),

						DIV(
							Class("button-group"),

							BUTTON(
								Class("ink-button green"),
								"left",
							),

							BUTTON(
								Class("ink-button green"), Attrs_("disabled", ""),
								"middle",
							),

							BUTTON(
								Class("ink-button green"),
								"right",
							),
						),

						DIV(
							Class("button-group"),

							BUTTON(
								Class("ink-button black"),
								"left",
							),

							BUTTON(
								Class("ink-button black"),
								"middle",
							),

							BUTTON(
								Class("ink-button black"), Attrs_("disabled", ""),
								"right",
							),
						),
					),

					DIV(
						Class("column-group half-top-space"),

						DIV(
							Class("all-33"),

							DIV(
								Class("column-group"),

								DIV(
									Class("all-50 align-center"),

									A(
										Class("ink-button"), Attrs_("href", "#"),
										"Button",
									),
								),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button green"),
										"Button",
									),
								),
							),
						),

						DIV(
							Class("all-33"),

							DIV(
								Class("column-group"),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button orange"),
										"Button",
									),
								),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button red"),
										"Button",
									),
								),
							),
						),

						DIV(
							Class("all-33"),

							DIV(
								Class("column-group"),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button blue"),
										"Button",
									),
								),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button black"),
										"Button",
									),
								),
							),
						),
					),

					DIV(
						Class("column-group half-vertical-space"),

						DIV(
							Class("all-33"),

							DIV(
								Class("column-group"),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button"), Attrs_("disabled", ""),
										"Button",
									),
								),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button green"), Attrs_("disabled", ""),
										"Button",
									),
								),
							),
						),

						DIV(
							Class("all-33"),

							DIV(
								Class("column-group"),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button orange"), Attrs_("disabled", ""),
										"Button",
									),
								),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button red"), Attrs_("disabled", ""),
										"Button",
									),
								),
							),
						),

						DIV(
							Class("all-33"),

							DIV(
								Class("column-group"),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button blue"), Attrs_("disabled", ""),
										"Button",
									),
								),

								DIV(
									Class("all-50 align-center"),

									BUTTON(
										Class("ink-button black"), Attrs_("disabled", ""),
										"Button",
									),
								),
							),
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
