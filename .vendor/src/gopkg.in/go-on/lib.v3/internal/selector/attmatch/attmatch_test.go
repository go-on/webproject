package attmatch

import (
	"testing"
)

func TestAll(t *testing.T) {
	tests := map[string]string{
		Exists("target").Selector():                  `[target]`,
		Equals("target", "_blank").Selector():        `[target="_blank"]`,
		Includes("target", "blank").Selector():       `[target~="blank"]`,
		Contains("target", "ank").Selector():         `[target*="ank"]`,
		BeginsWith("target", "_").Selector():         `[target^="_"]`,
		EndsWith("target", "nk").Selector():          `[target$="nk"]`,
		HyphenBeginsWith("target", "_bl").Selector(): `[target|="_bl"]`,
	}

	for real, expected := range tests {
		if real != expected {
			t.Errorf("got %#v, expected: %#v", real, expected)
		}
	}

}
