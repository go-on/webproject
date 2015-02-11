package wsi

import (
	"testing"
)

func TestMapViaJSON(t *testing.T) {
	var x = struct {
		A string `json:"a"`
		B int
		C string `json:",omitempty"`
		D bool   `json:"-"`
	}{
		A: "a",
		B: 2,
		D: true,
	}

	m, err := MapViaJSON(&x)

	if err != nil {
		t.Errorf(err.Error())
	}

	if m["a"] != "a" {
		t.Errorf("wrong value for x.A: expected %#v, got %#v", "a", m["a"])
	}

	if m["B"] != 2.0 {
		t.Errorf("wrong value for x.B: expected %v, got %v", 2.0, m["B"])
	}

	if _, has := m["C"]; has {
		t.Errorf("x.C should be omitted, but is: %v", m["C"])
	}

	if _, has := m["D"]; has {
		t.Errorf("x.D should be omitted, but is: %v", m["D"])
	}

	if len(m) != 2 {
		t.Errorf("%#v should have length of 2", m)
	}
}
