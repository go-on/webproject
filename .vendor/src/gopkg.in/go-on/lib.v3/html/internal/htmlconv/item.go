package htmlconv

import "fmt"

type itemType int

const (
	_                           = iota
	itemError          itemType = 1 << iota // error
	itemEOF                                 // end of the file
	itemTagName                             // the html tag
	itemAttributeKey                        // an attribute
	itemAttributeValue                      // an attribute
	itemText                                // plain text inside a tag
	itemEntity                              // a html entity
	itemClass                               // a css class
	itemId                                  // a css id
	itemStyle                               // a css style
	itemBracketOpen
	itemBracketClosed
	itemEqualSign
	itemSingleQuote
	itemDoubleQuote
	itemSlash
	itemAnd
	itemSemicolon
	itemCommentStart
	itemCommentEnd
	itemComment
	itemDocTypeStart
	itemDocTypeEnd
	itemDocType
	itemExclamationMark
)

func (it itemType) String() (s string) {
	switch it {
	case itemError:
		s = "itemError"
	case itemEOF:
		s = "itemEOF"
	case itemTagName:
		s = "itemTagName"
	case itemAttributeKey:
		s = "itemAttributeKey"
	case itemAttributeValue:
		s = "itemAttributeValue"
	case itemText:
		s = "itemText"
	case itemEntity:
		s = "itemEntity"
	case itemClass:
		s = "itemClass"
	case itemId:
		s = "itemId"
	case itemStyle:
		s = "itemStyle"
	case itemBracketOpen:
		s = "itemBracketOpen"
	case itemBracketClosed:
		s = "itemBracketClosed"
	case itemEqualSign:
		s = "itemEqualSign"
	case itemSingleQuote:
		s = "itemSingleQuote"
	case itemDoubleQuote:
		s = "itemDoubleQuote"
	case itemSlash:
		s = "itemSlash"
	case itemCommentStart:
		s = "itemCommentStart"
	case itemCommentEnd:
		s = "itemCommentEnd"
	case itemComment:
		s = "itemComment"
	case itemDocType:
		s = "itemDocType"
	case itemDocTypeStart:
		s = "itemDocTypeStart"
	case itemDocTypeEnd:
		s = "itemDocTypeEnd"
	case itemExclamationMark:
		s = "itemExclamationMark"
	}
	return

}

type item struct {
	typ itemType
	val string
}

func (i item) String() (s string) {
	switch i.typ {
	case itemError:
		// s ="-itemError"
		s = i.val

	case itemEOF:
		s = "\nitemEOF"
	case itemTagName:
		s = "\nitemTagName" + fmt.Sprintf(": %#v", i.val)
	case itemAttributeKey:
		s = "\nitemAttributeKey" + fmt.Sprintf(": %#v", i.val)
	case itemAttributeValue:
		s = "\nitemAttributeValue" + fmt.Sprintf(": %#v", i.val)
	case itemText:
		s = "\nitemText" + fmt.Sprintf(": %#v", i.val)
	case itemEntity:
		s = "\nitemEntity" + fmt.Sprintf(": %#v", i.val)
	case itemClass:
		s = "\nitemClass" + fmt.Sprintf(": %#v", i.val)
	case itemId:
		s = "\nitemId" + fmt.Sprintf(": %#v", i.val)
	case itemStyle:
		s = "\nitemStyle" + fmt.Sprintf(": %#v", i.val)
	case itemBracketOpen:
		s = "\nitemBracketOpen" + fmt.Sprintf(": %#v", i.val)
	case itemBracketClosed:
		s = "\nitemBracketClosed" + fmt.Sprintf(": %#v", i.val)
	case itemEqualSign:
		s = "\nitemEqualSign" + fmt.Sprintf(": %#v", i.val)
	case itemSingleQuote:
		s = "\nitemSingleQuote" + fmt.Sprintf(": %#v", i.val)
	case itemDoubleQuote:
		s = "\nitemDoubleQuote" + fmt.Sprintf(": %#v", i.val)
	case itemSlash:
		s = "\nitemSlash" + fmt.Sprintf(": %#v", i.val)
	case itemCommentStart:
		s = "\nitemCommentStart" + fmt.Sprintf(": %#v", i.val)
	case itemCommentEnd:
		s = "\nitemCommentEnd" + fmt.Sprintf(": %#v", i.val)
	case itemComment:
		s = "\nitemComment" + fmt.Sprintf(": %#v", i.val)
	case itemDocType:
		s = "\nitemDocType" + fmt.Sprintf(": %#v", i.val)
	case itemDocTypeStart:
		s = "\nitemDocTypeStart" + fmt.Sprintf(": %#v", i.val)
	case itemDocTypeEnd:
		s = "\nitemDocTypeEnd" + fmt.Sprintf(": %#v", i.val)

	}

	return
}
