package logic

import (
	"testing"

	. "gopkg.in/go-on/lib.v2/html"
)

func TestTimes(t *testing.T) {
	type person struct{ firstname, lastname string }

	persons := []person{{"Donald", "Duck"}, {"Peter", "Pan"}}

	got := UL(
		TIMES_(len(persons),
			func(no int) interface{} { return LI(persons[no].firstname + " " + persons[no].lastname) },
		),
	).String()

	expected := `<ul><li>Donald Duck</li><li>Peter Pan</li></ul>`

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}

func TestIf(t *testing.T) {
	var switched bool

	on := func() bool { return switched }

	// InputCheckbox(name, ...)
	chbx := func() string { return InputCheckbox("chk", IF_(on, Checked_)).String() }

	switched = true
	got := chbx()
	expected := "<input type=\"checkbox\" name=\"chk\" checked=\"checked\" />"

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}

	switched = false
	got = chbx()
	expected = "<input type=\"checkbox\" name=\"chk\" />"

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}

func TestIfElse(t *testing.T) {
	var switched bool

	on := func() bool { return switched }

	// InputCheckbox(name, ...)
	chbx := func() string { return InputCheckbox("chk", IF_ELSE_(on, Checked_, Value_("off"))).String() }

	switched = true
	got := chbx()
	expected := "<input type=\"checkbox\" name=\"chk\" checked=\"checked\" />"

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}

	switched = false
	got = chbx()
	expected = "<input type=\"checkbox\" name=\"chk\" value=\"off\" />"

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}
