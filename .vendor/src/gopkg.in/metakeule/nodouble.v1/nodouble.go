package nodouble

import (
	"fmt"
	"runtime"
)

type info struct {
	version int
	file    string
	line    int
}

var pkgs = map[string]info{}

func AddPackage(path string, version int) {
	_, file, line, _ := runtime.Caller(0)
	i, has := pkgs[path]
	if has {
		panic(fmt.Sprintf("can't load package %s version %v as requested in %s:%v; already loaded version %v as requested from %s:%v",
			path, version, file, line,
			i.version, i.file, i.line,
		))
	}
	pkgs[path] = info{version, file, line}
}
