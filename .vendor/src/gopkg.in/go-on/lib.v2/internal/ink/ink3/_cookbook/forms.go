
	package main

	import (
		 "fmt"
		. "gopkg.in/go-on/lib.v2/types"
		. "gopkg.in/go-on/lib.v2/html"
		. "gopkg.in/go-on/lib.v2/html/internal/element"	   
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

"Forms",
),

META(
Attrs_("content", "", "name", "description"), 
),

META(
Attrs_("content", "ink, cookbook, recipes", "name", "author"), 
),

META(
Attrs_("name", "HandheldFriendly", "content", "True"), 
),

META(
Attrs_("name", "MobileOptimized", "content", "320"), 
),

META(
Attrs_("content", "width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0", "name", "viewport"), 
),
Comment(" Place favicon.ico and apple-touch-icon(s) here  "),

LINK(
Attrs_("rel", "shortcut icon", "href", "../img/favicon.ico"), 
),

LINK(
Attrs_("rel", "apple-touch-icon-precomposed", "href", "../img/touch-icon.57.png"), 
),

LINK(
Attrs_("href", "../img/touch-icon.72.png", "rel", "apple-touch-icon-precomposed", "sizes", "72x72"), 
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
Attrs_("src", "../js/autoload.js", "type", "text/javascript"), 
),

STYLE(
Attrs_("type", "text/css"), 
"body {\n                background: #ededed;\n            }\n            header {\n                padding: 2em 0;\n                margin-bottom: 2em;\n            }\n            header h1 {\n                font-size: 2em;\n            }\n            header h1 small:before  {\n                content: \"|\";\n                margin: 0 0.5em;\n                font-size: 1.6em;\n            }\n            footer {\n                background: #ccc;\n            }",
),
),

