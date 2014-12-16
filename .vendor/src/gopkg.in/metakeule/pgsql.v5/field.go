package pgsql

import (
	"fmt"
	"runtime"
)

type Flag int

const (
	_                    = iota
	hasDefaults     Flag = 1 << iota
	NullAllowed          // field may have null values
	PrimaryKey           // field is primary key
	Indexed              // field is indexed
	Unique               // field is unique
	Serial               // field is a serial field
	UuidGenerate         // generate a uuid
	OnDeleteCascade      // fkey is on delete cascade (default: restrict)
)

type SelectionArray []interface{}

func Selection(o ...interface{}) SelectionArray {
	return SelectionArray(o)
}

func backtrace() (btr []string) {
	for i := 0; i < 100; i++ {
		pc, file, line, _ := runtime.Caller(2 + i)
		if file == "" {
			continue
		}
		f := runtime.FuncForPC(pc)
		if f != nil {
			btr = append(btr, fmt.Sprintf("%v: %v\n\t%v()", file, line, f.Name()))
			continue
		}
		btr = append(btr, fmt.Sprintf("%v: %v", file, line))
	}
	return
}

type Field struct {
	Name        string
	flags       Flag
	Default     Sqler
	Type        Type
	Table       *Table
	ForeignKey  *Field
	Selection   SelectionArray
	Validations []FieldValidator
	queryField  string // name of a field in a query struct that should match this field
}

func NewField(name string, options ...interface{}) *Field {
	ø := &Field{
		Name:        name,
		flags:       hasDefaults,
		Validations: []FieldValidator{},
	}
	ø.AddValidator(&TypeValidator{ø})
	ø.Add(options...)
	if len(ø.Selection) > 0 {
		ø.AddValidator(&SelectionValidator{ø})
	}
	return ø
}

func (ø *Field) InSelection(value interface{}) bool {
	if ø.Selection == nil {
		return true
	}
	asString := ToString(value)
	for _, s := range ø.Selection {
		if ToString(s) == asString {
			return true
		}
	}
	return false
}

// sets the queryField, allows chaining
func (ø *Field) SetQueryField(f string) *Field {
	if ø.queryField != "" {
		panic("queryField already set to " + ø.queryField + ", can't change")
	}
	ø.queryField = f
	return ø
}

func (ø *Field) QueryField() string {
	return ø.queryField
}

func (ø *Field) AddValidator(v ...FieldValidator) {
	for _, val := range v {
		ø.Validations = append(ø.Validations, val)
	}
}

// return the value in a typed fashion converted to
// the required postgres type
func (ø *Field) Value(val interface{}) (tv *TypedValue, err error) {
	if val == nil {
		if ø.Is(NullAllowed) {
			tv = &TypedValue{PgType: ø.Type}
			return
		} else {
			err = nullNotAllowedError(ø, val)
			return
		}
	}
	tv = &TypedValue{PgType: ø.Type}
	e := Convert(val, tv)
	if e != nil {
		err = convertError(ø, val, e)
	}
	return
}

func (ø *Field) MustValue(val interface{}) (tv *TypedValue) {
	var e error
	tv, e = ø.Value(val)
	if e != nil {
		panic(e.Error())
	}
	return
}

func (ø *Field) Validate(value interface{}) (err error) {
	if _, ok := value.(*pgInterpretedString); ok {
		return
	}
	for _, v := range ø.Validations {
		err = v.Validate(value)
		if err != nil {
			err = fieldError(ø, err)
			return
		}
	}
	return
}

func (ø *Field) Add(options ...interface{}) {
	for _, option := range options {
		switch v := option.(type) {
		case *Table:
			ø.Table = v
		case Type:
			ø.Type = v
		case Flag:
			ø.flags = ø.flags | v
		case *Field:
			ø.ForeignKey = v
			ø.Type = v.Type
		case SelectionArray:
			ø.Selection = v
		default:
			if val, ok := v.(FieldValidator); ok {
				ø.Validations = append(ø.Validations, val)
				continue
			}
			if sqler, ok := v.(Sqler); ok {
				ø.Default = sqler
			} else {
				panic("unknown type for field " + fmt.Sprintf("%v\n", v))
			}
		}
	}
}

// checks if a given flag is set, e.g.
//
// 	Is(NullAllowed)
//
// checks is null is allowed
func (ø *Field) Is(f Flag) bool {
	return ø.flags&f != 0
}

func (ø *Field) Sql() SqlType {
	if ø.Table == nil {
		return Sql(fmt.Sprintf(`"%s"`, ø.Name))
	}
	return Sql(fmt.Sprintf(`%s."%s"`, ø.Table.Sql(), ø.Name))
}
