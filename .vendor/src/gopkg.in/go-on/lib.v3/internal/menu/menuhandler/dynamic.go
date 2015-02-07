package menuhandler

import (
	"gopkg.in/go-on/lib.v3/internal/menu"
	"net/http"
)

type dynHandler struct {
	requestMenu RequestMenu
	writerTo    menu.WriterTo
	depth       int
}

func NewDynamic(requestMenu RequestMenu, depth int, writerTo menu.WriterTo) http.Handler {
	return &dynHandler{
		requestMenu: requestMenu,
		writerTo:    writerTo,
		depth:       depth,
	}
}

func (dm *dynHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wt := dm.writerTo.WriterTo(dm.requestMenu.Menu(r), dm.depth, r.URL.Path)
	if wt != nil {
		wt.WriteTo(w)
	}
}

type dynSubHandler struct {
	requestMenu RequestMenu
	writerTo    menu.WriterTo
	// depth of the top menu where the menu starts
	fromDepth int
	// display depth for the submenu
	depth int
}

// NewDynamicSub creates a static menu handler that is a sub tree of
// the given root, starting at fromDepth until toDepth.
// The menu is only written to the ResponseWriter if the is a tree
// at the given fromDepth for a menu item that matches the request path
func NewDynamicSub(requestMenu RequestMenu, fromDepth int, toDepth int, writerTo menu.WriterTo) http.Handler {
	depth := toDepth - fromDepth
	if depth < 1 {
		panic("toDepth must be larger than fromDepth")
	}

	return &dynSubHandler{
		requestMenu: requestMenu,
		writerTo:    writerTo,
		fromDepth:   fromDepth,
		depth:       depth,
	}
}

func (rs *dynSubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub := rs.requestMenu.Menu(r).RootAt(rs.fromDepth, r.URL.Path)
	// if sub != nil {
	wt := rs.writerTo.WriterTo(sub, rs.depth, r.URL.Path)
	if wt != nil {
		wt.WriteTo(w)
	}
	// }
}
