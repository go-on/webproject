package benchmark

import (
	"bytes"
	"gopkg.in/go-on/builtin.v1"
	"gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/html/internal/element/compiler"
	"gopkg.in/go-on/lib.v2/internal/template/placeholder"

	. "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/internal/template"
	"gopkg.in/go-on/lib.v2/types"
	ph "gopkg.in/go-on/lib.v2/types/placeholder"
	tt "html/template"
	"testing"
)

var (
	_description_ = ph.New(types.HTMLString("description"))
	_plaintext_   = ph.New(types.Text("plaintext"))

	PlaceholderMap = map[string]ph.Placeholder{
		"<p>desc</p>": _description_,
		"<escaped>":   _plaintext_,
	}

	PlaceholderT = compiler.ElementTemplate("placeholderT", element.Elements(
		DIV(_description_),
		H1(_plaintext_),
	))

	PlaceholderN *template.Template
	PlaceholderM *template.Template

	PlaceholderMapM map[string]ph.Placeholder
)

func PrepareNPlaceholder() {
	t := SimpleTemplate(_description_, _plaintext_)
	for i := 0; i < 2500; i++ {
		t = SimpleTemplate(t, _plaintext_)
	}
	PlaceholderN = t.Compile("placeholderN")
}

func SimpleTemplate(description, plaintext builtin.Stringer) *element.Element {
	return element.Elements(DIV(
		description,
	),
		H1(plaintext),
	)
}

func TemplateTCreator() *element.Element {
	return SimpleTemplate(
		types.HTMLString(`{{.description}}`),
		types.HTMLString(`{{.plaintext}}`),
	)
}

func NTemplate() *element.Element {
	t := SimpleTemplate(types.HTMLString("<p>desc</p>"), types.Text("<escaped>"))
	for i := 0; i < 2500; i++ {
		t = SimpleTemplate(t, types.Text("<escaped>"))
	}
	return t
}

func MTemplate() *element.Element {
	t := element.Elements()
	for i := 0; i < 5000; i++ {
		t.MustAdd(SimpleTemplate(types.HTMLString(fmt.Sprintf("<p>desc%d</p>", i)), types.Text(fmt.Sprintf("<escaped%d>", i))))
	}
	return t
}

func PrepareMPlaceholder() {
	t := element.Elements()
	PlaceholderMapM = map[string]ph.Placeholder{}
	for i := 0; i < 5000; i++ {
		phDesc := html.Htmlf("description%d", i).Placeholder()
		phPlain := html.Textf("plaintext%d", i).Placeholder()
		PlaceholderMapM[fmt.Sprintf("<p>desc%d</p>", i)] = phDesc
		PlaceholderMapM[fmt.Sprintf("<escaped%d>", i)] = phPlain
		t.MustAdd(SimpleTemplate(phDesc, phPlain))
	}
	PlaceholderM = t.Compile("placeholderM")
}

func TemplateMCreator() *html.Element {
	t := html.Elements()
	for i := 0; i < 5000; i++ {
		t.MustAdd(SimpleTemplate(types.HTMLString(fmt.Sprintf("{{.description%d}}", i)), types.HTMLString(fmt.Sprintf("{{.plaintext%d}}", i))))
	}
	return t
}

func TemplateNCreator() *html.Element {
	t := SimpleTemplate(
		html.Html(`{{.description}}`),
		html.Html(`{{.plaintext}}`))
	for i := 0; i < 2500; i++ {
		t = SimpleTemplate(t, types.HTMLString(`{{.plaintext}}`))
	}
	return t
}

var (
	templ     = NewTemplate()
	TemplateT = TemplateTCreator().String()
	TemplateM = TemplateMCreator().String()
	TemplateN = TemplateNCreator().String()
	//	Expected  = `<a href="http://example.com?a=b&&amp;c=d">a &lt;link&gt;</a>`
	Expected  = SimpleTemplate(html.Html("<p>desc</p>"), html.Text("<escaped>")).String()
	ExpectedM = MTemplate().String()
	ExpectedN = NTemplate().String()
)

var (
	TemlateMap = map[string]interface{}{
		"description": tt.HTML("<p>desc</p>"),
		"plaintext":   "<escaped>",
	}
)

