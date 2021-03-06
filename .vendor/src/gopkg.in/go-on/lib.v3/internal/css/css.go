package css

import (
	"fmt"
	// . "gopkg.in/go-on/lib.v3/html"
	. "gopkg.in/go-on/lib.v3/internal/selector"
	. "gopkg.in/go-on/lib.v3/types"
	"strings"
)

type Import string

func (ø Import) String() string {
	return fmt.Sprintf("@import %#v;\n", string(ø))
}

type nest RuleStruct

func Nest(xs ...interface{}) (r *nest) {
	rl, err := Rule(xs...)
	if err != nil {
		panic(err.Error())
	}
	n := nest(*rl)
	r = &n
	return
}

type RuleStruct struct {
	Comment     string
	Selector    Selector
	Styles      []Styler
	nested      []*RuleStruct
	children    []*RuleStruct
	cssStringer []CSSSer
	Parent      *RuleStruct
}

func Rule(xs ...interface{}) (r *RuleStruct, err error) {
	r = &RuleStruct{}
	r.Styles = []Styler{}
	r.nested = []*RuleStruct{}
	r.children = []*RuleStruct{}
	r.cssStringer = []CSSSer{}
	// r.Selector = SelectorString("")
	for _, x := range xs {

		switch v := x.(type) {
		case []Styler:
			// we don't want to handle *RuleStruct and *rules here,
			// since there would be different ways to handle them (Next, Embed, Compose )
			// and we don't want to have a implicit default way
			for _, fd := range v {
				r.Styles = append(r.Styles, fd)
			}

			//r.Style(v...)
			//case []FormatDefinition:
			//	r.Styles = append(r.Styles, v...)
			//case FormatDefinitions:
			//	fs := []FormatDefinition(v)
			//	r.Styles = append(r.Styles, fs...)
		case []Style:
			for _, fd := range v {
				r.Styles = append(r.Styles, fd)
			}
			// fs := []FormatDefinition(v)
			// r.Style(fs...)
		case Styler:
			r.Styles = append(r.Styles, v)
		case Style:
			r.Styles = append(r.Styles, v)
		case string:
			r.Comment = v
		case Comment:
			r.Comment = string(v)
		case *nest:
			rls := RuleStruct(*v)
			r.nested = append(r.nested, &rls)
		case *RuleStruct:
			r.children = append(r.children, v)
		case CSSSer:
			r.cssStringer = append(r.cssStringer, v)
		case Selector:
			r.Selector = v

		default:
			fmt.Printf("%T not handled\n", x)
			err = fmt.Errorf("%T is not an allowed type", x)
			return
		}
	}
	return
}

// for each selector, my selectors is prefixed and
// my rules are applied
func (ø *RuleStruct) ForEach(c SelectorAdder, sel ...Selector) (*RuleStruct, error) {
	comb := c.Add(ø.Selector)
	all := Each()
	for _, s := range sel {
		all.Add(comb.Add(s))
	}
	return Rule(all, ø.Styles)
}

func (ø *RuleStruct) String() string {
	styles := []string{}
	// fmt.Println(ø)
	for _, st := range ø.Styles {
		styles = append(styles, st.Style())
	}
	comment := ""
	if ø.Comment != "" {
		comment = fmt.Sprintf("/* %s */\n", ø.Comment)
	}
	var my string
	if ø.Selector != nil {
		strs := []string{}
		strs = append(strs, ø.Selector.Selector())
		nested := []string{}
		for _, nr := range ø.nested {
			ns := nr.String()
			nssp := strings.Split(ns, "\n")
			for _, nsi := range nssp {
				nested = append(nested, nsi)
			}
		}
		my = fmt.Sprintf("%s%s {\n\t%s\n\t%s\n}", comment, strings.Join(strs, ",\n"), strings.Join(styles, "\n\t"), strings.Join(nested, "\n\t"))
	} else {
		my = ""
	}
	all := []string{my}

	for _, cs := range ø.cssStringer {
		all = append(all, cs.CSS())
	}

	for _, child := range ø.children {
		if ø.Selector != nil {
			nu := child.Embed(ø.Selector)
			all = append(all, nu.String())
		} else {
			all = append(all, child.String())
		}
	}
	return strings.Join(all, "\n")
}

