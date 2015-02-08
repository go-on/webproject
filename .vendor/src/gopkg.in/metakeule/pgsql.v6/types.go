package pgsql

import (
	"fmt"
	"gopkg.in/metakeule/fmtdate.v1"
	"strconv"
	"time"
	// 	"encoding/xml"
	"gopkg.in/metakeule/typeconverter.v2"
	"strings"
)

type SqlType string

func Sql(s string) SqlType {
	return SqlType(s)
}

// converts a string with formatting rules to sql string via fmt.Sprintf
func Sqlf(s string, v ...interface{}) SqlType {
	return Sql(fmt.Sprintf(s, v...))
}

type Sqler interface {
	Sql() SqlType
}

func (ø SqlType) Sql() SqlType {
	return ø
}

func (ø SqlType) String() string {
	return string(ø)
}

func (ø SqlType) TypedValue() *TypedValue {
	return &TypedValue{TextType, ø, true}
}

var TypeNames = map[Type]string{
	TextType:        "text",
	IntType:         "int",
	FloatType:       "float",
	BoolType:        "bool",
	TimeStampTZType: "timestamptz",
	TimeStampType:   "timestamp",
	DateType:        "date",
	TimeType:        "time",
	XmlType:         "xml",
	HtmlType:        "text",
	IntsType:        "integer[]",
	StringsType:     "character varying[]",
	BoolsType:       "boolean[]",
	//FloatsType:      "double precision[]",
	FloatsType:       "float[]",
	TimeStampsTZType: "timestamptz[]",
	UuidType:         "uuid",
	LtreeType:        "ltree",
	TriggerType:      "trigger",
	JsonType:         "json",
}

// uuid NOT NULL DEFAULT uuid_generate_v4()

var TypeCompatibles = map[Type][]Type{
	IntType:          []Type{IntType},
	IntsType:         []Type{IntsType, TextType},
	StringsType:      []Type{StringsType, TextType},
	BoolsType:        []Type{BoolsType, TextType},
	FloatsType:       []Type{FloatsType, TextType},
	TimeStampsTZType: []Type{TimeStampsTZType, TextType},
	FloatType:        []Type{IntType, FloatType},
	TextType:         []Type{TextType, XmlType},
	BoolType:         []Type{BoolType},
	DateType:         []Type{TextType, DateType, TimeType, TimeStampTZType, TimeStampType},
	TimeType:         []Type{TextType, DateType, TimeType, TimeStampTZType, TimeStampType},
	XmlType:          []Type{XmlType, TextType},
	HtmlType:         []Type{HtmlType, TextType},
	UuidType:         []Type{UuidType, TextType},
	LtreeType:        []Type{LtreeType, TextType},
	TriggerType:      []Type{TriggerType, TextType},
	TimeStampTZType:  []Type{TextType, DateType, TimeType, TimeStampTZType, TimeStampType},
	TimeStampType:    []Type{TextType, DateType, TimeType, TimeStampTZType, TimeStampType},
	JsonType:         []Type{JsonType, TextType},
}

func (ø Type) IsCompatible(other Type) bool {
	if (IsVarChar(ø) && IsVarChar(other)) ||
		(IsVarChar(ø) && other == TextType) ||
		(ø == TextType && IsVarChar(other)) {
		return true
	}
	compatible := TypeCompatibles[ø]
	for _, comp := range compatible {
		if comp == other {
			return true
		}
	}
	return false
}

type Type int

const (
	NullType Type = iota + 256 // starting from 256 to allow VarChar to have numbers from 1 to 255
	IntType
	FloatType
	TextType
	BoolType
	DateType
	TimeType
	XmlType
	TimeStampTZType
	TimeStampType
	IntsType
	StringsType
	HtmlType
	UuidType
	LtreeType
	TriggerType
	BoolsType
	FloatsType
	TimeStampsTZType
	JsonType
)

var TypeConverter = NewTypeConverter()

func Convert(in interface{}, out interface{}) (err error) { return TypeConverter.Convert(in, out) }

