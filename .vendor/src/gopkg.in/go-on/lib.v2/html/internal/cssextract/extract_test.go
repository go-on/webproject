package cssextract

import (
	"testing"
)

type testCase struct {
	Style   string
	Ids     []string
	Classes []string
}

var testCases = []testCase{
	{
		`.a { display: none; }`,
		[]string{},
		[]string{"a"},
	},
	{
		`.a-ha42 { display: none; }`,
		[]string{},
		[]string{"a-ha42"},
	},
	{
		`.a.b { display: none; }`,
		[]string{},
		[]string{"a", "b"},
	},
	{
		`#hu.a.b { display: none; }`,
		[]string{"hu"},
		[]string{"a", "b"},
	},
	{
		`#hu.a #ho.b { display: none; }`,
		[]string{"hu", "ho"},
		[]string{"a", "b"},
	},
	{
		`table #hu.a.c > #ho.b.c #hu { display: none; font-weight:bold; }`,
		[]string{"hu", "ho"},
		[]string{"a", "b", "c"},
	},
	{
		`table.a { 
			display: none; 
			font-weight:bold; 
			background-image: url('image.jpg');
		}`,
		[]string{},
		[]string{"a"},
	},
	{
		`@media screen and (min-width: 650px) {
  		.ho-ho {
		    width: auto;
		    max-width: 1200px;
		    margin: 0 auto;
		    padding: 0 1.5em;
		  }
		 }
		@font-face {
		  font-weight: normal;
		  font-style: normal;
		  background-image: url('image.jpg');
		}
		 `,
		[]string{},
		[]string{"ho-ho"},
	},
	{
		`/*
		  .hiho
		 */
		 table.a {
		  /* .hu */ 
			display: none; 
			font-weight:bold; 
			background-image: url('image.jpg');
		}`,
		[]string{},
		[]string{"a"},
	},

	{
		`p.x[data-y^=".ho"] .z { font-weight: bold; }`,
		[]string{},
		[]string{"x", "z"},
	},
}

func TestExtract(t *testing.T) {

	for n, tc := range testCases {
		p := Parse(tc.Style)

		if len(p.Classes) != len(tc.Classes) {
			t.Errorf("expected %d classes, got: %d in testcase %d", len(tc.Classes), len(p.Classes), n)
		}

		if len(p.Ids) != len(tc.Ids) {
			t.Errorf("expected %d ids, got: %d in testcase %d", len(tc.Ids), len(p.Ids), n)
		}

		for _, id := range tc.Ids {
			if !hasId(p, id) {
				t.Errorf("missing id %#v in testcase %d", id, n)
			}
		}

		for _, class := range tc.Classes {
			if !hasClass(p, class) {
				t.Errorf("missing class %#v in testcase %d", class, n)
			}
		}

	}

}