// adds given styles
func (ø *RuleStruct) Style(styles ...Style) *RuleStruct {
	for _, st := range styles {
		ø.Styles = append(ø.Styles, st)
	}
	return ø
}

func (ø *RuleStruct) Nest(xs ...interface{}) (i *RuleStruct, err error) {
	i, err = Rule(xs...)
	if err != nil {
		return
	}
	ø.nested = append(ø.nested, i)
	return
}

func (ø *RuleStruct) _inner(sa SelectorAdder, sel Selector, xs ...interface{}) (i *RuleStruct, err error) {
	i, err = Rule(xs...)
	if err != nil {
		return
	}
	i.Selector = sa.Add(ø.Selector).Add(sel)
	return
}

func (ø *RuleStruct) Descendant(sel Selector, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Descendant(), sel, xs...)
}

func (ø *RuleStruct) Child(sel Selector, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Child(), sel, xs...)
}

func (ø *RuleStruct) DirectFollows(sel Selector, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(DirectFollows(), sel, xs...)
}

func (ø *RuleStruct) Follows(sel Selector, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Follows(), sel, xs...)
}

func (ø *RuleStruct) Each(sel Selector, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Each(), sel, xs...)
}

// returns a copy
func (ø *RuleStruct) Copy() (newrule *RuleStruct) {
	newStyles := []Styler{}
	for _, st := range ø.Styles {
		newStyles = append(newStyles, st)
	}

	newnested := []*RuleStruct{}
	for _, st := range ø.nested {
		newnested = append(newnested, st)
	}

	newchildren := []*RuleStruct{}
	for _, st := range ø.children {
		newchildren = append(newchildren, st)
	}

	newrule = &RuleStruct{
		Comment:  ø.Comment,
		Styles:   newStyles,
		Selector: ø.Selector,
		nested:   newnested,
		children: newchildren,
		Parent:   ø.Parent,
	}
	return
}

// returns a copy that is embedded in the selector
func (ø *RuleStruct) Embed(selector Selector) (newrule *RuleStruct) {
	newrule = ø.Copy()
	newSelector := SelectorString(selector.Selector() + " " + ø.Selector.Selector())
	newrule.Selector = newSelector
	return
}

// returns a copy that is a composition of this rule with the styles
// of other rules
func (ø *RuleStruct) Compose(parents ...*RuleStruct) (newrule *RuleStruct) {
	newrule = ø.Copy()
	for _, parent := range parents {
		for _, st := range parent.Styles {
			newrule.Styles = append(newrule.Styles, st)
		}
	}
	return
}

type rules struct {
	rules []*RuleStruct
}

func Rules() *rules {
	r := []*RuleStruct{}
	return &rules{r}
}

func (ø *rules) Add(r *RuleStruct) {
	ø.rules = append(ø.rules, r)
}

func (ø *rules) New(xs ...interface{}) (r *RuleStruct, err error) {
	r, err = Rule(xs...)
	if err != nil {
		return
	}
	ø.Add(r)
	return
}

func (ø *rules) String() string {
	rules := []string{}

	for _, r := range ø.rules {
		rules = append(rules, r.String())
	}

	return strings.Join(rules, "\n\n")
}

// returns a copy with all rules embedded with selector
func (ø *rules) Embed(selector Selector) *rules {
	newRules := []*RuleStruct{}
	for _, r := range ø.rules {
		newRules = append(newRules, r.Embed(selector))
	}
	return &rules{newRules}
}

/*
type Css []Stringer

func (ø Css) String() string {
	rules := []string{}

	for _, r := range ø {
		rules = append(rules, r.String())
	}

	return strings.Join(rules, "\n\n")
}
*/

func Css(xs ...interface{}) *RuleStruct {
	r, ſ := Rule(xs...)
	if ſ != nil {
		panic(ſ.Error())
	}
	return r
}
