package cssextract

type itemType int

const (
	_                  = iota
	itemError itemType = 1 << iota // error
	itemEOF                        // end of the file

	itemCurlyBraceOpen
	itemCurlyBraceClosed
	itemColon
	itemSemicolon
	itemStyleProperty
	itemStyleValue
	itemStar
	itemSelector
	itemDot
	itemHash

	itemClass
	itemId
	itemTag
	itemPseudo
	itemAttr
	itemAtSelector
	itemStyles

	itemSquareBracketOpen
	itemSquareBracketClosed
	itemCommentStart
	itemCommentEnd
	itemComment
)

func (it itemType) String() (s string) {
	switch it {

	case itemError:
		s = "itemError"
	case itemEOF:
		s = "itemEOF"
	case itemCurlyBraceOpen:
		s = "itemCurlyBraceOpen"
	case itemCurlyBraceClosed:
		s = "itemCurlyBraceClosed"
	case itemColon:
		s = "itemColon"
	case itemSemicolon:
		s = "itemSemicolon"
	case itemStyleProperty:
		s = "itemStyleProperty"
	case itemStyleValue:
		s = "itemStyleValue"
	case itemStar:
		s = "itemStar"
	case itemSelector:
		s = "itemSelector"
	case itemDot:
		s = "itemDot"
	case itemHash:
		s = "itemHash"
	case itemClass:
		s = "itemClass"
	case itemId:
		s = "itemId"
	case itemTag:
		s = "itemTag"
	case itemPseudo:
		s = "itemPseudo"
	case itemAttr:
		s = "itemAttr"
	case itemAtSelector:
		s = "itemAtSelector"
	case itemStyles:
		s = "itemStyles"
	case itemSquareBracketOpen:
		s = "itemSquareBracketOpen"
	case itemSquareBracketClosed:
		s = "itemSquareBracketClosed"
	case itemCommentStart:
		s = "temCommentStart"
	case itemCommentEnd:
		s = "itemCommentEnd"
	case itemComment:
		s = "itemComment"
	}
	return
}

type item struct {
	typ itemType
	val string
}
