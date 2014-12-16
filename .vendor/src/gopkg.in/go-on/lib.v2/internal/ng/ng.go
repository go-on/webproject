package ng

import (
	"bytes"
	"fmt"

	// "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/types"
)

//	This package contains structs based on angualar js

//func (a attr) ClassString() string { return fmt.Sprintf(`%s: %s;`, a.key, a.value) }

/*
type Module struct {
	Name         string
	Dependencies []string
}
*/

//type E string

/*
func (e E) String() string { return string(e) }

type Expression interface {
	String() string
}
*/

var ANGULAR_VERSION = "1.2.2"

//type Template string // any string which can contain {{}} markup.

// take module
func App(module string) types.Attribute { return types.Attribute{`ng-app`, module} }

// take expression
func Bind(expr string) types.Attribute       { return types.Attribute{`ng-bind`, expr} }
func BindHtml(expr string) types.Attribute   { return types.Attribute{`ng-bind-html`, expr} }
func Blur(expr string) types.Attribute       { return types.Attribute{`ng-blur`, expr} }
func Change(expr string) types.Attribute     { return types.Attribute{`ng-change`, expr} }
func Checked(expr string) types.Attribute    { return types.Attribute{`ng-checked`, expr} }
func Class(expr string) types.Attribute      { return types.Attribute{`ng-class`, expr} }
func ClassEven(expr string) types.Attribute  { return types.Attribute{`ng-class-even`, expr} }
func ClassOdd(expr string) types.Attribute   { return types.Attribute{`ng-class-odd`, expr} }
func Click(expr string) types.Attribute      { return types.Attribute{`ng-click`, expr} }
func Controller(expr string) types.Attribute { return types.Attribute{`ng-controller`, expr} }
func Copy(expr string) types.Attribute       { return types.Attribute{`ng-copy`, expr} }
func Csp(expr string) types.Attribute        { return types.Attribute{`ng-csp`, expr} }
func Cut(expr string) types.Attribute        { return types.Attribute{`ng-cut`, expr} }
func Dblclick(expr string) types.Attribute   { return types.Attribute{`ng-dblclick`, expr} }
func Disabled(expr string) types.Attribute   { return types.Attribute{`ng-disabled`, expr} }
func Focus(expr string) types.Attribute      { return types.Attribute{`ng-focus`, expr} }
func Hide(expr string) types.Attribute       { return types.Attribute{`ng-hide`, expr} }
func If(expr string) types.Attribute         { return types.Attribute{`ng-if`, expr} }
func Init(expr string) types.Attribute       { return types.Attribute{`ng-init`, expr} }
func Keydown(expr string) types.Attribute    { return types.Attribute{`ng-keydown`, expr} }
func Keypress(expr string) types.Attribute   { return types.Attribute{`ng-keypress`, expr} }
func Keyup(expr string) types.Attribute      { return types.Attribute{`ng-keyup`, expr} }
func Mousedown(expr string) types.Attribute  { return types.Attribute{`ng-mousedown`, expr} }
func Mouseenter(expr string) types.Attribute { return types.Attribute{`ng-mouseenter`, expr} }
func Mouseleave(expr string) types.Attribute { return types.Attribute{`ng-mouseleave`, expr} }
func Mousemove(expr string) types.Attribute  { return types.Attribute{`ng-mousemove`, expr} }
func Mouseover(expr string) types.Attribute  { return types.Attribute{`ng-mouseover`, expr} }
func Mouseup(expr string) types.Attribute    { return types.Attribute{`ng-mouseup`, expr} }
func Open(expr string) types.Attribute       { return types.Attribute{`ng-open`, expr} }
func Paste(expr string) types.Attribute      { return types.Attribute{`ng-paste`, expr} }
func Readonly(expr string) types.Attribute   { return types.Attribute{`ng-readonly`, expr} }
func Selected(expr string) types.Attribute   { return types.Attribute{`ng-selected`, expr} }
func Show(expr string) types.Attribute       { return types.Attribute{`ng-show`, expr} }
func Style(expr string) types.Attribute      { return types.Attribute{`ng-style`, expr} }
func Submit(expr string) types.Attribute     { return types.Attribute{`ng-submit`, expr} }
func Switch(expr string) types.Attribute     { return types.Attribute{`ng-switch`, expr} }

