package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "gopkg.in/metakeule/pgsql.v6"
	"gopkg.in/metakeule/pgsql.v6/pgsqlfat"
	"gopkg.in/go-on/builtin.v1/db"
	"gopkg.in/go-on/lib.v3/internal/fat"
	"gopkg.in/go-on/lib.v3/internal/meta"
	"gopkg.in/go-on/router.v2"
	"gopkg.in/metakeule/fmtdate.v1"
	// "net/url"
	"reflect"
	"strings"
	"time"
)

type CRUD struct {
	table        *Table
	prototype    interface{}
	type_        reflect.Type
	type_string  string
	fields       map[string]map[string]bool
	primaryKey   *Field
	pKeyIsString bool
	*pgsqlfat.Registry
}

func NewCRUD(registry *pgsqlfat.Registry, proto interface{}) (c *CRUD) {
	c = &CRUD{
		Registry: registry,
		table:    registry.TableOf(proto),
		// table:       table,
		prototype:   proto,
		type_:       reflect.TypeOf(proto),
		type_string: pgsqlfat.TypeString(proto),
		fields:      map[string]map[string]bool{},
	}
	err := c.scanFields()
	if err != nil {
		panic(fmt.Sprintf("can't scan fields of %#v (%T): %s", proto, proto, err))
	}
	return
}

func (r *CRUD) Mount(d db.DB, rt *router.Router, mountPoint string, options *options) *Mounter {
	return Mount(r, d, rt, mountPoint, options)
}

func (r *CRUD) scanFields() (err error) {
	var hasDeleteField bool
	//fn := func(fld reflect.StructField, vl reflect.Value, tag string) {
	fn := func(fld *meta.Field) {
		if err != nil {
			return
		}
		tag := fld.Type.Tag.Get("rest")
		if tag == "" {
			return
		}

		methods := map[string]bool{}

		if strings.Contains(tag, "C") {
			methods["C"] = true
		}

		if strings.Contains(tag, "D") {
			if hasDeleteField {
				err = fmt.Errorf("more than one delete field (key) is not supported")
				return
			}
			methods["D"] = true
			hasDeleteField = true
		}

		if strings.Contains(tag, "L") {
			methods["L"] = true
		}

		if strings.Contains(tag, "U") {
			methods["U"] = true
		}

		if strings.Contains(tag, "R") {
			methods["R"] = true
		}

		if len(methods) == 0 {
			return
		}

		//ff := r.field(fld.Name)
		ff := r.field(fld.Type.Name)
		if ff == nil {
			err = fmt.Errorf("can't find field for table %s field %s\n", r.typeString(), fld.Type.Name)
			return
		}

		// pgsql.flags
		if strings.Contains(fld.Type.Tag.Get("db"), "PKEY") {
			if r.primaryKey != nil {
				err = fmt.Errorf("can't have more than one primary key %s and %s\n", r.primaryKey, fld.Type.Name)
				return
			}
			r.primaryKey = r.field(fld.Type.Name)
		}

		r.fields[fld.Type.Name] = methods
	}

	var stru *meta.Struct
	stru, err = meta.StructByValue(reflect.ValueOf(r.prototype))
	stru.Each(fn)
	//meta.Struct.EachTag(r.prototype, "rest", fn)

	if r.primaryKey == nil {
		err = fmt.Errorf("has not primary key, add db:\"PKEY\"")
	}

	pkType := r.primaryKey.Type
	if !pkType.IsCompatible(IntType) {
		if pkType.IsCompatible(TextType) {
			r.pKeyIsString = true
			return
		}
		err = fmt.Errorf("primary key %s (%s) is not compatible to int or string", r.primaryKey.Name, pkType.String())
	}

	return
}

func (r *CRUD) newObject() interface{} {
	return fat.New(r.prototype, reflect.New(r.type_.Elem()).Interface())
}
func (r *CRUD) newObjects(num int) []interface{} { return make([]interface{}, num) }
func (r *CRUD) typeString() string {
	return r.type_string
}
func (c *CRUD) field(fld string) *Field { return c.Field(c.typeString(), fld) }

