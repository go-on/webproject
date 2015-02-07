package pgsql

import (
	"fmt"
	"reflect"

	"gopkg.in/go-on/lib.v3/internal/template"
	"gopkg.in/go-on/lib.v3/internal/template/placeholder"
	// "github.com/metakeule/fastreplace"
	// "github.com/metakeule/templ"

	"strings"
)

type Limit int

func (ø Limit) Sql() SqlType {
	if ø == 0 {
		return Sql("")
	}
	return Sql(fmt.Sprintf("\nLIMIT %v", ø))
}

type Offset int

func (ø Offset) Sql() (s SqlType) {
	if ø == 0 {
		s = Sql("")
	} else {
		s = Sql(fmt.Sprintf("\nOffset %v", ø))
	}
	return
}

func Escape(in string) (out string) {
	out = strings.Replace(in, "$userinput$", "", -1)
	out = "$userinput$" + out + "$userinput$"
	return
}

func (ø SqlType) Placeholder() Placeholder {
	t := template.NewPlaceholder("sql."+ø.String(), handleSql)
	// t := templ.NewPlaceholder("sql."+ø.String(), handleSql)
	return typedPlaceholder{t}
}

func handleSql(in interface{}) (out string) {
	switch v := in.(type) {
	case string:
		out = v
	case Sqler:
		out = v.Sql().String()
	case Stringer:
		out = v.String()
	default:
		out = fmt.Sprintf("%v", v)
		//panic("unsupported type: " + fmt.Sprintf("%v (%T)", v, v))
	}
	return
}

var Transformer = map[string]func(interface{}) string{
	"value":    escapeValue,
	"sql":      handleSql,
	"%search":  escapeSearchEnd,
	"search%":  escapeSearchStart,
	"%search%": escapeSearchBoth,
}

func escapeValue(in interface{}) (out string) {
	switch v := in.(type) {
	case string:
		out = Escape(v)
	case TypedValue:
		out = Escape(v.Value.String())
	case Sqler:
		out = v.Sql().String()
	case Stringer:
		out = Escape(v.String())
	default:
		out = Escape(fmt.Sprintf("%v", v))
	}
	return
}

type Placeholder interface {
	placeholder.Setter
	Sqler
	Set(val interface{}) placeholder.Setter
	Setf(format string, val ...interface{}) placeholder.Setter
	// String() string
}

type typedPlaceholder struct{ placeholder.Placeholder }

func (ø typedPlaceholder) Sql() SqlType {
	return Sql("@@" + ø.Placeholder.Name() + "@@")
	//return Sql(ø.Placeholder.String())
}

type SearchStart string

func (ø SearchStart) Placeholder() Placeholder {
	t := template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø), escapeSearchStart)
	return typedPlaceholder{t}
}

type SearchEnd string

func (ø SearchEnd) Placeholder() Placeholder {
	t := template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø), escapeSearchEnd)
	return typedPlaceholder{t}
}

type SearchBoth string

func (ø SearchBoth) Placeholder() Placeholder {
	t := template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø), escapeSearchBoth)
	return typedPlaceholder{t}
}

func (ø *Field) Placeholder() Placeholder {
	fn := func(in interface{}) (out string) {
		tv := ø.MustValue(in)
		return tv.Sql().String()
	}
	t := template.NewPlaceholder(ø.Table.Name+"."+ø.Name, fn)
	return typedPlaceholder{t}
}

func (ø *TypedValue) Placeholder() Placeholder {
	t := template.NewPlaceholder(ø.PgType.String(), escapeValue)
	return typedPlaceholder{t}
}

func escapeSearchStart(in interface{}) (out string) {
	var inString string
	switch v := in.(type) {
	case string:
		inString = v
	case TypedValue:
		inString = v.Value.String()
	case Sqler:
		inString = v.Sql().String()
	case Stringer:
		inString = v.String()
	default:
		inString = fmt.Sprintf("%v", v)
	}
	return Escape(inString + "%")
}

func escapeSearchEnd(in interface{}) (out string) {
	var inString string
	switch v := in.(type) {
	case string:
		inString = v
	case TypedValue:
		inString = v.Value.String()
	case Sqler:
		inString = v.Sql().String()
	case Stringer:
		inString = v.String()
	default:
		inString = fmt.Sprintf("%v", v)
	}
	return Escape("%" + inString)
}

