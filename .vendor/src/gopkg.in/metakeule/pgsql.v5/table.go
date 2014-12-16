package pgsql

import (
	"fmt"
	//	. "gopkg.in/metakeule/nil.v1"
	"strings"
)

type foreignKey struct {
	Field           *Field
	Reference       *Field
	OnDeleteCascade bool
}

func (ø *foreignKey) Sql() SqlType {
	casc := ""
	if ø.OnDeleteCascade {
		casc = `ON UPDATE CASCADE ON DELETE CASCADE`
	}
	s := `CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s) MATCH SIMPLE %s`
	return Sql(fmt.Sprintf(s,
		`"`+ø.Field.Table.Name+"_fk_"+ø.Field.Name+`"`,
		`"`+ø.Field.Name+`"`,
		ø.Reference.Table.Sql(),
		`"`+ø.Reference.Name+`"`,
		casc,
	))
}

/*
	CREATE [ UNIQUE ] INDEX [ CONCURRENTLY ] [ name ] ON table [ USING method ]
    ( { column | ( expression ) } [ COLLATE collation ] [ opclass ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ] [, ...] )
    [ WITH ( storage_parameter = value [, ... ] ) ]
    [ TABLESPACE tablespace ]
    [ WHERE predicate ]
*/
type Index struct {
	Name         string
	Unique       bool
	Concurrently bool
	Fields       []*Field
	Table        *Table
	Method       string
}

func (ø *Index) Sql() SqlType {
	return Sql(ø.Name)
}

func (ø *Index) Create() SqlType {
	s := `CREATE %sINDEX %s%s ON %s %s(%s)`
	unique, concurr, using := "", "", ""
	if ø.Unique {
		unique = " UNIQUE "
	}
	if ø.Concurrently {
		concurr = " CONCURRENTLY "
	}
	if ø.Method != "" {
		using = " USING " + ø.Method
	}

	fields := []string{}
	for _, f := range ø.Fields {
		fields = append(fields, `"`+f.Name+`"`)
	}

	return Sql(fmt.Sprintf(s, unique, concurr, `"`+ø.Name+`"`, ø.Table.Sql(), using,
		strings.Join(fields, ",")))
}

type unique struct {
	Fields []*Field
	Name   string
}

func (ø *unique) Sql() SqlType {
	s := `CONSTRAINT %s UNIQUE (%s)`
	fields := []string{}
	for _, f := range ø.Fields {
		fields = append(fields, `"`+f.Name+`"`)
	}
	return Sql(fmt.Sprintf(s, `"`+ø.Name+`"`, strings.Join(fields, ",")))
}

type Table struct {
	Name   string
	Schema *Schema
	Fields []*Field
	//	PrimaryKeySeq Sqler
	PrimaryKey  []*Field
	Validations []RowValidator
	Constraints []Sqler
}

func NewTable(name string, options ...interface{}) *Table {
	t := &Table{
		Name:   name,
		Fields: []*Field{},
	}
	for _, option := range options {
		switch v := option.(type) {
		case *Schema:
			t.Schema = v
		case *Field:
			t.AddField(v)
		default:
			if val, ok := v.(RowValidator); ok {
				t.Validations = append(t.Validations, val)
			}
		}
	}
	return t
}

func (ø *Table) NewField(name string, options ...interface{}) (field *Field) {
	field = NewField(name, options...)
	ø.AddField(field)
	return
}

func (ø *Table) AddValidator(v ...RowValidator) {
	for _, val := range v {
		ø.Validations = append(ø.Validations, val)
	}
}

func (ø *Table) IsPrimaryKey(f *Field) (is bool) {
	for _, pk := range ø.PrimaryKey {
		if pk == f {
			return true
		}
	}
	return false
}

// TODO: handle somehow validations for placeholders
func (ø *Table) Validate(values map[*Field]interface{}) (errs map[Sqler]error) {
	errs = map[Sqler]error{}
	//pkey := ø.PrimaryKey
	//for _, f := range ø.Fields {
	for f := range values {
		// don't validate empty ids
		if ø.IsPrimaryKey(f) && values[f] == nil {
			continue
		}

		err := f.Validate(values[f])
		//fmt.Printf("VALIDATING %s: %s\n", f.Name, err)
		if err != nil {
			errs[f] = err
		}
	}
	for _, val := range ø.Validations {
		//		fmt.Println("custom validation")
		err := val.ValidateRow(values)
		if err != nil {
			errs[ø] = err
		}
	}
	/*
		if len(errs) > 0 {
			errs[Sql("backtrace")] = fmt.Errorf(strings.Join(backtrace(), "\n"))
		}
	*/
	return
}