func ToSql(i interface{}) Sqler {
	if s, ok := i.(Sqler); ok {
		return s
	}
	out := &TypedValue{}
	err := Convert(i, out)
	if err != nil {
		panic("can't convert to sql: " + err.Error())
	}
	return out
}

var TypeDefaults = map[Type]interface{}{
	TextType:         "",
	IntType:          0,
	FloatType:        float32(0),
	BoolType:         false,
	TimeStampTZType:  time.Now(),
	TimeStampType:    time.Now(),
	DateType:         time.Now(),
	TimeType:         time.Now(),
	XmlType:          "",
	HtmlType:         "",
	IntsType:         []int{},
	StringsType:      []string{},
	UuidType:         "",
	BoolsType:        []bool{},
	FloatsType:       []float64{},
	TimeStampsTZType: []time.Time{},
	JsonType:         "{}",
	// LtreeType:       "ltree",
	// TriggerType:     "trigger",
}

func (ø Type) Default() interface{} {
	return TypeDefaults[ø]
}

func (ø Type) String() string { return TypeNames[ø] }
func (ø Type) Type() Type     { return ø }

var intInstance = int(0)
var int32Instance = int32(0)
var int64Instance = int64(0)
var intsInstance = []int{}
var float64Instance = float64(0)
var float32Instance = float32(0)
var stringInstance = string("")
var stringsInstance = []string{}
var jsonInstance = typeconverter.Json("")
var boolInstance = bool(true)
var timeInstance = time.Time{}
var mapInstance = map[string]interface{}{}
var arrInstance = []interface{}{}
var typedValueInstance = TypedValue{}
var sqlInstance = Sql("")
var typeInstance = Type(0)
var boolsInstance = []bool{}
var floatsInstance = []float64{}
var timesInstance = []time.Time{}

type intsStringer []int

func (ø intsStringer) String() string {
	s := `{%s}`
	str := []string{}
	for _, i := range ø {
		str = append(str, fmt.Sprintf("%v", i))
	}
	return fmt.Sprintf(s, strings.Join(str, ","))
}

func (ø intsStringer) Ints() []int {
	return []int(ø)
}

type Intser interface {
	Ints() []int
}

type stringsStringer []string

func (ø stringsStringer) String() string {
	s := `{%s}`
	str := []string{}
	for _, i := range ø {
		str = append(str, fmt.Sprintf(`"%v"`, i))
	}
	return fmt.Sprintf(s, strings.Join(str, ","))
}

func (ø stringsStringer) Strings() []string {
	return []string(ø)
}

type Stringser interface {
	Strings() []string
}

type boolsStringer []bool

func (ø boolsStringer) String() string {
	s := `{%s}`
	str := []string{}
	for _, i := range ø {
		str = append(str, fmt.Sprintf("%v", i))
	}
	return fmt.Sprintf(s, strings.Join(str, ","))
}

func (ø boolsStringer) Bools() []bool {
	return []bool(ø)
}

type Boolser interface {
	Bools() []bool
}

type floatsStringer []float64

func (ø floatsStringer) String() string {
	s := `{%s}`
	str := []string{}
	for _, i := range ø {
		str = append(str, fmt.Sprintf("%v", i))
	}
	return fmt.Sprintf(s, strings.Join(str, ","))
}

func (ø floatsStringer) Floats() []float64 {
	return []float64(ø)
}

type Floatser interface {
	Floats() []float64
}

type timesStringer []time.Time

func (ø timesStringer) String() string {
	s := `{%s}`
	str := []string{}
	for _, i := range ø {
		str = append(str, fmt.Sprintf("%v", i.Format(time.RFC3339)))
	}
	return fmt.Sprintf(s, strings.Join(str, ","))
}

func (ø timesStringer) Times() []time.Time {
	return []time.Time(ø)
}

type Timeser interface {
	Times() []time.Time
}

