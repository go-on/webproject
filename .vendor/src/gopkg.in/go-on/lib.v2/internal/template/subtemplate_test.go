package template

import (
	"gopkg.in/go-on/lib.v2/internal/replacer"
	"testing"
)

var (
	fruit      = NewPlaceholder("fruit")
	li         = New("item").MustAdd("<li>", fruit, "</li>").Parse()
	ul         = New("list").MustAdd("<ul>", li, "</ul>").Parse()
	listOutput = "<ul><li>Apple</li><li>Pear</li></ul>"
)

func TestSubTemplate(t *testing.T) {
	_ = replacer.P
	// println(li.Replace(fruit.Set("Apple")).String())

	all := li.New()

	li.ReplaceTo(all.Buffer, fruit.Set("Apple"))
	li.ReplaceTo(all.Buffer, fruit.Set("Pear"))

	if r := ul.Replace(all).String(); r != listOutput {
		t.Errorf("Error in setting: expected\n\t%#v\ngot\n\t%#v\n", listOutput, r)
	}
}
