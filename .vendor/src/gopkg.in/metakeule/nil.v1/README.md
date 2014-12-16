nil
===

Simple helpers to deal with nil values in go

[![Build Status](https://secure.travis-ci.org/metakeule/nil.png)](http://travis-ci.org/metakeule/nil)

Example
-------

```go
package main

import (
    "fmt"
    "github.com/metakeule/nil"
)

func main() {
    var x interface{}
    fmt.Println("is nil ", nil.Is(x))
    nil.MustNot(x) // panics
}
```

or

```go
package main

import (
    "fmt"
    . "github.com/metakeule/nil"
)

func main() {
    var x interface{}
    fmt.Println("is nil ", IsNil(x))
    Nø(x) // panics
}
```

Documentation
-------------

see http://godoc.org/github.com/metakeule/nil