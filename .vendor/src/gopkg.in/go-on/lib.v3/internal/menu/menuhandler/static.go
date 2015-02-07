package menuhandler

import (
	"gopkg.in/go-on/lib.v3/internal/menu"
	"net/http"
)

type handler struct {
	root     *menu.Node
	writerTo menu.WriterTo
	depth    int
}

func (rm *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wt := rm.writerTo.WriterTo(rm.root, rm.depth, r.URL.Path)
	if wt != nil {
		wt.WriteTo(w)
	}
}

// NewMenuHandler creates a http.Handler that writes the given menu
// to the ResponseWriter, formatting it with the given menuHtml (to the given depth)
// while the menuHtml may handle the current request path in a special way.
func NewStatic(root *menu.Node, depth int, writerTo menu.WriterTo) http.Handler {
	return &handler{
		root:     root,
		writerTo: writerTo,
		depth:    depth,
	}
}

// SubMenuByPath
type subHandler struct {
	root     *menu.Node
	writerTo menu.WriterTo
	// depth of the top menu where the menu starts
	fromDepth int
	// display depth for the submenu
	depth int
}

func (rs *subHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub := rs.root.RootAt(rs.fromDepth, r.URL.Path)
	// sub.Leaf = nil
	// if sub != nil {
	wt := rs.writerTo.WriterTo(sub, rs.depth, r.URL.Path)
	if wt != nil {
		wt.WriteTo(w)
	}
	// }
}

// NewStaticSub creates a static menu handler that is a sub tree of
// the given root, starting at fromDepth until toDepth.
// The menu is only written to the ResponseWriter if the is a tree
// at the given fromDepth for a menu item that matches the request path
func NewStaticSub(root *menu.Node, fromDepth int, toDepth int, writerTo menu.WriterTo) http.Handler {
	depth := toDepth - fromDepth
	if depth < 0 {
		panic("toDepth must not be smaller than fromDepth")
	}

	return &subHandler{
		root:      root,
		writerTo:  writerTo,
		fromDepth: fromDepth,
		depth:     depth,
	}
}
