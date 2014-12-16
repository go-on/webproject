package fat

import (
	"fmt"
	//"gopkg.in/metakeule/meta.v5"
	"gopkg.in/go-on/lib.v2/internal/meta"
	"reflect"
	"strings"
)

func StructType(østruct interface{}) string {
	ty := reflect.TypeOf(østruct).Elem()
	return "*" + ty.PkgPath() + "." + ty.Name()
}

// the first match wins
var matcher = []string{
	"[string]string", "[string]int", "[string]time", "[string]float", "[string]bool",
	"[]string", "[]int", "[]time", "[]float", "[]bool",
	"string", "int", "time", "float", "bool",
}

func findType(tag string) (typ string) {
	for _, t := range matcher {
		if strings.Contains(tag, t) {
			return t
		}
	}
	return
}

// sets all attributes of a struct that are of type *Field
// to a *Field with the Type set to what is given in the tag "fat.type"
// with defaults set to what is given in the tag "fat.default"
// and with enums set to what is in the tag "fat.enum", separated by pipe symbols (|)
func Proto(østruct interface{}) (østru interface{}) {
	structtype := StructType(østruct)
	//fn := func(field reflect.StructField, val reflect.Value) {
	fn := func(field *meta.Field) {
		if _, ok := field.Value.Interface().(*Field); ok {
			ty := findType(field.Type.Tag.Get("type"))
			if ty == "" {
				panic(fmt.Sprintf("struct %s has no valid type tag for field %s", structtype, field.Type.Name))
			}

			f := newSpec(structtype, østruct, field.Type.Name, ty)
			def := field.Type.Tag.Get("default")
			if def != "" {
				d := f.fieldSpec.new()
				err := d.Scan(def)
				if err != nil {
					panic(fmt.Sprintf("default value %s for field %s in struct %s is not of type %s",
						def, field.Type.Name, structtype, d.Typ()))
				}
				f.default_ = d
			}
			enum := field.Type.Tag.Get("enum")
			if enum != "" {
				enumVals := strings.Split(enum, "|")
				enums := make([]Type, len(enumVals))
				for i, en := range enumVals {
					e := f.fieldSpec.new()
					err := e.Scan(en)
					if err != nil {
						panic(fmt.Sprintf("enum value %s for field %s in struct %s is not of type %s",
							en, field.Type.Name, structtype, e.Typ()))
					}
					enums[i] = e
				}
				f.enum = enums
			}
			field.Value.Set(reflect.ValueOf(f))
		}
	}

	st, err := meta.StructByValue(reflect.ValueOf(østruct))
	if err != nil {
		panic(err.Error())
	}
	st.Each(fn)
	// meta.
	// meta.Struct.Each(fn)
	// meta.Struct.EachRaw(østruct, fn)
	return østruct
}

// prefills the given newstruct based on the given prototype
func New(øprototype interface{}, ønewstruct interface{}) (ønew interface{}) {
	prototype := reflect.TypeOf(øprototype).String()
	newtype := reflect.TypeOf(ønewstruct).String()
	if prototype != newtype {
		panic(fmt.Sprintf("prototype (%s) does and new type (%s) are not the same", prototype, newtype))
	}
	proto := reflect.ValueOf(øprototype).Elem()
	//fn := func(field reflect.StructField, val reflect.Value) {
	fn := func(field *meta.Field) {
		if _, ok := field.Value.Interface().(*Field); ok {
			field.Value.Set(
				reflect.ValueOf(
					proto.FieldByName(field.Type.Name).
						Interface().(*Field).
						New(ønewstruct),
				),
			)
		}
	}
	// meta.Struct.EachRaw(ønewstruct, fn)
	st, err := meta.StructByValue(reflect.ValueOf(ønewstruct))
	if err != nil {
		panic(err.Error())
	}
	st.Each(fn)
	return ønewstruct
}
