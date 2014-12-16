package menuhandler

import (
	"gopkg.in/go-on/lib.v2/internal/menu"
	"net/http"
)

type RequestMenu interface {
	Menu(*http.Request) *menu.Node
}