BODY(


DIV(
Class("ink-grid"), 
Comment("[if lte IE 9 ]>\n            <div class=\"ink-alert basic\" role=\"alert\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p>\n                    <strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n                </p>\n            </div>\n            <![endif]"),
Comment(" Add your site or application content here "),

HEADER(


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
Class("column-group gutters"), 

FORM(
Class("ink-form all-50 small-100 tiny-100"), Attrs_("action", ""), 

FIELDSET(


DIV(
Class("control-group required column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "first-name"), 
"Name",
),

DIV(
Class("control all-40"), 

INPUT(
Id("first-name"), Attrs_("type", "text"), 
),

P(
Class("tip"), 
"First Name",
),
),

DIV(
Class("control all-40"), 

INPUT(
Id("last-name"), Attrs_("type", "text"), 
),

P(
Class("tip"), 
"Last Name",
),
),
),

DIV(
Class("control-group required column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "email"), 
"Email",
),

DIV(
Class("control all-80"), 

INPUT(
Id("email"), Attrs_("type", "text"), 
),
),
),

DIV(
Class("control-group column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "area"), 
"Description",
),

DIV(
Class("control all-80"), 

TEXTAREA(
Id("area"), 
),
),
),

DIV(
Class("control-group column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "file-input"), 
"File input",
),

DIV(
Class("control all-80"), 

DIV(
Class("input-file"), 

INPUT(
Id("file-input"), Attrs_("type", "file", "name", ""), 
),
),
),
),
),
),

DIV(
Class("all-50 small-100 tiny-100"), 

P(

"Chuck Norris once kicked a baby elephant into puberty. Crop circles are Chuck Norris' way of telling the world that sometimes corn needs to lie the f*ck down.",
),

IMG(
Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),
),
),

FORM(
Class("ink-form column-group gutters"), Attrs_("action", "#"), 

FIELDSET(
Class("all-33 small-100 tiny-100"), 

LEGEND(

"Fieldset title",
),

DIV(
Class("control-group"), 

LABEL(
Attrs_("for", "name2"), 
"Name",
),

DIV(
Class("control"), 

INPUT(
Id("name2"), Attrs_("type", "text"), 
),
),
),

DIV(
Class("control-group"), 

LABEL(
Attrs_("for", "phone2"), 
"Phone",
),

DIV(
Class("control"), 

INPUT(
Id("phone2"), Attrs_("type", "text"), 
),
),
),

DIV(
Class("control-group"), 

LABEL(
Attrs_("for", "email2"), 
"Email",
),

DIV(
Class("control"), 

INPUT(
Id("email2"), Attrs_("type", "text"), 
),
),
),
),

FIELDSET(
Class("all-33 small-100 tiny-100"), 

LEGEND(

"Rock Bands",
),

DIV(
Class("control-group"), 

P(
Class("label"), 
"Please select one or more options",
),

UL(
Class("control unstyled"), 

LI(


INPUT(
Id("cb1"), Attrs_("name", "cb1", "value", "", "type", "checkbox"), 
),

LABEL(
Attrs_("for", "cb1"), 
"Chimaira",
),
),

LI(


INPUT(
Id("cb2"), Attrs_("type", "checkbox", "name", "cb2", "value", ""), 
),

LABEL(
Attrs_("for", "cb2"), 
"Metallica",
),
),

LI(


INPUT(
Id("cb3"), Attrs_("type", "checkbox", "name", "cb3", "value", ""), 
),

LABEL(
Attrs_("for", "cb3"), 
"Motörhead",
),
),

LI(


INPUT(
Id("cb4"), Attrs_("name", "cb4", "value", "", "type", "checkbox"), 
),

LABEL(
Attrs_("for", "cb4"), 
"Pantera",
),
),

LI(


INPUT(
Id("cb5"), Attrs_("type", "checkbox", "name", "cb5", "value", ""), 
),

LABEL(
Attrs_("for", "cb5"), 
"Slayer",
),
),

LI(


INPUT(
Id("cb6"), Attrs_("type", "checkbox", "name", "cb6", "value", ""), 
),

LABEL(
Attrs_("for", "cb6"), 
"Switchtense",
),
),
),
),
),

FIELDSET(
Class("all-33 small-100 tiny-100"), 

LEGEND(

"Pick a card",
),

DIV(
Class("control-group"), 

P(
Class("label"), 
"Please select one of these options",
),

UL(
Class("control unstyled"), 

LI(


INPUT(
Id("rb1"), Attrs_("type", "radio", "name", "rb", "value", ""), 
),

LABEL(
Attrs_("for", "rb1"), 
"Ace of Spades",
),
),

LI(


INPUT(
Id("rb2"), Attrs_("type", "radio", "name", "rb", "value", ""), 
),

LABEL(
Attrs_("for", "rb2"), 
"Queen of Diamonds",
),
),

LI(


INPUT(
Id("rb3"), Attrs_("type", "radio", "name", "rb", "value", ""), 
),

LABEL(
Attrs_("for", "rb3"), 
"Queen of Spades",
),
),

LI(


INPUT(
Id("rb4"), Attrs_("value", "", "type", "radio", "name", "rb"), 
),

LABEL(
Attrs_("for", "rb4"), 
"Jack of Hearts",
),
),

LI(


INPUT(
Id("rb5"), Attrs_("name", "rb", "value", "", "type", "radio"), 
),

LABEL(
Attrs_("for", "rb5"), 
"King of Clubs",
),
),

LI(


INPUT(
Id("rb6"), Attrs_("type", "radio", "name", "rb", "value", ""), 
),

LABEL(
Attrs_("for", "rb6"), 
"King of Diamonds",
),
),
),
),
),
),

FORM(
Class("ink-form"), Attrs_("action", "#"), 

H4(

"Inline form with inline fields",
),

P(

"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla gravida lacus purus. Integer turpis enim, condimentum non pellentesque vel, consequat vitae diam. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.",
),

DIV(
Class("column-group gutters vertical-space"), 

DIV(
Class("control-group all-33 small-100 tiny-100"), 

DIV(
Class("column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "name3"), 
"Name",
),

DIV(
Class("control all-80"), 

INPUT(
Id("name3"), Attrs_("type", "text"), 
),
),
),
),

DIV(
Class("control-group all-33 small-100 tiny-100"), 

DIV(
Class("column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "phone3"), 
"Phone",
),

DIV(
Class("control all-80"), 

INPUT(
Id("phone3"), Attrs_("type", "text"), 
),
),
),
),

DIV(
Class("control-group all-33 small-100 tiny-100"), 

DIV(
Class("column-group gutters"), 

LABEL(
Class("all-20 align-right"), Attrs_("for", "email3"), 
"Email",
),

DIV(
Class("control all-80"), 

INPUT(
Id("email3"), Attrs_("type", "text"), 
),
),
),
),
),
),

FORM(
Class("ink-form"), Attrs_("action", "#"), 

H4(

"Block form with inline fields",
),

P(

"Proin nibh nulla, consequat vitae aliquet nec, consequat consectetur quam. Morbi diam dui, ornare vel elementum ut, pharetra at urna. Proin vel purus orci, vel euismod lorem. In hac habitasse platea dictumst. Donec eu scelerisque velit. Suspendisse velit lectus, ultrices vitae luctus vel, lobortis non metus.",
),

DIV(
Class("column-group gutters vertical-space"), 

DIV(
Class("control-group all-33 small-100 tiny-100"), 

LABEL(
Attrs_("for", "name4"), 
"Name",
),

DIV(
Class("control"), 

INPUT(
Id("name4"), Attrs_("type", "text"), 
),
),
),

DIV(
Class("control-group all-33 small-100 tiny-100"), 

LABEL(
Attrs_("for", "phone4"), 
"Phone",
),

DIV(
Class("control"), 

INPUT(
Id("phone4"), Attrs_("type", "text"), 
),
),
),

DIV(
Class("control-group all-33 small-100 tiny-100"), 

LABEL(
Attrs_("for", "email4"), 
"Email",
),

DIV(
Class("control"), 

INPUT(
Id("email4"), Attrs_("type", "text"), 
),
),
),
),
),

FORM(
Class("ink-form"), Attrs_("action", "#"), 

H4(

"Appended and prepended buttons and symbols",
),

P(

"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla gravida lacus purus. Integer turpis enim, condimentum non pellentesque vel, consequat vitae diam. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.",
),

FIELDSET(


DIV(
Class("column-group gutters vertical-space"), 

DIV(
Class("control-group all-50 small-100 tiny-100"), 

DIV(
Class("column-group gutters"), 

DIV(
Class("control all-100 small-100 tiny-100 append-button"), Attrs_("role", "search"), 

SPAN(


INPUT(
Id("name5"), Attrs_("type", "text"), 
),
),

BUTTON(
Class("ink-button"), 

I(
Class("icon-search"), 
),
"Search",
),
),
),

DIV(
Class("column-group gutters"), 

DIV(
Class("control all-100 small-100 tiny-100 prepend-button"), Attrs_("role", "search"), 

INPUT(
Class("ink-button"), Attrs_("type", "submit", "value", "Search"), 
),

SPAN(


INPUT(
Id("phone5"), Attrs_("type", "text"), 
),
),
),
),
),

DIV(
Class("control-group all-50 small-100 tiny-100"), 

DIV(
Class("column-group gutters"), 

DIV(
Class("control all-100 small-100 tiny-100 append-symbol"), 

SPAN(


INPUT(
Id("email5"), Attrs_("type", "text"), 
),

I(
Class("fa fa-envelope-o"), 
),
),
),
),

DIV(
Class("column-group gutters"), 

DIV(
Class("control all-100 small-100 tiny-100 prepend-symbol"), 

SPAN(


INPUT(
Id("email6"), Attrs_("type", "text"), 
),

I(
Class("fa fa-envelope-o"), 
),
),
),
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
		