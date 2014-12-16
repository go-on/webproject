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

					"Modal gallery",
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
					Attrs_("rel", "shortcut icon", "href", "../img/favicon.ico"),
				),

				LINK(
					Attrs_("rel", "apple-touch-icon-precomposed", "href", "../img/touch-icon.57.png"),
				),

				LINK(
					Attrs_("rel", "apple-touch-icon-precomposed", "sizes", "72x72", "href", "../img/touch-icon.72.png"),
				),

				LINK(
					Attrs_("href", "../img/touch-icon.114.png", "rel", "apple-touch-icon-precomposed", "sizes", "114x114"),
				),

				LINK(
					Attrs_("href", "../img/splash.320x460.png", "media", "screen and (min-device-width: 200px) and (max-device-width: 320px) and (orientation:portrait)", "rel", "apple-touch-startup-image"),
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
					Attrs_("type", "text/css", "href", "../css/font-awesome.min.css", "rel", "stylesheet"),
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

					"body {\n                background: #ededed;\n            }\n\n            .panel {\n                border-radius: 2px;\n                -webkit-box-shadow: #dddddd 0 1px 1px 0;\n                -moz-box-shadow: #dddddd 0 1px 1px 0;\n                box-shadow: #dddddd 0 1px 1px 0;\n                padding: 1em;\n                border: 1px solid #BBB;\n                background-color: #FFF;\n            }\n\n            #thumbs_carousel .slide {\n                opacity: 0.6;\n            }\n            #thumbs_carousel .slide.active {\n                opacity: 1;\n            }",
				),
			),

			BODY(

				DIV(
					Class("ink-grid vertical-space"),

					H1(

						"Modal Gallery",
					),

					DIV(
						Class("panel"),

						DIV(
							Id("main_carousel"), Class("ink-carousel"),

							UL(
								Id("trigger-modal"), Class("stage column-group gutters"),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/1"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/2"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/3"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/4"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/5"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/6"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/7"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),

								LI(
									Class("slide all-100"),

									FIGURE(
										Class("ink-image"),

										A(
											Attrs_("href", "#"),

											IMG(
												Attrs_("src", "http://lorempixel.com/1650/928/city/8"),
											),
										),

										FIGCAPTION(
											Class("dark"),
											"This caption is placed after the image",
										),
									),
								),
							),

							NAV(
								Id("pagination_main"), Class("ink-navigation vertical-space"), Attrs_("data-next-label", "", "data-previous-label", ""),

								UL(
									Class("pagination chevron"),
								),
							),
						),
						Comment(" THUMBNAILS "),

						DIV(
							Id("thumbs_carousel"), Class("ink-carousel"),

							UL(
								Class("stage column-group gutters unstyled"),

								LI(
									Class("slide all-20 active"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/1"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/2"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/3"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/4"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/5"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/6"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/7"),
										),
									),
								),

								LI(
									Class("slide all-20"),

									A(
										Attrs_("href", "#"),

										IMG(
											Attrs_("src", "http://lorempixel.com/480/270/city/8"),
										),
									),
								),
							),

							NAV(
								Id("pagination_thumbs"), Class("ink-navigation top-space"), Attrs_("data-next-label", "", "data-previous-label", ""),

								UL(
									Class("pagination dotted"),
								),
							),
						),
						Comment(" MODAL "),

						DIV(
							Class("ink-shade fade"),

							DIV(
								Id("modal-window"), Class("ink-modal fade"), Attrs_("role", "dialog", "aria-hidden", "true", "aria-labelled-by", "modal-title", "data-trigger", "#trigger-modal", "data-width", "95%", "data-height", "95%"),

								DIV(
									Class("modal-header"),

									BUTTON(
										Class("modal-close ink-dismiss"),
									),

									H2(
										Id("modal-title"),
										"Use arrow left/right to navigate",
									),
								),

								DIV(
									Id("modalContent"), Class("modal-body"),
									Comment(" CAROUSEL MODAL "),

									DIV(
										Id("modal_carousel"), Class("ink-carousel"),

										UL(
											Class("stage column-group half-gutters unstyled"),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/1"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/2"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/3"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/4"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/5"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/6"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/7"),
												),
											),

											LI(
												Class("slide all-100"),

												IMG(
													Attrs_("src", "http://lorempixel.com/1650/928/city/8"),
												),
											),
										),

										NAV(
											Id("pagination_modal"), Class("ink-navigation"), Attrs_("data-next-label", "", "data-previous-label", ""),

											UL(
												Class("pagination chevron"),
											),
										),
									),
								),
							),
						),
					),
				),

				SCRIPT(
					Attrs_("type", "text/javascript", "src", "./gallery-modal.js"),
				),
			),
		),
	),
)

func main() {
	fmt.Println(elements)
}
