package main

import (
	"net/http"

	. "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/internal/template/placeholder"
	. "gopkg.in/go-on/lib.v2/types"
	ph "gopkg.in/go-on/lib.v2/types/placeholder"
	"gopkg.in/go-on/router.v2"
)

var (
	header = ph.New(Text("header"))
	body   = ph.New(HTMLString("body"))
	layout = HTML5(
		HTML(
			HEAD(
				TITLE(header),
			),
			BODY(
				NAV(
					A(Attrs_("href", "/"), "navigate to /"), BR(),
					A(Attrs_("href", "/app"), "navigate to /app"), BR(),
					A(Attrs_("href", "/other"), "navigate to /other"), BR(),
				),
				H1(header),
				PRE(body),
			),
		),
	).Template()
)

type App struct {
	URL     string
	Setters []placeholder.Setter
}

func (m *App) setURL(req *http.Request) {
	m.URL = req.URL.Path
}

func (m *App) setHeader() {
	m.Setters = append(m.Setters, header.Setf("header at <%#v>", m.URL))
}

func (m *App) setBody() {
	m.Setters = append(m.Setters, body.Set(PRE("body at", m.URL)))
}

func (m *App) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m.setURL(req)
	m.setHeader()
	m.setBody()
	layout.Replace(m.Setters...).WriteTo(rw)

	// you can also make a switch on req.Method or switch req.URL.Fragment for subroutes
}

func NewApp() http.Handler {
	return &App{}
}

var Router = router.New()

type RouterFunc func() http.Handler

func (rf RouterFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rf().ServeHTTP(rw, req)
}

func main() {
	appRouterFunc := RouterFunc(NewApp)

	Router.GET("/", layout)
	Router.GET("/app", appRouterFunc)
	Router.GET("/other", appRouterFunc)

	Router.Mount("/", nil)

	http.ListenAndServe(":8085", nil)
}
