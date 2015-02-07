replacer
========

fast and simple templating for go

[![Build Status](https://secure.travis-ci.org/go-on/replacer.png)](http://travis-ci.org/go-on/replacer)

If you need to simply replace placeholders in a template without escaping or logic,
replacer might be for you.

For the typical scenario - your template never changes on runtime -, replacer is faster than using (strings|bytes).Replace(r)() or regexp.ReplaceAllStringFunc() or the text/template package.

Performance
-----------

Runing benchmarks in the benchmark directory, I get the following results:
        
replacing 5000 placeholders that occur 1x in the template

    BenchmarkNaiveM          1  5468533647 ns/op 5507.8x (strings.Replace)
    BenchmarkNaive2M       500     5858220 ns/op    5.9x (strings.Replacer)
    BenchmarkRegM          100    27763655 ns/op   28.0x (regexp.ReplaceAllStringFunc)
    BenchmarkByteM        1000     2673126 ns/op    2.7x (bytes.Replace)
    BenchmarkTemplateM     500     6208385 ns/op    6.3x (template.Execute)
    BenchmarkReplacerM    2000      992872 ns/op    1.0x (replacer.Replace)

replacing 2 placeholders that occur 2500x in the template

    BenchmarkNaive         500     3466000 ns/op    7.3x (strings.Replace)
    BenchmarkNaive2       1000     2387591 ns/op    5.1x (strings.Replacer)
    BenchmarkReg           100    22783618 ns/op   48.2x (regexp.ReplaceAllStringFunc)
    BenchmarkByte         1000     2173147 ns/op    4.6x (bytes.Replace)
    BenchmarkTemplate      500     5700397 ns/op   12.1x (template.Execute)
    BenchmarkReplacer     5000      472550 ns/op    1.0x (replacer.Replace)

replacing 2 placeholders that occur 1x in the template, parsing template each time

    BenchmarkOnceNaive      500     3471216 ns/op   12.9x (strings.Replace)
    BenchmarkOnceNaive2    1000     2399079 ns/op    8.9x (strings.Replacer)
    BenchmarkOnceReg        100    22886067 ns/op   85.2x (regexp.ReplaceAllStringFunc)
    BenchmarkOnceByte      1000     2449898 ns/op    9.1x (bytes.Replace)
    BenchmarkOnceTemplate     1  1525545312 ns/op 5677.5x (template.Execute)
    BenchmarkOnceReplacer 10000      268698 ns/op    1.0x (replacer.Replace)

replacing 1 placeholder that occur 1x in the template, parsing template each time

    BenchmarkOnceSingleNaive    1000000    1044 ns/op   1.3x (strings.Replace)
    BenchmarkOnceSingleNaive2    500000    5884 ns/op   7.1x (strings.Replacer)
    BenchmarkOnceSingleReg       200000    9796 ns/op  11.8x (regexp.ReplaceAllStringFunc)
    BenchmarkOnceSingleByte     1000000    1532 ns/op   1.9x (bytes.Replace)
    BenchmarkOnceSingleTemplate  100000   24383 ns/op  29.5x (template.Execute)
    BenchmarkOnceSingleReplacer 2000000     827 ns/op   1.0x (replacer.Replace)


Usage
-----

```go
package main

import (
    "fmt"
    "github.com/go-on/replacer"
)

func main() {
    template := replacer.Placeholder("name").String() + " - " + replacer.Placeholder("animal").String()

    res := replacer.NewTemplateString(template).ReplaceStrings(
        replacer.MapStrings(
            "animal", "Duck",
            "name", "Donald",
        ),
    )

    // the returned replacement is a []byte
    fmt.Printf("%s\n", res)
}
```

results in

```
Donald - Duck
```

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

see LICENSE file.

GoDoc
-----

see http://godoc.org/github.com/go-on/replacer


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/go-on/replacer/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

