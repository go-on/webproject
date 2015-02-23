package wsi

import (
	"net/http"
)

// Scanner is a more comfortable scanner that works similar to sql.Rows
type Scanner interface {

	// Next should return false if there are no rows left or if any error happened before
	Next() bool

	// Scan allows scanning by column name instead of column position
	// If an error did happen, every successing call to Scan should return that error without doing any scanning
	Scan(vals ...interface{}) error

	// Columns returns the columns of the query
	Columns() []string

	// Error should return the first error that did happen
	Error() error

	// should close the scan
	Close() error
}

// Validater is a fallback/default validater for POST, PUT and PATCH requests
type Validater interface {
	Validate() map[string]error
}

// POSTValidater validates data of POST requests
type POSTValidater interface {
	ValidatePOST() map[string]error
}

// PUTValidater validates data of PUT requests
type PUTValidater interface {
	ValidatePUT() map[string]error
}

// PATCHValidater validates data of PATCH requests
type PATCHValidater interface {
	ValidatePATCH() map[string]error
}

type StreamEncoder interface {
	Encode(interface{}) error
	Finish()
}

type RequestDecoder interface {
	// decodes the given http request to the given interface.
	// must not close the request body
	Decode(*http.Request, interface{}) error
}
