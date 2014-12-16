package match

import (
	"testing"

	. "gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/types"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %v, should be %v\n", is, shouldbe)
}

func TestPositionMatcher(t *testing.T) {
	e := NewElement("p")
	m := &PositionMatcher{Element: e}

	if m.Matches(NewElement("a")) || m.Found {
		t.Errorf("incorrect position matcher matches, got: %v, expected %s", true, false)
	}

	_ = m.Matches(NewElement("a"))

	if !m.Matches(e) || !m.Found {
		t.Errorf("incorrect position matcher matches, got: %v wanted: %v", false, true)
	}

	_ = m.Matches(NewElement("a"))

	if m.Pos != 2 {
		t.Errorf("incorrect position matcher position first round, got: %#v, wanted: %#v", m.Pos, 2)
	}
}

func TestFieldMatcher(t *testing.T) {
	var m FieldMatcher
	if !m.Matches(NewElement("input", FormField)) {
		err(t, "incorrect field matcher matches", false, true)
	}
	if m.Matches(NewElement("a")) {
		err(t, "incorrect field matcher matches", true, false)
	}
}

func TestClassMatcher(t *testing.T) {
	c := types.Class("fine")
	a1 := NewElement("a")
	a1.Add(c)
	if !New(c).Matches(a1) {
		err(t, "incorrect class matcher matches", false, true)
	}
	a2 := NewElement("a")
	if New(c).Matches(a2) {
		err(t, "incorrect class matcher matches", true, false)
	}
}

func TestNotMatcher(t *testing.T) {
	cl := types.Class("fine")
	c := Not(New(cl))
	a1 := NewElement("a")
	a1.Add(cl)
	if c.Matches(a1) {
		err(t, "incorrect not matcher matches", true, false)
	}
	a2 := NewElement("a")
	if !c.Matches(a2) {
		err(t, "incorrect not matcher matches", false, true)
	}
}

func TestIdMatcher(t *testing.T) {
	i := types.Id("fine")
	a1 := NewElement("a")
	a1.Add(i)
	if !New(i).Matches(a1) {
		err(t, "incorrect id matcher matches", false, true)
	}
	a2 := NewElement("a")
	if New(i).Matches(a2) {
		err(t, "incorrect id matcher matches", true, false)
	}
}

func TestOrMatcher(t *testing.T) {
	i := Or(New(types.Id("fine")), New(types.Class("well")))
	a1 := NewElement("a")
	a1.Add(types.Id("fine"))

	if !i.Matches(a1) {
		err(t, "incorrect or id matcher matches", false, true)
	}

	a2 := NewElement("a")
	a2.Add(types.Class("well"))
	if !i.Matches(a2) {
		err(t, "incorrect or class matcher matches", false, true)
	}
	a3 := NewElement("a")
	if i.Matches(a3) {
		err(t, "incorrect or matcher matches", true, false)
	}
}

func TestAndMatcher(t *testing.T) {
	i := And(New(types.Id("fine")), New(types.Class("well")))
	a1 := NewElement("a")
	a1.Add(types.Id("fine"), types.Class("well"))

	if !i.Matches(a1) {
		err(t, "incorrect and matcher matches", false, true)
	}
	a2 := NewElement("a")
	a2.Add(types.Class("well"))

	if i.Matches(a2) {
		err(t, "incorrect and  class matcher matches", true, false)
	}
	a3 := NewElement("a")
	a3.Add(types.Id("fine"))

	if i.Matches(a3) {
		err(t, "incorrect and id matcher matches", true, false)
	}
}

func TestTagMatcher(t *testing.T) {
	m := types.Tag("a")
	if !New(m).Matches(NewElement("a")) {
		err(t, "incorrect tag matcher matches", false, true)
	}
	if New(m).Matches(NewElement("b")) {
		err(t, "incorrect tag matcher matches", true, false)
	}
}

func TestAttrMatcher(t *testing.T) {
	m := types.Attribute{"width", "200"}
	div1 := NewElement("div")
	div1.Add(types.Attribute{"width", "200"})
	if !New(m).Matches(div1) {
		err(t, "incorrect Attr matcher matches", false, true)
	}
	div2 := NewElement("div")
	if New(m).Matches(div2) {
		err(t, "incorrect Attr matcher matches", true, false)
	}
}

/*
func TestAttrsMatcher(t *testing.T) {
	m := Attrs("width", "200", "height", "300")
	div1 := NewElement("div")
	div1.Add(Attrs("width", "200", "height", "300"))
	if !m.Matches(div1) {
		err(t, "incorrect Attrs matcher matches", false, true)
	}
	div2 := NewElement("div")
	div2.Add(Attr("width", "200"))
	if m.Matches(div2) {
		err(t, "incorrect Attrs matcher matches", true, false)
	}
}
*/

/*
func TestStyleMatcher(t *testing.T) {
	m := Style("width", "200")
	if !m.Matches(Div(Style("width", "200"))) {
		err(t, "incorrect Style matcher matches", false, true)
	}
	if m.Matches(Div()) {
		err(t, "incorrect Style matcher matches", true, false)
	}
}

func TestStylesMatcher(t *testing.T) {
	m := Style("width", "200", "height", "300")
	if !m.Matches(Div(Style("width", "200", "height", "300"))) {
		err(t, "incorrect Styles matcher matches", false, true)
	}
	if m.Matches(Div(Style("width", "200"))) {
		err(t, "incorrect Styles matcher matches", true, false)
	}
}
*/
func TestHtmlMatcher(t *testing.T) {
	a := NewElement("a")
	a.Add("hiho")
	m := types.HTMLString(a.String())
	div := NewElement("div")
	div.Add(a)
	if !New(m).Matches(div) {
		err(t, "incorrect Html matcher matches", false, true)
	}

	a2 := NewElement("a")
	a2.Add("hoho")
	div2 := NewElement("div")
	div2.Add(a2)
	if New(m).Matches(div2) {
		err(t, "incorrect Html matcher matches", true, false)
	}
}
