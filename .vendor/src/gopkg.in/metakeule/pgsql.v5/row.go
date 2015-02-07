package pgsql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gopkg.in/go-on/lib.v3/internal/meta"

	// "gopkg.in/metakeule/meta.v5"
	"gopkg.in/metakeule/typeconverter.v2"
)

type PreValidate func(*Row) error
type PostValidate func(*Row) error
type PreGet func(*Row) error
type PostGet func(*Row) error
type PreInsert func(*Row) error
type PostInsert func(*Row) error
type PreUpdate func(*Row) error
type PostUpdate func(*Row) error
type PreDelete func(*Row) error
type PostDelete func(*Row) error

type Row struct {
	*Table
	DB           DB
	Tx           *sql.Tx
	values       map[*Field]*TypedValue
	aliasValues  map[*AsStruct]*TypedValue
	setErrors    []error
	Debug        bool
	LastSql      string
	PreValidate  []PreValidate
	PostValidate []PostValidate
	PreGet       []PreGet
	PostGet      []PostGet
	PreInsert    []PreInsert
	PostInsert   []PostInsert
	PreUpdate    []PreUpdate
	PostUpdate   []PostUpdate
	PreDelete    []PreDelete
	PostDelete   []PostDelete
}

func NewRow(db DB, table *Table, hooks ...interface{}) (ø *Row) {
	ø = &Row{
		Table:        table,
		setErrors:    []error{},
		PreValidate:  []PreValidate{},
		PostValidate: []PostValidate{},
		PreGet:       []PreGet{},
		PostGet:      []PostGet{},
		PreInsert:    []PreInsert{},
		PostInsert:   []PostInsert{},
		PreUpdate:    []PreUpdate{},
		PostUpdate:   []PostUpdate{},
		PreDelete:    []PreDelete{},
		PostDelete:   []PostDelete{},
	}

	tx, ok := db.(*sql.Tx)
	if ok {
		ø.SetTransaction(tx)
	} else {
		ø.SetDB(db.(DB))
	}

	for _, h := range hooks {
		switch hook := h.(type) {
		case PreValidate:
			ø.PreValidate = append(ø.PreValidate, hook)
		case PostValidate:
			ø.PostValidate = append(ø.PostValidate, hook)
		case PreGet:
			ø.PreGet = append(ø.PreGet, hook)
		case PostGet:
			ø.PostGet = append(ø.PostGet, hook)
		case PreInsert:
			ø.PreInsert = append(ø.PreInsert, hook)
		case PostInsert:
			ø.PostInsert = append(ø.PostInsert, hook)
		case PreUpdate:
			ø.PreUpdate = append(ø.PreUpdate, hook)
		case PostUpdate:
			ø.PostUpdate = append(ø.PostUpdate, hook)
		case PreDelete:
			ø.PreDelete = append(ø.PreDelete, hook)
		case PostDelete:
			ø.PostDelete = append(ø.PostDelete, hook)
		default:
			panic("unknown hook type: " + fmt.Sprintf("%#v", hook))
		}
	}
	ø.clearValues()
	return
}

// the parameters should be given in pairs of
// *Field and *value
func (ø *Row) Get(o ...interface{}) {
	for i := 0; i < len(o); i = i + 2 {
		switch field := o[i].(type) {
		case *Field:
			//field := o[i].(*Field)
			res := o[i+1]
			if ø.values[field] == nil {
				continue
			}
			err := Convert(ø.values[field], res)
			if err != nil {
				panic(
					"can't convert " +
						field.Name +
						" of table " +
						field.Table.Name +
						" to " +
						fmt.Sprintf("%#v", res) +
						": " +
						err.Error())
			}
		case *AsStruct:
			//as := o[i].(*AsStruct)
			res := o[i+1]
			if ø.aliasValues[field] == nil {
				continue
			}
			err := Convert(ø.aliasValues[field], res)
			if err != nil {
				panic(
					"can't convert " +
						field.As +
						" of alias field " +
						field.Sql() +
						" to " +
						fmt.Sprintf("%#v", res) +
						": " +
						err.Error())
			}
		default:
			panic(fmt.Sprintf("unsupported type %#v\n", o[i]))
		}
	}
}

