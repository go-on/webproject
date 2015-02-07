package templfat

import (
	"reflect"

	"gopkg.in/go-on/lib.v3/internal/fat"
	"gopkg.in/go-on/lib.v3/internal/meta"
	"gopkg.in/go-on/lib.v3/internal/template"
	"gopkg.in/go-on/lib.v3/internal/template/placeholder"
)

/*
   Support placeholders for fat fields
*/

// create a placeholder from a fat Field
func Placeholder(øField *fat.Field) placeholder.Placeholder {
	return template.NewPlaceholder(øField.Path())
}

// create a single setter from a fat Field
func Setter(øField *fat.Field) placeholder.Setter {
	t := template.NewPlaceholder(øField.Path())
	tph := t.(template.TemplatePlaceholder)
	tph.Value = øField.String()
	return tph
}

// create a slice of Setters from a struct with fat Fields
// only the attributes that are fat Fields are respected
func Setters(østruct interface{}) []placeholder.Setter {
	phs := []placeholder.Setter{}
	fn := func(field *meta.Field) {
		f, ok := field.Value.Interface().(*fat.Field)
		if ok {
			phs = append(phs, Setter(f))
		}
	}

	stru, err := meta.StructByValue(reflect.ValueOf(østruct))
	if err != nil {
		panic(err.Error())
	}
	stru.Each(fn)

	//	meta.Struct.Each(østruct, fn)
	return phs
}
