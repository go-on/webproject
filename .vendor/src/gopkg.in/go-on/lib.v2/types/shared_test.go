package types

import (
	// "fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	tests := map[string]string{
		Id("id-main").String():                      "id-main",
		Class("class-main").String():                "class-main",
		Comment("comment a").String():               "<!-- comment a -->",
		Descr("descr").String():                     "descr",
		HTMLString("<p>hu</p>").String():            "<p>hu</p>",
		(Attribute{"target", "_blank"}).String():    `target="_blank"`,
		(Attribute{"", "disabled"}).String():        "disabled",
		(Style{"color", "red"}).String():            `color:red;`,
		(Style{"display", "none"}).CSS():            `display:none;`,
		(Style{"float", "left"}).Val("right").CSS(): `float:right;`,
		Tag("div").String():                         "div",
		Text("hello world").String():                "hello world",
		Id("id-main").Selector():                    "#id-main",
		Class("class-main").Selector():              ".class-main",
		Tag("div").Selector():                       "div",
		EscapeHTML(`&'<>"hiho&'<>"`):                "&amp;&#39;&lt;&gt;&#34;hiho&amp;&#39;&lt;&gt;&#34;",
	}

	for result, should := range tests {
		if result != should {
			t.Errorf("expected: %#v, got %#v", should, result)
		}
	}
}
