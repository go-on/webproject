package htmlfat

import (
	. "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/internal/fat"
	"strings"
	"testing"
)

func stripWhiteSpace(in string) string {
	return strings.Replace(strings.Replace(strings.Replace(in, "\n", "", -1), "\t", "", -1), " ", "", -1)
}

type Person struct {
	FirstName *fat.Field `type:"string text"`
	LastName  *fat.Field `type:"string"`
	Vita      *fat.Field `type:"string html"`
}

var PERSON = fat.Proto(&Person{}).(*Person)

func NewPerson() *Person { return fat.New(PERSON, &Person{}).(*Person) }

func init() {
	Register(PERSON)
}

func TestHTMLFat(t *testing.T) {

	ul := UL("\n",
		LI(Placeholder(PERSON.FirstName)), "\n",
		LI(Placeholder(PERSON.LastName)), "\n",
		LI(Placeholder(PERSON.Vita)), "\n",
	)

	details := ul.Template("details")

	paul := NewPerson()
	paul.FirstName.Set("<Pa>ul")
	paul.LastName.Set("Pa<n>zer")
	paul.Vita.Set("<p>hier die vita</p>")

	expected := `
		<ul>
			<li>&lt;Pa&gt;ul</li>
			<li>Pa&lt;n&gt;zer</li>
			<li><p>hier die vita</p></li>
		</ul>
	`

	expected = stripWhiteSpace(expected)

	got := stripWhiteSpace(details.Replace(Setters(paul)...).String())

	if got != expected {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}
