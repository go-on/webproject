package selector

import (
	. "gopkg.in/go-on/lib.v3/types"
)

type combinator struct {
	selectors []Selector
	operator  string
}

func (ø combinator) Selector() string {
	var s string
	var pre string
	for _, sel := range ø.selectors {
		s += pre + sel.Selector()
		pre = ø.operator
	}
	return s
}

func (ø combinator) Add(s Selector) SelectorAdder {
	return combinator{append(ø.selectors, s), ø.operator}
}

type SelectorAdder interface {
	Selector() string
	Add(Selector) SelectorAdder
}

// F element descendant of an E element
func Descendant(selectors ...Selector) combinator { return combinator{selectors, " "} }

// F element child of an E element
func Child(selectors ...Selector) combinator { return combinator{selectors, " > "} }

// F element immediately preceded by an E element
func DirectFollows(selectors ...Selector) combinator { return combinator{selectors, " + "} }

// F element preceded by an E element
func Follows(selectors ...Selector) combinator { return combinator{selectors, " ~ "} }

// for each given selector the rules apply
func Each(selectors ...Selector) combinator { return combinator{selectors, ",\n"} }

/*
var ParentSelector = SelectorString("&")

func Super(sel ...Selector) Selector {
	return Selectors(ParentSelector, sel...)
}
*/

func Context(ctx Selector, inner1 Selector, inner ...Selector) (r Selector) {
	var s string
	inner = append(inner, inner1)
	var pre string
	for _, i := range inner {
		s += pre + ctx.Selector() + " " + i.Selector()
		pre = ",\n"
	}
	return SelectorString(s)
}

func ContextString(ctx string, inner1 Selector, inner ...Selector) (r Selector) {
	return Context(SelectorString(ctx), inner1, inner...)
}
