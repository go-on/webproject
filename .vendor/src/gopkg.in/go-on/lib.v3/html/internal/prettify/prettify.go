package prettify

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"text/scanner"

	"gopkg.in/go-on/wrap.v2"
)

type xmlScanner struct {
	tabNo        int
	scanner      *scanner.Scanner
	inClosingTag bool
	result       bytes.Buffer
}

func newXmlScanner(xml string) *xmlScanner {
	x := &xmlScanner{scanner: &scanner.Scanner{}}
	x.scanner.Init(strings.NewReader(xml))
	return x
}

func (x *xmlScanner) next() rune {
	rn := x.scanner.Next()
	if rn == scanner.EOF {
		return scanner.EOF
	}

	switch rn {
	case '\n':
		x.result.WriteRune(' ')
	case '<':
		x.inClosingTag = x.scanner.Peek() == '/'
		if x.inClosingTag {
			if x.tabNo > 0 {
				x.tabNo--
			}
		}
		x.result.WriteRune(rn)
	case '>':
		nextTag := x.scanner.Peek() == '<'
		if nextTag && !x.inClosingTag {
			x.tabNo++
		}

		if nextTag {
			x.result.WriteString(">\n" + strings.Repeat("\t", x.tabNo))
		} else {
			x.result.WriteRune(rn)
		}

	default:
		x.result.WriteRune(rn)
	}
	return rn
}

var reg = regexp.MustCompile(`>\s*<`)

// Prettify indents the given xml/html string.
// Be warned that it may change your html in incompatible ways, since it inserts
// whitespace between tags but inline elements and inline blocks behave differently if there is whitespace
// around them. So only use it to prettify the output.
// Only tags that are siblings are indented. If there is text around, they stay the same.
func Prettify(xml string) string {
	xml = reg.ReplaceAllString(xml, `><`)
	sc := newXmlScanner(xml)
	tok := sc.next()
	for tok != scanner.EOF {
		tok = sc.next()
	}
	return sc.result.String()
}

type wrapper struct{}

var Wrap = wrapper{}

func (wr wrapper) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			buf := wrap.NewBuffer(rw)
			next.ServeHTTP(buf, req)
			pretty := Prettify(buf.BodyString())
			buf.FlushHeaders()
			buf.FlushCode()
			fmt.Fprint(rw, pretty)
		},
	)
}
