package placeholder

import (
	"testing"
)

/*
var Escaper = template.Escaper{
	"text":     handleStrings(html.EscapeString, true),
	"":         handleStrings(html.EscapeString, true),
	"html":     handleStrings(idem, true),
	"comment":  handleStrings(idem, true),
	"px":       units("%vpx"),
	"%":        units("%v%%"),
	"em":       units("%vem"),
	"pt":       units("%vpt"),
	"urlparam": handleStrings(url.QueryEscape, false),
}
*/

func TestEscaper(t *testing.T) {

	tests := map[string]string{
		Escaper["text"]("h<i>u</i>"):                 `h&lt;i&gt;u&lt;/i&gt;`,
		Escaper["html"]("h<i>u</i>"):                 `h<i>u</i>`,
		Escaper["comment"]("a comment"):              `a comment`,
		Escaper["px"](10):                            `10px`,
		Escaper["%"](11):                             `11%`,
		Escaper["em"](12):                            `12em`,
		Escaper["pt"](13):                            `13pt`,
		Escaper["urlparam"]("http://www.google.com"): `http%3A%2F%2Fwww.google.com`,
	}

	for got, expected := range tests {
		if got != expected {
			t.Errorf("got: %#v expected: %#v", got, expected)
		}
	}

}
