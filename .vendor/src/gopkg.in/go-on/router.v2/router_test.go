package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-on/router.v2/route"
	"gopkg.in/go-on/wrap.v2"

	"gopkg.in/go-on/wrap-contrib.v2/wraps"

	"gopkg.in/go-on/method.v1"
)

type routetest struct {
	path    string
	method  method.Method
	handler http.Handler
}

var routeDefs = []routetest{
	{"/", method.GET, write("root")},
	{"/a.html", method.GET, write("A")},
	{"/b.html", method.GET, write("B")},
	// {"/x.html", "", 404},
	{"/a/x.html", method.GET, write("AX")},
	{"/a/x.html", method.POST, write("AX")},
	{"/a/b.html", method.GET, write("AB")},
	{"/b/x.html", method.GET, write("BX")},
	{"/:sth/x.html", method.GET, writeParams("SthX", "sth")},
	{"/only/get", method.GET, write("only-get")},
}

func makeRouter(mw ...wrap.Wrapper) *Router {
	router := New(mw...)
	for _, r := range routeDefs {
		router.HandleMethod(r.path, r.handler, r.method)
	}
	return router
}

func TestEachRoute(t *testing.T) {
	router := makeRouter()
	router.Mount("/", nil)

	findDef := func(path string) *routetest {
		for _, def := range routeDefs {
			if def.path == path {
				return &def
			}
		}
		return nil
	}

	numTotal := 0
	router.EachRoute(func(mountPoint string, rt *route.Route) {
		if mountPoint == "" {
			t.Errorf("has empty mountpoint")
			return
		}
		if rt == nil {
			t.Errorf("route is nil: %#v", mountPoint)
			return
		}
		def := findDef(mountPoint)
		if def == nil {
			t.Errorf("no route definition found for: %#v", mountPoint)
			return
		}
		numTotal += len(route.Options(rt)) - 2 // minus OPTIONS and HEAD
	})

	if numTotal != len(routeDefs) {
		t.Errorf("unexpected number of routes: %d (should be %d)", numTotal, len(routeDefs))
	}

}