func escapeSearchBoth(in interface{}) (out string) {
	var inString string
	switch v := in.(type) {
	case string:
		inString = v
	case TypedValue:
		inString = v.Value.String()
	case Sqler:
		inString = v.Sql().String()
	case Stringer:
		inString = v.String()
	default:
		inString = fmt.Sprintf("%v", v)
	}
	return Escape("%" + inString + "%")
}

type CompiledQuery struct {
	//freplace *fastreplace.FReplace
	//*template.Template
	*template.Template
	Query Query
}

/*
type CompiledQueryInstance struct {
	inst fastreplace.Replacer
}

func (ø *CompiledQueryInstance) Assign(key string, sql Sqler) {
	ø.inst.AssignString(key, sql.Sql().String())
}

func (ø *CompiledQueryInstance) AssignValue(key string, val string) {
	userinput := strings.Replace(val, "$userinput$", "", -1)
	ø.inst.AssignString(key, "$userinput$"+userinput+"$userinput$")
}

func (ø *CompiledQueryInstance) Sql() SqlType {
	return Sql(ø.inst.String())
}

func (ø *CompiledQuery) Instance() (c *CompiledQueryInstance) {
	return &CompiledQueryInstance{ø.freplace.Instance()}
}
*/

func Compile(name string, q Query) (c *CompiledQuery, ſ error) {
	t := template.New(name)
	t.MustAdd(q.String())
	t.Parse()

	// ſ = t.Parse()
	//t, ſ := template.New(q.String())
	//t.Strict = true
	c = &CompiledQuery{
		Query:    q,
		Template: t,
	}
	/*
		if ſ != nil {
			t = nil
			return
		}
	*/
	return
}

// panics on error
func MustCompile(name string, q Query) *CompiledQuery {
	c, e := Compile(name, q)
	if e != nil {
		panic(e.Error())
	}
	return c
}

type Query interface {
	Sql() SqlType
	String() string
}

const (
	LeftJoinType JoinType = iota
	RightJoinType
	InnerJoinType
	FullJoinType
)

type JoinType int

//type Placeholder int

// you can use Placeholder(1), as value to update insert etc. in order
// to get a prepare statement, that you may execute later on
/*
func (ø Placeholder) Sql() SqlType {
	return Sql(fmt.Sprintf("$%v", ø))
}
*/

type Comparer struct {
	A    Sqler
	B    Sqler
	Sign string
}

func (ø *Comparer) Sql() SqlType {
	return Sql(fmt.Sprintf("%s %s %s", ø.A.Sql(), ø.Sign, ø.B.Sql()))
}

type nuller struct {
	obj  Sqler
	null bool
}

func (ø *nuller) Sql() SqlType {
	if ø.null {
		return Sql(fmt.Sprintf("%s IS NULL", ø.obj.Sql()))
	}
	return Sql(fmt.Sprintf("%s IS NOT NULL", ø.obj.Sql()))
}

func IsNull(s Sqler) *nuller    { return &nuller{s, true} }
func IsNotNull(s Sqler) *nuller { return &nuller{s, false} }

func Like(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), "LIKE"}
}

func Equals(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), "="}
}

func EqualsNot(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), "!="}
}

func GreaterThan(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), ">"}
}

func LessThan(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), "<"}
}

func GreaterThanEqual(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), ">="}
}

func LessThanEqual(a interface{}, b interface{}) *Comparer {
	return &Comparer{ToSql(a), ToSql(b), "<="}
}

type InComparer struct {
	A  Sqler
	Bs []Sqler
}

func In(a interface{}, bs ...interface{}) *InComparer {
	bs_converted := []Sqler{}
	for _, e := range bs {
		bs_converted = append(bs_converted, ToSql(e))
	}
	return &InComparer{ToSql(a), bs_converted}
}

func (ø *InComparer) Sql() SqlType {
	bs := []string{}
	for _, b := range ø.Bs {
		bs = append(bs, string(b.Sql()))
	}
	return Sql(fmt.Sprintf("%s In(%s)", ø.A.Sql(), strings.Join(bs, ", ")))
}

type Condition struct {
	Conditions []Sqler
	Sign       string
}

func Or(sqls ...Sqler) *Condition {
	return &Condition{sqls, "OR"}
}

func And(sqls ...Sqler) *Condition {
	return &Condition{sqls, "AND"}
}

func (ø *Condition) Sql() (s SqlType) {
	if len(ø.Conditions) == 0 {
		s = Sql("")
	} else {
		w := []string{}
		for _, cond := range ø.Conditions {
			w = append(w, "("+string(cond.Sql())+")")
		}
		s = Sql(strings.Join(w, " "+ø.Sign+" "))
	}
	return
}