// take string
func BindTemplate(s string) types.Attribute { return types.Attribute{`ng-bind-template`, s} }
func Form(s string) types.Attribute         { return types.Attribute{`ng-form`, s} }
func Model(s string) types.Attribute        { return types.Attribute{`ng-model`, s} }
func Include(s string) types.Attribute      { return types.Attribute{`ng-include`, s} }
func List(s string) types.Attribute         { return types.Attribute{`ng-list`, s} }
func Value(s string) types.Attribute        { return types.Attribute{`ng-value`, s} }
func SwitchWhen(s string) types.Attribute   { return types.Attribute{`ng-switch-when`, s} }

// take template
func Href(template string) types.Attribute   { return types.Attribute{`ng-href`, template} }
func Src(template string) types.Attribute    { return types.Attribute{`ng-src`, template} }
func SrcSet(template string) types.Attribute { return types.Attribute{`ng-srcset`, template} }

// take nothing
func Cloak() types.Attribute       { return types.Attribute{`ng-cloak`, `ng-cloak`} }
func NonBindable() types.Attribute { return types.Attribute{`ng-non-bindable`, `ng-non-bindable`} }
func Transclude() types.Attribute  { return types.Attribute{`ng-transclude`, `ng-transclude`} }
func RepeatEnd() types.Attribute   { return types.Attribute{`ng-repeat-end`, `ng-repeat-end`} }
func SwitchDefault() types.Attribute {
	return types.Attribute{`ng-switch-default`, `ng-switch-default`}
}

type repeat struct{ name, tempVar, collection string }

func (r repeat) Attr() types.Attribute {
	return types.Attribute{r.name, fmt.Sprintf("%s in %s", r.tempVar, r.collection)}
}

func Repeat(tempVar string, collection string) types.Attribute {
	r := repeat{`ng-repeat`, tempVar, collection}
	return r.Attr()
}

func RepeatStart(tempVar string, collection string) types.Attribute {
	r := repeat{`ng-repeat-start`, tempVar, collection}
	return r.Attr()
}

type repeatKeyVal struct{ name, tempKey, tempVal, collection string }

func (r repeatKeyVal) Attr() types.Attribute {
	return types.Attribute{r.name, fmt.Sprintf("(%s, %s) in %s", r.tempKey, r.tempVal, r.collection)}
}

func RepeatKeyVal(tempKey string, tempVal string, collection string) types.Attribute {
	r := repeatKeyVal{`ng-repeat`, tempKey, tempVal, collection}
	return r.Attr()
}

func RepeatStartKeyVal(tempKey string, tempVal string, collection string) types.Attribute {
	r := repeatKeyVal{`ng-repeat-start`, tempKey, tempVal, collection}
	return r.Attr()
}

type repeatTrackBy struct {
	name       string
	tempVar    string
	collection string
	trackBy    string
}

func (r repeatTrackBy) Attr() types.Attribute {
	return types.Attribute{`ng-repeat`, fmt.Sprintf("%s in %s track by %s", r.tempVar, r.collection, r.trackBy)}
}

func RepeatTrackBy(tempVar, collection, trackBy string) types.Attribute {
	r := repeatTrackBy{`ng-repeat`, tempVar, collection, trackBy}
	return r.Attr()
}

func RepeatStartTrackBy(tempVar, collection, trackBy string) types.Attribute {
	r := repeatTrackBy{`ng-repeat-start`, tempVar, collection, trackBy}
	return r.Attr()
}

func Script(id string, elem *element.Element) types.HTMLString {
	return types.HTMLString(fmt.Sprintf(`<script type="text/ng-template" id="%s">%s</script>`, id, elem.String()))
}

type OptionsArray struct {
	Array     string //     Array / object: an expression which evaluates to an Array / object to iterate over.
	Value     string //     Value: local variable which will refer to each item in the Array or each property Value of object during iteration.
	Label     string //     Label: The result of this expression will be the Label for <option> element. The expression will most likely refer to the Value variable (e.g. Value.propertyName).
	Select    string //     select: The result of this expression will be bound to the model of the parent <select> element. If not specified, select expression will default to Value.
	Group     string //     group: The result of this expression will be used to group options using the <optgroup> DOM element.
	TrackExpr string //     trackexpr: Used when working with an Array of objects. The result of this expression will be used to identify the objects in the Array. The trackexpr will most likely refer to the Value variable (e.g. Value.propertyName).
}

