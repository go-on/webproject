package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResp struct {
	Success          bool
	ErrorMessage     string            `json:",omitempty"`
	ValidationErrors map[string]string `json:",omitempty"`
	status           int
}

type okCreated string

func (k okCreated) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	setJsonContentType(wr)
	resp := struct {
		Success bool
		Created string
	}{true, string(k)}
	wr.WriteHeader(201)
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(wr, "%s", b)
}

func serveError(err error, wr http.ResponseWriter, rq *http.Request) {
	switch e := err.(type) {
	case *json.SyntaxError, *json.InvalidUTF8Error, *json.UnsupportedTypeError:
		ErrNoJson.ServeHTTP(wr, rq)
	case *json.InvalidUnmarshalError:
		ErrInvalidJson.ServeHTTP(wr, rq)
	case *validationError:
		e.ServeHTTP(wr, rq)
		// NewValidationErrorHandler(e.validationErrors).ServeHTTP(wr, rq)
	case notFound:
		ErrNotFound.ServeHTTP(wr, rq)
	case fieldNotAllowed:
		e.ServeHTTP(wr, rq)
	default:
		ErrServer.ServeHTTP(wr, rq)
	}
}

func serveSuccess(wr http.ResponseWriter, rq *http.Request) {
	setJsonContentType(wr)
	resp := struct {
		Success bool
	}{true}
	wr.WriteHeader(200)
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(wr, "%s", b)
}

func (resp *jsonResp) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	setJsonContentType(wr)
	b, _ := json.MarshalIndent(resp, "", "  ")
	wr.WriteHeader(resp.status)
	fmt.Fprintf(wr, "%s", b)
}

func (resp *jsonResp) Error() string {
	return resp.ErrorMessage
}

var ErrNotFound = &jsonResp{
	status:           404,
	ErrorMessage:     "Not found",
	ValidationErrors: map[string]string{},
	Success:          false,
}

var ErrServer = &jsonResp{
	status:           500,
	ErrorMessage:     "Internal Server Error",
	ValidationErrors: map[string]string{},
	Success:          false,
}

var ErrNoJson = &jsonResp{
	status:           415, // Unsupported Media Type
	ErrorMessage:     "no json",
	ValidationErrors: map[string]string{},
	Success:          false,
}

var ErrInvalidJson = &jsonResp{
	status:           422, // Unprocessable Entity
	ErrorMessage:     "invalid json",
	ValidationErrors: map[string]string{},
	Success:          false,
}

func NewError(statusCode int, err error) http.Handler {
	return &jsonResp{
		status:           statusCode,
		ErrorMessage:     err.Error(),
		ValidationErrors: map[string]string{},
		Success:          false,
	}
}

/*
func NewValidationErrorHandler(validationErrors map[string]error) http.Handler {

	m := make(map[string]string, len(validationErrors))

	for k, err := range validationErrors {
		m[k] = err.Error()
	}

	return &jsonResp{
		status:           422,
		ErrorMessage:     "validation error",
		ValidationErrors: m,
		Success:          false,
	}
}
*/

func (ve *validationError) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m := make(map[string]string, len(ve.validationErrors))

	for k, err := range ve.validationErrors {
		m[k] = err.Error()
	}

	jr := &jsonResp{
		status:           422,
		ErrorMessage:     "validation error",
		ValidationErrors: m,
		Success:          false,
	}

	jr.ServeHTTP(rw, req)
}

type validationError struct {
	validationErrors map[string]error
}

func (*validationError) Error() string {
	return "validation error"
}

type notFound struct{}

func (notFound) Error() string {
	return "not found"
}

type fieldNotAllowed struct{}

func (fieldNotAllowed) Error() string {
	return "field not allowed"
}

func (ne fieldNotAllowed) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	jr := &jsonResp{
		status:       422,
		ErrorMessage: "field not allowed",
		Success:      false,
	}

	jr.ServeHTTP(rw, req)
}
