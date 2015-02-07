package cssextract

import (
	"fmt"
	"regexp"
	"strings"
)

/*
modelled after Lexical Scanning in Go - (Rob Pike), http://www.youtube.com/watch?v=HxaD_trXwRE
*/

var (
	curlyBraceOpen      = '{'
	curlyBraceClosed    = '}'
	dot                 = '.'
	colon               = ':'
	semiColon           = ';'
	kleenestar          = '*'
	at                  = '@'
	hash                = '#'
	squareBracketOpen   = '['
	squareBracketClosed = ']'
	slash               = '/'

	isOnlySpace                = regexp.MustCompile(`^\s+$`)
	isLetter                   = regexp.MustCompile("[a-zA-Z]")
	isLetterOrDash             = regexp.MustCompile("[-a-zA-Z]")
	isLetterOrDashOrUnderScore = regexp.MustCompile("[-_a-zA-Z0-9]")

	lexDebugging    = false
	parserDebugging = false
)

func isSpace(r rune) bool { return strings.IndexRune(" \n\t", r) >= 0 }

type Stringer interface {
	String() string
}

func debugLex(s string) {
	if lexDebugging {
		fmt.Printf("#STATE %s#\n", s)
	}
}

type Parser struct {
	Classes []string
	Ids     []string
}

func hasClass(p *Parser, class string) bool {
	for _, c := range p.Classes {
		if c == class {
			return true
		}
	}
	return false
}

func hasId(p *Parser, id string) bool {
	for _, i := range p.Ids {
		if i == id {
			return true
		}
	}
	return false
}

func Parse(s string) (p *Parser) {
	p = &Parser{}
	finish := make(chan bool, 1)

	_, items := Lex("test", s, finish)

	for {
		select {
		case i := <-items:
			if parserDebugging {
				fmt.Printf("%s: %#v\n", i.typ.String(), i.val)
			}

			switch i.typ {
			case itemId:
				if !hasId(p, i.val) {
					p.Ids = append(p.Ids, i.val)
				}
			case itemClass:
				if !hasClass(p, i.val) {
					p.Classes = append(p.Classes, i.val)
				}
			}

			// fmt.Println(i.String())
			// p.Handle(i)
		case _ = <-finish:
			close(finish)
			return
		default:
		}
	}

	return
}
