package wrapstesting

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"gopkg.in/go-on/wrap.v2"
)

type reverseProxy struct {
	*httputil.ReverseProxy
}

func ReverseProxy(rev *httputil.ReverseProxy) wrap.Wrapper {
	return &reverseProxy{rev}
}

func ReverseProxyByUrl(urlbase string) wrap.Wrapper {
	u, err := url.Parse(urlbase)

	if err != nil {
		panic(err.Error())
	}

	return &reverseProxy{httputil.NewSingleHostReverseProxy(u)}
}

// ServeHandle serves the request via the ReverseProxy, ignoring any inner handler
func (rp *reverseProxy) ServeHTTPNext(inner http.Handler, rw http.ResponseWriter, req *http.Request) {
	rp.ServeHTTP(rw, req)
}

// Wrap wraps the given inner handler with the returned handler
func (rp *reverseProxy) Wrap(next http.Handler) http.Handler {
	return wrap.NextHandler(rp).Wrap(next)
}