var (
	TemlateMapM = map[string]interface{}{}
)

func PrepareM() {
	for i := 0; i < 5000; i++ {
		kDescr := fmt.Sprintf("description%d", i)
		vDescr := fmt.Sprintf("<p>desc%d</p>", i)
		kPlaintext := fmt.Sprintf("plaintext%d", i)
		vPlaintext := fmt.Sprintf("<escaped%d>", i)
		TemlateMapM[kDescr] = tt.HTML(vDescr)
		TemlateMapM[kPlaintext] = vPlaintext
	}
}

func TestReplace(t *testing.T) {
	templ.Parse(TemplateT)
	var tbf bytes.Buffer
	if templ.Replace(TemlateMap, &tbf); tbf.String() != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), Expected)
	}

	res := SimpleTemplate(html.Html("<p>desc</p>"), html.Text("<escaped>")).String()
	if res != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "go-on/template", res, Expected)
	}

	setters := []placeholder.Setter{}

	for k, v := range PlaceholderMap {
		setters = append(setters, v.Set(k))
	}

	res = PlaceholderT.Replace(setters...).String()
	if res != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, Expected)
	}
}

func init() {
	PrepareM()
	PrepareMPlaceholder()
	PrepareNPlaceholder()
}

func TestReplaceN(t *testing.T) {
	templ.Parse(TemplateN)
	var tbf bytes.Buffer
	if templ.Replace(TemlateMap, &tbf); tbf.String() != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), ExpectedN)
	}

	res := NTemplate().String()

	if res != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedN)
	}

	setters := []placeholder.Setter{}

	for k, v := range PlaceholderMap {
		setters = append(setters, v.Set(k))
	}

	res = PlaceholderN.Replace(setters...).String()
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

	res := MTemplate().String()

	if res != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedM)
	}

	setters := []placeholder.Setter{}

	for k, v := range PlaceholderMapM {
		setters = append(setters, v.Set(k))
	}

	res = PlaceholderM.Replace(setters...).String()
	if res != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedM)
	}
}

func BenchmarkTemplateStandardLibN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tbf bytes.Buffer
		templ.Parse(TemplateN)
		templ.Replace(TemlateMap, &tbf)
		tbf.Reset()
	}
}

func BenchmarkTemplateStandardLibParsedN(b *testing.B) {
	b.StopTimer()
	var tbf bytes.Buffer
	templ.Parse(TemplateN)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		templ.Replace(TemlateMap, &tbf)
		tbf.Reset()
	}
}

func BenchmarkHtmlGoOnN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NTemplate().String()
	}
}

func BenchmarkHtmlGoOnCompiledN(b *testing.B) {
	b.StopTimer()
	setters := []placeholder.Setter{}

	for k, v := range PlaceholderMap {
		setters = append(setters, v.Set(k))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		PlaceholderN.Replace(setters...).String()
	}
}

func BenchmarkTemplateStandardLibM(b *testing.B) {
	var tbf bytes.Buffer
	for i := 0; i < b.N; i++ {
		templ.Parse(TemplateM)
		templ.Replace(TemlateMapM, &tbf)
		tbf.Reset()
	}
}

func BenchmarkTemplateStandardLibParsedM(b *testing.B) {
	b.StopTimer()
	var tbf bytes.Buffer
	templ.Parse(TemplateM)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		templ.Replace(TemlateMapM, &tbf)
		tbf.Reset()
	}
}

func BenchmarkHtmlGoOnM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MTemplate().String()
	}
}

func BenchmarkHtmlGoOnCompiledM(b *testing.B) {
	b.StopTimer()
	setters := []placeholder.Setter{}

	for k, v := range PlaceholderMapM {
		setters = append(setters, v.Set(k))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		PlaceholderM.Replace(setters...).String()
	}
}

func BenchmarkOnceTemplateStandardLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		templ.Parse(TemplateT)
		var tbf bytes.Buffer
		templ.Replace(TemlateMap, &tbf)
	}
}

func BenchmarkOnceHtmlGoOn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleTemplate(html.Html("<p>desc</p>"), html.Text("<escaped>")).String()
	}
}
