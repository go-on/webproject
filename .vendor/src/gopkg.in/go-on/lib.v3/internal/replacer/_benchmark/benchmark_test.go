package benchmark

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/go-on/replacer"
)

var (
	StringT   = "a string with " + replacer.P("replacement1").String() + " and " + replacer.P("replacement2").String() + " that c@ntinues"
	TemplateT = "a string with {{.replacement1}} and {{.replacement2}} that c@ntinues"
	ByteT     = []byte(StringT)
	Expected  = "a string with repl1 and repl2 that c@ntinues"

	StringN   = ""
	TemplateN = ""
	ByteN     = []byte{}
	ExpectedN = ""

	StringM   = ""
	TemplateM = ""
	ByteM     = []byte{}
	ExpectedM = ""
)

var (
	Map1 = map[string]string{
		"␣‣replacement1∎␣": "repl1",
	}
	Map = map[string]string{
		"␣‣replacement1∎␣": "repl1",
		"␣‣replacement2∎␣": "repl2",
	}

	Strings = []string{replacer.P("replacement1").String(), "repl1", replacer.P("replacement2").String(), "repl2"}

	Strings1 = []string{replacer.P("replacement1").String(), "repl1"}

	StringMap = map[string]string{
		"replacement1": "repl1",
		"replacement2": "repl2",
	}

	StringMap1 = map[string]string{
		"replacement1": "repl1",
	}

	ByteMap = map[string][]byte{
		"␣‣replacement1∎␣": []byte("repl1"),
		"␣‣replacement2∎␣": []byte("repl2"),
	}

	ByteMap1 = map[string][]byte{
		"␣‣replacement1∎␣": []byte("repl1"),
	}

	MapM       = map[string]string{}
	StringMapM = map[string]string{}
	ByteMapM   = map[string][]byte{}
	StringsM   = []string{}
)

var (
	mapperNaive = &Naive{}
	naive2      = &Naive2{}
	mapperReg   = &Regexp{Regexp: regexp.MustCompile("(␣‣[^∎]+∎␣)")}
	byts        = &Bytes{}
	//	repl        = replacer.New()
	templ = NewTemplate()
)

func PrepareM() {
	MapM = map[string]string{}
	ByteMapM = map[string][]byte{}
	StringMapM = map[string]string{}
	StringsM = []string{}
	s := []string{}
	r := []string{}
	t := []string{}
	for i := 0; i < 5000; i++ {
		s = append(s, fmt.Sprintf(`a string with ␣‣replacement%v∎␣`, i))
		t = append(t, fmt.Sprintf(`a string with {{.replacement%v}}`, i))
		r = append(r, fmt.Sprintf("a string with repl%v", i))
		key := fmt.Sprintf("replacement%v", i)
		val := fmt.Sprintf("repl%v", i)
		MapM["␣‣"+key+"∎␣"] = val
		ByteMapM["␣‣"+key+"∎␣"] = []byte(val)
		StringMapM[key] = val
		StringsM = append(StringsM, "␣‣"+key+"∎␣", val)
	}
	StringM = strings.Join(s, "")
	TemplateM = strings.Join(t, "")
	ExpectedM = strings.Join(r, "")
	ByteM = []byte(StringM)
}

func PrepareN() {
	s := []string{}
	r := []string{}
	t := []string{}
	for i := 0; i < 2500; i++ {
		s = append(s, StringT)
		r = append(r, Expected)
		t = append(t, TemplateT)
	}
	TemplateN = strings.Join(t, "")
	StringN = strings.Join(s, "")
	ExpectedN = strings.Join(r, "")
	ByteN = []byte(StringN)
}

func TestReplace(t *testing.T) {
	mapperNaive.Map = Map
	mapperNaive.Template = StringT
	if r := mapperNaive.Replace(); r != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "mapperNaive", r, Expected)
	}

	naive2.Replacements = Strings
	naive2.Template = StringT
	if r := naive2.Replace(); r != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "naive2", r, Expected)
	}

	mapperReg.Map = Map
	mapperReg.Template = StringT
	mapperReg.Setup()
	if r := mapperReg.Replace(); r != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "mapperReg", r, Expected)
	}

	byts.Map = ByteMap
	byts.Parse(StringT)
	if r := byts.Replace(); string(r) != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "byts", string(r), Expected)
	}

	templ.Parse(TemplateT)
	var tbf bytes.Buffer
	if templ.Replace(StringMap, &tbf); tbf.String() != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), Expected)
	}

	repl := replacer.NewTemplateBytes(ByteT)

	//	var bf bytes.Buffer
	res := string(repl.ReplaceStrings(replacer.StringsMap(StringMap)))
	if res != Expected {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, Expected)
	}
}

