package main

import (
	"fmt"
	"gopkg.in/metakeule/fmtdate.v1"
	"net/http"
	"time"

	. "gopkg.in/go-on/lib.v3/html"
	. "gopkg.in/go-on/lib.v3/types"
)

func path(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, req.URL.String())
}

func datetime(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, fmtdate.Format("hh:mm:ss", time.Now()))
}

func main() {
	layout := HTML5(
		HTML(
			HEAD(TITLE("hi nadja")),
			BODY(
				H1("<die url>", path),
				AHref("http://abc.de", "Abc.de"),
				DIV(
					Id("tester"),

					Style{"background-color", "yellow"},
					datetime,
				),
				PRE("<h1>h1</h1>"),
				SCRIPT(`alert("</script>");`),
			),
		),
	)

	http.ListenAndServe("localhost:7878", layout)
}