var fatField *fat.Field
var fatFieldNil = reflect.ValueOf(fatField)

/*
func pgsqlfat.TrimCurly(in string) string {
	in = strings.Replace(in, "{", "", -1)
	in = strings.Replace(in, "}", "", -1)
	// in = strings.Trim(in, `"`)
	return in
}
*/

/*
ScanFieldToStruct scans a field of the *Row from a database query into the given fatstruct and
returns the value that would be inserted into a json representation or an error, if scan or
validation fails
*/
func (c *CRUD) scanFieldToStruct(queryRow *Row, structField reflect.Value, dbField *Field) (jsonVal interface{}, err error) {
	fatField := structField.Interface().(*fat.Field)
	var stringInMap bool

	switch dbField.Type {
	case TimeType, DateType, TimeStampType, TimeStampTZType:
		stringInMap = true
		var t time.Time
		queryRow.Get(dbField, &t)
		err = fatField.Set(t)
	case IntsType:
		err = fatField.ScanString("[" +
			pgsqlfat.TrimCurly(queryRow.GetString(dbField)) +
			"]")
	case StringsType:
		ss := []string{}
		s__ := pgsqlfat.TrimCurly(queryRow.GetString(dbField))
		if s__ != "" {
			s_s := strings.Split(s__, ",")
			for _, sss := range s_s {
				ss = append(ss, strings.Trim(sss, `" `))
			}
		}
		err = fatField.Set(fat.Strings(ss...))
	case JsonType:
		err = fatField.ScanString(queryRow.GetString(dbField))
	case BoolsType:
		s__ := pgsqlfat.TrimCurly(queryRow.GetString(dbField))
		var bs []bool
		if s__ != "" {
			vls := strings.Split(s__, ",")
			bs = make([]bool, len(vls))
			for j, bstri := range vls {
				switch strings.TrimSpace(bstri) {
				case "t":
					bs[j] = true
				case "f":
					bs[j] = false
				default:
					return nil, fmt.Errorf("%s is no []bool", queryRow.GetString(dbField))
				}
			}
		}
		err = fatField.Set(fat.Bools(bs...))
	case FloatsType:
		err = fatField.ScanString("[" +
			pgsqlfat.TrimCurly(queryRow.GetString(dbField)) +
			"]")
	case TimeStampsTZType:
		//var t []time.Time
		var ts string
		queryRow.Get(dbField, &ts)
		s__ := pgsqlfat.TrimCurly(ts)
		var tms []time.Time
		if s__ != "" {
			vls := strings.Split(s__, ",")
			tms = make([]time.Time, len(vls))
			for j, tmsStri := range vls {
				tm, e := fmtdate.Parse("YYYY-MM-DD hh:mm:ss+00", strings.Trim(tmsStri, `"`))
				if e != nil {
					return nil, fmt.Errorf("can't parse time %s: %s", tmsStri, e.Error())
				}
				tms[j] = tm
			}
		}
		// fmt.Printf("times: %#v\n", tms)
		err = fatField.Set(fat.Times(tms...))
		/*
			err = fatField.ScanString("[" +
				pgsqlfat.TrimCurly(queryRow.GetString(dbField)) +
				"]")
		*/
	default:
		err = fatField.Scan(queryRow.GetString(dbField))
	}

	if err != nil {
		return nil, err
	}
	errs := fatField.Validate()
	if len(errs) > 0 {
		var errStr bytes.Buffer
		for _, e := range errs {
			errStr.WriteString(e.Error() + "\n")
		}
		return nil, fmt.Errorf("Can't set field %s: %s", fatField.Name(), errStr.String())
	}

	if stringInMap {
		jsonVal = fatField.String()
	} else {
		jsonVal = fatField.Get()
	}

	return
}

