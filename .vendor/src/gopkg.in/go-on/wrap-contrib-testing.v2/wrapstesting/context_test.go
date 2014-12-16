package wrapstesting

import (
	"fmt"
	"net/http"
	"testing"

	"gopkg.in/go-on/wrap.v2"
	"gopkg.in/go-on/wrap-contrib.v2/helper"
	"gopkg.in/go-on/wrap-contrib.v2/wraps"
)

type ctx struct {
	path string
	http.ResponseWriter
}

type blindctx struct {
	http.ResponseWriter
}

func mkBlindCtx(rw http.ResponseWriter, req *http.Request) http.ResponseWriter {
	return &blindctx{ResponseWriter: rw}
}

func mkCtx(rw http.ResponseWriter, req *http.Request) http.ResponseWriter {
	return &ctx{ResponseWriter: rw}
}

func (c *ctx) Prepare(w http.ResponseWriter, r *http.Request) {
	c.path = r.URL.Path
}

func (c *ctx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("~" + c.path + "~"))
}

func (c *ctx) check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("#" + c.path + "#"))
}

func check(w http.ResponseWriter, r *http.Request) {
	c := &ctx{}
	MustUnWrap(w, &c)
	w.Write([]byte("#" + c.path + "#"))
}

func TestContextHandlerMethod(t *testing.T) {
	r := wrap.New(
		Context(mkBlindCtx),
		Context(mkCtx),
		//wraps.After(http.HandlerFunc(check)),
		wraps.After(HandlerMethod((*ctx).check)),
		wraps.Before(HandlerMethod((*ctx).Prepare)),
		wrap.Handler(HandlerMethod((*ctx).ServeHTTP)),
	)

	rw, req := helper.NewTestRequest("GET", "/path")
	r.ServeHTTP(rw, req)
	err := helper.AssertResponse(rw, "~/path~#/path#", 200)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestContextUnwrapIdentical(t *testing.T) {
	c1 := ctx{path: "x"}
	c2 := &ctx{}

	err := UnWrap(&c1, &c2)

	if err != nil {
		panic(err.Error())
	}

	if c2.path != "x" {
		t.Errorf("c2.path should be x, but is: %#v", c2.path)
	}

	c3 := &ctx{}
	err = UnWrap(c1, &c3)

	if c3.path != "x" {
		t.Errorf("c3.path should be x, but is: %#v", c3.path)
	}

	if err != nil {
		panic(err.Error())
	}
}

func TestContextUnwrapNested(t *testing.T) {
	c1 := blindctx{&ctx{path: "x"}}
	c2 := &ctx{}

	err := UnWrap(&c1, &c2)
	if err != nil {
		panic(err.Error())
	}
	if c2.path != "x" {
		t.Errorf("c2.path should be x, but is: %#v", c2.path)
	}

	c3 := &ctx{}

	err = UnWrap(c1, &c3)
	if err != nil {
		panic(err.Error())
	}
	if c3.path != "x" {
		t.Errorf("c3.path should be x, but is: %#v", c3.path)
	}
}

func TestContextUnwrapError(t *testing.T) {
	_ = fmt.Println
	c1 := blindctx{}
	c2 := &ctx{}
	err := UnWrap(&c1, &c2)
	// fmt.Println(err.Error())
	if err == nil {
		t.Error("unwrap (1) should result in error, but does not")
	}
	rw, _ := helper.NewTestRequest("GET", "/path")
	c1 = blindctx{rw}
	err = UnWrap(&c1, &c2)
	// fmt.Println(err.Error())
	if err == nil {
		t.Error("unwrap (2) should result in error, but does not")
	}
}

func TestContext2HandlerMethod(t *testing.T) {
	r := wrap.New(
		Context(mkBlindCtx),
		Context(mkCtx),
		wraps.After(http.HandlerFunc(check)),
		wraps.Before(HandlerMethod((*ctx).Prepare)),
		wrap.Handler(HandlerMethod((*ctx).ServeHTTP)),
	)

	rw, req := helper.NewTestRequest("GET", "/path")
	r.ServeHTTP(rw, req)
	err := helper.AssertResponse(rw, "~/path~#/path#", 200)

	if err != nil {
		t.Error(err.Error())
	}

}
