package html

import (
	"fmt"

	. "gopkg.in/go-on/lib.v3/html/internal/element"
)

func A(objects ...interface{}) (t *Element) {
	t = NewElement("a", Inline)
	t.Add(objects...)
	return
}

func ABBR(objects ...interface{}) (t *Element) {
	t = NewElement("abbr")
	t.Add(objects...)
	return
}

func ACRONYM(objects ...interface{}) (t *Element) {
	t = NewElement("acronym")
	t.Add(objects...)
	return
}

func ADDRESS(objects ...interface{}) (t *Element) {
	t = NewElement("address")
	t.Add(objects...)
	return
}

func AREA(objects ...interface{}) (t *Element) {
	t = NewElement("area", SelfClosing)
	t.Add(objects...)
	return
}

func ARTICLE(objects ...interface{}) (t *Element) {
	t = NewElement("article")
	t.Add(objects...)
	return
}

func ASIDE(objects ...interface{}) (t *Element) {
	t = NewElement("aside")
	t.Add(objects...)
	return
}

func AUDIO(objects ...interface{}) (t *Element) {
	t = NewElement("audio")
	t.Add(objects...)
	return
}

func B(objects ...interface{}) (t *Element) {
	t = NewElement("b", Inline)
	t.Add(objects...)
	return
}

func BASE(objects ...interface{}) (t *Element) {
	t = NewElement("base", SelfClosing)
	t.Add(objects...)
	return
}

func BDO(objects ...interface{}) (t *Element) {
	t = NewElement("bdo")
	t.Add(objects...)
	return
}

func BDI(objects ...interface{}) (t *Element) {
	t = NewElement("bdi", Inline)
	t.Add(objects...)
	return
}

func BIG(objects ...interface{}) (t *Element) {
	t = NewElement("big")
	t.Add(objects...)
	return
}

func BLOCKQUOTE(objects ...interface{}) (t *Element) {
	t = NewElement("blockquote")
	t.Add(objects...)
	return
}

func BODY(objects ...interface{}) (t *Element) {
	t = NewElement("body")
	// t.ParentTags = []Tag{"html"}
	t.Add(objects...)
	return
}

func BR(objects ...interface{}) (t *Element) {
	t = NewElement("br", SelfClosing)
	t.Add(objects...)
	return
}

// FormField ??
func BUTTON(objects ...interface{}) (t *Element) {
	t = NewElement("button", FormField)
	t.Add(objects...)
	return
}

func CANVAS(objects ...interface{}) (t *Element) {
	t = NewElement("canvas")
	t.Add(objects...)
	return
}

func CAPTION(objects ...interface{}) (t *Element) {
	t = NewElement("caption")
	t.Add(objects...)
	return
}

func CITE(objects ...interface{}) (t *Element) {
	t = NewElement("cite")
	t.Add(objects...)
	return
}

func CODE(objects ...interface{}) (t *Element) {
	t = NewElement("code")
	t.Add(objects...)
	return
}

func COL(objects ...interface{}) (t *Element) {
	t = NewElement("col", SelfClosing)
	t.Add(objects...)
	return
}

func COLGROUP(objects ...interface{}) (t *Element) {
	t = NewElement("colgroup")
	t.Add(objects...)
	return
}

func COMMAND(objects ...interface{}) (t *Element) {
	t = NewElement("command", SelfClosing|Inline)
	t.Add(objects...)
	return
}

func DATA(objects ...interface{}) (t *Element) {
	t = NewElement("data", Inline)
	t.Add(objects...)
	return
}

func DATALIST(objects ...interface{}) (t *Element) {
	t = NewElement("datalist", Inline)
	t.Add(objects...)
	return
}

func DD(objects ...interface{}) (t *Element) {
	t = NewElement("dd")
	t.Add(objects...)
	return
}

func DEL(objects ...interface{}) (t *Element) {
	t = NewElement("del")
	t.Add(objects...)
	return
}

func DETAILS(objects ...interface{}) (t *Element) {
	t = NewElement("details")
	t.Add(objects...)
	return
}

func DFN(objects ...interface{}) (t *Element) {
	t = NewElement("dfn")
	t.Add(objects...)
	return
}

func DIV(objects ...interface{}) (t *Element) {
	t = NewElement("div")
	t.Add(objects...)
	return
}

