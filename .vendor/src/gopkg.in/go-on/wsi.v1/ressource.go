package wsi

import (
	"net/http"
)

// QueryFunc makes the sql query and returns a Scanner. If an error is returned, QueryFunc must write
// to the reponsewriter (set the status code etc). If no error is returned QueryFunc must not write
// to the response write. specific headers are the exception and may be set.
type QueryFunc func(QueryOptions, http.ResponseWriter, *http.Request) (Scanner, error)

// ExecFunc makes the sql exec and writes to the response writer. If must return an error, if
// some happened, so that the error may be passed to the general error handler
type ExecFunc func(Mapper, http.ResponseWriter, *http.Request) error

type RessourceFunc func() Mapper

func (r RessourceFunc) Exec(e ExecFunc) Exec {
	if e == nil {
		panic("ExecFunc can't be nil")
	}
	return Exec{mapperFn: r, fn: e, dec: JSONDecoder}
}

func (r RessourceFunc) Query(q QueryFunc) Query {
	if q == nil {
		panic("QueryFunc can't be nil")
	}
	return Query{encFn: NewJSONStreamer, mapperFn: r, fn: q}
}

type Ressource struct {
	RessourceFunc
	ErrorHandler func(r *http.Request, err error)
}

func (rs Ressource) ServeQuery(q QueryFunc, w http.ResponseWriter, r *http.Request) {
	qq := rs.RessourceFunc.Query(q)
	if rs.ErrorHandler != nil {
		qq = qq.SetErrorCallback(rs.ErrorHandler)
	}
	qq.ServeHTTP(w, r)
}

func (rs Ressource) ServeExec(e ExecFunc, w http.ResponseWriter, r *http.Request) {
	ee := rs.RessourceFunc.Exec(e)
	if rs.ErrorHandler != nil {
		ee = ee.SetErrorCallback(rs.ErrorHandler)
	}
	ee.ServeHTTP(w, r)
}

func NewRessource(fn func() Mapper) Ressource {
	return Ressource{RessourceFunc: fn}
}