func Where(sql Sqler) *WhereStruct {
	return &WhereStruct{sql}
}

type WhereStruct struct {
	Inner Sqler
}

func (ø WhereStruct) Sql() (s SqlType) {
	if ø.Inner == nil {
		s = Sql("")
	} else {
		s = Sql("\nWHERE \n\t" + ø.Inner.Sql().String())
	}
	return
}

func LeftJoin(from *Field, to *Field, as string) *JoinStruct {
	return &JoinStruct{to.Table, as, LeftJoinType, Equals(from, &FieldInJoin{to, as})}
}

func RightJoin(from *Field, to *Field, as string) *JoinStruct {
	return &JoinStruct{to.Table, as, RightJoinType, Equals(from, &FieldInJoin{to, as})}
}

func Join(from *Field, to *Field, as string) *JoinStruct {
	return &JoinStruct{to.Table, as, InnerJoinType, Equals(from, &FieldInJoin{to, as})}
}

func FullJoin(from *Field, to *Field, as string) *JoinStruct {
	return &JoinStruct{to.Table, as, FullJoinType, Equals(from, &FieldInJoin{to, as})}
}

var JoinSql = map[JoinType]string{
	InnerJoinType: "JOIN",
	LeftJoinType:  "LEFT JOIN",
	RightJoinType: "RIGHT JOIN",
	FullJoinType:  "FULL OUTER JOIN",
}

type FieldInJoin struct {
	*Field
	As string
}

func (ø *FieldInJoin) Sql() SqlType {
	return Sql(fmt.Sprintf(`"%s"."%s"`, ø.As, ø.Name))
}

type JoinStruct struct {
	Table *Table
	As    string
	Type  JoinType
	On    *Comparer
}

func (js *JoinStruct) Field(f *Field) *AsStruct {
	return JoinField(js, f)
}

func JoinField(js *JoinStruct, f *Field) *AsStruct {
	fij := &FieldInJoin{f, js.As}
	as := fmt.Sprintf(`%s.%s`, fij.As, fij.Name)
	return &AsStruct{Sqler: fij, As: as, Type: f.Type}
}

func (ø *JoinStruct) Sql() SqlType {
	return Sql(fmt.Sprintf("%s %s \"%s\" ON (%s)", JoinSql[ø.Type], ø.Table.Sql(), ø.As, ø.On.Sql()))
}

func As(sq Sqler, as string, typ Type) *AsStruct {
	return &AsStruct{Sqler: sq, As: as, Type: typ}
}

type AsStruct struct {
	Sqler      Sqler
	As         string
	Type       Type
	queryField string // name of a field in a query struct that should match this field
}

// sets the queryField, allows chaining
func (ø *AsStruct) SetQueryField(f string) *AsStruct {
	if ø.queryField != "" {
		panic("queryField already set to " + ø.queryField + ", can't change")
	}
	ø.queryField = f
	return ø
}

func (ø *AsStruct) QueryField() string {
	return ø.queryField
}

func (ø *AsStruct) Sql() string {
	return fmt.Sprintf("%s as \"%s\"", ø.Sqler.Sql(), ø.As)
}

type InsertQuery struct {
	Table *Table
	Sets  []map[*Field]interface{}
}

func InsertMap(table *Table, m map[*Field]interface{}) Query {
	return &InsertQuery{
		Table: table,
		Sets:  []map[*Field]interface{}{m},
	}
}

func Insert(table *Table, first_row SetArray, rows ...SetArray) Query {
	i := &InsertQuery{
		Table: table,
		Sets:  []map[*Field]interface{}{(&first_row).Map()},
	}
	for _, r := range rows {
		i.Sets = append(i.Sets, (&r).Map())
	}
	return i
}

func (ø *InsertQuery) fieldsAndValues() (fields string, values string, err error) {
	fi := []string{}
	va := []string{}
	fieldorder := []*Field{}
	for k, _ := range ø.Sets[0] {
		fieldorder = append(fieldorder, k)
		fi = append(fi, string(k.Sql()))
	}

	for _, r := range ø.Sets {
		ro := []string{}
		for _, k := range fieldorder {
			v := r[k]
			if k.Is(NullAllowed) && v == nil {
				ro = append(ro, "null")
				continue
			}
			tv := TypedValue{PgType: k.Type}
			e := Convert(v, &tv)
			if e != nil {
				err = convertError(k, v, e)
				return
			}
			sql := tv.Sql()
			ro = append(ro, string(sql))
		}
		rs := strings.Join(ro, ", ")
		va = append(va, "("+rs+")")
	}
	fields = strings.Join(fi, ",")
	values = strings.Join(va, ",\n\t")
	return
}