/*
var intsMatcher = regexp.MustCompile(`^\{("?[0-9]"?,)*("?[0-9]"?)\}$`)
var stringsMatcher = regexp.MustCompile(`^\{([^,]*,)*([^,]*)\}$`)

func typedValForString(t string) (tv *TypedValue) {
	tv = &TypedValue{}
	if t[0:1] == "{" && t[len(t)-1:len(t)] == "}" {
		if len(intsMatcher.FindStringSubmatch(t)) > 0 {
			tv.PgType = IntsType
			tv.Value = intsStringer(stringToInts(t))
			return
		}
		if len(stringsMatcher.FindStringSubmatch(t)) > 0 {
			tv.PgType = StringsType
			tv.Value = stringsStringer(stringToStrings(t))
			return
		}
	}
	tv.PgType = TextType
	tv.Value = typeconverter.String(t)
	return
}
*/

type pgInterpretedString struct {
	typeconverter.StringType
}

func NewPgInterpretedString(s string) (ip *pgInterpretedString) {
	ip = &pgInterpretedString{}
	ip.StringType = typeconverter.String(s)
	return
}

func (ø *pgInterpretedString) Int() (i int) {
	str := ø.StringType.String()
	i, ſ := strconv.Atoi(str)
	if ſ != nil {
		panic(ſ.Error())
	}
	return
	// return stringToInts(ø.StringType.String())
}

func (ø *pgInterpretedString) Ints() (i []int) {
	str := ø.StringType.String()
	inner := str[1 : len(str)-1]
	a := strings.Split(inner, ",")
	for _, s := range a {
		ii, ſ := strconv.Atoi(strings.Trim(s, `"`))
		if ſ == nil {
			i = append(i, ii)
		}
	}
	return
	// return stringToInts(ø.StringType.String())
}

func (ø *pgInterpretedString) Strings() (ses []string) {
	str := ø.StringType.String()
	inner := str[1 : len(str)-1]
	if inner == "" {
		return
	}
	a := strings.Split(inner, ",")
	for _, s := range a {
		// fmt.Printf("s: %#v\n", s)
		s_tr := strings.Trim(s, `"`)
		ses = append(ses, s_tr)
	}
	return
}

func (ø *pgInterpretedString) Bools() (bs []bool) {
	str := ø.StringType.String()
	inner := str[1 : len(str)-1]
	a := strings.Split(inner, ",")
	for _, s := range a {
		b, ſ := strconv.ParseBool(strings.Trim(s, `"`))
		if ſ == nil {
			bs = append(bs, b)
		}
	}
	return
}

func (ø *pgInterpretedString) Floats() (f []float64) {
	str := ø.StringType.String()
	inner := str[1 : len(str)-1]
	a := strings.Split(inner, ",")
	for _, s := range a {
		ff, ſ := strconv.ParseFloat(strings.Trim(s, `"`), 64)
		if ſ == nil {
			f = append(f, ff)
		}
	}
	return
	// return stringToInts(ø.StringType.String())
}

func (ø *pgInterpretedString) Times() (f []time.Time) {
	str := ø.StringType.String()
	inner := str[1 : len(str)-1]
	a := strings.Split(inner, ",")
	for _, s := range a {
		ff, ſ := time.Parse(time.RFC3339, strings.Trim(s, `"`))
		//ff, ſ := strconv.ParseFloat(strings.Trim(s, `"`), 64)
		if ſ == nil {
			f = append(f, ff)
		}
	}
	return
	// return stringToInts(ø.StringType.String())
}

//typeconverter.String

