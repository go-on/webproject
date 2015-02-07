package placeholder

import (
	"io"
	"net/http"
)

type Placeholder interface {
	Name() string
	Set(interface{}) Setter
	Setf(format string, vals ...interface{}) Setter
	SetString() string
	WriteTo(w io.Writer) (int64, error)
}

type Setter interface {
	io.WriterTo
	Name() string
	SetString() string
}

type PlaceholderHandler interface {
	Placeholder
	http.Handler
}

type placeholderHandler struct {
	Placeholder
	http.Handler
}

func NewPlaceholderHandler(ph Placeholder, h http.Handler) PlaceholderHandler {
	return placeholderHandler{ph, h}
}
