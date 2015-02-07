package main

import (
	"fmt"
	"net/http"

	. "gopkg.in/go-on/lib.v3/html"
	"gopkg.in/go-on/lib.v3/html/internal/element"
	. "gopkg.in/go-on/lib.v3/types"
	"gopkg.in/go-on/router.v2"
	"gopkg.in/go-on/router.v2/route/routehtml"
	"gopkg.in/go-on/router.v2/tea/t"
)

var noDecoration = Style{"text-decoration", "none"}

func menu(w http.ResponseWriter, r *http.Request) {
	m := UL()

	entries := []*element.Element{
		LI(routehtml.AHref(&betterRoute, nil, "without params")),
		LI(routehtml.AHref(&helloRoute, routehtml.Params(paramName, "<world>"), "with params")),
		LI(routehtml.AHref(&errorRoute, nil, "with error")),
		LI(AHref("/", "no route")),
	}

	var no = -1
	switch router.GetRouteId(r) {
	case betterRoute.Id:
		no = 0
	case helloRoute.Id:
		no = 1
	case errorRoute.Id:
		no = 2
	}

	if no != -1 {
		entries[no].Add(Class("active"))
	}

	for _, e := range entries {
		m.Add(e)
	}
	m.ServeHTTP(w, r)
}

func layout(body ...interface{}) http.Handler {
	return HTML5(
		HTML(
			HEAD(CssHref(static.MustURL("styles.css"), Media_("screen"))),
			BODY(
				NAV(menu),
				DIV(Class("main"), DIV(body...)),
			),
		),
	)
}

func headingParam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! The parameter is: %s", EscapeHTML(t.RouteParam(r, paramName)))
}
