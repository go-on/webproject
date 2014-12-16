package ink3

import (
	"fmt"

	"gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/types"
)

var V3_0_5 = CDN(3, 0, 5)

func CDN(major, minor, patch uint) cdn {
	return cdn{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

type cdn struct {
	Major, Minor, Patch uint
	CacheFunc           func(string) string
}

// relPath must not start with /
func (i cdn) File(relPath string) string {
	path := i.basePath() + relPath

	if i.CacheFunc == nil {
		return path
	}

	return i.CacheFunc(path)
}

func (i cdn) basePath() string {
	return fmt.Sprintf("//fastly.ink.sapo.pt/%d.%d.%d/", i.Major, i.Minor, i.Patch)
}

// relPath must not start with /
func (i cdn) CSSFile(relPath string) string  { return i.File("css/" + relPath) }
func (i cdn) FontFile(relPath string) string { return i.File("fonts/" + relPath) }
func (i cdn) ImgFile(relPath string) string  { return i.File("img/" + relPath) }
func (i cdn) JSFile(relPath string) string   { return i.File("js/" + relPath) }

func (i cdn) InkCSS() string         { return i.CSSFile("ink.min.css") }
func (i cdn) FontawesomeCSS() string { return i.CSSFile("font-awesome.min.css") }
func (i cdn) InkFlexCSS() string     { return i.CSSFile("ink-flex.min.css") }
func (i cdn) InkIECSS() string       { return i.CSSFile("ink-ie.min.css") }
func (i cdn) InkLegacyCSS() string   { return i.CSSFile("ink-legacy.min.css") }
func (i cdn) QuickStartCSS() string  { return i.CSSFile("quick-start.css") }

func (i cdn) InkAllJS() string         { return i.JSFile("ink-all.min.js") }
func (i cdn) HTML5ShivJS() string      { return i.JSFile("html5shiv.js") }
func (i cdn) HTML5PrintShivJS() string { return i.JSFile("html5shiv-printshiv.js") }
func (i cdn) ModernizrJS() string      { return i.JSFile("modernizr.js") }
func (i cdn) ModernizrAllJS() string   { return i.JSFile("modernizr-all.js") }
func (i cdn) AutoloadJS() string       { return i.JSFile("autoload.js") }
func (i cdn) HolderJS() string         { return i.JSFile("holder.js") }

func (i cdn) Head() types.HTMLer {
	return types.HTMLString(
		element.Elements(
			html.CharsetUtf8(),
			html.HttpEquiv("X-UA-Compatible", "IE=edge,chrome=1"),
			html.META(html.Attrs_("name", "HandheldFriendly", "content", "True")),
			html.META(html.Attrs_("name", "MobileOptimized", "content", "320")),
			html.Viewport("width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"),
			html.CssHref(i.InkFlexCSS()),
			html.CssHref(i.FontawesomeCSS()),
			types.Comment(`[if lt IE 9 ]><link rel="stylesheet" href="`+i.InkIECSS()+`" type="text/css" media="screen" title="no title" charset="utf-8"><![endif]`),
			html.JsSrc(i.ModernizrJS()),
			html.SCRIPT(
				html.Attrs_("type", "text/javascript"),
				`Modernizr.load({test: Modernizr.flexbox,nope : '`+i.InkLegacyCSS()+`'});`,
			),
			html.JsSrc(i.InkAllJS()),
			html.JsSrc(i.AutoloadJS()),
		).String(),
	)
}
