package wrapstesting

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/go-on/wrap.v2"
)

type logger struct{ *log.Logger }

func (l *logger) ServeHTTPNext(in http.Handler, w http.ResponseWriter, r *http.Request) {
	checked := wrap.NewPeek(w, func(ck *wrap.Peek) bool {
		ck.FlushHeaders()
		ck.FlushCode()
		return true
	})
	// l.Logger.Printf("ResponseWriter: %#v\nRequest: %#v\n", w, r)
	requestHeaders := fmt.Sprintf("%v", r.Header)

	// buf := helper.NewResponseBuffer(w)

	in.ServeHTTP(checked, r)

	if checked.HasChanged() {
		l.Printf(`
-- REQUEST --
%s %s
HEADERS
%s
-- RESPONSE --
STATUS CODE: %d
HEADERS
%s
`, r.Method, r.URL.Path, requestHeaders, checked.Code, checked.Header())
	}

}

func (l *logger) Wrap(next http.Handler) http.Handler {
	return wrap.NextHandler(l).Wrap(next)
}

func LOGGER(prefix string) wrap.Wrapper {
	return Logger(log.New(os.Stderr, prefix+" ", log.LstdFlags))
}

func Logger(l *log.Logger) wrap.Wrapper {
	return &logger{l}
}