func TestReplaceN(t *testing.T) {
	PrepareN()
	mapperNaive.Map = Map
	mapperNaive.Template = StringN
	if r := mapperNaive.Replace(); r != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "mapperNaive", r, ExpectedN)
	}

	naive2.Replacements = Strings
	naive2.Template = StringN
	if r := naive2.Replace(); r != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "naive2", r, ExpectedN)
	}

	mapperReg.Map = Map
	mapperReg.Template = StringN
	mapperReg.Setup()
	if r := mapperReg.Replace(); r != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "mapperReg", r, ExpectedN)
	}

	templ.Parse(TemplateN)
	var tbf bytes.Buffer
	if templ.Replace(StringMap, &tbf); tbf.String() != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), ExpectedN)
	}

	repl := replacer.NewTemplateBytes(ByteN)

	//	var bf bytes.Buffer
	res := string(repl.ReplaceStrings(replacer.StringsMap(StringMap)))
	if res != ExpectedN {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedN)
	}
}

func TestReplaceM(t *testing.T) {
	PrepareM()
	mapperNaive.Map = MapM
	mapperNaive.Template = StringM
	if r := mapperNaive.Replace(); r != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "mapperNaive", r, ExpectedM)
	}

	naive2.Replacements = StringsM
	naive2.Template = StringM
	if r := naive2.Replace(); r != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "naive2", r, ExpectedM)
	}

	mapperReg.Map = MapM
	mapperReg.Template = StringM
	mapperReg.Setup()
	if r := mapperReg.Replace(); r != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "mapperReg", r, ExpectedM)
	}

	naive2.Replacements = StringsM
	naive2.Template = StringM
	if r := naive2.Replace(); r != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "naive2", r, ExpectedM)
	}

	templ.Parse(TemplateM)
	var tbf bytes.Buffer
	if templ.Replace(StringMapM, &tbf); tbf.String() != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "template", tbf.String(), ExpectedM)
	}

	repl := replacer.NewTemplateBytes(ByteM)

	//	var bf bytes.Buffer

	res := string(repl.ReplaceStrings(replacer.StringsMap(StringMapM)))
	if res != ExpectedM {
		t.Errorf("unexpected result for %s: %#v, expected: %#v", "fastreplace2", res, ExpectedM)
	}
}

func BenchmarkNaive(b *testing.B) {
	b.StopTimer()
	mapperNaive.Map = Map
	mapperNaive.Template = StringN
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mapperNaive.Replace()
	}
}

func BenchmarkNaive2(b *testing.B) {
	b.StopTimer()
	naive2.Replacements = Strings
	naive2.Template = StringN
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		naive2.Replace()
	}
}

func BenchmarkReg(b *testing.B) {
	b.StopTimer()
	mapperReg.Map = Map
	mapperReg.Template = StringN
	mapperReg.Setup()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mapperReg.Replace()
	}
}

func BenchmarkByte(b *testing.B) {
	b.StopTimer()
	byts.Map = ByteMap
	byts.Parse(StringN)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		byts.Replace()
	}
}
func BenchmarkTemplate(b *testing.B) {
	b.StopTimer()
	templ.Parse(TemplateN)
	var tbf bytes.Buffer
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		templ.Replace(StringMap, &tbf)
		tbf.Reset()
	}

}

func BenchmarkReplacer(b *testing.B) {
	b.StopTimer()
	//  repl.ParseBytes(ByteN)
	repl := replacer.NewTemplateBytes(ByteN)
	//	var bf bytes.Buffer
	m := replacer.StringsMap(StringMap)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		repl.ReplaceStrings(m)
		// repl.MustReplaceStringsTo(&bf, m)
		//repl.ReplaceStrings(m)
		// bf.Reset()
	}
}

