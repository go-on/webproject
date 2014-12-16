package wrapstesting

import (
	// "fmt"
	"net/http"

	"gopkg.in/go-on/wrap.v2"
	"gopkg.in/go-on/wrap-contrib.v2/helper"
	"gopkg.in/go-on/wrap-contrib.v2/wraps"
	. "launchpad.net/gocheck"
)

type jsonSuite struct{}

var _ = Suite(&jsonSuite{})

func mkJsonCtx(rw http.ResponseWriter, req *http.Request) http.ResponseWriter {
	return &jsonCtx{ResponseWriter: rw}
}

type jsonCtx struct {
	Path                string
	http.ResponseWriter `json:"-"`
}

func (j *jsonCtx) Prepare(w http.ResponseWriter, r *http.Request) {
	j.Path = r.URL.Path
}

func (s *jsonSuite) TestJson(c *C) {

	r := wrap.New(
		Context(mkJsonCtx),
		wraps.Before(HandlerMethod((*jsonCtx).Prepare)),
		Json(&jsonCtx{}),
	)

	rw, req := helper.NewTestRequest("GET", "/hiho")
	r.ServeHTTP(rw, req)
	err := helper.AssertResponse(rw, `{"Path":"/hiho"}`, 200)

	c.Assert(err, Equals, nil)

	rw, req = helper.NewTestRequest("GET", "/huho")
	r.ServeHTTP(rw, req)
	err = helper.AssertResponse(rw, `{"Path":"/huho"}`, 200)

	c.Assert(err, Equals, nil)
}
