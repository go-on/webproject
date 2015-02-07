package benchmark

import (
	"bytes"
	"fmt"
	"github.com/go-on/template/placeholder"
	"html"
	tt "html/template"
	"strings"
	"testing"

	"github.com/go-on/replacer"
	"github.com/go-on/template"
)

func html_(name string) (t placeholder.Placeholder) {
	return template.NewPlaceholder(name)
}

func text_(name string) (t placeholder.Placeholder) {
	return template.NewPlaceholder(name, func(in interface{}) (out string) {
		return html.EscapeString(in.(string))
	})
}

var (
	_description_ = html_("description")
	_plaintext_   = text_("plaintext")

	SimpleTemplate = template.New("simple").MustAdd(
		`<div>`, _description_, `</div><h1>`, _plaintext_, `</h1>`,
	).Parse()

	NTemplate = template.New("n-template")
	MTemplate = template.New("m-template")

	StringT = `<div>` + replacer.Placeholder("description").String() + `</div><h1>` +
		replacer.Placeholder("plaintext").String() + `</h1>`
	TemplateT = `<div>{{.description}}</div><h1>{{.plaintext}}</h1>`
	//	Expected  = `<a href="http://example.com?a=b&&amp;c=d">a &lt;link&gt;</a>`
	Expected = `<div><p>desc</p></div><h1>&lt;escaped&gt;</h1>`
)

var (
	PlaceholderMap = map[string]placeholder.Placeholder{
		"description": _description_,
		"plaintext":   _plaintext_,
	}
	Map = map[string]string{
		"␣‣description∎␣": "<p>desc</p>",
		"␣‣plaintext∎␣":   "<escaped>",
	}

	StringMap = map[string]string{
		"description": "<p>desc</p>",
		"plaintext":   "<escaped>",
	}

	PlaceholderMapM = map[string]placeholder.Placeholder{}

	MapM       = map[string]string{}
	StringMapM = map[string]string{}
	TemlateMap = map[string]interface{}{
		"description": tt.HTML("<p>desc</p>"),
		"plaintext":   "<escaped>",
	}
	StringsM = []string{}
)

var (
	templ       = NewTemplate()
	StringM     string
	TemplateM   string
	ExpectedM   string
	TemlateMapM = map[string]interface{}{}
)

func PrepareM() {
	MapM = map[string]string{}
	StringMapM = map[string]string{}
	// StringsM = []string{}
	s := []string{}
	r := []string{}
	t := []string{}
	for i := 0; i < 5000; i++ {
		s = append(s, `<div>`+replacer.Placeholder(fmt.Sprintf("description%d", i)).String()+`</div><h1>`+
			replacer.Placeholder(fmt.Sprintf("plaintext%d", i)).String()+`</h1>`)
		t = append(t, fmt.Sprintf(`<div>{{.description%d}}</div><h1>{{.plaintext%d}}</h1>`, i, i))
		//r = append(r, fmt.Sprintf(`<a href="http://example.com?a=b&&amp;c=%d">a &lt;link%d&gt;</a>`, i,i))
		r = append(r, fmt.Sprintf(`<div><p>desc%d</p></div><h1>&lt;escaped%d&gt;</h1>`, i, i))

		kDescr := fmt.Sprintf("description%d", i)
		vDescr := fmt.Sprintf("<p>desc%d</p>", i)
		kPlaintext := fmt.Sprintf("plaintext%d", i)
		vPlaintext := fmt.Sprintf("<escaped%d>", i)
		MapM["␣‣"+kDescr+"∎␣"] = vDescr
		MapM["␣‣"+kPlaintext+"∎␣"] = vPlaintext
		PlaceholderMapM[kDescr] = html_(kDescr)
		PlaceholderMapM[kPlaintext] = text_(vPlaintext)
		StringMapM[kDescr] = vDescr
		StringMapM[kPlaintext] = vPlaintext
		TemlateMapM[kDescr] = tt.HTML(vDescr)
		TemlateMapM[kPlaintext] = vPlaintext
		MTemplate.MustAdd(
			`<div>`, html_(kDescr), `</div><h1>`, text_(vPlaintext), `</h1>`,
		)
		// StringsM = append(StringsM, "␣‣"+key+"∎␣", val)
	}
	StringM = strings.Join(s, "")
	TemplateM = strings.Join(t, "")
	ExpectedM = strings.Join(r, "")
	MTemplate.Parse()
	// ByteM = []byte(StringM)
}