func setFieldInStruct(vl reflect.Value, fieldName string, tagVal string, v *TypedValue, s interface{}) error {

	str, err := meta.StructByValue(meta.FinalValue(s))

	if err != nil {
		return err
	}

	tag, err2 := str.Tag(fieldName)
	if err2 != nil {
		return err2
	}

	// tag does match the given
	if tag != nil && strings.Contains(tag.Get("db.select"), tagVal) {
		err := Convert(v, vl.Addr().Interface())
		if err != nil {
			return fmt.Errorf("error in field %s: %s", fieldName, err.Error())
		}
	}
	return nil
}

//ro.Get(artist.Id, &a.Id, artist.FirstName, &a.FirstName, artist.LastName, &a.LastName, artist.GalleryArtist, &a.GalleryArtist, artist.Vita, &a.Vita, artist.Area, &ar)
func (ø *Row) GetStruct(tagVal string, s interface{}) (err error) {
	fv := meta.FinalValue(s)

	for f, v := range ø.values {
		if f.queryField == "" {
			panic("queryField not set for " + f.Name)
		}
		// a field with this name does exist
		if vl := fv.FieldByName(f.queryField); vl.IsValid() {
			err = setFieldInStruct(vl, f.queryField, tagVal, v, s)
			if err != nil {
				return
			}
			/*
				tag := meta.Struct.Tag(s, f.queryField)
				// tag does match the given
				if tag != nil && strings.Contains(tag.Get("db.select"), tagVal) {
					err := Convert(v, vl.Addr().Interface())
					if err != nil {
						return fmt.Errorf("error in field %s: %s", f.queryField, err.Error())
					}
				}
			*/
		}
	}

	for f, v := range ø.aliasValues {
		if vl := fv.FieldByName(f.queryField); vl.IsValid() {
			err = setFieldInStruct(vl, f.queryField, tagVal, v, s)
			if err != nil {
				return
			}
			continue
		}
		if vl := fv.FieldByName(f.As); vl.IsValid() {
			err = setFieldInStruct(vl, f.As, tagVal, v, s)
			if err != nil {
				return
			}
			continue
		}
		/*
			if vl := fv.FieldByName(f.As); vl.IsValid() {
				setFieldInStruct(vl, f.As, tagVal, v, s)
					tag := meta.Struct.Tag(s, f.As)
					// tag does match the given
					if tag != nil && strings.Contains(tag.Get("db.select"), tagVal) {
						err := Convert(v, vl.Addr().Interface())
						if err != nil {
							return fmt.Errorf("error in field %s: %s", f.As, err.Error())
						}
					}
			}

		*/
		/*
			else if {
				tag := meta.Struct.Tag(s, f.As)
				// tag does match the given
				if tag != nil && strings.Contains(tag.Get("db.select"), tagVal) {
					err := Convert(v, vl.Addr().Interface())
					if err != nil {
						return fmt.Errorf("error in field %s: %s", f.As, err.Error())
					}
				}
			}
		*/
	}
	return nil
}

func (ø *Row) GetString(field interface{}) (s string) {
	ø.Get(field, &s)
	return
}

func (ø *Row) ValidateAll() (errs map[Sqler]error) {
	vals := []string{}
	vs := ø.Values()

	for vf := range vs {
		vals = append(vals, vf.Name)
	}

	// fmt.Printf("values: %#v\n", vals)

	return ø.Table.Validate(ø.Values())
}

// the parameters should be given in pairs of
// *Field and value
func (ø *Row) Set(o ...interface{}) (err error) {
	err = ø.set(o...)
	if err != nil {
		return err
	}
	err = ø.validate()
	if err != nil {
		return err
	}
	return
}

// set field to null
func (ø *Row) SetNull(field *Field) {
	ø.values[field].Value = nil
}

// unset the given fields
func (ø *Row) Unset(o ...*Field) {
	for _, f := range o {
		delete(ø.values, f)
	}
}

