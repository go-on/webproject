package wsi

import (
	"net/http"
)

// ColumnsMapper maps columns to fields
type ColumnsMapper interface {
	// MapColumns must map sql query columns to pointer of fields of the object.
	// This method is used by WriteJSON in order to do a
	// search query, write the results back to the provided fields that correspond to the columns
	// and writing an array of json objects that are the serialization of the object.
	// Therefor MapColumns must be a pointer method and must set the columns to field pointers.
	//
	// Example
	//
	// type Person struct {
	//	  Id         int
	//	  Name       string
	// }
	//
	// func (p *Person) MapColumns(colToField map[string]interface{}) {
	//	colToField["id"] = &p.Id
	//	colToField["name"] = &p.Name
	// }
	//
	MapColumns(colToField map[string]interface{})
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

// QueryFunc is a function that returns a Scanner (with the help of the Query function) for the given
// search parameters that are
//   - orderBy strings, such as "name ASC" or "id DESC"
//   - offset: number of skipped entries
//   - limit: limit of returned entries
//   - filter: any further filtering parameters
// How the values of the filter are used is up to the function
// type QueryFunc func(QueryOptions) (Scanner, error)

// ExecFunc is a function that runs the sql and writes to the responsewriter
// type ExecFunc func(map[string]interface{}, http.ResponseWriter)

type StreamEncoder interface {
	Encode(interface{}) error
	Finish()
}

type RequestDecoder interface {
	// decodes the given http request to the given interface.
	// must not close the request body
	Decode(*http.Request, interface{}) error
}