func NewTypeConverter() (ø *typeconverter.BasicConverter) {
	ø = typeconverter.New()

	inSwitch := func(from interface{}, to interface{}) (err error) {
		// fmt.Printf("-- convert %v (%T)-- to %T\n", from, from, to)
		switch t := from.(type) {
		//case Placeholder:
		//	err = ø.Output.Dispatch(to, &TypedValue{TextType, t, true})
		case int:
			err = ø.Output.Dispatch(to, &TypedValue{IntType, typeconverter.Int(t), false})
		case int32:
			err = ø.Output.Dispatch(to, &TypedValue{IntType, typeconverter.Int(int(t)), false})
		case int64:
			err = ø.Output.Dispatch(to, &TypedValue{IntType, typeconverter.Int64(t), false})
		case float64:
			err = ø.Output.Dispatch(to, &TypedValue{FloatType, typeconverter.Float(t), false})
		case float32:
			err = ø.Output.Dispatch(to, &TypedValue{FloatType, typeconverter.Float32(t), false})
		case string:
			// err = ø.Output.Dispatch(to, typedValForString(t))
			//fmt.Printf("as interpreted string: %#v\n", t)
			err = ø.Output.Dispatch(to, &TypedValue{TextType, NewPgInterpretedString(t), false})
			//err = ø.Output.Dispatch(to, &TypedValue{FloatType, typeconverter.String(t)})
		case bool:
			err = ø.Output.Dispatch(to, &TypedValue{BoolType, typeconverter.Bool(t), false})
		case time.Time:
			err = ø.Output.Dispatch(to, &TypedValue{TimeStampTZType, typeconverter.Time(t), false})
		case []int:
			err = ø.Output.Dispatch(to, &TypedValue{IntsType, intsStringer(t), false})
		case []string:
			// fmt.Printf("strings: %#v\n", t)
			err = ø.Output.Dispatch(to, &TypedValue{StringsType, stringsStringer(t), false})
		case []bool:
			err = ø.Output.Dispatch(to, &TypedValue{BoolsType, boolsStringer(t), false})
		case []float64:
			err = ø.Output.Dispatch(to, &TypedValue{FloatsType, floatsStringer(t), false})
		case []time.Time:
			err = ø.Output.Dispatch(to, &TypedValue{TimeStampsTZType, timesStringer(t), false})
		case *int:
			err = ø.Output.Dispatch(to, &TypedValue{IntType, typeconverter.Int(*t), false})
		case *int32:
			err = ø.Output.Dispatch(to, &TypedValue{IntType, typeconverter.Int(int(*t)), false})
		case *int64:
			err = ø.Output.Dispatch(to, &TypedValue{IntType, typeconverter.Int64(*t), false})
		case *float64:
			err = ø.Output.Dispatch(to, &TypedValue{FloatType, typeconverter.Float(*t), false})
		case *float32:
			err = ø.Output.Dispatch(to, &TypedValue{FloatType, typeconverter.Float32(*t), false})
		case *string:
			//fmt.Printf("as interpreted *string: %#v\n", *t)
			err = ø.Output.Dispatch(to, &TypedValue{TextType, NewPgInterpretedString(*t), false})
			// err = ø.Output.Dispatch(to, &TypedValue{TextType, typeconverter.String(*t)})
			// err = ø.Output.Dispatch(to, typedValForString(*t))
		case *bool:
			err = ø.Output.Dispatch(to, &TypedValue{BoolType, typeconverter.Bool(*t), false})
		case *time.Time:
			err = ø.Output.Dispatch(to, &TypedValue{TimeStampTZType, typeconverter.Time(*t), false})
		case *TypedValue:
			err = ø.Output.Dispatch(to, t)
		case TypedValue:
			err = ø.Output.Dispatch(to, &t)
		case SqlType:
			err = ø.Output.Dispatch(to, t)
		case Type:
			err = ø.Output.Dispatch(to, t)
		default:
			err = ø.Output.Dispatch(to, &TypedValue{TextType, from.(typeconverter.Stringer), false})
		}
		return
	}

	ø.Input.SetHandler(intInstance, inSwitch)
	ø.Input.SetHandler(int32Instance, inSwitch)
	ø.Input.SetHandler(int64Instance, inSwitch)
	ø.Input.SetHandler(float64Instance, inSwitch)
	ø.Input.SetHandler(float32Instance, inSwitch)
	ø.Input.SetHandler(stringInstance, inSwitch)
	ø.Input.SetHandler(boolInstance, inSwitch)
	ø.Input.SetHandler(timeInstance, inSwitch)
	ø.Input.SetHandler(&intInstance, inSwitch)
	ø.Input.SetHandler(&int32Instance, inSwitch)
	ø.Input.SetHandler(&int64Instance, inSwitch)
	ø.Input.SetHandler(&float64Instance, inSwitch)
	ø.Input.SetHandler(&float32Instance, inSwitch)
	ø.Input.SetHandler(&stringInstance, inSwitch)
	ø.Input.SetHandler(&boolInstance, inSwitch)
	ø.Input.SetHandler(&timeInstance, inSwitch)
	ø.Input.SetHandler(jsonInstance, inSwitch)
	ø.Input.SetHandler(typedValueInstance, inSwitch)
	ø.Input.SetHandler(&typedValueInstance, inSwitch)
	ø.Input.SetHandler(typeInstance, inSwitch)
	ø.Input.SetHandler(intsInstance, inSwitch)
	ø.Input.SetHandler(stringsInstance, inSwitch)
	ø.Input.SetHandler(boolsInstance, inSwitch)
	ø.Input.SetHandler(floatsInstance, inSwitch)
	ø.Input.SetHandler(timesInstance, inSwitch)

	outSwitch := func(out interface{}, in interface{}) (err error) {
		// fmt.Printf("in: %#v (%T) out: %#v (%T)\n", in, in, out, out)
		switch t := out.(type) {
		case *TypedValue:
			iTyped := in.(Valuer).TypedValue()
			oTyped := out.(*TypedValue)
			// fmt.Printf("in: %#v (%#v) out: %#v (%#v)\n", iTyped, iTyped.String(), oTyped, oTyped.String())
			if int(oTyped.Type()) == 0 {
				*oTyped = *iTyped
				return
			}

			// fmt.Printf("iTyped %#v oTyped %#v\n", iTyped.Type().String(), oTyped.Type().String())
			if iTyped.Type() == oTyped.Type() {
				*oTyped = *iTyped
			} else {
				if oTyped.Type().IsCompatible(iTyped.Type()) {
					oTyped.Value = iTyped.Value
				} else {
					return fmt.Errorf("value %s type %s is incompatible with type %s\n%s", iTyped.Value.String(), iTyped.Type().String(), oTyped.Type().String(), strings.Join(backtrace(), "\n"))
				}
			}
		case *bool:
			*out.(*bool) = in.(*TypedValue).Value.(typeconverter.Booler).Bool()
		case *int:
			//fmt.Printf("typed values: %#v\n", in.(*TypedValue))
			*out.(*int) = in.(*TypedValue).Value.(typeconverter.Inter).Int()
		case *int64:
			*out.(*int64) = int64(in.(*TypedValue).Value.(typeconverter.Inter).Int())
		case *string:
			*out.(*string) = in.(typeconverter.Stringer).String()
		case *float64:
			*out.(*float64) = in.(*TypedValue).Value.(typeconverter.Floater).Float()
		case *float32:
			*out.(*float32) = float32(in.(*TypedValue).Value.(typeconverter.Floater).Float())
		case *time.Time:
			*out.(*time.Time) = in.(*TypedValue).Value.(typeconverter.Timer).Time()
		case *[]int:
			*out.(*[]int) = in.(*TypedValue).Value.(Intser).Ints()
		case *[]string:
			// fmt.Printf("stringser %T\n", in.(*TypedValue).Value)
			*out.(*[]string) = in.(*TypedValue).Value.(Stringser).Strings()
		case *[]bool:
			*out.(*[]bool) = in.(*TypedValue).Value.(Boolser).Bools()
		case *[]float64:
			*out.(*[]float64) = in.(*TypedValue).Value.(Floatser).Floats()
		case *[]time.Time:
			*out.(*[]time.Time) = in.(*TypedValue).Value.(Timeser).Times()
		case *SqlType:
			*out.(*SqlType) = in.(Sqler).Sql()
		case *Type:
			*out.(*Type) = in.(Typer).Type()
		default:
			return fmt.Errorf("can't convert to %#v: no converter found", t)
		}
		return
	}

	ø.Output.SetHandler(&intInstance, outSwitch)
	ø.Output.SetHandler(&int32Instance, outSwitch)
	ø.Output.SetHandler(&int64Instance, outSwitch)
	ø.Output.SetHandler(&float64Instance, outSwitch)
	ø.Output.SetHandler(&float32Instance, outSwitch)
	ø.Output.SetHandler(&stringInstance, outSwitch)
	ø.Output.SetHandler(&boolInstance, outSwitch)
	ø.Output.SetHandler(&timeInstance, outSwitch)
	ø.Output.SetHandler(&jsonInstance, outSwitch)
	ø.Output.SetHandler(&typedValueInstance, outSwitch)
	ø.Output.SetHandler(&typeInstance, outSwitch)
	ø.Output.SetHandler(&sqlInstance, outSwitch)
	ø.Output.SetHandler(&intsInstance, outSwitch)
	ø.Output.SetHandler(&stringsInstance, outSwitch)
	ø.Output.SetHandler(&boolsInstance, outSwitch)
	ø.Output.SetHandler(&floatsInstance, outSwitch)
	ø.Output.SetHandler(&timesInstance, outSwitch)
	return
}

