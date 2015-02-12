package wsi

import (
	"net/http"
)

// QueryFunc makes the sql query and returns a Scanner. If an error is returned, QueryFunc must write
// to the reponsewriter (set the status code etc). If no error is returned QueryFunc must not write
// to the response write. specific headers are the exception and may be set.
type QueryFunc func(limit, offset int, w http.ResponseWriter, r *http.Request) (Scanner, error)

// ExecFunc makes the sql exec and writes to the response writer. If must return an error, if
// some happened, so that the error may be passed to the general error handler
type ExecFunc func(map[string]interface{}, http.ResponseWriter, *http.Request) error

type Ressource struct {
	RessourceFunc func() interface{}
	ErrorHandler  func(r *http.Request, err error)
}

func (rs Ressource) ServeQuery(q QueryFunc, w http.ResponseWriter, r *http.Request) {
	rs.Query(q).ServeHTTP(w, r)
}

func (rs Ressource) ServeExec(e ExecFunc, w http.ResponseWriter, r *http.Request) {
	rs.Exec(e).ServeHTTP(w, r)
}

func (rs Ressource) Exec(e ExecFunc) Exec {
	if e == nil {
		panic("ExecFunc can't be nil")
	}
	ee := Exec{mapperFn: rs.RessourceFunc, fn: e, dec: JSONDecoder}
	if rs.ErrorHandler != nil {
		ee = ee.SetErrorCallback(rs.ErrorHandler)
	}
	return ee
}

func (rs Ressource) Query(q QueryFunc) Query {
	if q == nil {
		panic("QueryFunc can't be nil")
	}
	qq := Query{encFn: NewJSONStreamer, mapperFn: rs.RessourceFunc, fn: q}
	if rs.ErrorHandler != nil {
		qq = qq.SetErrorCallback(rs.ErrorHandler)
	}
	return qq
}
