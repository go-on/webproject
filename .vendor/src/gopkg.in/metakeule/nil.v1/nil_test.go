package nil

import (
	"testing"
)

type s struct{}

func (ss *s) IsNil() bool { return IsNil(ss) }

func TestIsNil(t *testing.T) {
	var empty interface{}
	var nilInter interface{}
	var emptyStruct *s
	var nilStruct *s
	nilInter = nil
	nilStruct = nil
	cases := map[string]bool{
		"empty interface":         IsNil(empty),
		"nil interface":           IsNil(nilInter),
		"empty pointer to struct": IsNil(emptyStruct),
		"nil pointer to struct":   IsNil(nilStruct),
		"empty pointer method":    emptyStruct.IsNil(),
		"nil pointer method":      nilStruct.IsNil(),
	}

	for test, result := range cases {
		if !result {
			t.Errorf("%s should be nil, but is not", test)
		}
	}

}
