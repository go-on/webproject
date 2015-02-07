package replacer

import "net/http"

type Setter struct {
	t *Template
	m map[Placeholder][]byte
}

func newSetter(t *Template) *Setter {
	return &Setter{t, map[Placeholder][]byte{}}
}

func (i *Setter) SetString(ph Placeholder, s string) {
	i.m[ph] = []byte(s)
}

func (i *Setter) SetBytes(ph Placeholder, b []byte) {
	i.m[ph] = b
}

func (i *Setter) Bytes() []byte {
	return i.t.ReplaceBytes(i.m)
}

func (i *Setter) String() string {
	return string(i.t.ReplaceBytes(i.m))
}

func (i *Setter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var last int
	for _, place := range i.t.Places {
		w.Write(i.t.Original[last:place.Pos])
		w.Write(i.m[place.Placeholder])
		last = place.Pos
	}
	w.Write(i.t.Original[last:i.t.Length])
}
