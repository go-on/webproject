package pgsqlfat

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gopkg.in/go-on/lib.v3/internal/fat"
	"gopkg.in/go-on/lib.v3/internal/meta"
	"gopkg.in/metakeule/fmtdate.v1"
	. "gopkg.in/metakeule/pgsql.v5"
)

var fatField *fat.Field
var fatFieldNil = reflect.ValueOf(fatField)

/*
  only the fields that are *fat.Field and not nil are chosen to be set
	from the row. fields that are not set in the row, are set to nil
*/
func (r *Registry) FromRow(row *Row, øptrToFatStruct interface{}) (err error) {
	fn := func(field *meta.Field) {
		if err != nil {
			return
		}

		if field.Value.IsNil() {
			return
		}

		ff, isFat := field.Value.Interface().(*fat.Field)

		if !isFat {
			return
		}

		dbField := r.FieldOf(ff)

		if row.Values()[dbField] != nil {
			fatField := field.Value.Interface().(*fat.Field)
			err = scanFieldToStruct(row, fatField, dbField)
			return
		}

		field.Value.Set(fatFieldNil)
		//		ff.Set(fatFieldNil)
		/*
			if dbField.Is(NullAllowed) {
				ff.Set(fatFieldNil)
			}
		*/
	}

	var stru *meta.Struct
	stru, err = meta.StructByValue(reflect.ValueOf(øptrToFatStruct))

	if err == nil {
		stru.Each(fn)
	}
	return
}

func scanFieldToStruct(row *Row, fatField *fat.Field, dbField *Field) (err error) {
	switch dbField.Type {

	case TimeType, DateType, TimeStampType, TimeStampTZType:
		var t time.Time
		row.Get(dbField, &t)
		err = fatField.Set(t)

	case IntsType:
		err = fatField.ScanString("[" + TrimCurly(row.GetString(dbField)) + "]")

	case StringsType:
		ss := []string{}
		s_s := strings.Split(TrimCurly(row.GetString(dbField)), ",")
		for _, sss := range s_s {
			ss = append(ss, strings.Trim(sss, `" `))
		}
		err = fatField.Set(fat.Strings(ss...))

	case JsonType:
		err = fatField.ScanString(row.GetString(dbField))

	case BoolsType:
		vls := strings.Split(TrimCurly(row.GetString(dbField)), ",")
		bs := make([]bool, len(vls))

		for j, bstri := range vls {
			switch strings.TrimSpace(bstri) {
			case "t":
				bs[j] = true
			case "f":
				bs[j] = false
			default:
				return fmt.Errorf("%s is no []bool", row.GetString(dbField))
			}
		}
		err = fatField.Set(fat.Bools(bs...))

	case FloatsType:
		err = fatField.ScanString("[" +
			TrimCurly(row.GetString(dbField)) +
			"]")

	case TimeStampsTZType:

		var ts string
		row.Get(dbField, &ts)
		vls := strings.Split(TrimCurly(ts), ",")
		tms := make([]time.Time, len(vls))
		for j, tmsStri := range vls {
			tm, e := fmtdate.Parse("YYYY-MM-DD hh:mm:ss+00", strings.Trim(tmsStri, `"`))
			if e != nil {
				return fmt.Errorf("can't parse time %s: %s", tmsStri, e.Error())
			}
			tms[j] = tm
		}
		err = fatField.Set(fat.Times(tms...))

	default:
		err = fatField.Scan(row.GetString(dbField))
	}

	if err != nil {
		return err
	}
	errs := fatField.Validate()
	if len(errs) > 0 {
		var errStr bytes.Buffer
		for _, e := range errs {
			errStr.WriteString(e.Error() + "\n")
		}
		return fmt.Errorf("Can't set field %s: %s", fatField.Name(), errStr.String())
	}

	return nil
}

func TrimCurly(in string) string {
	in = strings.Replace(in, "{", "", -1)
	in = strings.Replace(in, "}", "", -1)
	return in
}
