package html

import (
	"fmt"

	"gopkg.in/go-on/lib.v3/types"
)

const (
	Hidden   = "hidden"
	Checkbox = "checkbox"
)

func Rel_(relation string) types.Attribute  { return types.Attribute{"rel", relation} }
func Href_(url string) types.Attribute      { return types.Attribute{"href", url} }
func Type_(ty string) types.Attribute       { return types.Attribute{"type", ty} }
func Media_(md string) types.Attribute      { return types.Attribute{"media", md} }
func Src_(src string) types.Attribute       { return types.Attribute{"src", src} }
func Name_(n string) types.Attribute        { return types.Attribute{"name", n} }
func Content_(n string) types.Attribute     { return types.Attribute{"content", n} }
func Charset_(c string) types.Attribute     { return types.Attribute{"charset", c} }
func CharsetUtf8_() types.Attribute         { return types.Attribute{"charset", "utf-8"} }
func Lang_(l string) types.Attribute        { return types.Attribute{"lang", l} }
func Role_(l string) types.Attribute        { return types.Attribute{"role", l} }
func Value_(v string) types.Attribute       { return types.Attribute{"value", v} }
func Alt_(alttext string) types.Attribute   { return types.Attribute{"alt", alttext} }
func Title_(text string) types.Attribute    { return types.Attribute{"title", text} }
func Method_(m string) types.Attribute      { return types.Attribute{"method", m} }
func Action_(a string) types.Attribute      { return types.Attribute{"action", a} }
func For_(f string) types.Attribute         { return types.Attribute{"for", f} }
func Width_(f string) types.Attribute       { return types.Attribute{"width", f} }
func Height_(f string) types.Attribute      { return types.Attribute{"height", f} }
func OnSubmit_(js string) types.Attribute   { return types.Attribute{"onsubmit", "javascript:" + js} }
func OnClick_(js string) types.Attribute    { return types.Attribute{"onclick", "javascript:" + js} }
func Enctype_(f string) types.Attribute     { return types.Attribute{"enctype", f} }
func Target_(f string) types.Attribute      { return types.Attribute{"target", f} }
func DataToggle_(f string) types.Attribute  { return types.Attribute{"data-toggle", f} }
func DataTarget_(f string) types.Attribute  { return types.Attribute{"data-target", f} }
func DataId_(f string) types.Attribute      { return types.Attribute{"data-id", f} }
func Style_(f string) types.Attribute       { return types.Attribute{"style", f} }
func Placeholder_(f string) types.Attribute { return types.Attribute{"placeholder", f} }

// RDFa
func About_(a string) types.Attribute    { return types.Attribute{"about", a} }
func TypeOf_(a string) types.Attribute   { return types.Attribute{"typeof", a} }
func Property_(a string) types.Attribute { return types.Attribute{"property", a} }

// vars
var Checked_ = types.Attribute{"checked", "checked"}
var Selected_ = types.Attribute{"selected", "selected"}
var Disabled_ = types.Attribute{"disabled", "disabled"}
var Required_ = types.Attribute{"required", "required"}
var MultiPart_ = Enctype_("multipart/form-data")
var TargetBlank_ = Target_("_blank")

func Attrs_(pairs ...string) []types.Attribute {
	l := len(pairs)
	if l%2 != 0 {
		panic("len pairs must be even")
	}
	a := make([]types.Attribute, l/2)

	for i := 0; i < l; i += 2 {
		a[i/2] = types.Attribute{pairs[i], pairs[i+1]}
	}
	return a
}

func Classf_(format string, i ...interface{}) types.Class {
	return types.Class(fmt.Sprintf(format, i...))
}
