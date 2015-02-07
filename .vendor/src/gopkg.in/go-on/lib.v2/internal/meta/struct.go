package meta

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	Value  *reflect.Value
	Type   reflect.StructField
	Struct *Struct
}

func (f *Field) Set(v reflect.Value) error {
	if !f.Value.CanSet() {
		return fmt.Errorf("%s can't be set to %#v", f.Type.Name, f.Value.Interface())
	}
	if !v.Type().AssignableTo(f.Type.Type) {
		return fmt.Errorf("can't set field %s of type %s with value of type %s (not assignable)", f.Type.Name, f.Type.Type, v.Type())
	}
	f.Value.Set(v)
	return nil
}

type Struct struct {
	Value *reflect.Value
}

// t must be a of Kind reflect.Struct
func StructByType(t reflect.Type) (*Struct, error) {
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is no struct type", t)
	}
	val := reflect.New(t)
	ref := reflect.New(val.Type())
	ref.Elem().Set(val)
	rf := ref.Elem()
	return &Struct{&rf}, nil
}

// val must be the reflect.Value of a reference to a struct
func StructByValue(val reflect.Value) (*Struct, error) {
	if val.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%s is no pointer to a struct", val.Type().String())
	}
	if val.IsNil() {
		return nil, fmt.Errorf("%s is nil pointer", val.Type().String())
	}
	if val.Type().Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is no pointer to a struct", val.Type().String())
	}
	return &Struct{&val}, nil
}

// replace the value of the struct, val must be a pointer to a struct
func (s *Struct) Set(val reflect.Value) error {
	if val.IsNil() {
		return fmt.Errorf("setting to nil is not allowed")
	}
	if !s.Value.CanSet() {
		return fmt.Errorf("can't replace %s", s.Value.Type())
	}

	if !val.Type().AssignableTo(s.Value.Type()) {
		return fmt.Errorf("can't set %s to value of type %s (not assignable)", s.Value.Type(), val.Type())
	}
	s.Value.Set(val)
	return nil
}

func (s *Struct) Field(name string) (*Field, error) {
	fld, exists := s.Value.Elem().Type().FieldByName(name)
	if !exists {
		return nil, fmt.Errorf("not an exported field: %s", name)
	}

	v := s.Value.Elem().FieldByName(name)
	return &Field{&v, fld, s}, nil
}

func (s *Struct) Each(fn func(field *Field)) {
	elem := s.Value.Elem().NumField()
	for i := 0; i < elem; i++ {
		v := s.Value.Elem().Field(i)
		fn(&Field{&v, s.Value.Elem().Type().Field(i), s})
	}
	return
}

// get every field and its tag for a given tag key, empty tags and tags with value "-" are ignored
func (s *Struct) EachTag(tagKey string, fn func(field *Field, tagVal string)) {
	f := func(field *Field) {
		tagVal := field.Type.Tag.Get(tagKey)
		if tagVal != "" && tagVal != "-" {
			fn(field, tagVal)
		}
	}
	s.Each(f)
}

// get every field and its tag for a given tag key, tags with value "-" are ignored
func (s *Struct) EachTagWithEmpty(tagKey string, fn func(field *Field, tagVal string)) {
	f := func(field *Field) {
		tagVal := field.Type.Tag.Get(tagKey)
		if tagVal != "-" {
			fn(field, tagVal)
		}
	}
	s.Each(f)
}

// returns a struct tag for a field
func (s *Struct) Tag(field string) (*reflect.StructTag, error) {
	if !(s.Value.Kind() == reflect.Struct) {
		return nil, fmt.Errorf("%T is not a struct", s)
	}
	f, exists := s.Value.Type().FieldByName(field)
	if !exists {
		return nil, fmt.Errorf("field %s does not exist in %s", field, s.Value.Interface())
	}
	return &f.Tag, nil
}

// returns all struct tags
func (s *Struct) Tags() (tags map[string]*reflect.StructTag, err error) {
	tags = map[string]*reflect.StructTag{}

	// ft := ø.FinalType(s)
	if !(s.Value.Kind() == reflect.Struct) {
		// Panicf("%s is not a struct / pointer to a struct", Inspect(s))
		return nil, fmt.Errorf("%s is not a struct / pointer to a struct", s.Value.Interface())
	}
	//elem := ft.NumField()
	elem := s.Value.NumField()
	for i := 0; i < elem; i++ {
		f := s.Value.Type().Field(i)
		if string(f.Tag) != "" {
			tags[f.Name] = &f.Tag
		}
	}
	return
}

// ToMap converts a given struct to a map. If tag is not "", it follows the
// same semantics as json encoding, but for the given tag.
// Panics on errors
func (s *Struct) ToMap(tag string) map[string]interface{} {
	sm := &structMap{s, tag, map[string]interface{}{}}
	sm.ToMap()
	return sm.m
}

type structMap struct {
	*Struct
	tag string
	m   map[string]interface{}
}

// toMap ignores the tag
func (s *structMap) toMapNotTagged() {
	s.Each(s.setField)
}

func (s *structMap) setField(field *Field) {
	s.m[field.Type.Name] = field.Value.Interface()
}

// toMapTagged requires tag to be set to something other than ""
func (s *structMap) toMapTagged() {
	s.EachTagWithEmpty(s.tag, s.setTaggedField)
}

func (s *structMap) ToMap() {
	if s.tag == "" {
		s.toMapNotTagged()
		return
	}
	s.toMapTagged()
}

func (s *structMap) setTaggedField(field *Field, tagVal string) {
	if tagVal == "" {
		// same as if tag == ""
		s.setField(field)
		return
	}

	tvs := strings.Split(tagVal, ",")
	fieldName := tvs[0]

	var omitempty bool
	if len(tvs) > 1 && tvs[1] == "omitempty" {
		omitempty = true
	}

	// omit empty values as required by omitempty tag
	if IsZero(*field.Value) && omitempty {
		return
	}

	if fieldName == "" {
		fieldName = field.Type.Name
	}

	s.m[fieldName] = field.Value.Interface()
}

func IsZero(v reflect.Value) bool {
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