func DIALOG(objects ...interface{}) (t *Element) {
	t = NewElement("dialog")
	t.Add(objects...)
	return
}

func DL(objects ...interface{}) (t *Element) {
	t = NewElement("dl")
	t.Add(objects...)
	return
}

func DT(objects ...interface{}) (t *Element) {
	t = NewElement("dt")
	t.Add(objects...)
	return
}

func EM(objects ...interface{}) (t *Element) {
	t = NewElement("em", Inline)
	t.Add(objects...)
	return
}

func EMBED(objects ...interface{}) (t *Element) {
	t = NewElement("embed", SelfClosing)
	t.Add(objects...)
	return
}

func FIELDSET(objects ...interface{}) (t *Element) {
	t = NewElement("fieldset")
	t.Add(objects...)
	return
}

func FIGCAPTION(objects ...interface{}) (t *Element) {
	t = NewElement("figcaption", Inline)
	t.Add(objects...)
	return
}

func FIGURE(objects ...interface{}) (t *Element) {
	t = NewElement("figure")
	t.Add(objects...)
	return
}

func FOOTER(objects ...interface{}) (t *Element) {
	t = NewElement("footer")
	t.Add(objects...)
	return
}

func FORM(objects ...interface{}) (t *Element) {
	t = NewElement("form")
	//t.MustAdd(Attr("enctype", "multipart/form-data", "method", "post"))
	t.Add(objects...)
	return
}

func FRAME(objects ...interface{}) (t *Element) {
	t = NewElement("frame", SelfClosing)
	t.Add(objects...)
	return
}

func FRAMESET(objects ...interface{}) (t *Element) {
	t = NewElement("frameset")
	t.Add(objects...)
	return
}

// make a heading of level n (h1, h2, ...)
func H(n int, objects ...interface{}) (t *Element) {
	t = NewElement(fmt.Sprintf("h%v", n))
	t.Add(objects...)
	return
}

func H1(objects ...interface{}) (t *Element) {
	t = NewElement("h1")
	t.Add(objects...)
	return
}

func H2(objects ...interface{}) (t *Element) {
	t = NewElement("h2")
	t.Add(objects...)
	return
}

func H3(objects ...interface{}) (t *Element) {
	t = NewElement("h3")
	t.Add(objects...)
	return
}

func H4(objects ...interface{}) (t *Element) {
	t = NewElement("h4")
	t.Add(objects...)
	return
}

func H5(objects ...interface{}) (t *Element) {
	t = NewElement("h5")
	t.Add(objects...)
	return
}

func H6(objects ...interface{}) (t *Element) {
	t = NewElement("h6")
	t.Add(objects...)
	return
}

func HEAD(objects ...interface{}) (t *Element) {
	t = NewElement("head", IdForbidden|ClassForbidden|Invisible)
	// t.ParentTags = []Tag{"html"}
	t.Add(objects...)
	return
}

func HEADER(objects ...interface{}) (t *Element) {
	t = NewElement("header")
	t.Add(objects...)
	return
}

func HGROUP(objects ...interface{}) (t *Element) {
	t = NewElement("hgroup")
	t.Add(objects...)
	return
}

func HR(objects ...interface{}) (t *Element) {
	t = NewElement("hr", SelfClosing)
	t.Add(objects...)
	return
}

func HTML(objects ...interface{}) (t *Element) {
	t = NewElement("html")
	// t.ParentTags = []Tag{"doc"}
	t.Add(objects...)
	return
}

func I(objects ...interface{}) (t *Element) {
	t = NewElement("i", Inline)
	t.Add(objects...)
	return
}

func IFRAME(objects ...interface{}) (t *Element) {
	t = NewElement("iframe", SelfClosing, Invisible)
	t.Add(objects...)
	return
}

func IMG(objects ...interface{}) (t *Element) {
	t = NewElement("img", Inline, SelfClosing)
	t.Add(objects...)
	return
}

func INPUT(objects ...interface{}) (t *Element) {
	t = NewElement("input", FormField, SelfClosing, Inline)
	t.Add(objects...)
	return
}

func INS(objects ...interface{}) (t *Element) {
	t = NewElement("ins")
	t.Add(objects...)
	return
}

func KBD(objects ...interface{}) (t *Element) {
	t = NewElement("kbd")
	t.Add(objects...)
	return
}

// FormField?
func KEYGEN(objects ...interface{}) (t *Element) {
	t = NewElement("keygen", SelfClosing|Inline|FormField)
	t.Add(objects...)
	return
}