func (ø *Row) set(o ...interface{}) (err error) {
	for i := 0; i < len(o); i = i + 2 {
		field := o[i].(*Field)
		var tv *TypedValue
		vl := o[i+1]
		placeh, ok := vl.(Placeholder)
		if ok {
			//fmt.Println("is placeholder: " + placeh.String())
			//tv, err = field.Value(placeh.String())
			//fmt.Printf("pgtype %s value %#v\n", tv.PgType.String(), tv.Value)
			//tv.PgType = VarChar(255)
			//tv = &TypedValue{TextType, &pgInterpretedString{StringType: typeconverter.StringType(placeh.String())}, true}
			tv = &TypedValue{TextType, &pgInterpretedString{StringType: typeconverter.StringType("@@" + placeh.Name() + "@@")}, true}
			//fmt.Printf("pgtype %s value %#v\n", tv.PgType.String(), tv.Value)
			//tv.PgType = VarChar(125)
			//tv.dontChange = true
		} else {
			tv, err = field.Value(vl)
		}
		/*
			if o[i+1] == nil {
				if field.Is(NullAllowed) {
					ø.values[field] = &TypedValue{PgType: field.Type}
					continue
				} else {
					e := fmt.Errorf("error when setting field %s to value %#v: Null is not allowed for this field\n", field.Sql(), o[i+1])
					ø.setErrors = append(ø.setErrors, e)
					return e
				}

			}
			tv := &TypedValue{PgType: field.Type}
			err = Convert(o[i+1], tv)
		*/
		if err != nil {
			ø.setErrors = append(ø.setErrors, err)
			return
		}
		ø.values[field] = tv
	}
	return
}

var _ = strings.Replace

func (ø *Row) validate() (err error) {
	if len(ø.setErrors) > 0 {
		errString := []string{}
		for _, e := range ø.setErrors {
			errString = append(errString, fmt.Sprintf("\t%s", e.Error()))
		}
		return fmt.Errorf("%s\n", strings.Join(errString, "\n"))
	}

	errs := ø.ValidateAll()
	if len(errs) > 0 {
		errString := []string{}
		for k, e := range errs {
			errString = append(errString, fmt.Sprintf("\tValidation Error in %s: %s\n", k.Sql(), e.Error()))
		}
		return &ValidationError{Err: fmt.Errorf("%s\n", strings.Join(errString, "\n"))}
	}
	return
}

func (ø *Row) Validate() (err error) {
	for _, hook := range ø.PreValidate {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	err = ø.validate()
	if err != nil {
		return
	}
	for _, hook := range ø.PostValidate {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	return nil
}

func (ø *Row) Save() (err error) {
	if len(ø.PrimaryKey) != 1 {
		panic("save should only be called for single primary keys, try update or insert directly")
	}
	err = ø.Validate()
	if err != nil {
		return fmt.Errorf("Can't save row of %s:\n%s\n", ø.Table.Sql(), err.Error())
	}
	ø.setErrors = []error{}
	if ø.HasId() {
		err = ø.Update()
	} else {
		err = ø.Insert()
	}
	return
}

func (ø *Row) HasId() bool {
	for _, pkey := range ø.PrimaryKey {
		if ø.values[pkey].IsNil() {
			return false
		}
	}
	return true
}

func (ø *Row) Fill(m map[string]interface{}) error {
	ø.values = map[*Field]*TypedValue{}
	for k, v := range m {
		e := ø.set(ø.Table.Field(k), v)

		if e != nil {
			fmt.Printf("error while filling %s with %v: %s\n", k, v, e.Error())
			return e
		}
	}
	err := ø.validate()
	if err != nil {
		return err
	}
	return nil
}

func (ø *Row) Properties() (m map[string]interface{}) {
	m = map[string]interface{}{}
	for field, val := range ø.values {
		m[field.Name] = val.Value
	}
	for as, val := range ø.aliasValues {
		m[as.As] = val.Value
	}
	return
}

func (ø *Row) SetDebug() *Row {
	ø.Debug = true
	return ø
}

func (ø *Row) UnsetDebug() *Row {
	ø.Debug = false
	return ø
}

func (ø *Row) AsStrings() (m map[string]string) {
	m = map[string]string{}
	// fmt.Printf("values: %#v\n", ø.values)
	for field, val := range ø.values {
		// fmt.Printf("key: %#v val: %#v\n", field.Sql(), val)
		var s string
		err := Convert(val, &s)
		if err != nil {
			panic(convertError(field, val, err).Error())
			// panic("can't convert to string: " + err.Error())
		}
		m[field.Name] = s
	}
	for as, val := range ø.aliasValues {
		var s string
		err := Convert(val, &s)
		if err != nil {
			err = &ValidationError{
				Err:     fmt.Errorf("can't convert %#v to string for alias %s", val, as.As),
				Details: err,
			}
			panic(err.Error())
			// panic("can't convert to string: " + err.Error())
		}
		m[as.As] = s
	}
	return
}

func (ø *Row) Reset() {
	ø.clearValues()
	ø.setErrors = []error{}
}

func (ø *Row) clearValues() {
	ø.values = map[*Field]*TypedValue{}
	if len(ø.PrimaryKey) > 0 {
		for _, pkey := range ø.PrimaryKey {
			ø.values[pkey] = &TypedValue{PgType: pkey.Type}
		}
	}
	ø.aliasValues = map[*AsStruct]*TypedValue{}
}

// vals must be in the order of ø.PrimaryKey
func (ø *Row) SetId(vals ...string) (err error) {
	for i, val := range vals {
		ø.values[ø.PrimaryKey[i]] = &TypedValue{ø.PrimaryKey[i].Type, NewPgInterpretedString(val), false}
		/*
			err = Convert(val, ø.values[ø.PrimaryKey[i]])
			if err != nil {
				return
			}
		*/
	}
	return
}

func (ø *Row) Id() (vals []SqlType) {
	//var idVal SqlType
	vals = []SqlType{}
	for _, pkey := range ø.PrimaryKey {
		var val SqlType
		err := Convert(ø.values[pkey], &val)
		if err != nil {
			panic("can't get id for table " + ø.Table.Name + ": " + err.Error())
		}
		vals = append(vals, val)
	}
	return
}

type Rows struct {
	*sql.Rows // inherits from *sql.Rows and is fully compatible
	row       *Row
	Fields    []interface{}
}

// scan the result into a row
func (ø *Rows) ScanRow() (row *Row, ſ error) {
	ſ = ø.row.Scan(ø.Rows, ø.Fields...)
	if ſ == nil {
		row = ø.row
	}
	return
}

// scan the result into a struct
func (ø *Rows) ScanStruct(tagVal string, s interface{}) (ſ error) {
	ſ = ø.row.Scan(ø.Rows, ø.Fields...)
	if ſ != nil {
		return
	}
	return ø.row.GetStruct(tagVal, s)
}

// call fn on each row
func (ø *Row) Each(fn func(*Row) error, options ...interface{}) (err error) {
	var rows *Rows
	rows, err = ø.Find(options...)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var r *Row
		r, err = rows.ScanRow()
		if err != nil {
			return
		}
		err = fn(r)
		if err != nil {
			return
		}
	}
	return
}

// return the first result
func (ø *Row) Any(options ...interface{}) (r *Row, err error) {
	var rows *Rows
	opts := []interface{}{Limit(1)}
	opts = append(opts, options...)
	rows, err = ø.Find(opts...)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		r, err = rows.ScanRow()
	} else {
		err = fmt.Errorf("no rows found")
	}
	return
}