/*
ScanToStruct scans the *Row from a database query into the given fatstruct.
It returns error if the scan could not be done or if the validation for the fatstruct fails.
*/
func (c *CRUD) scanToStruct(queryRow *Row, taggedFields []string, fatstruc interface{}) error {
	structV := reflect.ValueOf(fatstruc).Elem()
	for _, taggedField := range taggedFields {
		dbField := c.field(taggedField)
		structField := structV.FieldByName(taggedField)
		queryVal := queryRow.Values()[dbField]
		if queryVal != nil {
			_, err := c.scanFieldToStruct(queryRow, structField, dbField)
			if err != nil {
				return err
			}
		} else {
			if dbField.Is(NullAllowed) {
				structField.Set(fatFieldNil)
			}
		}
	}
	return nil
}

/*
ScanToStructAndMap scans the *Row from a database query into the given fatstruct and into the
jsonMap for json output. It returns error if the scan could not be done or if the validation
for the fatstruct fails.
*/
func (c *CRUD) scanToStructAndMap(queryRow *Row, taggedFields []string, fatstruc interface{}, jsonMap map[string]interface{}) error {
	structV := reflect.ValueOf(fatstruc).Elem()
	for _, taggedField := range taggedFields {
		dbField := c.field(taggedField)
		structField := structV.FieldByName(taggedField)
		queryVal := queryRow.Values()[dbField]
		if queryVal != nil {
			mapVal, err := c.scanFieldToStruct(queryRow, structField, dbField)
			if err != nil {
				return err
			}
			jsonMap[taggedField] = mapVal
		} else {
			if dbField.Is(NullAllowed) {
				structField.Set(fatFieldNil)
			}
			jsonMap[taggedField] = nil
		}
	}
	return nil
}

