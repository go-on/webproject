package main

import (
	"fmt"
	"os"

	"gopkg.in/go-on/lib.v3/internal/menu"
	"gopkg.in/go-on/lib.v3/internal/menu/menuhtml"
	"gopkg.in/go-on/lib.v3/types"
)

func main() {
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

	ul := menuhtml.NewUL(types.Class("menu-open"), types.Class("menu-active"), types.Class("menu-sub"))

	// allows to mount a menu that was made in a different way
	subA := m.FindByPath("$sub_a")
	// fmt.Printf("subA: %#v\n", subA)
	m.FindByPath("~replacement").Edges = subA.Edges
	// allows to make an alias
	//m.FindByPath("~replacement").Leaf = menu.Item("/newrepl", "New Repl")
	ul.WriterTo(m, 2, "/a/b/a").WriteTo(os.Stdout)

	fmt.Println("\n#################\n")

	// fmt.Println(UL.Html(m, 2, "/a/a/a").String())
	_ = fmt.Print
	// allows to show a submenu at another place within the layout
	openSub := m.RootAt(1, "/a/a/a")
	// fmt.Printf("opensub: %#v\n", openSub)
	if openSub != nil {
		ul.WriterTo(openSub, 2, "/a/a/a").WriteTo(os.Stdout)

		//fmt.Println(menuhtml.UL.Html(openSub, 2, "/a/a/a"))
	}
}
