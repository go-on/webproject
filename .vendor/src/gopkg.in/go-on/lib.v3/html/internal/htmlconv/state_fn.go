package htmlconv

import "strings"

type stateFn func(*lexer) stateFn

func lexEntity(l *lexer) stateFn {
	d("lexEntity")
	for {

		// fmt.Printf("EEEE %#v...", l.input[l.pos:])

		if strings.HasPrefix(l.input[l.pos:], string(seminColon)) {
			if l.pos > l.start {
				// fmt.Printf("%#v...", string(r))
				l.emit(itemEntity)
				//l.backup()
				return lexSemicolon
			}
			return l.errorf("invalid entity")
		}

		switch r := l.next(); {
		case r == eof:
			return l.errorf("invalid entity - EOF reached")
		case r == seminColon:
			if l.pos > l.start {
				// l.pos += len(string(seminColon))
				l.emit(itemEntity)
				return lexSemicolon
			}
			// return lexText
			return l.errorf("invalid entity")

		case r == bracketOpen:
			if l.pos > l.start {
				l.emit(itemEntity)
			}
			return l.errorf("invalid entity")
			/*
					l.backup()
				if l.pos > l.start {
					// l.pos += len(string(seminColon))
					l.emit(itemEntity)
					return lexBracketOpen
				}
				return l.errorf("invalid entity")
			*/

		default:
			if !isLetter.MatchString(string(r)) {
				return l.errorf("invalid entity - no letter: %#v", string(r))
			}

			// fmt.Printf("%#v...", string(r))
		}
	}
}

func lexComment(l *lexer) stateFn {
	d("lexComment")
	for {
		if strings.HasPrefix(l.input[l.pos:], "-->") {
			if l.pos > l.start {
				l.emit(itemComment)
			}
			l.pos += len("-->")
			l.emit(itemCommentEnd)
			return lexText // next state
		}

		l.next()
	}
}

func lexDoctype(l *lexer) stateFn {
	d("lexDoctype")
	for {
		if strings.HasPrefix(l.input[l.pos:], ">") {
			if l.pos > l.start {
				l.emit(itemDocType)
			}
			l.pos += len(">")
			l.emit(itemDocTypeEnd)
			return lexText // next state
		}
		l.next()
	}
}

func lexText(l *lexer) stateFn {
	d("lexText")
	// inspace.MatchString(s)
	for {

		if strings.HasPrefix(l.input[l.pos:], string(bracketOpen)) {
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexBracketOpen // next state
		}

		switch r := l.next(); {
		case r == eof:
			if l.pos > l.start {
				l.emit(itemText)
			}
			l.emit(itemEOF)
			return nil
		case r == andSign:
			l.backup()
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexAnd
			// l.pos += len(string(andSign))
			// return lexEntity
		case r == bracketOpen:
			l.backup()
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexBracketOpen
			// default:
			//panic(fmt.Errorf("unknown, ...))
			// return l.errorf("unknown token in lexText", ...)
		}

		// l.emit(itemEOF)
		// return nil
	}
	return nil
}

func lexTagNameEnd(l *lexer) stateFn {
	d("lexTagNameEnd")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(bracketClosed)) {
			if l.pos > l.start {
				l.emit(itemTagName)
			} else {
				return l.errorf("tag name missing")
			}
			return lexBracketClosed
		}

		switch r := l.next(); {
		case r == eof:
			return l.errorf("unclosed tag")
		case isSpace(r):
			l.ignore()
		default:
			if !isLetter.MatchString(string(r)) {
				return l.errorf("invalid tag name")
			}
		}
	}
}

func lexSlash(l *lexer) stateFn {
	d("lexSlash")
	for {
		switch r := l.next(); {
		case r == eof:
			return l.errorf("unclosed tag")
		case isSpace(r):
			l.ignore()
		case r == bracketClosed:
			l.emit(itemBracketClosed)
			return lexText
		default:
			l.backup()
			return lexTagNameEnd
		}

	}
}

func lexExclamationMark(l *lexer) stateFn {
	d("lexExclamationMark")
	l.emit(itemExclamationMark)
	/*
		l.pos += len(string(exclamationMark))
	*/
	if strings.HasPrefix(l.input[l.pos:], "DOCTYPE") ||
		strings.HasPrefix(l.input[l.pos:], "doctype") {
		return lexDoctype
	}

	if strings.HasPrefix(l.input[l.pos:], "--") {
		l.pos += len("--")
		l.emit(itemCommentStart)
		return lexComment // next state
	}

	return l.errorf("<! must be followed by DOCTYPE or --")
}

func lexBracketOpen(l *lexer) stateFn {
	d("lexBracketOpen")
	l.pos += len(string(bracketOpen))
	l.emit(itemBracketOpen)
	for {
		switch r := l.next(); {
		case r == exclamationMark:
			return lexExclamationMark
		case r == eof:
			return l.errorf("unclosed tag")
		case isSpace(r):
			l.ignore()
		case r == slash:
			l.emit(itemSlash)
			return lexSlash
		default:
			l.backup()
			return lexTagNameStart
		}

	}
}

func lexAnd(l *lexer) stateFn {
	d("lexAnd")
	l.pos += len(string(andSign))
	l.emit(itemAnd)
	return lexEntity // now inside < >
}

func lexSemicolon(l *lexer) stateFn {
	d("lexSemicolon")
	l.pos += len(string(seminColon))
	l.emit(itemSemicolon)
	return lexText // now inside < >
}

