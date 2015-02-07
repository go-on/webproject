package meta

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSliceArg(t *testing.T) {
	fn := func(a ...interface{}) string {
		var b bytes.Buffer

		for _, aa := range a {
			fmt.Fprintf(&b, "%T-", aa)
		}

		return b.String()
	}

	sl := []int{2, 3, 4}

	r := fn(SliceArg(sl)...)
	exp := "int-int-int-"
	if r != exp {
		t.Errorf("SliceArg not working, expected: %s, got: %s\n", exp, r)
	}
}
