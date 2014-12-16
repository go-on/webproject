
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
Attrs_("content", "IE=edge,chrome=1", "http-equiv", "X-UA-Compatible"), 
),

TITLE(

"Article page",
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
Attrs_("href", "../img/touch-icon.114.png", "rel", "apple-touch-icon-precomposed", "sizes", "114x114"), 
),

LINK(
Attrs_("href", "../img/splash.320x460.png", "media", "screen and (min-device-width: 200px) and (max-device-width: 320px) and (orientation:portrait)", "rel", "apple-touch-startup-image"), 
),

LINK(
Attrs_("rel", "apple-touch-startup-image", "href", "../img/splash.768x1004.png", "media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:portrait)"), 
),

LINK(
Attrs_("href", "../img/splash.1024x748.png", "media", "screen and (min-device-width: 481px) and (max-device-width: 1024px) and (orientation:landscape)", "rel", "apple-touch-startup-image"), 
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
Attrs_("type", "text/css"), 
"body {\n                background: #ededed;\n            }\n\n            header h1 small:before  {\n                content: \"|\";\n                margin: 0 0.5em;\n                font-size: 1.6em;\n            }\n\n            article header{\n                padding: 0;\n                overflow: hidden;\n            }\n\n            article footer {\n                background: none;\n            }\n\n            article {\n                padding-bottom: 2em;\n                border-bottom: 1px solid #ccc;\n            }\n\n            .date {\n                float: right;\n            }\n\n            summary {\n                font-weight: 700;\n                line-height: 1.5;\n            }\n\n            footer {\n                background: #ccc;\n            }",
),
),

