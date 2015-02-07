package replacer

import (
	"net/http"
)

func StringsMap(in map[string]string) map[Placeholder]string {
	res := make(map[Placeholder]string, len(in))
	for k, v := range in {
		res[Placeholder(k)] = v
	}
	return res
}

func BytesMap(in map[string][]byte) map[Placeholder][]byte {
	res := make(map[Placeholder][]byte, len(in))
	for k, v := range in {
		res[Placeholder(k)] = v
	}
	return res
}

func MapStrings(pairs ...string) map[Placeholder]string {
	l := len(pairs)
	if l%2 != 0 {
		panic("len pairs must be even")
	}
	m := make(map[Placeholder]string, l/2)

	for i := 0; i < l; i += 2 {
		m[Placeholder(pairs[i])] = pairs[i+1]
	}
	return m
}

func MapHandlers(pairs ...interface{}) map[string]http.Handler {
	l := len(pairs)
	if l%2 != 0 {
		panic("len pairs must be even")
	}
	m := make(map[string]http.Handler, l/2)

	for i := 0; i < l; i += 2 {
		m[pairs[i].(string)] = pairs[i+1].(http.Handler)
	}
	return m
}
