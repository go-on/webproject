package method

import (
	"reflect"
	"testing"
)

func TestSomething(t *testing.T) {
	ty := reflect.TypeOf(GET).Name()

	if ty != "Method" {
		t.Errorf("type of GET is not Method, but %#v", ty)
	}
}
