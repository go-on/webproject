
	package main

	import (
		 "fmt"
		. "gopkg.in/go-on/lib.v3/types"
		. "gopkg.in/go-on/lib.v3/html"
		. "gopkg.in/go-on/lib.v3/html/internal/element"	   
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

"Sticky footer",
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
"html, body {\n                height: 100%;\n                background: #f0f0f0;\n            }\n            .wrap {\n                min-height: 100%;\n                height: auto !important;\n                height: 100%;\n                margin: 0 auto -120px;\n                overflow: auto;\n            }\n            .push, footer {\n                height: 120px;\n                margin-top: 0;\n            }\n            footer {\n                background: #ccc;\n                border: 0;\n            }\n            footer * {\n                line-height: inherit;\n            }\n            .top-menu {\n                background: #1a1a1a;\n            }",
),
),

BODY(

Comment("[if lte IE 9 ]>\n        <div class=\"ink-grid\">\n            <div class=\"ink-alert basic\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p>\n                    <strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n                </p>\n            </div>\n        </div>\n        <![endif]"),

DIV(
Class("wrap"), 

DIV(
Class("top-menu"), 

NAV(
Class("ink-navigation ink-grid"), 

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
Class("ink-grid vertical-space"), 

H1(

"heading",
),

P(

"And though of all men the moody captain of the Pequod was the least given to that sort of shallowest assumption; and though the only homage he ever exacted, was implicit, instantaneous obedience; though he required no man to remove the shoes from his feet ere stepping upon the quarter-deck; and though there were times when, owing to peculiar circumstances connected with events hereafter to be detailed, he addressed them in unusual terms, whether of condescension or IN TERROREM, or otherwise; yet even Captain Ahab was by no means unobservant of the paramount forms and usages of the sea.",
),
),

DIV(
Class("push"), 
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
		