func (ø *InsertQuery) fieldsAndValuesInsert() (fields string, values string, err error) {
	fi := []string{}
	va := []string{}
	fieldorder := []*Field{}
	for k, _ := range ø.Sets[0] {
		fieldorder = append(fieldorder, k)
		fi = append(fi, `"`+k.Name+`"`)
	}

	for _, r := range ø.Sets {
		ro := []string{}
		for _, k := range fieldorder {
			v := r[k]
			if v == nil {
				if k.Is(NullAllowed) {
					ro = append(ro, "null")
					continue
				} else {
					err = nullNotAllowedError(k, v)
					return
				}

			}
			tv, ok := v.(*TypedValue)
			if !ok {
				tv = &TypedValue{PgType: k.Type}
				e := Convert(v, tv)
				if e != nil {
					err = convertError(k, v, e)
					return
				}
			}

			sql := tv.Sql()
			ro = append(ro, string(sql))
		}
		rs := strings.Join(ro, ", ")
		va = append(va, "("+rs+")")
	}
	fields = strings.Join(fi, ",")
	values = strings.Join(va, ",\n\t")
	return
}

func (ø *InsertQuery) Sql() (s SqlType) {
	var err error
	s, err = ø.SqlError()

	if err != nil {
		panic(err.Error())
	}
	return s
}

func (ø *InsertQuery) SqlError() (s SqlType, e error) {
	t := ø.Table
	//currval := Sql(fmt.Sprintf("SELECT\n\tcurrval('%s')", t.PrimaryKeySeq))
	//SELECT currval('\"#{@table.schema.name}\".\"#{@table.primary_key_seq}\"') as id;"
	fi, va, err := ø.fieldsAndValuesInsert()
	if err != nil {
		//panic(err)
		e = err
		return
	}
	//s = Sql(fmt.Sprintf("INSERT INTO \n\t%s (%s) \nVALUES \n\t%s;\n%s", t.Sql(), fi, va, (&AsStruct{currval, "id"}).Sql()))
	stmt := fmt.Sprintf("INSERT INTO \n\t%s (%s) \nVALUES \n\t%s", t.Sql(), fi, va)
	if len(t.PrimaryKey) == 1 {
		stmt += " RETURNING " + t.PrimaryKey[0].Name
	}
	stmt += ";"
	s = Sql(stmt)
	return
}

func (ø *InsertQuery) String() string {
	return ø.Sql().String()
}

type UpdateQuery struct {
	Table  *Table
	Where  *WhereStruct
	Limit  Limit
	Set    map[*Field]interface{}
	SetSql []Sqler
}

func Update(table *Table, options ...interface{}) Query {
	u := &UpdateQuery{
		Limit: Limit(0),
		Table: table,
		Where: &WhereStruct{},
	}
	for _, option := range options {
		switch v := option.(type) {
		case *WhereStruct:
			u.Where = v
		case WhereStruct:
			u.Where = &v
		case Limit:
			u.Limit = v
		case SetArray:
			u.Set = (&v).Map()
		case *SetArray:
			u.Set = v.Map()
		case map[*Field]interface{}:
			u.Set = v
		default:
			sqler := option.(Sqler)
			u.SetSql = append(u.SetSql, sqler)
		}
	}
	return u
}

func (ø *UpdateQuery) setString() (set string, err error) {
	sets := []string{}
	for k, v := range ø.Set {
		// fmt.Printf("in update %s: %v (%T)\n", k.Name, v, v)
		var valstr SqlType
		typedv, ok := v.(*TypedValue)
		if k.Is(NullAllowed) && (v == nil || ok && typedv.Value == nil) {
			valstr = Sql("Null")
		} else {
			tv := typedv
			if !ok {
				tv = &TypedValue{PgType: k.Type}
				e := Convert(v, tv)
				if e != nil {
					err = convertError(k, v, e)
					return
				}
			}

			valstr = tv.Sql()
		}
		sets = append(sets, fmt.Sprintf(`"%s" = %s`, k.Name, valstr))
	}

	for _, sql := range ø.SetSql {
		sets = append(sets, sql.Sql().String())
	}
	set = strings.Join(sets, ",\n\t")
	return
}

