package meta

import (
	"reflect"
	"testing"
)

func TestPtrAssoc(t *testing.T) {
	s := "string"

	ptr, err := PointerByValue(reflect.ValueOf(&s))
	if err != nil {
		t.Errorf("can't create pointer: %s", err)
	}

	var target string
	tgt := &target
	err = ptr.Assoc(reflect.ValueOf(&tgt))

	if err != nil {
		t.Errorf("can't assoc: %s", err)
	}

	*tgt = "huho"

	if s != "huho" {
		t.Errorf("association did not work")
	}
}
