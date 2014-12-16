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
