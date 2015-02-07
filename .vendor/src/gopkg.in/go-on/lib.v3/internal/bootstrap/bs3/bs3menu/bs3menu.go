package bs3menu

import (
	"fmt"
	"io"
	// "fmt"
	. "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	"gopkg.in/go-on/lib.v3/internal/bootstrap/bs3"
	// "gopkg.in/go-on/lib.v3/html/h"
	"gopkg.in/go-on/lib.v3/internal/menu"
	"gopkg.in/go-on/lib.v3/internal/menu/menuhtml"
	"gopkg.in/go-on/lib.v3/types"
)

type navMenu struct {
	baseClass         types.Class
	dropDown          bool
	dataToggleTab     bool
	additionalClasses []types.Class
}

func (n navMenu) Item(path, text string, isActive, hasChildren bool, depth int) *element.Element {
	e := LI()
	var link *element.Element
	if path == "_" {
		link = AHref("#", text)
		e.Add(bs3.Disabled)
	} else {
		if n.dropDown {
			e.Add(bs3.Dropdown)
			link = AHref(path, text+" ")
			if hasChildren && depth < 2 {
				// fmt.Printf("hasChildren: %v\n", hasChildren)
				link.Add(SPAN(bs3.Caret))
				link.Add(bs3.Dropdown_toggle, types.Attribute{"data-toggle", "dropdown"})
			}
		} else {
			link = AHref(path, text)
		}
	}
	if text == "---" {
		e.Add(bs3.Divider)
	} else {
		if n.dataToggleTab {
			link.Add(types.Attribute{"data-toggle", "tab"})
		}
		e.Add(link)
	}
	if isActive {
		e.Add(bs3.Active)
	}
	return e
}

func (n navMenu) List(depth int, text string) *element.Element {
	e := UL(bs3.Nav)
	if n.dropDown && depth == 1 {
		e.Add(bs3.Dropdown_menu)
	}
	if depth == 0 {
		e.Add(n.baseClass)

		for _, cl := range n.additionalClasses {
			e.Add(cl)
		}

	}
	return e
}

func (n navMenu) ClassOpen() types.Class {
	return types.Class("active")
}

type navDropdown struct{}

func (n navDropdown) Item(path, text string, isActive, hasChildren bool, depth int) *element.Element {
	e := LI()
	if path != "" {
		link := AHref(path, text)
		if path == "_" {
			e.Add(bs3.Disabled)
			link = AHref("#", text)
		}
		e.Add(link)
	} else {
		if text == "---" {
			e.Add(bs3.Divider)
		} else {
			e.Add(text)
		}
	}
	if isActive {
		e.Add(bs3.Active)
	}
	return e
}

func (n navDropdown) List(depth int, text string) *element.Element {
	e := UL(bs3.Dropdown_menu) //html.Attr("role", "menu"),

	return e
}

func (n navDropdown) ClassOpen() types.Class {
	return types.Class("active")
}

func Dropdown() menu.WriterTo {
	return menuhtml.New(navDropdown{})
}

type navDropdownBtn struct {
	baseClass    types.Class
	fallbackText string
	textFormat   string
}

// WriterTo(root *Node, depth int, path string) io.WriterTo

func (n navDropdownBtn) WriterTo(root *menu.Node, depth int, path string) io.WriterTo {
	if depth > 0 {
		return nil
	}

	if root == nil {
		return BUTTON(
			bs3.Btn, bs3.Dropdown_toggle,
			n.baseClass,
			bs3.Disabled,
			n.fallbackText+" ", SPAN(bs3.Caret),
		)
	}

	var text string
	if root.Leaf != nil && root.Leaf.String() != "" && n.textFormat != "" {
		text = fmt.Sprintf(n.textFormat, root.Leaf.String())
	}

	if text == "" {
		text = n.fallbackText
	}

	return BUTTON(
		bs3.Btn, bs3.Dropdown_toggle,
		n.baseClass,
		// bs3.bs3.Disabled,
		types.Attribute{"data-toggle", "dropdown"},
		types.Attribute{"type", "button"},
		text+" ", SPAN(bs3.Caret),
	)

}

