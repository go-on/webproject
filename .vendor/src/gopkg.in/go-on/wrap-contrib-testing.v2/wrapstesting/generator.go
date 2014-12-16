package wrapstesting

import (
	"net/http"

	"gopkg.in/go-on/wrap.v2"
)

type Generater interface {
	New(w http.ResponseWriter, r *http.Request) http.ResponseWriter
}

type GeneraterFunc func(w http.ResponseWriter, r *http.Request) http.ResponseWriter

func (gf GeneraterFunc) New(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	return gf(w, r)
}

// a generator serves a request by creating a new request handler
// on each request that serves the current request
type generator struct {
	Generater
}

func (ø generator) Wrap(in http.Handler) (out http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in.ServeHTTP(ø.Generater.New(w, r), r)
	})
}

func Generator(gen Generater) wrap.Wrapper {
	return generator{gen}
}

func GeneratorFunc(fn func(w http.ResponseWriter, r *http.Request) http.ResponseWriter) wrap.Wrapper {
	return generator{GeneraterFunc(fn)}
}
