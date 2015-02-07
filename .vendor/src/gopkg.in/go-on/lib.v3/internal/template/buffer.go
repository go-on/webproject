package template

import (
	"bytes"
	"fmt"

	// "github.com/go-on/template/placeholder"
	// "io"
	"net/http"
	// "reflect"
	// "strings"

	// "github.com/go-on/replacer"
)

// a named buffer, fullfills the Setter interface
type Buffer struct {
	*bytes.Buffer
	name string
}

func newBuffer(name string) *Buffer {
	bf := &Buffer{}
	bf.name = name
	bf.Buffer = &bytes.Buffer{}
	return bf
}

// ServeHTTP writes the content of the buffer to the responsewriter, each time
// it is called. It does not clean the buffer as would be done with WriteTo(ResponseWriter)
func (bf *Buffer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(bf.Bytes())
}

func (b *Buffer) Print() {
	fmt.Print(b.Buffer.String())
}

func (b *Buffer) Println() {
	fmt.Println(b.Buffer.String())
}

func (b *Buffer) Name() string {
	return b.name
}

func (b *Buffer) SetString() string {
	return b.Buffer.String()
}