func (s OptionsArray) String() string {
	var selectstr, groupstr, trackstr string
	if s.Select != "" {
		selectstr = fmt.Sprintf("%s as ", s.Select)
	}
	if s.Group != "" {
		groupstr = fmt.Sprintf(" group by %s", s.Group)
	}

	if s.TrackExpr != "" {
		trackstr = fmt.Sprintf(" track by %s", s.TrackExpr)
	}
	return fmt.Sprintf("%s%s%s for %s in %s%s", selectstr, s.Label, groupstr, s.Value, s.Array, trackstr)
}

type OptionsObject struct {
	Object string //     Array / object: an expression which evaluates to an Array / object to iterate over.
	Value  string //     Value: local variable which will refer to each item in the Array or each property Value of object during iteration.
	Key    string //     key: local variable which will refer to a property name in object during iteration.
	Label  string //     Label: The result of this expression will be the Label for <option> element. The expression will most likely refer to the Value variable (e.g. Value.propertyName).
	Select string //     select: The result of this expression will be bound to the model of the parent <select> element. If not specified, select expression will default to Value.
	Group  string //     group: The result of this expression will be used to group options using the <optgroup> DOM element.
}

func (s OptionsObject) String() string {
	var selectstr, groupstr string
	if s.Select != "" {
		selectstr = fmt.Sprintf("%s as ", s.Select)
	}
	if s.Group != "" {
		groupstr = fmt.Sprintf(" group by %s", s.Group)
	}
	return fmt.Sprintf("%s%s%s for (%s, %s) in %s", selectstr, s.Label, groupstr, s.Key, s.Value, s.Object)
}

type Decorator struct {
	Model         string         // ng-model="{string}"
	Required      string         // ng-required
	MinLength     int            // ng-minlength
	MaxLength     int            // ng-maxlength
	Pattern       string         // ng-pattern
	Change        string         // ng-change
	True          string         // ng-true-value
	False         string         // ng-false-value
	NoTrim        bool           // ng-trim="{boolean}"
	OptionsArray  *OptionsArray  // ng-options="{comprehension_expression}"
	OptionsObject *OptionsObject // ng-options="{comprehension_expression}"
}

func (d Decorator) AttrString() string {
	attrs := map[string]string{}
	if d.Model != "" {
		attrs["ng-model"] = d.Model
	}
	if d.Required != "" {
		attrs["ng-required"] = d.Required
	}
	if d.MinLength != 0 {
		attrs["ng-minlength"] = fmt.Sprintf("%d", d.MinLength)
	}
	if d.MaxLength != 0 {
		attrs["ng-maxlength"] = fmt.Sprintf("%d", d.MaxLength)
	}
	if d.Pattern != "" {
		attrs["ng-pattern"] = d.Pattern
	}
	if d.Change != "" {
		attrs["ng-change"] = d.Change
	}
	if d.True != "" {
		attrs["ng-true-value"] = d.True
	}
	if d.False != "" {
		attrs["ng-false-value"] = d.False
	}
	if d.NoTrim {
		attrs["ng-trim"] = "false"
	}
	if d.OptionsArray != nil {
		attrs["ng-options"] = d.OptionsArray.String()
	}
	if d.OptionsObject != nil {
		attrs["ng-options"] = d.OptionsObject.String()
	}

	var buffer bytes.Buffer
	for k, v := range attrs {
		buffer.WriteString(fmt.Sprintf(` %s="%s"`, k, v))
	}
	return buffer.String()
}

type Pluralize struct {
	Count  string
	When   string
	Offset int
}

func (p Pluralize) Attrs() []types.Attribute {
	attrs := []types.Attribute{
		types.Attribute{"ng-pluralize", "ng-pluralize"},
		types.Attribute{"count", p.Count},
		types.Attribute{"when", p.When},
	}
	if p.Offset != 0 {
		attrs = append(attrs, types.Attribute{"offset", fmt.Sprintf("%d", p.Offset)})
	}
	return attrs
}
