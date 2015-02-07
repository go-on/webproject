package types

/*
	This should become the central place for reused interfaces within go-on

	This package should have no external dependencies and be imported from libraries, such
	as gopherjs with minimal overhead.
*/

type Id string

func (id Id) String() string { return string(id) }

func (id Id) Selector() string { return "#" + string(id) }

type Descr string

func (d Descr) String() string { return string(d) }

type Comment string

func (c Comment) String() string {
	return "<!-- " + string(c) + " -->"
}

type HTMLer interface {
	HTML() string
}

type HTMLString string

func (h HTMLString) String() string { return string(h) }

func (h HTMLString) HTML() string {
	return string(h)
}

type JavaScripter interface {
	JavaScript() string
}

type JavaScriptString string

func (js JavaScriptString) String() string {
	return string(js)
}

func (js JavaScriptString) JavaScript() string {
	return string(js)
}

type CSSString string

func (cs CSSString) String() string {
	return string(cs)
}

func (cs CSSString) CSS() string {
	return string(cs)
}

type CSSSer interface {
	CSS() string
}

type Class string

func (c Class) String() string { return string(c) }

func (c Class) Selector() string { return "." + string(c) }

type Attribute struct {
	Key   string
	Value string
}

func (a Attribute) String() string {
	// a boolean attribute, e.g. disabled
	if a.Key == "" {
		return a.Value
	}
	return a.Key + `="` + EscapeHTML(a.Value) + `"`
}

type Styler interface {
	Style() string
}

type Style struct {
	Property, Value string
}

func (s Style) Val(v string) Style { s.Value = v; return s }
func (s Style) String() string     { return s.Property + ":" + s.Value + ";" }
func (s Style) CSS() string        { return s.String() }
func (s Style) Style() string      { return s.String() }

type Text string

func (t Text) String() string { return string(t) }

type Tag string

func (t Tag) String() string   { return string(t) }
func (t Tag) Selector() string { return t.String() }

type Selector interface {
	Selector() string
}

type SelectorString string

func (ø SelectorString) Selector() string { return string(ø) }

// combine several selectors to one
func Selectors(sel1 Selector, selects ...Selector) Selector {
	var s string
	for _, sel := range selects {
		s += sel.Selector()
	}
	return SelectorString(s)
}

/*
type Element interface {
	Tag
	HTMLString

	Id() Id
	SetId(Id) error

	Classes() []Class
	SetClass(classes ...Class)
	HasClass(class Class) bool
	AddClass(classes ...Class) error
	RemoveClass(class Class)

	Attributes() map[string]string
	SetAttribute(k, v string)
	RemoveAttribute(k string)

	Add(objects ...interface{}) (err error)
	MustAdd(objects ...interface{})

	WriteTo(w Writer) (n int64, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type TemplateSetter interface {
	WriteTo(w Writer) (n int64, err error)
	Name() string
	SetString() string
}


type HTMLPlaceholder interface {
	TemplateSetter
	Set(val interface{}) TemplateSetter
	Setf(format string, val ...interface{}) TemplateSetter
	String() string
	Type() interface{}
}
*/

type Route interface {
	Definition() string // was: Route() string
	HasParams() bool
	URL(params ...string) (string, error)
	URLMap(params map[string]string) (string, error)
	MustURL(params ...string)
	MustURLMap(params map[string]string)
}
