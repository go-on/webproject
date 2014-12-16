package main

import (
	"fmt"

	"gopkg.in/go-on/lib.v2/internal/ink/ink3"

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
			Lang_("en"),
			// Attrs_("lang", "en"),
			HEAD(
				CharsetUtf8(),
				/*
					META(
						Attrs_("charset", "utf-8"),
					),
				*/

				HttpEquiv("X-UA-Compatible", "IE=edge,chrome=1"),
				/*
					META(
						 Attrs_("http-equiv", "X-UA-Compatible", "content", "IE=edge,chrome=1"),
					),
				*/
				TITLE(
					"3 Boxes",
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
				Viewport("width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"),
				/*
					META(
						Attrs_("name", "viewport", "content", "width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"),
					),
				*/
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
					Attrs_("media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:landscape)", "rel", "apple-touch-startup-image", "href", "../img/splash.1024x748.png"),
				),
				Comment(" load inks CSS "),
				CssHref("../css/ink-flex.min.css"),
				/*
					LINK(
						Attrs_("rel", "stylesheet", "type", "text/css", "href", "../css/ink-flex.min.css"),
					),
				*/
				CssHref("../css/font-awesome.min.css"),
				/*
					LINK(
						Attrs_("href", "../css/font-awesome.min.css", "rel", "stylesheet", "type", "text/css"),
					),
				*/
				Comment(" load inks CSS for IE8 "),
				Comment("[if lt IE 9 ]>\n            <link rel=\"stylesheet\" href=\"../css/ink-ie.min.css\" type=\"text/css\" media=\"screen\" title=\"no title\" charset=\"utf-8\">\n        <![endif]"),
				Comment(" test browser flexbox support and load legacy grid if unsupported "),
				JsSrc("../js/modernizr.js"),
				/*
					SCRIPT(
						Attrs_("type", "text/javascript", "src", "../js/modernizr.js"),
					),
				*/
				SCRIPT(
					Attrs_("type", "text/javascript"),
					"Modernizr.load({\n              test: Modernizr.flexbox,\n              nope : '../css/ink-legacy.min.css'\n            });",
				),
				Comment(" load inks javascript files "),
				JsSrc("../js/holder.js"),
				/*
					SCRIPT(
						Attrs_("type", "text/javascript", "src", "../js/holder.js"),
					),
				*/
				JsSrc("../js/ink-all.min.js"),
				/*
					SCRIPT(
						Attrs_("src", "../js/ink-all.min.js", "type", "text/javascript"),
					),
				*/
				JsSrc("../js/autoload.js"),
				/*
					SCRIPT(
						Attrs_("type", "text/javascript", "src", "../js/autoload.js"),
					),
				*/
				STYLE(
					Attrs_("type", "text/css"),
					"body {\n                background: #ededed;\n            }\n\n            header h1 small:before  {\n                content: \"|\";\n                margin: 0 0.5em;\n                font-size: 1.6em;\n            }\n\n            footer {\n                background: #ccc;\n            }",
				),
			),

			BODY(
				DIV(
					Class("ink-grid"),
					Comment("[if lte IE 9 ]>\n            <div class=\"ink-alert basic\" role=\"alert\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p>\n                    <strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n                </p>\n            </div>\n            <![endif]"),
					Comment(" Add your site or application content here "),
					HEADER(
						ink3.Vertical_space,
						// Class("vertical-space"),
						H1(
							"site name",
							SMALL(
								"smaller text",
							),
						),
						NAV(
							ink3.Ink_navigation,
							// Class("ink-navigation"),
							UL(
								ink3.Horizontal, ink3.Black, ink3.Menu,
								// Class("menu horizontal black"),
								LI(
									ink3.Active,
									// Class("active"),
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
						ink3.Column_group, ink3.Gutters,
						// Class("column-group gutters"),
						DIV(
							ink3.All_100,
							// Class("all-100"),
							H2(
								"heading",
							),
							IMG(
								Attrs_("src", "holder.js/1650x928/auto/ink", "alt", ""),
							),
						),
					),

					DIV(
						ink3.Column_group, ink3.Gutters,
						// Class("column-group gutters"),
						DIV(
							ink3.All_50, ink3.Small_100, ink3.Tiny_100,
							// Class("all-50 small-100 tiny-100"),
							H3(
								"heading",
							),
							IMG(
								Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""),
							),
							P(
								ink3.Quarter_top_space,
								// Class("quarter-top-space"),
								"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
								E_mdash,
								"well, because it was scarlet. There is no other word for it.\"",
							),
						),
						DIV(
							ink3.All_50, ink3.Small_100, ink3.Tiny_100,
							// Class("all-50 small-100 tiny-100"),
							H3(
								"heading",
							),
							IMG(
								Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""),
							),
							P(
								// Class("quarter-top-space"),
								ink3.Quarter_top_space,
								"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
								E_mdash,
								"well, because it was scarlet. There is no other word for it.\"",
							),
						),
					),
				),

				FOOTER(
					ink3.Clearfix,
					// Class("clearfix"),
					DIV(
						ink3.Ink_grid,
						// Class("ink-grid"),
						UL(
							ink3.Unstyled, ink3.Inline, ink3.Half_vertical_space,
							// Class("unstyled inline half-vertical-space"),
							LI(
								ink3.Active,
								// Class("active"),
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
							ink3.Note,
							// Class("note"),
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