func (ø *UpdateQuery) Sql() (s SqlType) {
	var err error
	s, err = ø.SqlError()

	if err != nil {
		panic(err)
	}
	return
}

func (ø *UpdateQuery) SqlError() (s SqlType, e error) {
	t := ø.Table
	sets, err := ø.setString()
	if err != nil {
		e = err
		return
		// panic(err)
	}
	s = Sql(
		fmt.Sprintf(
			"UPDATE \n\t%s \nSET \n\t%s %s %s",
			t.Sql(),
			sets,
			ø.Where.Sql(),
			ø.Limit.Sql()))
	return
}

func (ø *UpdateQuery) String() string {
	return ø.Sql().String()
}

type DeleteQuery struct {
	Table *Table
	Where *WhereStruct
	Limit Limit
}

func Delete(options ...interface{}) Query {
	d := &DeleteQuery{
		Where: &WhereStruct{},
		Limit: Limit(0)}
	for _, option := range options {
		switch v := option.(type) {
		case *WhereStruct:
			d.Where = v
		case WhereStruct:
			d.Where = &v
		case Limit:
			d.Limit = v
		case *Table:
			d.Table = v
		}
	}
	return d
}

func (ø *DeleteQuery) Sql() (s SqlType) {
	s = Sql(
		fmt.Sprintf(
			"DELETE \n\t \nFROM \n\t%s %s %s",
			ø.Table.Sql(),
			ø.Where.Sql(),
			ø.Limit.Sql()))
	return
}

func (ø *DeleteQuery) String() string {
	return ø.Sql().String()
}

type Direction bool

var ASC = Direction(true)
var DESC = Direction(false)

func (ø Direction) Sql() (s SqlType) {
	if ø {
		s = Sql("ASC")
	} else {
		s = Sql("DESC")
	}
	return
}

func OrderBy(os ...interface{}) (o []*OrderByStruct) {
	o = []*OrderByStruct{}
	for i := 0; i < len(os); i = i + 2 {
		f := os[i].(Sqler)
		d := os[i+1].(Direction)
		o = append(o, &OrderByStruct{f, d})
	}
	return
}

type OrderByStruct struct {
	Sqler
	Direction Direction
}

func (ø *OrderByStruct) Sql() SqlType {
	return Sql(ø.Sqler.Sql().String() + " " + ø.Direction.Sql().String())
}

func GroupBy(f ...*Field) GroupByArray {
	return GroupByArray(f)
}

type GroupByArray []*Field

func (ø GroupByArray) Sql() SqlType {
	g := []string{}
	for _, f := range ø {
		g = append(g, string(f.Sql()))
	}
	return Sql("\nGROUP BY " + strings.Join(g, ","))
}

type Distinct bool

func (ø Distinct) Sql() (s SqlType) {
	if ø {
		s = Sql("DISTINCT")
	} else {
		s = Sql("")
	}
	return
}

type SelectQuery struct {
	Distinct        Distinct
	Table           Sqler
	Where           *WhereStruct
	Limit           Limit
	Joins           []*JoinStruct
	Fields          []*Field
	FieldsWithAlias []*AsStruct
	Offset          Offset
	OrderBy         []*OrderByStruct
	GroupBy         GroupByArray
}

func Select(options ...interface{}) Query {
	s := &SelectQuery{
		Distinct:        Distinct(false),
		Joins:           []*JoinStruct{},
		Fields:          []*Field{},
		FieldsWithAlias: []*AsStruct{},
		Limit:           Limit(0),
		Where:           &WhereStruct{},
		OrderBy:         []*OrderByStruct{},
	}
	for _, option := range options {
		switch v := option.(type) {
		case *WhereStruct:
			s.Where = v
		case WhereStruct:
			s.Where = &v
		case Limit:
			s.Limit = v
		case *JoinStruct:
			s.Joins = append(s.Joins, v)
		case *Field:
			s.Fields = append(s.Fields, v)
		case []*Field:
			for _, fld := range v {
				s.Fields = append(s.Fields, fld)
			}
			//s.Fields = append(s.Fields, v)
		case *AsStruct:
			s.FieldsWithAlias = append(s.FieldsWithAlias, v)
		case JoinStruct:
			s.Joins = append(s.Joins, &v)
		case Field:
			s.Fields = append(s.Fields, &v)
		case AsStruct:
			s.FieldsWithAlias = append(s.FieldsWithAlias, &v)
		case Offset:
			s.Offset = v
		case *OrderByStruct:
			s.OrderBy = append(s.OrderBy, v)
		case []*OrderByStruct:
			s.OrderBy = v
		case OrderByStruct:
			s.OrderBy = append(s.OrderBy, &v)
		case GroupByArray:
			s.GroupBy = v
		case Distinct:
			s.Distinct = v
		default:
			if sqler, ok := v.(Sqler); ok {
				s.Table = sqler
			} else {
				panic("unknown select option " + fmt.Sprintf("%#v", v))
			}
		}
	}
	if s.Table == nil {
		panic("no table to select from")
	}
	return s
}

