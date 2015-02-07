package filter

import "fmt"

type filter struct {
	name   string
	params string
}

func (f filter) String() string {
	if f.params == "" {
		return fmt.Sprintf(" | %s", f.name)
	}
	return fmt.Sprintf(" | %s:%s", f.name, f.params)
}

func Currency(format string) filter { return filter{"date", fmt.Sprintf("'%s'", format)} }
func Date(format string) filter     { return filter{"date", fmt.Sprintf("'%s'", format)} }
func Json(format string) filter     { return filter{"json", ""} }
func LimitToString(l string) filter { return filter{"limitTo", l} }
func LimitToInt(l int) filter       { return filter{"limitTo", fmt.Sprintf("%d", l)} }
func LowerCase() filter             { return filter{"lowercase", ""} }
func Uppercase() filter             { return filter{"uppercase", ""} }
func NumberString(s string) filter  { return filter{"number", s} }
func NumberInt(i int) filter        { return filter{"number", fmt.Sprintf("%d", i)} }
func OrderBy(expr string) filter    { return filter{"orderBy", expr} }

type specialFilter struct {
	expression string
	comparator string
}

func (s specialFilter) String() string {
	return fmt.Sprintf(" | filter:%s:%s", s.expression, s.comparator)
}

func Filter(expression string, comparator string) specialFilter {
	return specialFilter{expression, comparator}
}
