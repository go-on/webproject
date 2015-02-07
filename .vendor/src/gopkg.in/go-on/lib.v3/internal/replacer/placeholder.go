package replacer

import "fmt"

type Placeholder string

func (p Placeholder) String() string {
	return fmt.Sprintf("\u2423\u2023%s\u220e\u2423", string(p))
}

func (p Placeholder) Name() string {
	return string(p)
}

// shortcut
func P(s string) Placeholder {
	return Placeholder(s)
}
