
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

"Carousel",
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
Attrs_("media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:portrait)", "rel", "apple-touch-startup-image", "href", "../img/splash.768x1004.png"), 
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

"body {\n                background: #ededed;\n            }\n\n            .panel {\n                border-radius: 2px;\n                -webkit-box-shadow: #dddddd 0 1px 1px 0;\n                -moz-box-shadow: #dddddd 0 1px 1px 0;\n                box-shadow: #dddddd 0 1px 1px 0;\n                padding: 1em;\n                border: 1px solid #BBB;\n                background-color: #FFF;\n            }",
),
),

BODY(


DIV(
Class("ink-grid vertical-space"), 
Comment("[if lte IE 9 ]>\n    <div class=\"ink-alert basic\" role=\"alert\">\n        <button class=\"ink-dismiss\">&times;</button>\n        <p>\n            <strong>You are using an outdated Internet Explorer version.</strong>\n            Please <a href=\"http://browsehappy.com/\">upgrade to a modern browser</a> to improve your web experience.\n        </p>\n    </div>\n    <![endif]"),

DIV(
Class("panel"), 

DIV(
Id("car1"), Class("ink-carousel"), Attrs_("data-space-after-last-slide", "false"), 

UL(
Class("stage column-group half-gutters"), 

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/1"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/2"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"Devemos cultivar os nossos soft powers sagrados",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/3"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"Rumo ao inconseguimento de cenas",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/4"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/5"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/6"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/7"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/8"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/sports/9"), 
),

DIV(
Class("description"), 

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),
),
),

NAV(
Id("p1"), Class("ink-navigation"), 

UL(
Class("pagination black"), 
),
),
),

SCRIPT(

"Ink.requireModules(['Ink.UI.Carousel_1'], function(InkCarousel) {\n            new InkCarousel('#car1', {pagination: '#p1'});\n        });",
),

DIV(
Class("panel vertical-space"), 

DIV(
Id("car2"), Class("ink-carousel"), 

UL(
Class("stage column-group half-gutters unstyled"), 

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/1"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impactoagilidade decisória.",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/2"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"Devemos cultivar os nossos soft powers sagrados",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/3"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"Rumo ao inconseguimento de cenas",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/4"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/5"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/6"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/7"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-25 large-25 medium-33 small-50 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/400/200/city/8"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

NAV(
Id("p2"), Class("ink-navigation"), Attrs_("data-next-label", "", "data-previous-label", ""), 

UL(
Class("pagination dotted"), 
),
),
),
),

SCRIPT(

"Ink.requireModules(['Ink.UI.Carousel_1'], function(InkCarousel) {\n            new InkCarousel('#car2', {pagination: '#p2'})\n        });",
),

DIV(
Class("panel vertical-space"), 

DIV(
Id("car3"), Class("ink-carousel"), 

UL(
Class("stage column-group half-gutters unstyled"), 

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/1"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impactoagilidade decisória.",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/2"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"Devemos cultivar os nossos soft powers sagrados",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/3"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"Rumo ao inconseguimento de cenas",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/4"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/5"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/6"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/7"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),

LI(
Class("slide xlarge-100 large-100 medium-100 small-100 tiny-100"), 

IMG(
Class("half-bottom-space"), Attrs_("src", "http://lorempixel.com/1400/675/nightlife/8"), 
),

H4(
Class("no-margin"), 
"Highlight Title",
),

H5(
Class("slab"), 
"Baby Doe",
),

P(

"É importante questionar o quanto a constante divulgação das informações ainda não demonstrou convincentemente que vai participar na mudança do impacto na agilidade decisória.",
),
),
),

NAV(
Id("p3"), Class("ink-navigation"), Attrs_("data-next-label", "", "data-previous-label", ""), 

UL(
Class("pagination chevron"), 
),
),
),

SCRIPT(

"Ink.requireModules(['Ink.UI.Carousel_1'], function(InkCarousel) {\n                new InkCarousel('#car3', { pagination: '#p3', nextLabel: '', previousLabel: ''});\n            });",
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
		