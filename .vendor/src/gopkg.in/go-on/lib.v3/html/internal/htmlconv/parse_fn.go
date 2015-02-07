package htmlconv

import (
	"fmt"
	"strings"

	// tag "gopkg.in/go-on/lib.v3/html"
	tagProps "gopkg.in/go-on/lib.v3/html/internal/tagproperties"
	//ht "gopkg.in/go-on/lib.v3/html/internal/element"
)

func parseBracketClose(p *parser) parseFn {
	pp("parseBracketClose")
	return parseText
}

func parseTagClose(p *parser) parseFn {
	pp("parseTagClose")
	switch p.item.typ {
	case itemBracketClosed:
		return parseText
	case itemTagName:
		last := p.openedTags[len(p.openedTags)-1]
		closing := strings.TrimSpace(strings.ToUpper(p.item.val))
		if last != closing {
			panic(fmt.Sprintf("opening tag %s and closing tag %s do not match", last, closing))
		}
		p.openedTags = p.openedTags[:len(p.openedTags)-1]
		fmt.Fprint(p.buffer, "\n),")
		return parseBracketClose
	}
	return nil
}

/*
func parseDoctype(p *parser) parseFn {
	pp("parseDoctype")
	//inner := fmt.Sprintf("<!DOCTYPE %s>\n", p.item.val)
	p.mustcloseDocType = true
	if p.StripPrefixes {
		fmt.Fprintf(p.buffer, "\nNewDocType(%#v),", p.item.val)
	} else {
		fmt.Fprintf(p.buffer, "\ntag.NewDocType(%#v),", p.item.val)
	}
	return parseDoctypeEnd
}
*/

func parseDoctypeEnd(p *parser) parseFn {
	pp("parseDoctypeEnd")
	return parseText
}

func parseComment(p *parser) parseFn {
	pp("parseComment")
	if p.StripPrefixes {
		fmt.Fprintf(p.buffer, "\nComment(%#v),", p.item.val)
	} else {
		fmt.Fprintf(p.buffer, "\nhtml.Comment(%#v),", p.item.val)
	}
	return parseCommentEnd
}

func parseCommentEnd(p *parser) parseFn {
	pp("parseCommentEnd")
	return parseText
}

func parseAttributeKey(p *parser) parseFn {
	pp("parseAttributeKey")
	switch p.item.typ {
	case itemDoubleQuote:
	case itemSingleQuote:
	case itemEqualSign:
	case itemBracketClosed:
		p.printAttributes()
		if tagProps.IsSelfclosing(strings.ToLower(p.openedTags[len(p.openedTags)-1])) {
			fmt.Fprintf(p.buffer, "\n),")
			p.openedTags = p.openedTags[:len(p.openedTags)-1]
		}
		return parseText
	case itemAttributeValue:
		p.lastAttributes[p.lastAttributeKey] = p.item.val
		p.lastAttributeKey = ""
	case itemAttributeKey:
		p.lastAttributeKey = p.item.val
		p.lastAttributes[p.item.val] = ""
		// p.printAttr()
		// p.lastAttributeVal = ""
	}
	return parseAttributeKey
}

func parseExclamationMark(p *parser) parseFn {
	pp("parseExclamationMark")
	switch p.item.typ {
	case itemCommentStart:
		return parseComment
	case itemDocType:
		p.mustcloseDocType = true
		val := fmt.Sprintf(`<!%s>`, p.item.val)
		if p.StripPrefixes {
			fmt.Fprintf(p.buffer, "\nNewDocType(%#v,", val)
		} else {
			fmt.Fprintf(p.buffer, "\ntag.NewDocType(%#v,", val)
		}
		return parseDoctypeEnd
		// return parseDoctype
	default:
		panic("unexpected token " + p.item.typ.String())
	}
}

func parseTag(p *parser) parseFn {
	pp("parseTag")
	switch p.item.typ {
	case itemAttributeKey:
		p.lastAttributeKey = p.item.val
		p.lastAttributes[p.item.val] = ""
		return parseAttributeKey
	case itemExclamationMark:
		return parseExclamationMark
	case itemSlash:
		return parseTagClose
	case itemBracketClosed:
		if tagProps.IsSelfclosing(strings.ToLower(p.openedTags[len(p.openedTags)-1])) {
			fmt.Fprintf(p.buffer, "\n),")
			p.openedTags = p.openedTags[:len(p.openedTags)-1]
		}
		return parseText
	case itemTagName:
		openTag := strings.TrimSpace(strings.ToUpper(p.item.val))

		if p.StripPrefixes {
			fmt.Fprintf(p.buffer, "\n\n%s(\n", openTag)
		} else {
			fmt.Fprintf(p.buffer, "\n\ntag.%s(\n", openTag)
		}

		p.openedTags = append(p.openedTags, openTag)
		p.lastAttributes = map[string]string{}
		//		return parseTagStart
		return parseTag
	}
	return nil
}

func parseText(p *parser) parseFn {
	pp("parseText")
	// fmt.Printf("%T\n", i.typ)
	//fmt.Print(i.String())
	//fmt.Printf("%s: %#v\n", p.item.typ.String(), p.item.val)

	switch p.item.typ {
	case itemBracketOpen:
		// return parseTagStart
		return parseTag
	case itemText:
		// p.buffer.WriteString()
		if inspace.MatchString(p.item.val) {
			if !p.TrimSpace {
				fmt.Fprintf(p.buffer, "\n\" \",")
			}
		} else {
			if p.TrimSpace {
				fmt.Fprintf(p.buffer, "\n%#v,", strings.TrimSpace(p.item.val))
			} else {
				fmt.Fprintf(p.buffer, "\n%#v,", p.item.val)
			}

		}
		return parseText
	case itemEOF:
		p.finished = true
		return parseText
	case itemAnd:
		return parseText
	case itemSemicolon:
		return parseText
	case itemEntity:
		if p.StripPrefixes {
			fmt.Fprintf(p.buffer, "\nE_%s,", p.item.val)
		} else {
			fmt.Fprintf(p.buffer, "\nentity.E_%s,", p.item.val)
		}

		// fmt.Printf("entity: %#v\n", p.item.val)
		return parseText
	case itemError:
		fmt.Printf("Error: %s\n", p.item.val)
		p.finished = true
		return parseText
	default:
		if !p.finished {
			panic("unexpected token " + p.item.typ.String())
		}
		return parseText
	}
}

type parseFn func(*parser) parseFn
