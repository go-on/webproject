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

					"Fixed width column",
				),

				META(
					Attrs_("content", "", "name", "description"),
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
					Attrs_("type", "text/javascript", "src", "../js/autoload.js"),
				),

				STYLE(

					"body {\n                background: #ededed;\n            }\n\n            header h1 small:before  {\n                content: \"|\";\n                margin: 0 0.5em;\n                font-size: 1.6em;\n            }\n\n            .highlight {\n                clear: none;\n            }\n\n            .highlight-wrapper {\n                overflow: hidden; /* this is really important to make the float work */\n            }\n\n            .pub {\n                width: 300px;\n            }\n\n            .pub a {\n                display: block;\n            }\n\n            footer {\n                background: #ccc;\n            }\n\n            @media screen and (min-width: 1261px) {\n                /* xlarge */\n                .highlight {\n                    /* create the needed space for the fixed width column */\n                    padding-right: 20.75em;\n                }\n            }\n\n            @media screen and (min-width: 961px) and (max-width: 1260px) {\n                /* large */\n                .highlight {\n                    /* create the needed space for the fixed width column */\n                    padding-right: 20.75em;\n                }\n            }\n\n            @media screen and (min-width: 641px) and (max-width: 960px) {\n                /* medium */\n                .highlight {\n                    padding-right: 0;\n                }\n            }\n\n            @media screen and (min-width: 321px) and (max-width: 640px) {\n                /* small */\n                .highlight {\n                    padding-right: 0;\n                }\n            }\n\n            @media screen and (max-width: 320px) {\n                /* tiny */\n                .highlight {\n                    padding-right: 0;\n                }\n            }",
				),
			),

			BODY(

				DIV(
					Class("ink-grid"),
					Comment("[if lte IE 9 ]>\n            <div class=\"ink-alert basic\" role=\"alert\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p>\n                    <strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n                </p>\n            </div>\n            <![endif]"),
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

									UL(
										Class("submenu"),

										LI(

											A(
												Attrs_("href", "#"),
												"One sub",
											),
										),

										LI(

											A(
												Attrs_("href", "#"),
												"Two sub",
											),
										),

										LI(

											A(
												Attrs_("href", "#"),
												"Three sub",
											),
										),

										LI(

											A(
												Attrs_("href", "#"),
												"Four sub",
											),
										),
									),
								),
							),
						),
					),

					SECTION(
						Class("pub hide-all show-large show-xlarge xlarge-push-right large-push-right"),

						DIV(
							Class("column-group gutters"),

							A(
								Class("all-100 mrec"), Attrs_("href", "#"),

								IMG(
									Attrs_("src", "../assets/js/holder.js/300x250/ink/auto/text: MREC PUB"),
								),
							),

							DIV(
								Class("all-50"),

								IMG(
									Attrs_("src", "../assets/js/holder.js/300x250/ink/auto/text: PUB"),
								),
							),

							DIV(
								Class("all-50"),

								IMG(
									Attrs_("src", "../assets/js/holder.js/300x250/ink/auto/text: PUB"),
								),
							),
						),

						P(

							"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
							E_mdash,
							"well, because it was scarlet. There is no other word for it.\"",
						),
					),

					SECTION(
						Class("highlight"),

						DIV(
							Class("highlight-wrapper"),

							DIV(
								Class("column-group gutters"),

								DIV(
									Class("xlarge-60 large-60 medium-50 small-100 tiny-100"),

									DIV(
										Class("image"),

										IMG(
											Attrs_("src", "../assets/js/holder.js/600x430/ink/auto/text: highlight"),
										),
									),
								),

								DIV(
									Class("xlarge-40 large-40 medium-50 small-100 tiny-100"),

									H2(

										"heading",
									),

									P(

										"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
										E_mdash,
										"well, because it was scarlet. There is no other word for it.\"",
									),
								),
							),

							DIV(
								Class("column-group gutters"),

								DIV(
									Class("all-50 small-100 tiny-100"),

									IMG(
										Attrs_("src", "../assets/js/holder.js/600x430/ink/auto/text: another story"),
									),
								),

								DIV(
									Class("all-50 small-100 tiny-100"),

									IMG(
										Attrs_("src", "../assets/js/holder.js/600x430/ink/auto/text: another story"),
									),
								),
							),

							DIV(
								Class("column-group gutters"),

								DIV(
									Class("all-33 small-100 tiny-100"),

									IMG(
										Attrs_("src", "../assets/js/holder.js/600x430/ink/auto/text: another story"),
									),
								),

								DIV(
									Class("all-33 small-100 tiny-100"),

									IMG(
										Attrs_("src", "../assets/js/holder.js/600x430/ink/auto/text: another story"),
									),
								),

								DIV(
									Class("all-33 small-100 tiny-100"),

									IMG(
										Attrs_("src", "../assets/js/holder.js/600x430/ink/auto/text: another story"),
									),
								),
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
									"About",
								),
							),

							LI(

								A(
									Attrs_("href", "#"),
									"Sitemap",
								),
							),

							LI(

								A(
									Attrs_("href", "#"),
									"Contacts",
								),
							),
						),

						P(
							Class("note"),
							"Identification of the owner of the copyright, either by name, abbreviation, or other designation by which it is generally known.",
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