type Typer interface {
	Type() Type
}

type TypedValue struct {
	PgType     Type
	Value      typeconverter.Stringer
	dontChange bool
}

type Valuer interface {
	TypedValue() *TypedValue
}

func (ø *TypedValue) TypedValue() *TypedValue {
	return ø
}

func (ø *TypedValue) IsNil() bool {
	if ø == nil || ø.Value == nil {
		return true
	}
	return false
}

func (ø *TypedValue) Sql() SqlType {
	if ø.IsNil() {
		return Sql("Null")
	}
	if ø.dontChange {
		return Sql(ø.Value.String())
	}
	val := Escape(ø.Value.String())
	return Sql(fmt.Sprintf("%s::%s", val, ø.PgType.String()))
}

func (ø *TypedValue) String() string { return ø.Value.String() }
func (ø *TypedValue) Type() Type     { return ø.PgType }

func VarChar(i int) Type {
	if i > 255 {
		panic("varchar may not be larger than 255")
	}
	if i < 1 {
		panic("varchar may not be smaller than 1")
	}
	t := Type(i)
	TypeNames[t] = fmt.Sprintf("varchar(%v)", i)
	return t
}

// is the type a varchar
func IsVarChar(t Type) bool {
	if i := int(t); i < 256 && i > 0 {
		return true
	}
	return false
}

