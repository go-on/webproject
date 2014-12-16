package wrapstesting

import (
	"fmt"
	"net/http"

	"gopkg.in/go-on/wrap.v2"
)

// ErrorWriter has a method WriteError to write error information to ResponseWriters
type ErrorWriter interface {

	// WriteError is like a http.Handler but gets as additional argument the error that happened
	WriteError(http.ResponseWriter, *http.Request, error)
}

/*
// ErrorResponse captures a response with an error
type ErrorResponse struct {
	*helper.ResponseBuffer
	Error error
}

// HandleError fulfills the github.com/go-on/queue.ErrHandler interface
// and always returns an error to stop the queue
func (e *ErrorResponse) HandleError(in error) (out error) {
	e.Error = in
	return in
}
*/

// HTTPStatusError is an error that is based on what was written to a http.ResponseWriter
// any status code >= 400 is considered an error
type HTTPStatusError struct {

	// Code has the http status code that was written to the http.ResponseWriter that was written to
	// it is always >= 400
	Code int

	// Header has the Header of the http.ResponseWriter that was written to
	Header http.Header
}

// Error fulfills the error interface
func (h HTTPStatusError) Error() string {
	return fmt.Sprintf("HTTP Status Error: Code %d", h.Code)
}

// errorWrapper is a github.com/go-on/wrap.Wrapper based on a ErrorWriter
type errorWrapper struct {
	ErrorWriter
}

// NewErrorWrapper creates a new "github.com/go-on/wrap.Wrapper
// for the given ErrorWriter
func NewErrorWrapper(errHandler ErrorWriter) wrap.Wrapper {
	return &errorWrapper{errHandler}
}

// Wrap returns an http.Handler that calls next.ServeHTTP with an ErrorResponse
// as ResponseWriter.
// The next http.Handler might make a type assertion to the
// github.com/go-on/queue.ErrHandler interface, in order to use the ErrorResponse as
// an error handler for a queue.
// Or the next handler might write a status code >= 400 and an error message to the body
// of the ErrorResponse to communicate the error.
// In both cases the ErrorWriter is called with the corresponding error, which in the latter case
// is a HTTPStatusError.
// That way the ErrorWriter might type switch on the given error to determine the correct action.
// Wrap fulfills the github.com/go-on/wrap.Wrapper interface.
func (e *errorWrapper) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//checked := helper.NewCheckedResponseWriter(w, func(ck *helper.CheckedResponseWriter) bool {
		bodywritten := false
		checked := wrap.NewPeek(w, func(ck *wrap.Peek) bool {
			if ck.Code >= 400 {
				e.WriteError(w, r, HTTPStatusError{ck.Code, ck.Header()})
				return true
			}
			ck.FlushHeaders()
			ck.FlushCode()
			bodywritten = true
			// ck.WriteHeadersTo(w)
			// ck.WriteCodeTo(w)
			return true
		})

		next.ServeHTTP(checked, r)

		if !bodywritten {
			if checked.Code >= 400 {
				e.WriteError(w, r, HTTPStatusError{checked.Code, checked.Header()})
				return
			}
		}

		checked.FlushMissing()
	})
}

// ValidationError is an interface for validation errors.
// Since validation often happens on a set of values, it is
// not possible to express which value has which error with the
// standard error interface.
// Therefor ValidationError requires an additional method ValidationErrors
// that returns an association of value names to
// error slices.
type ValidationError interface {
	Error() string

	// ValidationErrors returns a map that associates value names to
	// the corresponding error slices. The empty string
	// is used to refer to the complete value set as an entity.
	ValidationErrors() map[string][]error
}

// Validatable has a Validate method that returns a ValidationError error in case
//of an error, or nil otherwise.
type Validatable interface {

	// Validate does a validation and returns nil, if the validation was successful, otherwise
	// a ValidationError
	Validate() ValidationError
}