func DropdownButton(btnClass types.Class, textformat, fallbacktext string) menu.WriterTo {
	return navDropdownBtn{btnClass, fallbacktext, textformat}
}

func Button(btnClass types.Class, textformat, fallbacktext string) menu.WriterTo {
	return navBtn{btnClass, fallbacktext, textformat}
}

type breadcrumb struct{}

func (b breadcrumb) find(node *menu.Node, path string) []menu.Leaf {
	if node.Leaf != nil && node.Leaf.Path() == path {
		return []menu.Leaf{node.Leaf}
	}
	for _, edge := range node.Edges {
		l := b.find(edge, path)
		if len(l) > 0 {
			return append([]menu.Leaf{node.Leaf}, l...)
		}
	}

	return nil
}

func (b breadcrumb) WriterTo(root *menu.Node, depth int, path string) io.WriterTo {
	if root == nil {
		return nil
	}

	var l []menu.Leaf

	for _, edge := range root.Edges {
		l = b.find(edge, path)
		if len(l) > 0 {
			break
		}
	}

	bc := OL(bs3.Breadcrumb)

	for _, leaf := range l {
		if leaf != nil {
			li := LI(AHref(leaf.Path(), leaf.String()))
			if leaf.Path() == path {
				li.Add(bs3.Active)
			}
			bc.Add(li)
		}
	}

	return bc
}

func Breadcrumb() breadcrumb {
	return breadcrumb{}
}

type navBtn struct {
	baseClass    types.Class
	fallbackText string
	textFormat   string
}

// WriterTo(root *Node, depth int, path string) io.WriterTo
func (n navBtn) find(node *menu.Node, path string) (found bool) {
	if node.Leaf != nil && node.Leaf.Path() == path {
		return true
	}
	for _, edge := range node.Edges {
		found = n.find(edge, path)
		if found {
			return
		}
	}

	return false
}

func (n navBtn) WriterTo(root *menu.Node, depth int, path string) io.WriterTo {
	if depth > 0 {
		return nil
	}

	if root == nil {
		return BUTTON(
			bs3.Btn,
			n.baseClass,
			bs3.Disabled,
			n.fallbackText,
		)
	}

	text := ""

	for _, edge := range root.Edges {
		found := n.find(edge, path)
		if found {
			if edge.Leaf != nil {
				text = edge.Leaf.String()
			}
			break
		}
	}

	if text == "" {
		text = n.fallbackText
	}

	return BUTTON(
		bs3.Btn,
		n.baseClass,
		// bs3.bs3.Disabled,
		types.Attribute{"type", "button"},
		text,
	)

}

func Tabs(dropDown bool, datatoggle bool, classes ...types.Class) menu.WriterTo {
	return menuhtml.New(navMenu{bs3.Nav_tabs, dropDown, datatoggle, classes})
}

func Pills(dropDown bool, classes ...types.Class) menu.WriterTo {
	return menuhtml.New(navMenu{bs3.Nav_pills, dropDown, false, classes})
}

func NavBar(classes ...types.Class) menu.WriterTo {
	return menuhtml.New(navMenu{bs3.Navbar_nav, true, false, classes})
}

type listgroup struct {
}

func (lg listgroup) WriterTo(root *menu.Node, depth int, path string) io.WriterTo {
	d := DIV(bs3.List_group)

	if root == nil {
		return d
	}

	for _, edge := range root.Edges {
		if edge.Leaf != nil {
			if edge.Leaf.Path() != "" && edge.Leaf.Path() != "_" {
				link := AHref(edge.Leaf.Path(), edge.Leaf.String(), bs3.List_group_item)
				if edge.Leaf.Path() == path {
					link.Add(bs3.Active)
				}
				d.Add(link)
			} else {
				d.Add(SPAN(edge.Leaf.String(), bs3.List_group_item))
			}
		}
	}
	return d
}

func ListGroup() listgroup {
	return listgroup{}
}