func (ø *Table) createField(field *Field) string {
	t := field.Type.String()
	if field.ForeignKey != nil {
		t = field.ForeignKey.Type.String()
	}
	s := field.Name
	//if field.Is(PrimaryKey) {
	if field.Is(Serial) {
		s += " SERIAL"
	} else {
		s += " " + t
	}

	if field.Default != nil {
		s += " DEFAULT " + string(field.Default.Sql())
	}
	if !field.Is(NullAllowed) {
		s += " NOT NULL "
	}

	if field.Is(UuidGenerate) {
		s += " DEFAULT uuid_generate_v4()"
	}
	//	s += " PRIMARY KEY"
	//}

	/*
		if field.ForeignKey != nil {
			s += " REFERENCES " + string(field.ForeignKey.Table.Sql()) + `("` + field.ForeignKey.Name + `")`
			if field.Is(OnDeleteCascade) {
				s += " ON DELETE CASCADE"
			} else {
				s += " ON DELETE RESTRICT"
			}
		}
	*/
	return s
}

func (ø *Table) AddUnique(name string, fields ...*Field) {
	if len(fields) == 0 {
		panic("need a field for unique contraint")
	}
	uniq := &unique{}
	uniq.Name = name
	uniq.Fields = fields
	ø.Constraints = append(ø.Constraints, uniq)
}

func (ø *Table) Create() SqlType {
	fs := []string{}
	pkeys := []string{}

	for _, f := range ø.Fields {
		fs = append(fs, ø.createField(f))
		if f.Is(PrimaryKey) {
			pkeys = append(pkeys, f.Name)
		}
	}

	if len(pkeys) > 0 {
		fs = append(fs, fmt.Sprintf("PRIMARY KEY(%s)", strings.Join(pkeys, ",")))
	}

	for _, constr := range ø.Constraints {
		fs = append(fs, constr.Sql().String())
	}

	str := fmt.Sprintf(
		"CREATE TABLE %s \n(%s) ", ø.Sql(), strings.Join(fs, ",\n"))
	return Sql(str)
}

//func (ø *Table) Alter() Sql {
//}

func (ø *Table) Drop() SqlType {
	return Sql("DROP TABLE " + string(ø.Sql()))
}

func (ø *Table) DropCascade() SqlType {
	return Sql("DROP TABLE " + string(ø.Sql()) + " CASCADE")
}

func (ø *Table) AddForeignKeyConstraint(field *Field, onDeleteCascade bool) {
	fk := &foreignKey{}
	fk.Field = field
	fk.Reference = field.ForeignKey
	if onDeleteCascade {
		fk.OnDeleteCascade = true
	}
	ø.Constraints = append(ø.Constraints, fk)
}

func (ø *Table) AddField(fields ...*Field) {
	for _, f := range fields {
		ø.Fields = append(ø.Fields, f)
		if f.Is(PrimaryKey) {
			ø.PrimaryKey = append(ø.PrimaryKey, f)
		}
		if f.ForeignKey != nil {
			fk := &foreignKey{}
			fk.Field = f
			fk.Reference = f.ForeignKey
			if f.Is(OnDeleteCascade) {
				fk.OnDeleteCascade = true
			}
			ø.Constraints = append(ø.Constraints, fk)
		}
		f.Table = ø
	}
}

func (ø *Table) Field(name string) (f *Field) {
	for _, ff := range ø.Fields {
		if ff.Name == name {
			return ff
		}
	}
	return
}

func (ø *Table) QueryField(name string) (f *Field) {
	for _, ff := range ø.Fields {
		if ff.QueryField() == name {
			return ff
		}
	}
	return
}

func (ø *Table) Sql() SqlType {
	if ø.Schema == nil {
		return Sql(fmt.Sprintf(`"%s"`, ø.Name))
	}
	return Sql(fmt.Sprintf(`%s."%s"`, ø.Schema.Sql(), ø.Name))
}
