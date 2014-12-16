package menu

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func stripWhiteSpace(in string) string {
	return strings.Replace(strings.Replace(strings.Replace(in, "\n", "", -1), "\t", "", -1), " ", "", -1)
}

func TestJSON(t *testing.T) {
	// m.toJson().Text
	var m = &Node{
		Edges: []*Node{
			&Node{Leaf: Item("A", "")},
			&Node{
				Leaf: Item("B", "/b"),
				Edges: []*Node{
					&Node{Leaf: Item("BA", "/ba")},
					&Node{Leaf: Item("BB", "/bb")},
				},
			},
		},
	}

	expected := `{
		  "Subs": [
		    {
		      "Text": "A"
		    },
		    {
		      "Text": "B",
		      "Path": "/b",
		      "Subs": [
		        {
		          "Text": "BA",
		          "Path": "/ba"
		        },
		        {
		          "Text": "BB",
		          "Path": "/bb"
		        }
		      ]
		    }
		  ]
		}
`

	b, err := json.Marshal(m)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	s := fmt.Sprintf("%s", b)

	expected = stripWhiteSpace(expected)

	// expected = strings.Replace(strings.Replace(strings.Replace(expected, "\n", "", -1), "\t", "", -1), " ", "", -1)

	if s != expected {
		t.Errorf("json marshall expected: \n%s\n\ngot:\n%s\n", expected, s)
	}

	var m2 Node

	err = json.Unmarshal(b, &m2)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	b, err = json.Marshal(m2)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	s2 := fmt.Sprintf("%s", b)

	if s != s2 {
		t.Errorf("json unmarshall expected: \n%s\n\ngot:\n%s\n", s, s2)
	}

}

func TestFindByPath(t *testing.T) {
	a := &Node{Leaf: Item("A", "")}
	ba := &Node{Leaf: Item("BA", "/ba")}
	bb := &Node{Leaf: Item("BB", "/bb")}
	b := &Node{
		Leaf: Item("B", "/b"),
		Edges: []*Node{
			ba,
			bb,
		},
	}

	var m = &Node{
		Edges: []*Node{
			a,
			b,
		},
	}

	if m.FindByPath("/b") != b {
		t.Errorf("find path did not find /b")
	}

	if m.FindByPath("/ba") != ba {
		t.Errorf("find path did not find /ba")
	}

	if m.FindByPath("/bb") != bb {
		t.Errorf("find path did not find /bb")
	}
}

func TestFindByText(t *testing.T) {
	a := &Node{Leaf: Item("A", "")}
	ba := &Node{Leaf: Item("BA", "/ba")}
	bb := &Node{Leaf: Item("BB", "/bb")}
	b := &Node{
		Leaf: Item("B", "/b"),
		Edges: []*Node{
			ba,
			bb,
		},
	}

	var m = &Node{
		Edges: []*Node{
			a,
			b,
		},
	}

	if m.FindByText("B") != b {
		t.Errorf("find path did not find B")
	}

	if m.FindByText("BA") != ba {
		t.Errorf("find path did not find BA")
	}

	if m.FindByText("BB") != bb {
		t.Errorf("find path did not find BB")
	}
}

func TestRootAt(t *testing.T) {
	var m = &Node{
		Edges: []*Node{
			&Node{Leaf: Item("B", "")},
			&Node{
				Edges: []*Node{
					&Node{Leaf: Item("repl", "~replacement")},
					&Node{
						Edges: []*Node{
							&Node{Leaf: Item("AAA", "/a/a/a")},
							&Node{Leaf: Item("AAB", "/a/a/b")},
						},
						Leaf: Item("AA", "/a/a"),
					},
					&Node{
						Edges: []*Node{
							&Node{Leaf: Item("ABA", "/a/b/a")},
						},
						Leaf: Item("AB", "$sub_a"),
					},
				},
				Leaf: Item("A", "/a"),
			},
		},
	}

	r1 := m.RootAt(1, "/a/a")

	if r1 == nil {
		t.Errorf("no root for /a/a")
	}

	if r1.Leaf == nil {
		t.Errorf("leaf of root for /a/a is nil")
	}

	if r1.Leaf.Path() != "/a" || r1.Leaf.String() != "A" {
		t.Errorf("leaf should have path /a and Text A, but is: %#v", r1.Leaf)
	}

	r0 := m.RootAt(0, "/a/a")

	if len(r0.Edges) != 2 {
		t.Errorf("must have root at 0")
	}

	r2 := m.RootAt(2, "/a/a/b")

	if r2 == nil {
		t.Errorf("no root 2 for /a/a/b")
	}

	if r2.Leaf == nil {
		t.Errorf("leaf of root 2 for /a/a/b is nil")
	}

	if r2.Leaf.Path() != "/a/a" || r2.Leaf.String() != "AA" {
		t.Errorf("leaf should have path /a/a and Text AA, but is: %#v", r2.Leaf)
	}

}