func (ø *SelectQuery) fieldstr() (s string) {
	f := []string{}

	for _, field := range ø.Fields {
		f = append(f, string(field.Sql()))
	}

	for _, alias := range ø.FieldsWithAlias {
		f = append(f, alias.Sql())
	}

	s = strings.Join(f, ", \n\t")
	return
}

func (ø *SelectQuery) joins() (s SqlType) {
	if len(ø.Joins) == 0 {
		return Sql("")
	}
	str := []string{}
	for _, j := range ø.Joins {
		str = append(str, string(j.Sql().String()))
	}
	return Sql(strings.Join(str, "\n"))
}

func (ø *SelectQuery) group_by() (s SqlType) {
	if len(ø.GroupBy) == 0 {
		return Sql("")
	}
	str := []string{}
	for _, g := range ø.GroupBy {
		str = append(str, string(g.Sql().String()))
	}
	return Sql("\nGROUP BY " + strings.Join(str, ", "))
}

func (ø *SelectQuery) order_by() (s SqlType) {
	if len(ø.OrderBy) == 0 {
		return Sql("")
	}
	str := []string{}
	for _, o := range ø.OrderBy {
		str = append(str, o.Sql().String())
	}
	return Sql("\nORDER BY " + strings.Join(str, ", "))
}

type caseStruct struct {
	whens map[Sqler]Sqler
	Else  Sqler
}

func Case(sqls ...Sqler) (ø *caseStruct) {
	ø = &caseStruct{
		whens: map[Sqler]Sqler{},
	}
	for i := 0; i < len(sqls); i = i + 2 {
		ø.whens[sqls[i]] = sqls[i+1]
	}
	return
}

func (ø *caseStruct) Sql() SqlType {
	s := `
CASE
  %s
  %s
END
`
	whens := []string{}
	for k, v := range ø.whens {
		whens = append(whens, fmt.Sprintf(`WHEN %s THEN %s`, k.Sql().String(), v.Sql().String()))
	}
	el := ""
	if ø.Else != nil {
		el = " ELSE " + ø.Else.Sql().String()
	}

	return Sql(fmt.Sprintf(s, strings.Join(whens, "\n  "), el))
}

type callStruct struct {
	Name   string
	Params []Sqler
}

func Call(name string, params ...Sqler) (out *callStruct) {
	return &callStruct{name, params}
}

func (ø *callStruct) Sql() SqlType {
	params := []string{}
	for _, sq := range ø.Params {
		params = append(params, sq.Sql().String())
	}
	return Sql(fmt.Sprintf("%s(%s)", ø.Name, strings.Join(params, ", ")))
}

/*
	SELECT #{distinct}
		#{fields}
	FROM
		#{table}
	#{joins}
	#{where}
	#{group_by}
	#{order_by}
	#{limit}
	#{offset}"
*/
func (ø *SelectQuery) Sql() (s SqlType) {
	s = Sql(
		fmt.Sprintf(
			"SELECT %s \n\t%s \nFROM \n\t%s %s %s %s%s%s%s",
			ø.Distinct.Sql(),
			ø.fieldstr(),
			ø.Table.Sql(),
			ø.joins(),
			ø.Where.Sql(),
			ø.group_by(),
			ø.order_by(),
			ø.Limit.Sql(),
			ø.Offset.Sql()))
	return
}

func (ø *SelectQuery) String() string {
	return ø.Sql().String()
}

func Set(i ...interface{}) (out SetArray) {
	out = SetArray{}
	for _, o := range i {
		out = append(out, o)
	}
	return
}

type SetArray []interface{}

func (ø SetArray) Map() (m map[*Field]interface{}) {
	m = map[*Field]interface{}{}
	for i := 0; i < len(ø); i = i + 2 {
		m[ø[i].(*Field)] = ø[i+1]
	}
	return
}
