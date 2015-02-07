package menuhtml

import (
	"bytes"
	"strings"
	"testing"

	"gopkg.in/go-on/lib.v3/internal/menu"
	"gopkg.in/go-on/lib.v3/types"
)

func stripWhiteSpace(in string) string {
	return strings.Replace(strings.Replace(strings.Replace(in, "\n", "", -1), "\t", "", -1), " ", "", -1)
}

func TestMenuHTML(t *testing.T) {
	m := &menu.Node{
		Edges: []*menu.Node{
			&menu.Node{Leaf: menu.Item("B", "")},
			&menu.Node{
				Edges: []*menu.Node{
					&menu.Node{Leaf: menu.Item("repl", "~replacement")},
					&menu.Node{
						Edges: []*menu.Node{
							&menu.Node{Leaf: menu.Item("AAA", "/a/a/a")},
							&menu.Node{Leaf: menu.Item("AAB", "/a/a/b")},
						},
						Leaf: menu.Item("AA", "/a/a"),
					},
					&menu.Node{
						Edges: []*menu.Node{
							&menu.Node{Leaf: menu.Item("ABA", "/a/b/a")},
						},
						Leaf: menu.Item("AB", "$sub_a"),
					},
				},
				Leaf: menu.Item("A", "/a"),
			},
		},
	}

	ul := NewUL(types.Class("menu-open"), types.Class("menu-active"), types.Class("menu-sub"))

	// allows to mount a menu that was made in a different way
	subA := m.FindByPath("$sub_a")
	// fmt.Printf("subA: %#v\n", subA)
	m.FindByPath("~replacement").Edges = subA.Edges

	var buf bytes.Buffer
	ul.WriterTo(m, 5, "/a/a").WriteTo(&buf)

	expected1 := `
	<ul>
		<li><span>B</span></li>
		<li class="menu-open">
		 <a href="/a">A</a>
		 <ul class="menu-sub">
		 	<li>
		 		<span>repl</span>
		 		<ul class="menu-sub">
		 			<li><a href="/a/b/a">ABA</a></li>
		 		</ul>
		 	</li>
		 	<li class="menu-active">
		 		<a href="/a/a">AA</a>
		 		<ul class="menu-sub">
		 			<li><a href="/a/a/a">AAA</a></li>
		 			<li><a href="/a/a/b">AAB</a></li>
		 		</ul>
		 	</li>
		 	<li>
		 		<span>AB</span>
		 		<ul class="menu-sub">
		 			<li><a href="/a/b/a">ABA</a></li>
		 		</ul>
		 	</li>
		 </ul>
		</li>
	</ul>
`

	expected1 = stripWhiteSpace(expected1)
	got1 := stripWhiteSpace(buf.String())
	if got1 != expected1 {
		t.Errorf("expecting: %#v, got: %#v", expected1, got1)
	}

	buf = bytes.Buffer{}
	ul.WriterTo(m, 1, "/").WriteTo(&buf)

	expected2 := `
	<ul>
		<li><span>B</span></li>
		<li>
			<a href="/a">A</a>
			<ul class="menu-sub">
				<li><span>repl</span></li>
				<li><a href="/a/a">AA</a></li>
				<li><span>AB</span></li>
			</ul>
		</li>
	</ul>`

	expected2 = stripWhiteSpace(expected2)
	got2 := stripWhiteSpace(buf.String())
	if got2 != expected2 {
		t.Errorf("expecting: %#v, got: %#v", expected2, got2)
	}
}
