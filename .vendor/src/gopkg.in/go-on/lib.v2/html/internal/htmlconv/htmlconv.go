package htmlconv

import (
	"regexp"
	"strings"
)

/*
Converter assumes, input is valid html file in utf8 charset with no byte order marker

modelled after Lexical Scanning in Go - (Rob Pike), http://www.youtube.com/watch?v=HxaD_trXwRE
*/

var bracketOpen = '<'
var bracketClosed = '>'
var slash = '/'
var equalSign = '='
var singleQuote = '\''
var doubleQuote = '"'
var andSign = '&'
var seminColon = ';'
var exclamationMark = '!'

func isSpace(r rune) bool {
	return strings.IndexRune(" \n\t", r) >= 0
}

type Stringer interface {
	String() string
}

/*
	TODO:

	if all text is just space, i.e. \n\t\s, replace it with just one space
*/

var inspace = regexp.MustCompile(`^\s+$`)
var isLetter = regexp.MustCompile("[a-zA-Z0-9]")
var isLetterOrDash = regexp.MustCompile("[-a-zA-Z0-9]")

func d(s string) {
	// fmt.Printf("STATE %#v\n", s)
}