BODY(


DIV(
Class("ink-grid"), 
Comment("[if lte IE 9 ]>\n            <div class=\"ink-alert basic\" role=\"alert\">\n                <button class=\"ink-dismiss\">&times;</button>\n                <p>\n                    <strong>You are using an outdated Internet Explorer version.</strong>\n                    Please <a href=\"http://browsehappy.com/\">upgrade to a modern  browser</a> to improve your web experience.\n                </p>\n            </div>\n            <![endif]"),
Comment(" Add your site or application content here "),

HEADER(
Class("clearfix vertical-padding"), 

H1(
Class("logo xlarge-push-left large-push-left"), 
"site name",

SMALL(

"smaller text",
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

SECTION(
Class("column-group gutters article"), 

DIV(
Class("xlarge-70 large-70 medium-60 small-100 tiny-100"), 

ARTICLE(


HEADER(


H1(
Class("push-left"), 
"Article title",
),

P(
Class("push-right"), 
"Published:",

TIME(
Attrs_("pubdate", "pubdate"), 
"2009-10-09",
),
),
),

SUMMARY(

"But being in a great hurry to resume scolding the man in the purple Shirt, who was waiting for it in the entry, and seeming to hear nothing but the word \"clam,\" Mrs. Hussey hurried towards an open door leading to the kitchen, and bawling out \"clam for two,\" disappeared.",
),

FIGURE(
Class("ink-image vertical-space"), 

IMG(
Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),

FIGCAPTION(
Class("dark over-bottom"), 
"\"Because you wear silver shoes and have killed the Wicked Witch. Besides, you have white in your frock, and only witches and sorceresses wear white.\"",
),
),

P(


BLOCKQUOTE(

"\"Queequeg,\" said I, \"do you think that we can make out a supper for us both on one clam?\"",
),
),

P(

"However, a warm savory steam from the kitchen served to belie the apparently cheerless prospect before us. But when that smoking chowder came in, the mystery was delightfully explained. Oh, sweet friends! hearken to me. It was made of small juicy clams, scarcely bigger than hazel nuts, mixed with pounded ship biscuit, and salted pork cut up into little flakes; the whole enriched with butter, and plentifully seasoned with pepper and salt. Our appetites being sharpened by the frosty voyage, and in particular, Queequeg seeing his favourite fishing food before him, and the chowder being surpassingly excellent, we despatched it with great expedition: when leaning back a moment and bethinking me of Mrs. Hussey's clam and cod announcement, I thought I would try a little experiment. Stepping to the kitchen door, I uttered the word \"cod\" with great emphasis, and resumed my seat. In a few moments the savoury steam came forth again, but with a different flavor, and in good time a fine cod-chowder was placed before us.",
),

FOOTER(


P(


SMALL(

"Creative Commons Attribution-ShareAlike License",
),
),
),
),
),

SECTION(
Class("xlarge-30 large-30 medium-40 small-100 tiny-100"), 

H2(

"Related",
),

UL(
Class("unstyled"), 

LI(
Class("column-group half-gutters"), 

DIV(
Class("all-40 small-50 tiny-50"), 

IMG(
Attrs_("src", "holder.js/640x380/auto/ink", "alt", ""), 
),
),

DIV(
Class("all-60 small-50 tiny-50"), 

P(

"\"Where's them crabs, Hoo-Hoo?\" Edwin demanded. \"Granser's set upon  having a snack.\"",
),
),
),

LI(
Class("column-group half-gutters"), 

DIV(
Class("all-40 small-50 tiny-50"), 

IMG(
Attrs_("src", "holder.js/640x380/auto/ink", "alt", ""), 
),
),

DIV(
Class("all-60 small-50 tiny-50"), 

P(

"\"Where's them crabs, Hoo-Hoo?\" Edwin demanded. \"Granser's set upon  having a snack.\"",
),
),
),

LI(
Class("column-group half-gutters"), 

DIV(
Class("all-40 small-50 tiny-50"), 

IMG(
Attrs_("src", "holder.js/640x380/auto/ink", "alt", ""), 
),
),

DIV(
Class("all-60 small-50 tiny-50"), 

P(

"\"Where's them crabs, Hoo-Hoo?\" Edwin demanded. \"Granser's set upon  having a snack.\"",
),
),
),

LI(
Class("column-group half-gutters"), 

DIV(
Class("all-40 small-50 tiny-50"), 

IMG(
Attrs_("src", "holder.js/640x380/auto/ink", "alt", ""), 
),
),

DIV(
Class("all-60 small-50 tiny-50"), 

P(

"\"Where's them crabs, Hoo-Hoo?\" Edwin demanded. \"Granser's set upon  having a snack.\"",
),
),
),

LI(
Class("column-group half-gutters"), 

DIV(
Class("all-40 small-50 tiny-50"), 

IMG(
Attrs_("src", "holder.js/640x380/auto/ink", "alt", ""), 
),
),

DIV(
Class("all-60 small-50 tiny-50"), 

P(

"\"Where's them crabs, Hoo-Hoo?\" Edwin demanded. \"Granser's set upon  having a snack.\"",
),
),
),
),
),
),

SECTION(
Class("column-group gutters"), 

DIV(
Class("all-20 small-100 tiny-100"), 

H3(

"heading",
),

IMG(
Class("half-bottom-space"), Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),

P(

"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
E_mdash,
"well, because it was scarlet. There is no other word for it.\"",
),
),

DIV(
Class("all-20 small-100 tiny-100"), 

H3(

"heading",
),

IMG(
Class("half-bottom-space"), Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),

P(

"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
E_mdash,
"well, because it was scarlet. There is no other word for it.\"",
),
),

DIV(
Class("all-20 small-100 tiny-100"), 

H3(

"heading",
),

IMG(
Class("half-bottom-space"), Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),

P(

"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
E_mdash,
"well, because it was scarlet. There is no other word for it.\"",
),
),

DIV(
Class("all-20 small-100 tiny-100"), 

H3(

"heading",
),

IMG(
Class("half-bottom-space"), Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),

P(

"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
E_mdash,
"well, because it was scarlet. There is no other word for it.\"",
),
),

DIV(
Class("all-20 small-100 tiny-100"), 

H3(

"heading",
),

IMG(
Class("half-bottom-space"), Attrs_("src", "holder.js/1200x600/auto/ink", "alt", ""), 
),

P(

"\"Red is not the right word,\" was the reply. \"The plague was scarlet.  The whole face and body turned scarlet in an hour's time. Don't I  know? Didn't I see enough of it? And I am telling you it was scarlet  because",
E_mdash,
"well, because it was scarlet. There is no other word for it.\"",
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
		