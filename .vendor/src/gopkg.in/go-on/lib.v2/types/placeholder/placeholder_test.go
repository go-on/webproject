package placeholder

import (
	. "gopkg.in/go-on/lib.v2/types"
	// "fmt"
	"testing"
)

func TestPlaceholders(t *testing.T) {
	tests := map[string]string{
		"html <value>":       New(HTMLString("html")).Set("html <value>").SetString(),
		"text &lt;value&gt;": New(Text("text")).Set("text <value>").SetString(),
		"comment <value>":    New(Comment("comment")).Set("comment <value>").SetString(),
		"descr <value>":      New(Descr("descr")).Set("descr <value>").SetString(),
		"class <value>":      New(Class("class")).Set("class <value>").SetString(),
		// "<nil>":              Class("class").Placeholder().Set(nil).SetString(),
		"id <value>":       New(Id("id")).Set("id <value>").SetString(),
		`href="&lt;#&gt;"`: New(Attribute{"href", "special link"}).Set("<#>").SetString(),
		`color:<red>;`:     New(Style{"color", "special color"}).Set("<red>").SetString(),
		``:                 New(Style{"color", "special color"}).Set(nil).SetString(),
		`height:45;`:       New(Style{"height", "special height"}).Set(45).SetString(),
		`height:fun;`:      New(Style{"height", "special height"}).Set(Id("fun")).SetString(),
		"div":              New(Tag("special")).Set("div").SetString(),
		"␣‣Comment.comment∎␣": New(Comment("comment")).String(),
		"<!-- comment -->":    New(Comment("comment")).Type().(Comment).String(),
		"␣‣Text.a string∎␣":   New("a string").String(),
		"␣‣Text.42∎␣":         New(42).String(),
	}

	for should, result := range tests {
		if result != should {
			t.Errorf("expected: %#v, got %#v", should, result)
		}
	}
}
