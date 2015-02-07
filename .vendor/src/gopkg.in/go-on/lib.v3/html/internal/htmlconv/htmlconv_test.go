package htmlconv

import (
	"fmt"
	"strings"
	"testing"
)

var wsStripper = strings.NewReplacer(
	"\n", "",
	"\t", "",
	" ", "",
)

var tests = map[string]string{

	`<div></div>`:                      `DIV()`,
	`<div class="center"></div>`:       `DIV(Class("center"),)`,
	`<div><a href="#">click</a></div>`: `DIV(A(Attrs_("href","#"),"click",),)`,
	`<div>hello&nbsp;world</div>`:      `DIV("hello",E_nbsp,"world",)`,
	`<h1>hello&nbsp;world</h1>`:        `H1("hello",E_nbsp,"world",)`,
	`<hr/>`:                            `HR()`,
	`<hr/><hr/>`:                       `HR(),HR()`,
	`<hr><p><hr></p><div><br></div>`:   `HR(),P(HR(),),DIV(BR(),)`,

	`<!--<div>hello&nbsp;world</div>-->`: `Comment("<div>hello&nbsp;world</div>")`,

	// since empty space is stripped away via testing, we take the dash - instead of empty space here
	`<!DOCTYPE-html>`: `NewDocType("<!DOCTYPE-html>",)`,
	`<!DOCTYPE-html>
							   <html lang="de">
							    <head></head>
							    <body class="main">
							    </body>
							   </html>`: `NewDocType("<!DOCTYPE-html>",HTML(Attrs_("lang","de"),HEAD(),BODY(Class("main"),),),)`,

	`<meta http-equiv="X-UA-Compatible">`: `META(Attrs_("http-equiv","X-UA-Compatible"),)`,

	`<img src="../assets/js/holder.js/300x250/ink/auto/text: MREC PUB" />`: `IMG(Attrs_("src","../assets/js/holder.js/300x250/ink/auto/text: MREC PUB"),)`,
}

var p = Parser{TrimSpace: true, StripPrefixes: true}

func TestAll(t *testing.T) {
	for in, out := range tests {
		out = fmt.Sprintf("Elements(%s,)", out)
		res := wsStripper.Replace(p.Parse(in).String())
		out = wsStripper.Replace(out)

		if res != out {
			t.Errorf("failed to convert %#v, expected %#v, got %#v", in, out, res)
		}
	}

}