type NullTime struct {
	*time.Time
	Valid bool
}

func (n *NullTime) Scan(value interface{}) (err error) {
	//fmt.Printf("nulltime scan: %#v\n", value)
	if value == nil {
		n.Time = nil
		n.Valid = false
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		n.Time = &v
		n.Valid = true
	case *time.Time:
		n.Time = v
		n.Valid = true
	case string:
		if v == "" {
			n.Time = nil
			n.Valid = false
			return nil
		}
		n.Valid = true
		var t time.Time
		t, err = fmtdate.Parse("YYYY-MM-DD hh:mm:ss", v)
		n.Time = &t
	case []byte:
		s := string(v)
		if s == "" {
			n.Time = nil
			n.Valid = false
			return nil
		}
		n.Valid = true
		var t time.Time
		t, err = fmtdate.Parse("YYYY-MM-DD hh:mm:ss", s)
		n.Time = &t
	default:
		n.Time = nil
		n.Valid = false
		return fmt.Errorf("unsupported type for time: %T", value)
	}

	return
}

func (n *NullTime) TypedValue() *TypedValue {
	if n.Valid {
		return &TypedValue{TimeType, typeconverter.Time(*n.Time), false}
	}
	//	return &TypedValue{NullType, Sql("NULL"), false}
	return &TypedValue{NullType, nil, false}
}