func (c *CRUD) setFieldToStruct(jsonVal interface{}, taggedField string, fatStructVal reflect.Value, backupStruct reflect.Value) (err error) {
	structField := fatStructVal.FieldByName(taggedField)
	backupField := backupStruct.FieldByName(taggedField)
	newFatField := backupField.Interface().(*fat.Field)

	if jsonVal == nil {
		if newFatField.Default() != nil {
			structField.Set(backupField)
			return nil
		}

		dbField := c.field(taggedField)
		if dbField.Is(NullAllowed) {
			structField.Set(fatFieldNil)
			return nil
		}
		return fmt.Errorf("null not allowed")
	}

	//	fmt.Printf("json val is: %T\n", jsonVal)

	switch jsonValTyped := jsonVal.(type) {
	case map[string]interface{}:
		var bt []byte
		bt, err = json.Marshal(jsonValTyped)
		if err == nil {
			err = newFatField.ScanString(string(bt))
		}
	//case map[interface{}]interface{}:
	case []interface{}:
		var inputTypeOk bool
		switch newFatField.Typ() {
		case "[]float":
			flts := make([]float64, len(jsonValTyped))
			for i, intf := range jsonValTyped {
				flts[i], inputTypeOk = intf.(float64)
				if !inputTypeOk {
					//err = fmt.Errorf("is no float64: %v (%T)", intf, intf)
					err = fmt.Errorf("is no float")
					break
				}
			}
			err = newFatField.Set(fat.Floats(flts...))
		case "[]int":
			ints := make([]int64, len(jsonValTyped))

			for i, intf := range jsonValTyped {
				var fl float64
				fl, inputTypeOk = intf.(float64)
				ints[i] = toInt64(fl)
				if !inputTypeOk {
					//err = fmt.Errorf("is no float64: %v (%T)", intf, intf)
					err = fmt.Errorf("is no float")
					break
				}
			}
			err = newFatField.Set(fat.Ints(ints...))
		case "[]string":
			strings := make([]string, len(jsonValTyped))

			for i, intf := range jsonValTyped {
				strings[i], inputTypeOk = intf.(string)
				if !inputTypeOk {
					//err = fmt.Errorf("is no string: %v (%T)", intf, intf)
					err = fmt.Errorf("is no string")
					break
				}
			}
			err = newFatField.Set(fat.Strings(strings...))
		case "[]bool":
			bools := make([]bool, len(jsonValTyped))

			for i, intf := range jsonValTyped {
				bools[i], inputTypeOk = intf.(bool)
				if !inputTypeOk {
					//err = fmt.Errorf("is no bool: %v (%T)", intf, intf)
					err = fmt.Errorf("is no bool")
					break
				}
			}
			err = newFatField.Set(fat.Bools(bools...))
		case "[]time":
			times := make([]time.Time, len(jsonValTyped))
			for i, intf := range jsonValTyped {
				// fmt.Printf("[]Time: %#v\n", intf)
				var timestr string
				timestr, inputTypeOk = intf.(string)

				//times[i], inputTypeOk = intf.(time.Time)
				if !inputTypeOk {
					//err = fmt.Errorf("is no time: %v (%T)", intf, intf)
					err = fmt.Errorf("is no time")
					break
				}

				ti, e := time.Parse(time.RFC3339, timestr)
				if e != nil {
					//err = fmt.Errorf("can't parse time: %v: %s ", timestr, e.Error())
					err = fmt.Errorf("invalid time")
					break
				}
				times[i] = ti
			}
			err = newFatField.Set(fat.Times(times...))

		default:
			//err = fmt.Errorf("unsupported type: %#v", newFatField.Typ())
			err = fmt.Errorf("unsupported type")
		}

	case float64:
		switch newFatField.Typ() {
		case "float":
			err = newFatField.Set(jsonVal)
		case "int":
			err = newFatField.Set(int64(jsonValTyped))
		default:
			err = newFatField.Set(jsonVal)
		}
	case string:
		err = newFatField.ScanString(jsonValTyped)
		if err != nil {
			err = fmt.Errorf("is no %s", newFatField.Typ())
		}
	default:
		err = newFatField.Set(jsonVal)
	}
	errs := newFatField.Validate()
	if len(errs) > 0 {
		//var errStr bytes.Buffer
		errStrings := make([]string, len(errs))
		for j, e := range errs {
			errStrings[j] = e.Error()
			// errStr.WriteString(e.Error() + "\n")
		}

		//return fmt.Errorf("Can't set field %s: %s", taggedField, errStr.String())
		//return fmt.Errorf("error while setting")
		//return fmt.Errorf(errStr.String())
		return fmt.Errorf(strings.Join(errStrings, "\n"))
	}

	if err != nil {
		return err
	}

	structField.Set(backupField)
	return
}

func (c *CRUD) setMapToStruct(jsonMap map[string]interface{}, taggedFields []string, fatstruc interface{}) (errs map[string]error) {
	backupStruct := reflect.ValueOf(c.newObject()).Elem()
	fatStructVal := reflect.ValueOf(fatstruc).Elem()
	errs = map[string]error{}

	for _, taggedField := range taggedFields {
		//jsonVal, willSet := jsonMap[taggedField]
		jsonVal := jsonMap[taggedField]
		// if willSet {
		err := c.setFieldToStruct(jsonVal, taggedField, fatStructVal, backupStruct)
		if err != nil {
			errs[taggedField] = err
		}
		// }
	}
	return
}

// SetMap takes a map of fields to values and and sets the fatstruct based on the
// field names and values
// only the given fields are changed. if a value is nil, the field will be set to nil instead
// of a fat field
// since fatstruct might have nil values that will be set to non nil
// values, we need a backup object where every field is not nil,
// so we can take the field from it
func (c *CRUD) queryParams(method string) (query []interface{}, fields []string) {
	query = []interface{}{}
	for fld, has := range c.fields {
		if has[method] {
			query = append(query, c.field(fld))
			fields = append(fields, fld)
		}
	}
	return
}
