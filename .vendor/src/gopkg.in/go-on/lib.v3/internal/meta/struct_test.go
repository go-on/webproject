package meta

import (
	"fmt"
	"reflect"
	"testing"
)

type testStruct struct {
	A string `testtag:"x"`
	B string `testtag:"y"`
	S fmt.Stringer
}

type xx string

func (x xx) String() string {
	return string(x)
}

func TestToPtrSlice(t *testing.T) {
	x := testStruct{}

	stru, err := StructByValue(reflect.ValueOf(&x))

	if err != nil {
		t.Errorf(err.Error())
	}

	sl := stru.ToPtrSlice("testtag", []string{"S", "x"})

	if len(sl) != 2 {
		t.Errorf("len(sl) = %v // expected 2", len(sl))
	}

	*(sl[1].(*string)) = "a"

	if x.A != "a" {
		t.Errorf("wrong value for x.A: expected %#v, got %#v", "a", x.A)
	}

	*(sl[0].(*fmt.Stringer)) = xx("S")

	if x.S.String() != "S" {
		t.Errorf("wrong value for x.S.String(): expected %#v, got %#v", "S", x.S.String())
	}
}

func TestToPtrMap(t *testing.T) {
	x := testStruct{}

	stru, err := StructByValue(reflect.ValueOf(&x))

	if err != nil {
		t.Errorf(err.Error())
	}
	m := stru.ToPtrMap("testtag")

	*(m["x"].(*string)) = "a"

	if x.A != "a" {
		t.Errorf("wrong value for x.A: expected %#v, got %#v", "a", x.A)
	}

	*(m["S"].(*fmt.Stringer)) = xx("S")

	if x.S.String() != "S" {
		t.Errorf("wrong value for x.S.String(): expected %#v, got %#v", "S", x.S.String())
	}
}

func TestToMap(t *testing.T) {
	var x = struct {
		A string `testtag:"a"`
		B string `testtag:",omitempty"`
		C int
		D bool `testtag:"-"`
	}{
		A: "a",
		B: "",
		C: 5,
		D: true,
	}

	stru, err := StructByValue(reflect.ValueOf(&x))

	if err != nil {
		t.Errorf(err.Error())
	}

	m := stru.ToMap("testtag")

	if m["a"] != "a" {
		t.Errorf("wrong value for x.A: expected %#v, got %#v", "a", m["a"])
	}

	if m["C"] != 5 {
		t.Errorf("wrong value for x.C: expected %v, got %v", 5, m["C"])
	}

	if _, has := m["B"]; has {
		t.Errorf("x.B should be omitted, but is: %v", m["B"])
	}

	if _, has := m["D"]; has {
		t.Errorf("x.D should be omitted, but is: %v", m["D"])
	}

	if len(m) != 2 {
		t.Errorf("%#v should have length of 2", m)
	}
}

func TestEachTag(t *testing.T) {
	res := map[string]string{}

	fn := func(f *Field, tagVal string) {
		res[tagVal] = f.Value.String()
		f.Set(reflect.ValueOf(f.Value.String() + " changed"))
	}

	ts := testStruct{A: "valA", B: "valB"}
	vl := reflect.ValueOf(&ts)

	//s, err := StructByValue(&vl)
	s, err := StructByValue(vl)

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	s.EachTag("testtag", fn)

	resA := res["x"]
	resB := res["y"]

	if resA != "valA" {
		t.Errorf("did not set A")
	}

	if resB != "valB" {
		t.Errorf("did not set B")
	}

	if ts.B != "valB changed" {
		t.Errorf("could not change B")
	}

	f, _ := s.Field("A")
	err = f.Set(reflect.ValueOf("a changed"))
	if err != nil {
		t.Errorf("got error setting field: %s", err)
	}

	if ts.A != "a changed" {
		t.Errorf("could not change A")
	}

	f, _ = s.Field("S")
	var x fmt.Stringer
	x = xx("some text")
	err = f.Set(reflect.ValueOf(x))
	if err != nil {
		t.Errorf("got error setting field S: %s", err)
	}
	if ts.S != x {
		t.Errorf("could not change S")
	}

	s2, e2 := StructByType(reflect.TypeOf(ts))

	if e2 != nil {
		t.Errorf("can't create new val: %s", e2)
	}

	f, _ = s2.Field("A")
	f.Set(reflect.ValueOf("a set"))
	ts2 := s2.Value.Interface().(*testStruct)

	if ts2.A != "a set" {
		t.Errorf("can't build and set empty struct")
	}
}
