package selector

import (
	. "gopkg.in/go-on/lib.v3/types"
	"testing"
)

func TestSelectors(t *testing.T) {

	tests := map[string]string{
		Descendant(Tag("li"), Tag("a")).Selector():             `li a`,
		Child(Tag("li"), Tag("a")).Selector():                  `li > a`,
		Child(Tag("li"), Tag("a")).Add(Tag("em")).Selector():   `li > a > em`,
		DirectFollows(Tag("li"), Tag("a")).Selector():          `li + a`,
		Follows(Tag("li"), Tag("a")).Selector():                `li ~ a`,
		Each(Tag("li"), Tag("a")).Selector():                   "li,\na",
		ContextString(".main", Tag("li"), Tag("a")).Selector(): ".main a,\n.main li",
	}

	for got, expected := range tests {
		if got != expected {
			t.Errorf("got: %#v expected: %#v", got, expected)
		}
	}

}
