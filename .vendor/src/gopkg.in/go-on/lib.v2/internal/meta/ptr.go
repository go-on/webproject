package meta

import (
	"fmt"
	"reflect"
)

type Pointer struct {
	Value *reflect.Value
}

func PointerByValue(v reflect.Value) (*Pointer, error) {
	if v.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%s is no pointer", v.Type())
	}
	return &Pointer{&v}, nil
}

// Assoc associates targetPtrPtr with srcPtr so that
// targetPtrPtr is a pointer to srcPtr and
// targetPtr and srcPtr are pointing to the same address
func (p *Pointer) Assoc(targetPtrPtr reflect.Value) error {
	if targetPtrPtr.Kind() != reflect.Ptr || targetPtrPtr.Elem().Kind() != reflect.Ptr || targetPtrPtr.Elem().Elem().Type() != p.Value.Elem().Type() {
		return fmt.Errorf("%#v must be of type **%s", targetPtrPtr.Interface(), p.Value.Elem().Type())
	}
	targetPtrPtr.Elem().Set(*p.Value)
	return nil
}