func BenchmarkNaiveM(b *testing.B) {
	b.StopTimer()
	PrepareM()
	mapperNaive.Map = MapM
	mapperNaive.Template = StringM
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mapperNaive.Replace()
	}
}

func BenchmarkNaive2M(b *testing.B) {
	b.StopTimer()
	PrepareM()
	naive2.Replacements = StringsM
	naive2.Template = StringM
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		naive2.Replace()
	}
}

func BenchmarkRegM(b *testing.B) {
	b.StopTimer()
	PrepareM()
	mapperReg.Map = MapM
	mapperReg.Template = StringM
	mapperReg.Setup()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mapperReg.Replace()
	}
}

func BenchmarkByteM(b *testing.B) {
	b.StopTimer()
	PrepareM()
	byts.Map = ByteMap
	byts.Parse(StringM)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		byts.Replace()
	}
}

func BenchmarkTemplateM(b *testing.B) {
	b.StopTimer()
	PrepareM()
	templ.Parse(TemplateM)
	var tbf bytes.Buffer
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		templ.Replace(StringMapM, &tbf)
		tbf.Reset()
	}
}

func BenchmarkReplacerM(b *testing.B) {
	b.StopTimer()
	PrepareM()
	repl := replacer.NewTemplateBytes(ByteM)
	//var bf bytes.Buffer
	m := replacer.StringsMap(StringMapM)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		repl.ReplaceStrings(m)
		//repl.MustReplaceStringsTo(&bf, m)
		//bf.Reset()
	}
}

func BenchmarkOnceNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapperNaive.Map = Map
		mapperNaive.Template = StringN
		mapperNaive.Replace()
	}
}

func BenchmarkOnceNaive2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naive2.Replacements = Strings
		naive2.Template = StringN
		naive2.Replace()
	}
}

func BenchmarkOnceReg(b *testing.B) {
	mapperReg.Setup()
	for i := 0; i < b.N; i++ {
		mapperReg.Map = Map
		mapperReg.Template = StringN
		mapperReg.Replace()
	}
}

func BenchmarkOnceByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byts.Parse(StringN)
		byts.Map = ByteMap
		byts.Replace()
	}
}

func BenchmarkOnceTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		templ.Parse(TemplateN)
		var tbf bytes.Buffer
		templ.Replace(StringMap, &tbf)
	}
}

func BenchmarkOnceReplacer(b *testing.B) {
	m := replacer.StringsMap(StringMap)
	for i := 0; i < b.N; i++ {
		//var bf bytes.Buffer
		//replacer.NewTemplateBytes(ByteN).MustReplaceStringsTo(&bf, m)
		//replacer.NewTemplateBytes(ByteN).ReplaceStrings(m)
		replacer.NewTemplateBytes(ByteN).ReplaceStrings(m)
	}
}

func BenchmarkOnceSingleNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapperNaive.Map = Map1
		mapperNaive.Template = StringT
		mapperNaive.Replace()
	}
}

func BenchmarkOnceSingleNaive2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naive2.Replacements = Strings1
		naive2.Template = StringT
		naive2.Replace()
	}
}

func BenchmarkOnceSingleReg(b *testing.B) {
	mapperReg.Setup()
	for i := 0; i < b.N; i++ {
		mapperReg.Map = Map1
		mapperReg.Template = StringT
		mapperReg.Replace()
	}
}

func BenchmarkOnceSingleByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byts.Parse(StringT)
		byts.Map = ByteMap1
		byts.Replace()
	}
}

func BenchmarkOnceSingleTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		templ.Parse(TemplateT)
		var tbf bytes.Buffer
		templ.Replace(StringMap1, &tbf)
	}
}

func BenchmarkOnceSingleReplacer(b *testing.B) {
	//m := replacer.StringMap(StringMap1)
	m := replacer.BytesMap(ByteMap1)
	for i := 0; i < b.N; i++ {
		//var bf bytes.Buffer
		//replacer.NewTemplateBytes(ByteT).MustReplaceStringsTo(&bf, m)
		//replacer.NewTemplateBytes(ByteT).ReplaceStrings(m)
		replacer.NewTemplateBytes(ByteT).ReplaceBytes(m)
	}
}
