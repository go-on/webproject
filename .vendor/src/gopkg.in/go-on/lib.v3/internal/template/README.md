template
========

general purpose templating based on go-on/replacer

Performance
-----------

replacing 5000 placeholders that occur 1x in the template

    BenchmarkTemplateStandardLibM   20  52983126 ns/op   5.4x (html/template)
    BenchmarkTemplateGoOnM         200   9801185 ns/op   1.0x (go-on/template)

replacing 2 placeholders that occur 2500x in the template

    BenchmarkTemplateStandardLib    50  23802659 ns/op  25.3x (html/template)
    BenchmarkTemplateGoOn         2000    942400 ns/op   1.0x (go-on/template)

replacing 2 placeholders that occur 1x in the template, parsing template each time

    BenchmarkOnceTemplateStandardLib   50000  70096 ns/op 10.7x (html/template)
    BenchmarkOnceTemplateGoOn         500000   6574 ns/op  1.0x (go-on/template)

Usage
-----

```go
package main

import (
    "bytes"
    "fmt"
    "html"

    "github.com/go-on/template"
    "github.com/go-on/template/placeholder"
)

func Html(name string) (t placeholder.Placeholder) {
    return template.NewPlaceholder(name)
}

func Text(name string) (t placeholder.Placeholder) {
    return template.NewPlaceholder(
        name,
        func(in interface{}) (out string) { return html.EscapeString(in.(string)) },
    )
}

var (
    person   = Text("person")
    greeting = Html("greeting")
    T        = template.New("t").MustAdd("<h1>Hi, ", person, "</h1>", greeting).Parse()
)

func main() {
    fmt.Println(
        T.MustReplace(
            person.Set("S<o>meone"),
            greeting.Set("<div>Hi</div>"),
        ),
    )

    var buffer bytes.Buffer
    for i := 0; i < 10; i++ {
        T.MustReplaceTo(&buffer,
            person.Setf("Bugs <Bunny> %v", i+1),
            greeting.Set("<p>How are you?</p>\n"))
    }
    fmt.Println(buffer.String())
}
```

returns

```
<h1>Hi, S&lt;o&gt;meone</h1><div>Hi</div>
<h1>Hi, Bugs &lt;Bunny&gt; 1</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 2</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 3</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 4</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 5</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 6</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 7</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 8</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 9</h1><p>How are you?</p>
<h1>Hi, Bugs &lt;Bunny&gt; 10</h1><p>How are you?</p>
```