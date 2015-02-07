
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

"Quick start",
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
Attrs_("media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:landscape)", "rel", "apple-touch-startup-image", "href", "../img/splash.1024x748.png"), 
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

".screen-size-helper {\n                position: fixed;\n                bottom: 0;\n                left: 0;\n                right: 0;\n                width: 100%;\n                line-height: 1.6em;\n                font-size: 1em;\n                padding: 0.5333333333333333em 0.8em;\n                background: rgba(0, 0, 0, 0.7);\n                z-index: 100;\n            }\n            .screen-size-helper .title,\n            .screen-size-helper ul {\n                color: white;\n                text-shadow: 0 1px 0 #000000;\n            }\n            .screen-size-helper .title {\n                font-size: inherit;\n                line-height: inherit;\n                float: left;\n                text-transform: uppercase;\n                font-weight: 500;\n            }\n            .screen-size-helper ul {\n                float: right;\n                margin: 0;\n                padding: 0;\n                line-height: inherit !important;\n            }\n            .screen-size-helper ul li {\n                padding: 0;\n                margin: 0;\n                text-transform: uppercase;\n                font-weight: bold;\n                font-size: inherit !important;\n            }\n            .screen-size-helper ul li.tiny {\n              color: #0f75da;\n            }\n            .screen-size-helper ul li.small {\n              color: #4a9b17;\n            }\n            .screen-size-helper ul li.medium {\n              color: #ff9c00;\n            }\n            .screen-size-helper ul li.large {\n              color: #c91111;\n            }\n            .screen-size-helper ul li.xlarge {\n              color: white;\n            }",
),
),

BODY(

Comment("[if lte IE 9 ]>\n        <div class=\"ink-alert basic\" role=\"alert\">\n            <button class=\"ink-dismiss\">&times;</button>\n            <p>\n                <strong>You are using an outdated Internet Explorer version.</strong>\n                Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n            </p>\n        </div>\n        <![endif]"),
Comment(" Add your site or application content here "),

DIV(
Class("screen-size-helper"), 

P(
Class("title"), 
"Screen size:",
),

UL(
Class("unstyled"), 

LI(
Class("hide-all show-tiny tiny"), 
"TINY",
),

LI(
Class("hide-all show-small small"), 
"SMALL",
),

LI(
Class("hide-all show-medium medium"), 
"MEDIUM",
),

LI(
Class("hide-all show-large large"), 
"LARGE",
),

LI(
Class("hide-all show-xlarge xlarge"), 
"XLARGE",
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
		