func TestAddRoute(t *testing.T) {
	route := newRouteHandler(route.New("/route", method.GET))
	route.GETHandler = write("ROUTE")
	router := New()
	router.mustAddRouteHandler(route)
	router.Mount("/", nil)

	errMsg := assertResponse(method.GET, "/route", router, "GET ROUTE|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestBody(t *testing.T) {
	router := makeRouter()
	router.Mount("/", nil)

	corpus := map[string]string{
		"/":           "GET root|",
		"/only/get":   "GET only-get|",
		"/a.html":     "GET A|",
		"/b.html":     "GET B|",
		"/a/x.html":   "GET AX|",
		"/a/b.html":   "GET AB|",
		"/b/x.html":   "GET BX|",
		"/var/x.html": "GET SthX|sth:var,",
	}

	for path, expectedbody := range corpus {
		errMsg := assertResponse(method.GET, path, router, expectedbody, 200)
		if errMsg != "" {
			t.Error(errMsg)
		}
	}
}

func TestIsMounted(t *testing.T) {
	router := makeRouter()
	if router.IsMounted() {
		t.Errorf("router should not be mounted")
	}

	router.Mount("/", nil)
	if !router.IsMounted() {
		t.Errorf("router should be mounted")
	}
}

func TestNotFound(t *testing.T) {
	router := makeRouter()
	router.Mount("/", nil)
	errMsg := assertResponse(method.GET, "/notfound", router, "", 405)
	if errMsg != "" {
		t.Error(errMsg)
	}
	errMsg = assertResponse(method.POST, "/only/get", router, "", 405)

	router.NotFound = http.HandlerFunc(http.NotFound)
	errMsg = assertResponse(method.GET, "/notfound", router, "404 page not found", 404)
	if errMsg != "" {
		t.Error(errMsg)
	}
	errMsg = assertResponse(method.POST, "/only/get", router, "404 page not found", 404)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestMountPath(t *testing.T) {
	router1 := makeRouter()
	router1.Mount("/", nil)
	router2 := makeRouter()
	router2.Mount("/api1", nil)
	router3 := makeRouter()
	router3.Mount("/api/2/", nil)

	corpus := map[string]string{
		"/":      router1.MountPath(),
		"/api1":  router2.MountPath(),
		"/api/2": router3.MountPath(),
	}

	for expected, got := range corpus {
		if expected != got {
			t.Errorf("expected: %#v, got: %#v", expected, got)
		}
	}
}

func TestHandleMethods(t *testing.T) {
	router := New()
	router.HandleMethods("/get-post", write("get-post"), method.GET, method.POST)
	router.Mount("/", nil)

	errMsg := assertResponse(method.GET, "/get-post", router, "GET get-post|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.POST, "/get-post", router, "POST get-post|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestOptions(t *testing.T) {
	router := New()
	route := router.handleMethods("/options", write("options"), method.GET, method.POST, method.DELETE)
	SetOPTIONSHandler(route)
	// router.SetOPTIONSHandlers()
	router.Mount("/", nil)
	errMsg := assertResponseHeader(method.OPTIONS, "/options", router, "Allow", "OPTIONS,GET,HEAD,POST,DELETE")
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestMounted(t *testing.T) {
	router := makeRouter()
	router.Mount("/api/v1", nil)

	errMsg := assertResponse(method.GET, "/api/v1/b/x.html", router, "GET BX|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestHandle(t *testing.T) {
	router := New()
	router.Handle("/all", write("all"))
	router.Mount("/", nil)

	corpus := map[method.Method]string{
		method.GET:     "GET all|",
		method.POST:    "POST all|",
		method.PUT:     "PUT all|",
		method.PATCH:   "PATCH all|",
		method.DELETE:  "DELETE all|",
		method.OPTIONS: "OPTIONS all|",
	}

	for meth, expected := range corpus {
		errMsg := assertResponse(meth, "/all", router, expected, 200)
		if errMsg != "" {
			t.Error(errMsg)
		}
	}
}

func TestWrappers(t *testing.T) {
	router := makeRouter(wraps.Around(writeString("^"), writeString("$")))
	router.Mount("/", nil)
	errMsg := assertResponse(method.GET, "/b/x.html", router, "^GET BX|$", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestWrappersWithSub(t *testing.T) {
	router := New(wraps.Around(writeString("^"), writeString("$")))
	router.GET("/body", write("body"))

	sub := New(wraps.Around(writeString("A"), writeString("Z")))
	sub.GET("/body", write("sub-body"))
	sub.Mount("/sub", router)

	router.Mount("/", nil)
	errMsg := assertResponse(method.GET, "/body", router, "^GET body|$", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.GET, "/sub/body", router, "^AGET sub-body|Z$", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestSubRouter(t *testing.T) {
	sub := makeRouter()
	router := New()
	sub.Mount("/sub", router)
	//router.Handle("/sub", sub)
	router.Mount("/api/1", nil)
	errMsg := assertResponse(method.GET, "/api/1/sub/b/x.html", router, "GET BX|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestOptionsWithSub(t *testing.T) {
	sub := New()
	route3 := sub.handleMethods("/options", write("options"), method.GET, method.POST, method.DELETE)
	SetOPTIONSHandler(route3)

	router := New()
	route1 := router.handleMethods("/options", write("options"), method.GET, method.PUT)
	SetOPTIONSHandler(route1)
	router.Handle("/sub", sub)
	router.Mount("/api/1", nil)

	errMsg := assertResponseHeader(method.OPTIONS, "/api/1/options", router, "Allow", "OPTIONS,GET,HEAD,PUT")
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponseHeader(method.OPTIONS, "/api/1/sub", router, "Allow", "")
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponseHeader(method.OPTIONS, "/api/1/sub/options", router, "Allow", "OPTIONS,GET,HEAD,POST,DELETE")
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestRoute(t *testing.T) {
	router := New()
	route := router.GET("/route", write("ROUTE"))

	router.Mount("/", nil)

	rq, _ := http.NewRequest("GET", "/route", nil)
	got := router.RequestRoute(rq)

	if got != route {
		t.Errorf("wrong route")
	}
}

func TestHead(t *testing.T) {
	router := New()
	router.GET("/head", write("will be hidden by server"))
	router.Mount("/", nil)

	errMsg := assertResponse(method.HEAD, "/head", router, "HEAD will be hidden by server|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	router1 := New()
	router1.GET("/route", write("route"))
	router1.Mount("/app", mux)

	errMsg := assertResponse(method.GET, "/app/route", mux, "GET route|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
	router2 := New()
	router2.GET("/", write("root"))
	router2.Mount("/", mux)

	errMsg = assertResponse(method.GET, "/", mux, "GET root|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestServe(t *testing.T) {
	router := New()
	router.GET("/rewrite", write("rewrite"))
	router.Mount("/", nil)

	errMsg := assertResponse(method.GET, "/rewrite//", router, "GET rewrite|", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.GET, "/rewrite//", router.ServingHandler(), "<a href=\"/rewrite/\">Moved Permanently</a>.", 301)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestServeNotMounted(t *testing.T) {
	router := New()
	router.GET("/rewrite", write("rewrite"))

	errMsg := assertResponse(method.GET, "/rewrite//", router.ServingHandler(), "<a href=\"/rewrite/\">Moved Permanently</a>.", 301)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

func TestGetRouteId(t *testing.T) {
	router := New()
	route1 := router.GETFunc("/route1", wrap.NoOp)
	route2 := router.GETFunc("/route2", wrap.NoOp)
	router.Mount("/", nil)

	rq, _ := http.NewRequest("GET", "/route1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, rq)
	got := GetRouteId(rq)
	if got != route1.Id {
		t.Errorf("has wrong id: %#v, should have: %#v", got, route1.Id)
	}

	if got == route2.Id {
		t.Errorf("must not have id: %#v", route2.Id)
	}

	rq, _ = http.NewRequest("GET", "/route-not-existent", nil)
	got = GetRouteId(rq)

	if got != "" {
		t.Errorf(`route id should be "", but is %#v`, got)
	}
}

func TestEtagged(t *testing.T) {
	router := NewETagged()
	router.GET("/etag", writeString("etag"))
	router.Mount("/", nil)

	errMsg := assertResponseHeader(method.GET, "/etag", router, "Etag", "1872ade88f3013edeb33decd74a4f947")
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.GET, "/etag", router, "etag", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

// TestEtaggedWrapped tests that the wrappers do not affect the etag
func TestEtaggedWrapped(t *testing.T) {
	router := NewETagged(wraps.Around(writeString("^"), writeString("$")))
	router.GET("/etag", writeString("etag"))
	router.Mount("/", nil)

	errMsg := assertResponseHeader(method.GET, "/etag", router, "Etag", "1872ade88f3013edeb33decd74a4f947")
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.GET, "/etag", router, "^etag$", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}
}

// TestHandleRoute tests that the wrappers do not affect the etag
func TestHandleRouteMethods(t *testing.T) {
	rt := route.New("/a", method.GET)
	router := New()
	router.HandleRouteMethods(rt, writeString("A"), method.GET)
	router.Mount("/c", nil)

	errMsg := assertResponse(method.GET, "/c/a", router, "A", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

}

func TestHandleRouteFunc(t *testing.T) {
	rt := route.New("/a", method.GET, method.POST)
	router := New()
	router.HandleRouteFunc(rt, writeString("A").ServeHTTP)
	router.Mount("/c", nil)

	errMsg := assertResponse(method.GET, "/c/a", router, "A", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.POST, "/c/a", router, "A", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

}

func TestHandleRoute(t *testing.T) {
	rt := route.New("/a", method.GET)
	router := New()
	router.HandleRoute(rt, writeString("A"))
	router.Mount("/c", nil)

	errMsg := assertResponse(method.GET, "/c/a", router, "A", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

}

// TestHandleRoute tests that the wrappers do not affect the etag
func TestHandleRouteMethodsFunc(t *testing.T) {
	rt := route.New("/a", method.GET, method.PUT)
	router := New()
	router.HandleRouteMethodsFunc(rt, writeString("A").ServeHTTP, method.GET, method.PUT)
	router.Mount("/c", nil)

	errMsg := assertResponse(method.GET, "/c/a", router, "A", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

	errMsg = assertResponse(method.PUT, "/c/a", router, "A", 200)
	if errMsg != "" {
		t.Error(errMsg)
	}

}

/*
func TestNewVariant(t *testing.T) {
	a := route.NewRoute("/:x/a.html")
	a.GETHandler = writeParams("A", "x")
	a.POSTHandler = writeParams("A", "x")

	// fmt.Printf("route a is %p\n", a)

	rtr := New()
	rtr.MustAdd(a)
	rtr.Mount("/", nil)

	rec, req := newTestRequest("GET", "/y/a.html")
	rtr.ServeHTTP(rec, req)

	body := rec.Body.String()

	exp := "GET A|x:y,"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}

	rec, req = newTestRequest("POST", "/z/a.html")
	rtr.ServeHTTP(rec, req)

	body = rec.Body.String()
	exp = "POST A|x:z,"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}
	// assertResponse(c, rw, "A", 200)
}

func TestNewVariant2(t *testing.T) {
	a := route.NewRoute("/a.html")
	a.GETHandler = write("A")
	a.POSTHandler = write("A")

	b := route.NewRoute("/:sth/x.html")
	b.GETHandler = writeParams("B", "sth")

	// fmt.Printf("route a is %p\n", a)

	rtr := New()
	rtr.MustAdd(a)
	rtr.MustAdd(b)
	rtr.Mount("/", nil)

	rec, req := newTestRequest("GET", "/a.html")
	rtr.ServeHTTP(rec, req)

	body := rec.Body.String()
	exp := "GET A|"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}

	rec, req = newTestRequest("POST", "/a.html")
	rtr.ServeHTTP(rec, req)

	body = rec.Body.String()
	exp = "POST A|"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}

	rec, req = newTestRequest("GET", "/x/x.html")
	rtr.ServeHTTP(rec, req)

	body = rec.Body.String()
	exp = "GET B|sth:x,"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}
	// assertResponse(c, rw, "A", 200)
}

func TestNewMounted(t *testing.T) {
	a := route.NewRoute("/:x/:p/a/:b")
	a.GETHandler = writeParams("A", "x", "p", "b")
	a.POSTHandler = writeParams("A", "x", "p", "b")

	// fmt.Printf("route a is %p\n", a)

	rtr := New()
	rtr.MustAdd(a)
	rtr.Mount("/ho", nil)

	rec, req := newTestRequest("GET", "/ho/y/f/a/q")
	rtr.ServeHTTP(rec, req)

	body := rec.Body.String()
	exp := "GET A|x:y,p:f,b:q,"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}

	rec, req = newTestRequest("POST", "/ho/z/g/a/r")
	rtr.ServeHTTP(rec, req)

	body = rec.Body.String()
	exp = "POST A|x:z,p:g,b:r,"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}
	// assertResponse(c, rw, "A", 200)
}

func TestNewSub(t *testing.T) {
	zero := route.NewRoute("/zero")
	a := route.NewRoute("/:x/:p/a/:b")
	a.GETHandler = writeParams("A", "x", "p", "b")

	// fmt.Printf("route a is %p\n", a)

	rtr := New()
	rtr.MustAdd(a)
	zero.GETHandler = rtr

	outer := New()
	outer.MustAdd(zero)
	outer.Mount("/ho", nil)

	rec, req := newTestRequest("GET", "/ho/zero/y/f/a/q")
	outer.ServeHTTP(rec, req)

	body := rec.Body.String()
	exp := "GET A|x:y,p:f,b:q,"
	if body != exp {
		t.Errorf("expected %#v, got: %#v", exp, body)
	}

}
*/

/*
func TestFindParams(t *testing.T) {
	corpus := map[string]string{
		"abc/def/key/val1/mno/pqr//0xfun":  "val1",
		"key/val2/mno/pqr//0xfun":          "val2",
		"abc/def/xyz/val1/key/val3//0xfun": "val3",
	}

	for k, v := range corpus {
		start, end := findParam(k, "key")
		got := ""
		if start > -1 && end > -1 {
			got = k[start:end]
		}
		if got != v {
			t.Errorf("expected: %#v, got: %#v", v, got)
		}
	}

}
*/
