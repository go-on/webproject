package htmlconv

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

type parser struct {
	parser           parseFn
	item             item
	buffer           *bytes.Buffer
	openedTags       []string
	lastAttributes   map[string]string
	lastAttributeKey string
	*Parser
	*sync.Mutex
	finished         bool
	mustcloseDocType bool
}

var debugParser = false

type Parser struct {
	StripPrefixes bool
	TrimSpace     bool
}

func (pa Parser) Parse(s string) Stringer {
	finish := make(chan bool, 1)
	_, items := Lex("tester", s, finish)

	p := parse(&pa)

	for {
		select {
		case i := <-items:
			p.Handle(i)
		case _ = <-finish:
			close(finish)
			if p.mustcloseDocType {
				p.buffer.WriteString("\n),")
			}
			return p
		default:
		}
	}
	return nil
}

func pp(s string) {
	if debugParser {
		fmt.Println(s)
	}
}

func parse(p *Parser) *parser {
	var buf = bytes.Buffer{}
	return &parser{
		parser: parseText,
		buffer: &buf,
		Mutex:  &sync.Mutex{},
		Parser: p,
	}
}

func (p *parser) printAttributes() {
	attrs := []string{}
	id, hasId := p.lastAttributes["id"]

	if hasId {
		if p.StripPrefixes {
			fmt.Fprintf(p.buffer, "Id(%#v), ", id)
		} else {
			fmt.Fprintf(p.buffer, "html.Id(%#v), ", id)
		}
		delete(p.lastAttributes, "id")
	}

	class, hasClass := p.lastAttributes["class"]

	if hasClass {
		if p.StripPrefixes {
			fmt.Fprintf(p.buffer, "Class(%#v), ", class)
		} else {
			fmt.Fprintf(p.buffer, "html.Class(%#v), ", class)
		}
		delete(p.lastAttributes, "class")
	}

	for k, v := range p.lastAttributes {
		key := strings.TrimSpace(strings.ToLower(k))
		attrs = append(attrs, fmt.Sprintf("%#v", key), fmt.Sprintf("%#v", v))
	}

	if len(attrs) > 0 {
		if p.StripPrefixes {
			//fmt.Fprintf(p.buffer, "\nATTR(%#v, %#v),", p.lastAttributeKey, p.lastAttributeVal)
			fmt.Fprintf(p.buffer, "Attrs_(%s), ", strings.Join(attrs, ", "))
		} else {
			fmt.Fprintf(p.buffer, "html.Attrs_(%s),", strings.Join(attrs, ", "))
		}
	}

	p.lastAttributes = map[string]string{}
	p.lastAttributeKey = ""

	/*
		if tag.IsSelfclosing(strings.ToLower(p.openedTags[len(p.openedTags)-1])) {
			fmt.Fprintf(p.buffer, "),")
			p.openedTags = p.openedTags[:len(p.openedTags)-1]
		}
	*/
}

func (p *parser) Handle(i item) {
	// p.Lock()
	p.item = i
	if debugParser {

		fmt.Printf("handling: %s (%#v)\n", p.item.typ, p.item.val)
	}
	p.parser = p.parser(p)
	if debugParser {
		fmt.Printf("%s\n-----------\n\n", p.buffer.Bytes())
	}
	//

	if p.parser == nil {
		panic("unexpected nil parser")
	}
	// p.Unlock()
}
func (p *parser) String() string {
	if p.StripPrefixes {
		return fmt.Sprintf(
			`Elements(
			%s
		 )`,
			p.buffer.String(),
		)
	} else {
		return fmt.Sprintf(
			`html.Elements(
			%s
		 )`,
			p.buffer.String(),
		)
	}
}
