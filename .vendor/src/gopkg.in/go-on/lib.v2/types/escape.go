package types

import (
	"bytes"
	"strings"
)

const escapedChars = `&'<>"`

type writer interface {
	WriteString(string) (int, error)
}

//copied from http://golang.org/src/pkg/html/escape.go
// modified a bit (inlined escape function and remove error that weren't checked anyway)
func EscapeHTML(s string) string {
	i := strings.IndexAny(s, escapedChars)
	if i == -1 {
		return s
	}
	var w bytes.Buffer
	for i != -1 {
		w.WriteString(s[:i])
		var esc string
		switch s[i] {
		case '&':
			esc = "&amp;"
		case '\'':
			// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
			esc = "&#39;"
		case '<':
			esc = "&lt;"
		case '>':
			esc = "&gt;"
		case '"':
			// "&#34;" is shorter than "&quot;".
			esc = "&#34;"
		}
		s = s[i+1:]
		w.WriteString(esc)
		i = strings.IndexAny(s, escapedChars)
	}
	w.WriteString(s)
	return w.String()
}