func LABEL(objects ...interface{}) (t *Element) {
	t = NewElement("label")
	t.Add(objects...)
	return
}

func LEGEND(objects ...interface{}) (t *Element) {
	t = NewElement("legend")
	t.Add(objects...)
	return
}

func LI(objects ...interface{}) (t *Element) {
	t = NewElement("li")
	// t.ParentTags = []Tag{"ul", "ol"}
	t.Add(objects...)
	return
}

func LINK(objects ...interface{}) (t *Element) {
	t = NewElement("link", SelfClosing, IdForbidden, ClassForbidden, Invisible)
	// t.ParentTags = []Tag{"head", "body"}
	t.Add(objects...)
	return
}

func MAP(objects ...interface{}) (t *Element) {
	t = NewElement("map")
	t.Add(objects...)
	return
}

func MATH(objects ...interface{}) (t *Element) {
	t = NewElement("math", Inline)
	t.Add(objects...)
	return
}

func MENU(objects ...interface{}) (t *Element) {
	t = NewElement("menu")
	t.Add(objects...)
	return
}

func META(objects ...interface{}) (t *Element) {
	t = NewElement("meta", SelfClosing, IdForbidden, ClassForbidden, Invisible)
	// t.ParentTags = []Tag{"head"}
	t.Add(objects...)
	return
}

func METER(objects ...interface{}) (t *Element) {
	t = NewElement("meter", Inline)
	t.Add(objects...)
	return
}

func NAV(objects ...interface{}) (t *Element) {
	t = NewElement("nav")
	t.Add(objects...)
	return
}

func NOFRAMES(objects ...interface{}) (t *Element) {
	t = NewElement("noframes")
	t.Add(objects...)
	return
}

func NOSCRIPT(objects ...interface{}) (t *Element) {
	t = NewElement("noscript")
	t.Add(objects...)
	return
}

func OBJECT(objects ...interface{}) (t *Element) {
	t = NewElement("object", FormField)
	t.Add(objects...)
	return
}

func OL(objects ...interface{}) (t *Element) {
	t = NewElement("ol")
	t.Add(objects...)
	return
}

func OPTGROUP(objects ...interface{}) (t *Element) {
	t = NewElement("optgroup")
	t.Add(objects...)
	return
}

func OPTION(objects ...interface{}) (t *Element) {
	t = NewElement("option")
	// t.ParentTags = []Tag{"select"}
	t.Add(objects...)
	return
}

func OUTPUT(objects ...interface{}) (t *Element) {
	t = NewElement("output", Inline)
	t.Add(objects...)
	return
}

func P(objects ...interface{}) (t *Element) {
	t = NewElement("p")
	t.Add(objects...)
	return
}

func PARAM(objects ...interface{}) (t *Element) {
	t = NewElement("param", SelfClosing)
	t.Add(objects...)
	return
}

func PRE(objects ...interface{}) (t *Element) {
	t = NewElement("pre")
	t.Add(objects...)
	return
}

func PROGRESS(objects ...interface{}) (t *Element) {
	t = NewElement("progress", Inline)
	t.Add(objects...)
	return
}

func Q(objects ...interface{}) (t *Element) {
	t = NewElement("q")
	t.Add(objects...)
	return
}

func RP(objects ...interface{}) (t *Element) {
	t = NewElement("rp", Inline)
	t.Add(objects...)
	return
}

func RT(objects ...interface{}) (t *Element) {
	t = NewElement("rt", Inline)
	t.Add(objects...)
	return
}

func RUBY(objects ...interface{}) (t *Element) {
	t = NewElement("ruby", Inline)
	t.Add(objects...)
	return
}

func S(objects ...interface{}) (t *Element) {
	t = NewElement("s")
	t.Add(objects...)
	return
}

func SAMP(objects ...interface{}) (t *Element) {
	t = NewElement("samp")
	t.Add(objects...)
	return
}

func SCRIPT(objects ...interface{}) (t *Element) {
	//t = NewElement(Tag("script"), Invisible, IdForbidden, ClassForbidden, WithoutEscaping, JavascriptSpecialEscaping)
	t = NewElement("script", Invisible, ClassForbidden, WithoutEscaping, JavascriptSpecialEscaping)
	// t.ParentTags = []Tag{"head", "body"}
	t.Add(objects...)
	return
}