func (ø *Row) FindWithArgs(args []interface{}, options ...interface{}) (rows *Rows, err error) {
	sel := ø.SelectQuery(options...)
	r, err := ø.Query(sel, args...)

	if err != nil {
		return
	}

	rows = &Rows{
		Rows:   r,
		row:    ø,
		Fields: []interface{}{},
	}

	for _, f := range sel.Fields {
		rows.Fields = append(rows.Fields, f)
	}

	for _, aliasF := range sel.FieldsWithAlias {
		//rows.Fields = append(rows.Fields, NewField(aliasF.As, aliasF.Type))
		rows.Fields = append(rows.Fields, aliasF)
	}

	return
}

func (ø *Row) Find(options ...interface{}) (rows *Rows, err error) {
	sel := ø.SelectQuery(options...)
	r, err := ø.Query(sel)

	if err != nil {
		return
	}

	rows = &Rows{
		Rows:   r,
		row:    ø,
		Fields: []interface{}{},
	}

	for _, f := range sel.Fields {
		rows.Fields = append(rows.Fields, f)
	}

	for _, aliasF := range sel.FieldsWithAlias {
		//rows.Fields = append(rows.Fields, NewField(aliasF.As, aliasF.Type))
		rows.Fields = append(rows.Fields, aliasF)
	}

	return
}

// runs any and returns the result into the struct
func (ø *Row) Result(tagVal string, structPtr interface{}, findOptions ...interface{}) error {
	row, err := ø.Any(findOptions...)
	if err != nil {
		return err
	}
	return row.GetStruct(tagVal, structPtr)
}

/*
type NullBool struct {
   137		Bool  bool
   138		Valid bool // Valid is true if Bool is not NULL
   139	}
   140
   141	// Scan implements the Scanner interface.
   142	func (n *NullBool) Scan(value interface{}) error {
   143		if value == nil {
   144			n.Bool, n.Valid = false, false
   145			return nil
   146		}
   147		n.Valid = true
   148		return convertAssign(&n.Bool, value)
   149	}
   150
   151	// Value implements the driver Valuer interface.
   152	func (n NullBool) Value() (driver.Value, error) {
   153		if !n.Valid {
   154			return nil, nil
   155		}
   156		return n.Bool, nil
   157	}
*/

