package wrapstesting

import (
	"net/http"

	"gopkg.in/go-on/wrap.v2"
	"gopkg.in/go-on/wrap-contrib.v2/helper"
	"gopkg.in/go-on/wrap-contrib.v2/wraps"
	. "launchpad.net/gocheck"
)

type handleSuite struct{}

var _ = Suite(&handleSuite{})

type _handle struct {
	path string
	http.ResponseWriter
}

func (c *_handle) Prepare(w http.ResponseWriter, r *http.Request) {
	c.path = r.URL.Path
}

func (c *_handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("~" + c.path + "~"))
}

func mkHandle(rw http.ResponseWriter, req *http.Request) http.ResponseWriter {
	return _handle{ResponseWriter: rw}
}

func (s *handleSuite) TestContextHandlerMethod(c *C) {
	r := wrap.New(
		Context(mkHandle),
		wraps.Before(HandlerMethod((*_handle).Prepare)),
		ResponseWriterHandler,
	)

	rw, req := helper.NewTestRequest("GET", "/path")
	r.ServeHTTP(rw, req)
	err := helper.AssertResponse(rw, "~/path~", 200)

	c.Assert(err, Equals, nil)
}