func SECTION(objects ...interface{}) (t *Element) {
	t = NewElement("section")
	t.Add(objects...)
	return
}

func MAIN(objects ...interface{}) (t *Element) {
	t = NewElement("main")
	t.Add(objects...)
	return
}

func SELECT(objects ...interface{}) (t *Element) {
	t = NewElement("select", FormField, Inline)
	t.Add(objects...)
	return
}

func SMALL(objects ...interface{}) (t *Element) {
	t = NewElement("small")
	t.Add(objects...)
	return
}

func SOURCE(objects ...interface{}) (t *Element) {
	t = NewElement("source")
	t.Add(objects...)
	return
}

func SPAN(objects ...interface{}) (t *Element) {
	t = NewElement("span", Inline)
	t.Add(objects...)
	return
}

func STRONG(objects ...interface{}) (t *Element) {
	t = NewElement("strong", Inline)
	t.Add(objects...)
	return
}

func STYLE(objects ...interface{}) (t *Element) {
	t = NewElement("style", IdForbidden, ClassForbidden, Invisible, WithoutEscaping)
	// t.ParentTags = []Tag{"head", "body"}
	t.Add(objects...)
	return
}

func SUB(objects ...interface{}) (t *Element) {
	t = NewElement("sub")
	t.Add(objects...)
	return
}

func SUMMARY(objects ...interface{}) (t *Element) {
	t = NewElement("summary", Inline)
	t.Add(objects...)
	return
}

func SUP(objects ...interface{}) (t *Element) {
	t = NewElement("sup")
	t.Add(objects...)
	return
}

func SVG(objects ...interface{}) (t *Element) {
	t = NewElement("svg")
	t.Add(objects...)
	return
}

func TABLE(objects ...interface{}) (t *Element) {
	t = NewElement("table")
	t.Add(objects...)
	return
}

func TBODY(objects ...interface{}) (t *Element) {
	t = NewElement("tbody")
	// t.ParentTags = []Tag{"table"}
	t.Add(objects...)
	return
}

func TD(objects ...interface{}) (t *Element) {
	t = NewElement("td")
	// t.ParentTags = []Tag{"tr"}
	t.Add(objects...)
	return
}

func TEXTAREA(objects ...interface{}) (t *Element) {
	t = NewElement("textarea", FormField)
	t.Add(objects...)
	return
}

func TFOOT(objects ...interface{}) (t *Element) {
	t = NewElement("tfoot")
	t.Add(objects...)
	return
}

func TH(objects ...interface{}) (t *Element) {
	t = NewElement("th")
	// t.ParentTags = []Tag{"tr"}
	t.Add(objects...)
	return
}

func THEAD(objects ...interface{}) (t *Element) {
	t = NewElement("thead")
	// t.ParentTags = []Tag{"table"}
	t.Add(objects...)
	return
}

func TIME(objects ...interface{}) (t *Element) {
	t = NewElement("time", Inline)
	t.Add(objects...)
	return
}

func TITLE(objects ...interface{}) (t *Element) {
	t = NewElement("title", Invisible)
	// t.ParentTags = []Tag{"head"}
	t.Add(objects...)
	return
}

func TR(objects ...interface{}) (t *Element) {
	t = NewElement("tr")
	// t.ParentTags = []Tag{"tbody", "table", "thead"}
	t.Add(objects...)
	return
}

func TRACK(objects ...interface{}) (t *Element) {
	t = NewElement("track", SelfClosing)
	// t.ParentTags = []Tag{"audio", "video"}
	t.Add(objects...)
	return
}

func TT(objects ...interface{}) (t *Element) {
	t = NewElement("tt")
	t.Add(objects...)
	return
}

func U(objects ...interface{}) (t *Element) {
	t = NewElement("u")
	t.Add(objects...)
	return
}

func UL(objects ...interface{}) (t *Element) {
	t = NewElement("ul")
	t.Add(objects...)
	return
}

func VAR(objects ...interface{}) (t *Element) {
	t = NewElement("var")
	t.Add(objects...)
	return
}

func VIDEO(objects ...interface{}) (t *Element) {
	t = NewElement("video")
	t.Add(objects...)
	return
}

func WBR(objects ...interface{}) (t *Element) {
	t = NewElement("wbr", SelfClosing, Inline)
	t.Add(objects...)
	return
}
