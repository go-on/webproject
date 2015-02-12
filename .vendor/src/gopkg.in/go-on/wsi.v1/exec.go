package wsi

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Exec is a http.Handler that execs a ExecFunc
type Exec struct {
	mapperFn     func() interface{}
	fn           ExecFunc
	errorHandler func(*http.Request, error)
	dec          RequestDecoder
}

func (we Exec) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mapper := we.mapperFn()
	var err error
	if r.Body == nil {
		err = errors.New("empty body")
		w.WriteHeader(http.StatusBadRequest)
		if we.errorHandler != nil {
			we.errorHandler(r, err)
		}
		return
	}
	defer r.Body.Close()
	err = we.dec.Decode(r, mapper)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if we.errorHandler != nil {
			we.errorHandler(r, err)
		}
		return
	}

	switch r.Method {
	case "PUT":
		if val, ok := mapper.(PUTValidater); ok {
			errs := val.ValidatePUT()
			if len(errs) > 0 {
				w.WriteHeader(http.StatusBadRequest)
				ServeJSON(errsMarshaller(errs), w)
				return
			}
		} else {
			if val, ok := mapper.(Validater); ok {
				errs := val.Validate()
				if len(errs) > 0 {
					w.WriteHeader(http.StatusBadRequest)
					ServeJSON(errsMarshaller(errs), w)
					return
				}
			}
		}
	case "PATCH":
		if val, ok := mapper.(PATCHValidater); ok {
			errs := val.ValidatePATCH()
			if len(errs) > 0 {
				w.WriteHeader(http.StatusBadRequest)
				ServeJSON(errsMarshaller(errs), w)
				return
			}
		} else {
			if val, ok := mapper.(Validater); ok {
				errs := val.Validate()
				if len(errs) > 0 {
					w.WriteHeader(http.StatusBadRequest)
					ServeJSON(errsMarshaller(errs), w)
					return
				}
			}
		}
	case "POST":
		if val, ok := mapper.(POSTValidater); ok {
			errs := val.ValidatePOST()
			if len(errs) > 0 {
				w.WriteHeader(http.StatusBadRequest)
				ServeJSON(errsMarshaller(errs), w)
				return
			}
		} else {
			if val, ok := mapper.(Validater); ok {
				errs := val.Validate()
				if len(errs) > 0 {
					w.WriteHeader(http.StatusBadRequest)
					ServeJSON(errsMarshaller(errs), w)
					return
				}
			}
		}
	}

	var m map[string]interface{}
	m, err = MapSQL(mapper)
	if err != nil {
		if we.errorHandler != nil {
			we.errorHandler(r, err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = we.fn(m, w, r)
	if err != nil {
		if we.errorHandler != nil {
			we.errorHandler(r, err)
		}
	}
}

func (we Exec) SetDecoder(d RequestDecoder) Exec {
	we.dec = d
	return we
}

func (we Exec) SetErrorCallback(fn func(*http.Request, error)) Exec {
	we.errorHandler = fn
	return we
}

type errsMarshaller map[string]error

func (e errsMarshaller) MarshalJSON() ([]byte, error) {
	x := map[string]string{}

	for k, err := range e {
		x[k] = err.Error()
	}

	return json.Marshal(x)
}