//func (ø *Row) QueryByStruct()

func (ø *Row) Scan(row *sql.Rows, fields ...interface{}) (err error) {
	//ø.clearValues()
	ø.values = map[*Field]*TypedValue{}
	ø.aliasValues = map[*AsStruct]*TypedValue{}
	for _, hook := range ø.PreGet {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	scanF := []interface{}{}
	for _, field := range fields {

		switch f := field.(type) {
		case *Field:
			// make default values and append them
			switch f.Type {
			case IntType:
				if f.Is(NullAllowed) {
					var inNull sql.NullInt64
					scanF = append(scanF, &inNull)
				} else {
					var in int
					scanF = append(scanF, &in)
				}
			case FloatType:
				if f.Is(NullAllowed) {
					var flNull sql.NullFloat64
					scanF = append(scanF, &flNull)
				} else {
					var fl float32
					scanF = append(scanF, &fl)
				}
			case BoolType:
				if f.Is(NullAllowed) {
					var blNull sql.NullBool
					scanF = append(scanF, &blNull)
				} else {
					var bl bool
					scanF = append(scanF, &bl)
				}

			case TimeStampTZType, TimeStampType, DateType, TimeType:
				//f.Type.String()
				//				fmt.Printf("scanning field %v is timelike, nullallowed: %v\n", f.Type.String(), f.Is(NullAllowed))
				if f.Is(NullAllowed) {
					var tiNull NullTime
					scanF = append(scanF, &tiNull)
					//	var s
				} else {
					var t time.Time
					scanF = append(scanF, &t)
				}
			default:
				if f.Is(NullAllowed) {
					var sNull sql.NullString
					scanF = append(scanF, &sNull)
				} else {
					var s string
					scanF = append(scanF, &s)
				}
			}
		case *AsStruct:
			// make default values and append them
			switch f.Type {
			case IntType:
				var in int
				scanF = append(scanF, &in)
			case FloatType:
				var fl float32
				scanF = append(scanF, &fl)
			case BoolType:
				var bl bool
				scanF = append(scanF, &bl)
			case TimeStampTZType, TimeStampType, DateType, TimeType:
				var t time.Time
				scanF = append(scanF, &t)
			default:
				var s string
				scanF = append(scanF, &s)
			}
		}

	}
	err = row.Scan(scanF...)
	if err != nil {
		//	panic(err.Error())
		//		fmt.Printf("scanF had errors: %#v\n", err)
		return
	}
	for i, res := range scanF {
		fi := fields[i]
		switch f := fi.(type) {
		case *Field:
			tv := &TypedValue{PgType: f.Type}
			switch v := res.(type) {
			case *sql.NullInt64:
				if (*v).Valid {
					e := Convert((*v).Int64, tv)
					if e != nil {
						err = convertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *sql.NullFloat64:
				if (*v).Valid {
					e := Convert((*v).Float64, tv)
					if e != nil {
						err = convertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *sql.NullBool:
				if (*v).Valid {
					e := Convert((*v).Bool, tv)
					if e != nil {
						err = convertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *sql.NullString:
				if (*v).Valid {
					e := Convert((*v).String, tv)
					if e != nil {
						err = convertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *NullTime:
				if (*v).Valid {
					e := Convert((*v).Time, tv)
					if e != nil {
						err = convertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			default:
				e := Convert(v, tv)
				if e != nil {
					err = convertError(f, v, e)
					// err = e
					return
				}
			}
			ø.values[f] = tv
		case *AsStruct:
			tv := &TypedValue{PgType: f.Type}
			switch v := res.(type) {
			case *sql.NullInt64:
				if (*v).Valid {
					e := Convert((*v).Int64, tv)
					if e != nil {
						err = aliasConvertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *sql.NullFloat64:
				if (*v).Valid {
					e := Convert((*v).Float64, tv)
					if e != nil {
						err = aliasConvertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *sql.NullBool:
				if (*v).Valid {
					e := Convert((*v).Bool, tv)
					if e != nil {
						err = aliasConvertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			case *sql.NullString:
				if (*v).Valid {
					e := Convert((*v).String, tv)
					if e != nil {
						err = aliasConvertError(f, v, e)
						// err = e
						return
					}
				} else {
					continue
				}
			default:
				e := Convert(v, tv)
				if e != nil {
					err = aliasConvertError(f, v, e)
					// err = e
					return
				}
			}
			ø.aliasValues[f] = tv
		}
	}

	for _, hook := range ø.PostGet {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	return
}

func (ø *Row) Reload() (err error) {
	if !ø.HasId() {
		return fmt.Errorf("can't reload, have no id")
	}
	var ids []string

	for _, pk := range ø.PrimaryKey {
		var id string
		ø.Get(pk, &id)
		ids = append(ids, id)
	}
	// fmt.Printf("loaded ids to: %#v", ids)
	return ø.Load(ids...)
}

func (ø *Row) Load(ids ...string) (err error) {
	f := ø.Table.Fields
	err = ø.SetId(ids...)
	if err != nil {
		return
	}

	//ø.Debug = true

	var conds []Sqler
	is := ø.Id()

	for i, pk := range ø.PrimaryKey {
		conds = append(conds, Equals(pk, is[i]))
	}

	row, err := ø.Select(f, Where(And(conds...)))
	if err != nil {
		return
	}
	if !row.Next() {
		row.Close()
		return fmt.Errorf("no row for %v", ids)
	}

	fs := []interface{}{}
	for _, ff := range f {
		fs = append(fs, ff)
	}

	err = ø.Scan(row, fs...)
	row.Close()
	return
}

// runs load and puts the result into the struct
func (ø *Row) LoadStruct(tagVal string, structPtr interface{}, ids ...string) error {
	err := ø.Load(ids...)
	if err != nil {
		return err
	}
	return ø.GetStruct(tagVal, structPtr)
}

func (ø *Row) UpdateQuery(pkVals ...interface{}) Query {
	vals := ø.typedValues()
	// fmt.Println(vals)
	var cond []Sqler
	if len(ø.PrimaryKey) == 1 {
		ids := ø.Id()
		delete(vals, ø.PrimaryKey[0])
		cond = append(cond, Equals(ø.PrimaryKey[0], ids[0]))
	} else {
		if len(pkVals) != len(ø.PrimaryKey) {
			n := []string{}
			for _, pk := range ø.PrimaryKey {
				n = append(n, pk.Name)
			}
			panic(fmt.Sprintf("number of vals for multiple primary keys does not match: given: %v, needed: %v (%s)\n", len(pkVals), len(ø.PrimaryKey), strings.Join(n, ", ")))
		}
		for i, pk := range ø.PrimaryKey {
			cond = append(cond, Equals(pk, pkVals[i]))
		}
	}
	/*
		else {
			ids := ø.Id()
			for i, pkey := range ø.PrimaryKey {
				//delete(vals, pkey)
				cond = append(cond, Equals(pkey, ids[i]))
			}
		}
	*/

	//w := []*Condition{}
	//conditions = append(conditions, And(cond...))
	w := And(cond...)

	u := Update(
		ø.Table,
		vals,
		//Where(Equals(ø.PrimaryKey, ø.Id())))
		Where(w))
	return u
}

// has to be invoked directly
func (ø *Row) Update(pkVals ...interface{}) (err error) {
	err = ø.Validate()
	if err != nil {
		//return fmt.Errorf("Can't update row of %s:\n%s\n", ø.Table.Sql(), err.Error())
		return
	}
	ø.setErrors = []error{}

	for _, hook := range ø.PreUpdate {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	u := ø.UpdateQuery(pkVals...)
	// fmt.Println(u.Sql())
	_, err = ø.Exec(u)
	for _, hook := range ø.PostUpdate {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	return
}

func (ø *Row) Values() (vals map[*Field]interface{}) {
	vals = map[*Field]interface{}{}
	for k, v := range ø.values {
		if !v.IsNil() {
			vals[k] = v.Value
		}
	}
	return
}

func (ø *Row) AliasValues() (vals map[*AsStruct]interface{}) {
	vals = map[*AsStruct]interface{}{}
	for k, v := range ø.aliasValues {
		if !v.IsNil() {
			vals[k] = v.Value
		}
	}
	return
}

func (ø *Row) typedValues() (vals map[*Field]interface{}) {
	vals = map[*Field]interface{}{}
	//pkey := ø.PrimaryKey
	for k, v := range ø.values {
		if ø.IsPrimaryKey(k) && v.IsNil() {
			continue
		}
		vals[k] = v
	}
	return
}

func (ø *Row) InsertQuery() Query {
	return InsertMap(ø.Table, ø.typedValues())
}

func (ø *Row) Insert() (err error) {
	err = ø.Validate()
	if err != nil {
		// return fmt.Errorf("Can't insert row of %s:\n%s\n", ø.Table.Sql(), err.Error())
		return
	}
	ø.setErrors = []error{}
	for _, hook := range ø.PreInsert {
		err := hook(ø)
		if err != nil {
			return err
		}
	}
	u := ø.InsertQuery()
	//_, err = ø.Exec(u)

	/*
		r, err := ø.DB.Exec(u.String())
		if err != nil {
			return err
		}
	*/

	if len(ø.PrimaryKey) == 1 {
		var i string
		err := ø.DB.QueryRow(u.String()).Scan(&i)
		if err != nil {
			// fmt.Printf("Error while inserting: %v\n", err.Error())
			return err
		}
		tv := &TypedValue{ø.PrimaryKey[0].Type, NewPgInterpretedString(i), false}
		ø.values[ø.PrimaryKey[0]] = tv
	} else {
		_, err := ø.Exec(u)
		if err != nil {
			return err
		}
	}
	for _, hook := range ø.PostInsert {
		err := hook(ø)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ø *Row) deleteQuery() Query {
	cond := []Sqler{}

	ids := ø.Id()
	for i, pk := range ø.PrimaryKey {
		cond = append(cond, Equals(pk, ids[i]))
	}

	w := And(cond...)

	u := Delete(ø.Table, Where(w))
	return u
}

func (ø *Row) Delete() (err error) {
	for _, hook := range ø.PreDelete {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	u := ø.deleteQuery()
	_, err = ø.Exec(u)
	for _, hook := range ø.PostDelete {
		err = hook(ø)
		if err != nil {
			return
		}
	}
	return
}

func (ø *Row) SelectQuery(objects ...interface{}) (s *SelectQuery) {
	snew := make([]interface{}, len(objects)+1)
	snew[0] = ø.Table
	for i, o := range objects {
		snew[i+1] = o
	}
	s = Select(snew...).(*SelectQuery)
	if len(s.Fields) == 0 && len(s.FieldsWithAlias) == 0 {
		s.Fields = ø.Fields
	}
	return
}

func (ø *Row) Select(objects ...interface{}) (r *sql.Rows, err error) {
	s := ø.SelectQuery(objects...)
	r, err = ø.Query(s)
	return
}

/*
TODO: change signature to
	IsValid(f string, value interface{}) (error, bool)

then it may be called like this

  if err, is := IsValid(..) ; is {
		// handle validation error
  }

*/
func (ø *Row) IsValid(f string, value interface{}) bool {
	field := ø.Table.Field(f)
	tv := &TypedValue{PgType: field.Type}
	err := Convert(value, tv)
	if err != nil {
		return false
	}
	err = field.Validate(tv)
	if err == nil {
		return true
	}
	return false
}

/*
func (ø *Row) Begin() (tx *sql.Tx, ſ error) {
	if ø.isTransaction() {
		tx = ø.Tx
	} else {
		tx, ſ = ø.DB.Begin()
		ø.SetTransaction(tx)
	}
	return
}
*/

func (ø *Row) Rollback() (ſ error) {
	return ø.Tx.Rollback()
}

func (ø *Row) Commit() (ſ error) {
	return ø.Tx.Commit()
}

func (ø *Row) isTransaction() (is bool) {
	return ø.Tx != nil
}

func (ø *Row) setSearchPath() {
	if !ø.isTransaction() {
		if ø.Table.Schema != nil {
			schemaName := ø.Table.Schema.Name
			sql := `SET search_path = "` + schemaName + `"`
			if ø.Debug {
				fmt.Println(sql)
			}
			_, _ = ø.DB.Exec(sql)
		}
	}
}

func (ø *Row) Exec(query Query, args ...interface{}) (r sql.Result, err error) {
	ø.setSearchPath()
	//ø.Debug = true
	if ø.Debug {
		fmt.Println(query.String())
	}
	ø.LastSql = query.String()
	if ø.isTransaction() {
		r, err = ø.Tx.Exec(query.String(), args...)
	} else {
		r, err = ø.DB.Exec(query.String(), args...)
	}
	return
}

func (ø *Row) Prepare(query Query) (r *sql.Stmt, err error) {
	s := query.Sql()
	// ø.Debug = true
	if ø.Debug {
		fmt.Println(s.String())
	}
	ø.LastSql = s.String()
	ø.setSearchPath()
	if ø.isTransaction() {
		r, err = ø.Tx.Prepare(s.String())
	} else {
		r, err = ø.DB.Prepare(s.String())
	}
	return
}

func (ø *Row) Query(query Query, args ...interface{}) (r *sql.Rows, err error) {
	s := query.Sql()
	ø.LastSql = s.String()
	ø.setSearchPath()
	//ø.Debug = true
	if ø.Debug {
		fmt.Println(s.String())
	}
	if ø.isTransaction() {
		r, err = ø.Tx.Query(s.String(), args...)
	} else {
		r, err = ø.DB.Query(s.String(), args...)
	}
	return
}

func (ø *Row) QueryRow(query Query, args ...interface{}) (r *sql.Row) {
	//ø.Debug = true
	if ø.Debug {
		fmt.Println(query.String())
	}
	s := query.Sql()
	ø.LastSql = s.String()
	ø.setSearchPath()
	if ø.isTransaction() {
		r = ø.Tx.QueryRow(s.String(), args...)
	} else {
		r = ø.DB.QueryRow(s.String(), args...)
	}
	return
}

func (ø *Row) SetDB(db DB) {
	ø.DB = db
}

func (ø *Row) SetTransaction(tx *sql.Tx) {
	ø.Tx = tx
}

func (ø *Row) queryField(name string, opts ...interface{}) interface{} {
	f := ø.Table.QueryField(name)
	if f != nil {
		return f

	}
	for _, o := range opts {
		switch v := o.(type) {
		case *Field:
			if v.QueryField() == name {
				return v
			}
		case *AsStruct:
			if v.QueryField() == name {
				return v
			}
		}
	}
	return nil
}

// TODO make a compilable version that saves the infos about
// fieldnumbers etc and allows faster queriing
func (ø *Row) SelectByStructs(result interface{}, tagVal string, opts ...interface{}) (int, error) {
	/*
		if !meta.Slice.Check(result) {
			return 0, fmt.Errorf("result is no slice")
		}
	*/
	slic := reflect.ValueOf(result)
	l := slic.Len()
	if l == 0 {
		return 0, fmt.Errorf("result slice has length 0")
	}

	stru, err := meta.StructByValue(slic.Index(0))
	if err != nil {
		return 0, err
	}

	tags, err2 := stru.Tags()

	if err2 != nil {
		return 0, err2
	}

	// tags := meta.Struct.Tags(slic.Index(0).Interface())
	options := []interface{}{Limit(l)}
	order := []string{}

	for k, v := range tags {
		if t := v.Get("db.select"); t != "" && strings.Contains(t, tagVal) {
			fi := ø.queryField(k, opts...)
			if fi == nil {
				panic("can't find field " + k + " : no Queryfield property of given fields or aliases")
			}
			options = append(options, fi)
			order = append(order, k)
		}
	}

	options = append(options, opts...)
	rows, err := ø.Find(options...)
	if err != nil {
		return 0, fmt.Errorf("error in find: %s", err.Error())
	}
	i := 0
	errs := []string{}
	for rows.Next() {
		ro, e := rows.ScanRow()
		if e != nil {
			errs = append(errs, e.Error())
			continue
		}
		e = ro.GetStruct(tagVal, slic.Index(i).Addr().Interface())
		if e != nil {
			errs = append(errs, fmt.Sprintf("error while scanning row %v: %s", i, e.Error()))
		}
		i++
	}
	if len(errs) > 0 {
		return i, fmt.Errorf(strings.Join(errs, "\n"))
	}
	return i, nil
}

func (ø *Row) SelectByStruct(structPtr interface{}, tagVal string, opts ...interface{}) error {
	str, err := meta.StructByValue(meta.FinalValue(structPtr))
	if err != nil {
		return err
	}
	tags, err2 := str.Tags()
	if err2 != nil {
		return err2
	}
	options := []interface{}{}
	order := []string{}

	for k, v := range tags {
		if t := v.Get("db.select"); t != "" && strings.Contains(t, tagVal) {
			fi := ø.queryField(k, opts...)
			if fi == nil {
				panic("can't find field " + k + " : no Queryfield property of given fields or aliases")
			}
			options = append(options, fi)
			order = append(order, k)
		}
	}

	options = append(options, opts...)
	row, err := ø.Any(options...)
	if err != nil {
		return err
	}
	err = row.GetStruct(tagVal, structPtr)
	if err != nil {
		return err
	}
	return nil
}
