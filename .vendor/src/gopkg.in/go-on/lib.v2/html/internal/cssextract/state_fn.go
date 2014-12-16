package cssextract

import "strings"

type stateFn func(*lexer) stateFn

func lexText(l *lexer) stateFn {
	debugLex("lexText")
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(itemEOF)
			return nil
		case r == at:
			return lexAtDeclaration
		case isSpace(r):
			l.ignore()
		case r == '/':
			if l.peek() == kleenestar {
				l.next()
				l.emit(itemCommentStart)
				return lexCommentText
			} else {
				return l.errorf("/ not allowed here")
			}
		default:
			l.backup()
			return lexSelector
		}
	}
	return nil
}

func lexCommentText(l *lexer) stateFn {
	debugLex("lexCommentText")
	for {
		if strings.HasPrefix(l.input[l.pos:], "*/") {
			if l.pos > l.start {
				l.emit(itemComment)
			}
			l.pos += len(string("*/"))
			l.emit(itemCommentEnd)
			return lexText
		}
		l.next()
	}
}

func lexCommentStyles(l *lexer) stateFn {
	debugLex("lexCommentStyles")
	for {
		if strings.HasPrefix(l.input[l.pos:], "*/") {
			if l.pos > l.start {
				l.emit(itemComment)
			}
			l.pos += len(string("*/"))
			l.emit(itemCommentEnd)
			return lexStyles
		}
		l.next()
	}
}

func lexAtDeclaration(l *lexer) stateFn {
	debugLex("lexAtDeclaration")

	if len(l.input) > l.pos+5 {
		tempLower := strings.ToLower(l.input[l.pos : l.pos+5])

		if tempLower == "media" {
			return lexAtMedia
		}
	}

	for {
		if strings.HasPrefix(l.input[l.pos:], string(curlyBraceOpen)) {
			if l.pos > l.start {
				l.emit(itemAtSelector)
			}
			l.pos += len(string(curlyBraceOpen))
			l.emit(itemCurlyBraceOpen)
			return lexStyles
		}

		l.next()
	}
}

func lexAtMedia(l *lexer) stateFn {
	debugLex("lexAtMedia")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(curlyBraceOpen)) {
			if l.pos > l.start {
				l.emit(itemAtSelector)
			}
			l.pos += len(string(curlyBraceOpen))
			l.emit(itemCurlyBraceOpen)
			return lexText
		}
		l.next()
	}
}

func lexStyles(l *lexer) stateFn {
	debugLex("lexStyles")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(curlyBraceClosed)) {
			if l.pos > l.start {
				l.emit(itemStyles)
			}
			l.pos += len(string(curlyBraceClosed))
			l.emit(itemCurlyBraceClosed)
			return lexText
		}

		if strings.HasPrefix(l.input[l.pos:], "/*") {
			if l.pos > l.start {
				l.emit(itemCommentStart)
			}
			l.pos += len(string("/*"))
			l.emit(itemCommentStart)
			return lexCommentStyles
		}

		l.next()
	}
}

func lexAttribute(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], string(squareBracketClosed)) {
			if l.pos > l.start {
				l.emit(itemAttr)
			}
			l.pos += len(string(squareBracketClosed))
			l.emit(itemSquareBracketClosed)
			return lexSelector
		}
		switch r := l.next(); {
		case r == eof:
			l.emit(itemEOF)
			return nil
		case r == squareBracketClosed:
			l.backup()
			if l.pos > l.start {
				l.emit(itemAttr)
			}
			l.pos += len(string(squareBracketClosed))
			l.emit(itemSquareBracketClosed)
			return lexSelector
		}
	}
}

func lexSelector(l *lexer) stateFn {
	debugLex("lexSelector")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(curlyBraceOpen)) {
			if l.pos > l.start {
				l.emit(itemSelector)
			}
			l.pos += len(string(curlyBraceOpen))
			l.emit(itemCurlyBraceOpen)
			return lexStyles
		}
		switch r := l.next(); {
		case r == eof:
			l.emit(itemEOF)
			return nil
		case r == dot:
			l.backup()
			if l.pos > l.start {
				l.emit(itemSelector)
			}
			l.pos += len(string(dot))
			l.emit(itemDot)
			return lexClass
		case r == squareBracketOpen:
			l.backup()
			if l.pos > l.start {
				l.emit(itemSelector)
			}
			l.pos += len(string(squareBracketOpen))
			l.emit(itemSquareBracketOpen)
			return lexAttribute

		case r == hash:
			l.backup()
			if l.pos > l.start {
				l.emit(itemSelector)
			}
			l.pos += len(string(hash))
			l.emit(itemHash)
			return lexId
		case isSpace(r):
			l.ignore()
		case r == curlyBraceOpen:
			if l.pos > l.start {
				l.emit(itemSelector)
			}
			return lexStyles
		}
	}
}

func lexClass(l *lexer) stateFn {
	debugLex("lexClass")
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(itemEOF)
			return nil
		case r == curlyBraceOpen:

		default:
			if !isLetterOrDashOrUnderScore.MatchString(string(r)) {
				l.backup()
				if l.pos > l.start {
					l.emit(itemClass)
				}
				return lexSelector
			}
		}
	}
}

func lexId(l *lexer) stateFn {
	debugLex("lexId")
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(itemEOF)
			return nil
		case r == curlyBraceOpen:

		default:
			if !isLetterOrDashOrUnderScore.MatchString(string(r)) {
				l.backup()
				if l.pos > l.start {
					l.emit(itemId)
				}
				return lexSelector
			}
		}
	}
}