var (
	TemplateN string
	StringN   string
	ExpectedN string
)

func PrepareN() {
	s := []string{}
	r := []string{}
	t := []string{}
	for i := 0; i < 2500; i++ {
		s = append(s, StringT)
		r = append(r, Expected)
		t = append(t, TemplateT)
		NTemplate.MustAdd(
			`<div>`, _description_, `</div><h1>`, _plaintext_, `</h1>`,
		)
	}
	NTemplate.Parse()
	TemplateN = strings.Join(t, "")
	StringN = strings.Join(s, "")
	ExpectedN = strings.Join(r, "")
}

func TestReplace(t *testing.T) {
	templ.Parse(TemplateT)
	var tbf bytes.Buffer
	if templ.Replace(TemlateMap, &tbf); tbf.String() != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), Expected)
	}

	setters := []placeholder.Setter{}

	for k, v := range StringMap {
		ph := PlaceholderMap[k]
		setters = append(setters, ph.Set(v))
	}

	res := SimpleTemplate.Replace(setters...).String()
	if res != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "go-on/template", res, Expected)
	}
}

func init() {
	PrepareN()
	PrepareM()
}

func TestReplaceN(t *testing.T) {
	templ.Parse(TemplateN)
	var tbf bytes.Buffer
	if templ.Replace(TemlateMap, &tbf); tbf.String() != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), ExpectedN)
	}

	setters := []placeholder.Setter{}
	for k, v := range StringMap {
		ph := PlaceholderMap[k]
		setters = append(setters, ph.Set(v))
	}

	res := NTemplate.Replace(setters...).String()

	if res != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedN)
	}
}

func TestReplaceM(t *testing.T) {
	templ.Parse(TemplateM)
	var tbf bytes.Buffer
	if templ.Replace(TemlateMapM, &tbf); tbf.String() != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), ExpectedM)
	}

	setters := []placeholder.Setter{}

	for k, v := range StringMapM {
		ph := PlaceholderMapM[k]
		setters = append(setters, ph.Set(v))
	}

	res := MTemplate.Replace(setters...).String()

	if res != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedM)
	}
}

func BenchmarkTemplateStandardLib(b *testing.B) {
	b.StopTimer()
	templ.Parse(TemplateN)
	var tbf bytes.Buffer
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		templ.Replace(TemlateMap, &tbf)
		tbf.Reset()
	}
}

func BenchmarkTemplateGoOn(b *testing.B) {
	b.StopTimer()
	setters := []placeholder.Setter{}
	for k, v := range StringMap {
		ph := PlaceholderMap[k]
		setters = append(setters, ph.Set(v))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NTemplate.Replace(setters...)
	}
}

func BenchmarkTemplateStandardLibM(b *testing.B) {
	b.StopTimer()
	templ.Parse(TemplateM)
	var tbf bytes.Buffer
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		templ.Replace(TemlateMapM, &tbf)
		tbf.Reset()
	}
}

func BenchmarkTemplateGoOnM(b *testing.B) {
	b.StopTimer()
	setters := []placeholder.Setter{}
	for k, v := range StringMapM {
		ph := PlaceholderMapM[k]
		setters = append(setters, ph.Set(v))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MTemplate.Replace(setters...)
	}
}

func BenchmarkOnceTemplateStandardLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		templ.Parse(TemplateT)
		var tbf bytes.Buffer
		templ.Replace(TemlateMap, &tbf)
	}
}

func BenchmarkOnceTemplateGoOn(b *testing.B) {
	setters := []placeholder.Setter{}
	for k, v := range StringMap {
		ph := PlaceholderMap[k]
		setters = append(setters, ph.Set(v))
	}

	for i := 0; i < b.N; i++ {
		t := template.New("simple").MustAdd(
			`<div>`, _description_, `</div><h1>`, _plaintext_, `</h1>`,
		).Parse()

		t.Replace(setters...)
	}
}