func lexBracketClosed(l *lexer) stateFn {
	d("lexBracketClosed")
	l.pos += len(string(bracketClosed))
	l.emit(itemBracketClosed)
	return lexText // now inside < >
}

func lexTagNameStart(l *lexer) stateFn {
	d("lexTagNameStart")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(bracketClosed)) {
			if l.pos > l.start {
				l.emit(itemTagName)
			} else {
				return l.errorf("tag name missing")
			}
			return lexBracketClosed
		}

		switch r := l.next(); {
		case r == eof:
			return l.errorf("unclosed tag")
		case isSpace(r):
			l.backup()
			if l.pos > l.start {
				l.emit(itemTagName)
			} else {
				return l.errorf("tag name missing")
			}
			return lexTagAttributes
		case r == slash:
			n := l.peek()
			if n == bracketClosed {
				l.backup()
				if l.pos > l.start {
					l.emit(itemTagName)
				}
				l.next()
				return lexBracketClosed
			} else {
				return l.errorf("slash in opening tag")
			}
		case r == equalSign:
			return l.errorf("missing attribute key")
		default:
			if !isLetter.MatchString(string(r)) {
				return l.errorf("invalid tag name")
			}
		}
	}
}

func lexTagAttributes(l *lexer) stateFn {
	d("lexTagAttributes")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(bracketClosed)) {
			return lexBracketClosed
		}

		if strings.HasPrefix(l.input[l.pos:], string(slash)+string(bracketClosed)) {
			l.next()
			return lexBracketClosed
		}

		switch r := l.next(); {
		case r == eof:
			return l.errorf("unclosed tag")
		case isSpace(r):
			l.ignore()
		case r == bracketClosed:
			l.backup()
			return lexBracketClosed
		default:
			l.backup()
			if !isLetterOrDash.MatchString(string(r)) {
				// fmt.Println("invalid", string(r))
				return l.errorf("invalid attribute name")
			}

			return lexAttributeKey
		}
	}
}

func lexAttributeKey(l *lexer) stateFn {
	d("lexAttributeKey")
	for {
		if strings.HasPrefix(l.input[l.pos:], string(equalSign)) {
			if l.pos > l.start {
				l.emit(itemAttributeKey)
			}
			return lexEqualSign
		}
		switch r := l.next(); {
		case r == eof:
			return l.errorf("incomplete attribute and tag")
		case isSpace(r):
			l.backup()
			if l.pos > l.start {
				l.emit(itemAttributeKey)
			}
			return lexTagAttributes
		case r == bracketClosed:
			l.backup()
			if l.pos > l.start {
				l.emit(itemAttributeKey)
			}
			return lexText
		default:
			if !isLetterOrDash.MatchString(string(r)) {
				// fmt.Println("invalid", string(r))
				return l.errorf("invalid attribute name")
			}
		}
	}
}

func lexEqualSign(l *lexer) stateFn {
	d("lexEqualSign")
	l.pos += len(string(equalSign))
	l.emit(itemEqualSign)
	for {
		switch r := l.next(); {
		case r == eof:
			return l.errorf("incomplete attribute and tag")
		case isSpace(r):
			l.ignore()
		case r == singleQuote:
			l.emit(itemSingleQuote)
			return lexSingleQuote
		case r == doubleQuote:
			l.emit(itemDoubleQuote)
			return lexDoubleQuote
		default:
			return l.errorf("invalid attribute")
		}
	}
}

func lexSingleQuote(l *lexer) stateFn {
	d("lexSingleQuote")
	// if strings.HasPrefix(l.input[l.pos:], string(singleQuote)) {
	if l.pos > l.start {
		l.emit(itemAttributeValue)
	}
	l.pos += len(string(singleQuote))
	l.emit(itemSingleQuote)
	return lexTagAttributes
	// }

	for {
		switch r := l.next(); {
		case r == eof:
			return l.errorf("incomplete attribute value and tag")
		case r == doubleQuote:
			return l.errorf(`wrong attribute value quotes, starting with ' ending with "`)
		case r == equalSign:
			return l.errorf("= not allowed within attribute value")
		case r == singleQuote:
			l.backup()
			if l.pos > l.start {
				l.emit(itemAttributeValue)
			}
			l.pos += len(string(singleQuote))
			l.emit(itemSingleQuote)
			return lexTagAttributes
		case r == bracketClosed:
			return l.errorf("> not allowed within attribute value")
		}
	}
}

func lexDoubleQuote(l *lexer) stateFn {
	d("lexDoubleQuote")
	if strings.HasPrefix(l.input[l.pos:], string(doubleQuote)) {
		if l.pos > l.start {
			l.emit(itemAttributeValue)
		}
		l.pos += len(string(doubleQuote))
		l.emit(itemDoubleQuote)
		return lexTagAttributes
	}

	for {
		r := l.next()
		// fmt.Println("r ", string(r))
		// switch r := l.next(); {
		switch {
		// case r == eof:
		// return l.errorf("incomplete attribute value and tag (EOF)")
		case r == singleQuote:
			return l.errorf(`wrong attribute value quotes, starting with " ending with '`)
		case r == doubleQuote:
			l.backup()
			if l.pos > l.start {
				l.emit(itemAttributeValue)
			}
			l.pos += len(string(doubleQuote))
			l.emit(itemDoubleQuote)
			return lexTagAttributes
			//	case r == equalSign:
			//		return l.errorf("= not allowed within attribute value")
		case r == bracketClosed:
			return l.errorf("> not allowed within attribute value")
			// default:
			// fmt.Printf("%#v...", string(r))
		}
	}
}
