package wrapstesting

import (
	"net/http"
	"testing"

	"gopkg.in/go-on/wrap-contrib.v2/helper"
	"gopkg.in/go-on/wrap-contrib.v2/wraps"
	"gopkg.in/go-on/wrap.v2"
)

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
	return &_handle{ResponseWriter: rw}
}

func TestResponseWriterHandlerMethod(t *testing.T) {
	r := wrap.New(
		Context(mkHandle),
		wraps.Before(HandlerMethod((*_handle).Prepare)),
		ResponseWriterHandler,
	)

	rw, req := helper.NewTestRequest("GET", "/path")
	r.ServeHTTP(rw, req)
	err := helper.AssertResponse(rw, "~/path~", 200)

	if err != nil {
		t.Error(err)
	}
}
