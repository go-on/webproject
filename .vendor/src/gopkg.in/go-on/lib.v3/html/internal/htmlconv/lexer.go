package htmlconv

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// lexer holds the state of the scanner
type lexer struct {
	name        string    // used only for error reports
	input       string    // the string being scanned
	start       int       // start position of this item
	pos         int       // current position in the input
	width       int       // width of the last rune read
	items       chan item // channel of scanned items
	state       stateFn
	line        int
	linepos     int
	linePrev    int
	lineposPrev int
}

func (l *lexer) Run(finish chan bool) {
	for state := lexText; state != nil; {
		state = state(l)
	}
	// fmt.Println("closing run")
	close(l.items)
	finish <- true
}

// emit passes an item back to the client
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

var eof = rune('∎')

// var eof = rune('X')

// returns the next rune in the input
func (l *lexer) next() (rune_ rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune_, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	l.linePrev = l.line
	l.lineposPrev = l.linepos
	if rune_ == '\n' {
		l.line++
		l.linepos = 0
	} else {
		l.linepos++
	}

	return rune_
}

func (l *lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune
// can be called only once per call of next
func (l *lexer) backup() {
	rune_, _ := utf8.DecodeRuneInString(l.input[l.pos:])
	if rune_ == '\n' {
		l.line--
	}
	l.linepos = l.lineposPrev
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	start := l.pos - 5
	if start < 0 {
		start = 0
	}

	end := l.pos + 5

	if end > len(l.input) {
		end = len(l.input)
	}

	fmt.Printf(
		"Error in line %d at position %d: %s\ncontext:\n%s\n",
		l.line+1,
		l.linepos+1,
		fmt.Sprintf(format, args...),
		l.input[start:end],
	)

	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func Lex(name, input string, finish chan bool) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.Run(finish)
	return l, l.items
}
