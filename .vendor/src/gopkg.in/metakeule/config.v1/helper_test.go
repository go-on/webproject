package config

import (
	"fmt"
	"testing"
)

func TestValidateName(t *testing.T) {

	tests := []struct {
		name string
		err  error
	}{
		{"ab", nil},
		{"a1", nil},
		{"aa", nil},
		{"", InvalidNameError("")},
		{"a", InvalidNameError("a")},
		{"01", InvalidNameError("01")},
		{"A", InvalidNameError("A")},
		{"aA", InvalidNameError("aA")},
		{"a_a", InvalidNameError("a_a")},
	}

	for _, test := range tests {

		if got, want := ValidateName(test.name), test.err; got != want {
			t.Errorf("ValidateName(%v) = %v; want %v", test.name, got, want)
		}
	}

}

func ExampleConfig() {
	app := MustNew("testapp", "1.2.3", "help text")
	verbose := app.NewBool("verbose", "show verbose messages", Required)
	// real application would use
	// err := app.Run()
	app.mergeArgs(false, []string{"--verbose"})
	fmt.Printf("verbose: %v", verbose.Get())
	// Output: verbose: true
}
