package pseudo

import (
	"testing"
)

func TestAll(t *testing.T) {
	tests := map[string]string{
		Root().Selector():          `:root`,
		NthChild(4).Selector():     `:nth-child(4)`,
		NthLastChild(4).Selector(): `:nth-last-child(4)`,
		NthOfType(4).Selector():    `:nth-of-type(4)`,
		FirstChild().Selector():    `:first-child`,
		LastChild().Selector():     `:last-child`,
		FirstOfType().Selector():   `:first-of-type`,
		LastOfType().Selector():    `:last-of-type`,
		OnlyChild().Selector():     `:only-child`,
		OnlyOfType().Selector():    `:only-of-type`,
		Empty().Selector():         `:empty`,
		Link().Selector():          `:link`,
		Visited().Selector():       `:visited`,
		Active().Selector():        `:active`,
		Hover().Selector():         `:hover`,
		Focus().Selector():         `:focus`,
		Target().Selector():        `:target`,
		Lang("en").Selector():      `:lang(en)`,
		Enabled().Selector():       `:enabled`,
		Disabled().Selector():      `:disabled`,
		Checked().Selector():       `:checked`,
		FirstLine().Selector():     `::first-line`,
		FirstLetter().Selector():   `::first-letter`,
		Before().Selector():        `::before`,
		After().Selector():         `::after`,
		Not(Empty()).Selector():    `:not(:empty)`,
	}

	for real, expected := range tests {
		if real != expected {
			t.Errorf("got %#v, expected: %#v", real, expected)
		}
	}

}
