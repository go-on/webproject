package menuhandler

import (
	"gopkg.in/go-on/lib.v3/internal/menu"
	"net/http"
)

type RequestMenu interface {
	Menu(*http.Request) *menu.Node
}
