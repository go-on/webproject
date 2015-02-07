package pgsql

import (
	"fmt"

	"gopkg.in/go-on/lib.v3/internal/meta"
	// "gopkg.in/metakeule/meta.v5"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	tableType     = reflect.TypeOf(&Table{})
	fieldType     = reflect.TypeOf(&Field{})
	varcharRegexp = regexp.MustCompile(`^varchar\(([0-9]{0,3})\)$`)
)

type tableDefinition struct {
	Table    *Table
	Fields   map[string]*Field
	Uniques  map[string][]string
	finished bool
}

func (td *tableDefinition) MustFinish() {
	err := td.Finish()

	if err != nil {
		panic(err.Error())
	}

}

func (td *tableDefinition) Finish() error {
	if td.finished {
		return fmt.Errorf("Finish called twice")
	}
	for _, field := range td.Fields {
		td.Table.AddField(field)
	}

	for unique, fields := range td.Uniques {
		fs := make([]*Field, len(fields))

		for i, fld := range fields {
			f, ok := td.Fields[fld]
			if !ok {
				return fmt.Errorf("invalid unique %s: non existing field: %s", unique, fld)
			}
			fs[i] = f
		}
		td.Table.AddUnique(unique, fs...)
	}
	return nil
}

func TableDefinition(strPtr interface{}) (td *tableDefinition, err error) {
	td = &tableDefinition{}
	td.Fields = map[string]*Field{}
	td.Uniques = map[string][]string{}
	// maps fieldname to field
	st, err := meta.StructByValue(reflect.ValueOf(strPtr))
	if err != nil {
		return
	}
	st.Each(func(f *meta.Field) {

		field := f.Type
		val := f.Value

		switch field.Type {
		case tableType:
			tName := field.Name
			if name := field.Tag.Get("name"); name != "" {
				tName = name
			}
			td.Table = NewTable(tName)

			val.Set(reflect.ValueOf(td.Table))
			if un := field.Tag.Get("unique"); un != "" {
				for _, uniq := range strings.Split(un, ",") {
					td.Uniques[uniq] = strings.Split(uniq, "#")
				}
			}
		case fieldType:
			fName := field.Name
			if name := field.Tag.Get("name"); name != "" {
				fName = name
			}
			options := []interface{}{}

			type_ := field.Tag.Get("type")
			if type_ == "" {
				err = fmt.Errorf("no type tag set for field %s in table definition %T", field.Name, strPtr)
				return
			}

			var ty Type
			switch type_ {
			case "int":
				ty = IntType
			case "float":
				ty = FloatType
			case "text":
				ty = TextType
			case "bool":
				ty = BoolType
			case "timestamptz":
				ty = TimeStampTZType
			case "timestamp":
				ty = TimeType
			case "date":
				ty = DateType
			case "time":
				ty = TimeType
			case "xml":
				ty = XmlType
			case "integer[]":
				ty = IntsType
			case "character varying[]":
				ty = StringsType
			case "uuid":
				ty = UuidType
			case "ltree":
				ty = LtreeType
			case "trigger":
				ty = TriggerType
			default:
				md := varcharRegexp.FindStringSubmatch(type_)
				if len(md) == 2 {
					i, e := strconv.Atoi(md[1])
					if e != nil {
						err = fmt.Errorf("error in varchar type tag for field %s in table definition %T, can't parse integer", field.Name, strPtr)
						return
					}
					ty = VarChar(i)
				} else {
					err = fmt.Errorf("error in type tag for field %s in table definition %T, unknown type: %s", field.Name, strPtr, type_)
					return
				}
			}

			options = append(options, ty)

			if flags := field.Tag.Get("flag"); flags != "" {
				for _, fl := range strings.Split(flags, ",") {
					switch fl {
					case "null":
						options = append(options, NullAllowed)
					case "pkey":
						options = append(options, PrimaryKey)
					case "unique":
						options = append(options, Unique)
					case "index":
						options = append(options, Indexed)
					case "serial":
						options = append(options, Serial)
					case "uuidgenerate":
						options = append(options, UuidGenerate)
					default:
						err = fmt.Errorf("error in flags tag for field %s in table definition %T, unknown flag: %s", field.Name, strPtr, fl)
						return
					}
				}
			}

			if enum := field.Tag.Get("enum"); enum != "" {
				enums := strings.Split(enum, ",")
				sel := make([]interface{}, len(enums))
				for i, en := range enums {
					sel[i] = en
				}
				options = append(options, SelectionArray(sel))
			}

			f := NewField(fName, options...).SetQueryField(field.Name)
			td.Fields[fName] = f
			val.Set(reflect.ValueOf(f))
		}
	})
	// meta.Struct.EachRaw(strPtr, func(field reflect.StructField, val reflect.Value) {

	return
}

func MustTableDefinition(strPtr interface{}) *tableDefinition {
	td, err := TableDefinition(strPtr)

	if err != nil {
		panic(fmt.Sprintf("error in table definition for %T: %s\n", strPtr, err.Error()))
	}
	return td
}

/*
type area struct {
    TABLE               *Table  `name:"koelnart-area_focus" unique:"area#focus#work"`
    Area                *Field  `name:"area"     type:"varchar(255)" flag:"pkey" enum:"foto-own,foto-other"`
    Work                *Field  `name:"work"     type:"int"`
    Position            *Field  `name:"position" type:"int",         flag:"pkey"`
}

var Area = &area{}
*/

/*
func init() {
    td := MustTableDefinition(Area)
    // set foreign key
    td["work"].Add(work.Id, OnDeleteCascade)
    // set default value
    td["area"].Add(Sql("default area"))
    td.MustFinish()
}
*/
