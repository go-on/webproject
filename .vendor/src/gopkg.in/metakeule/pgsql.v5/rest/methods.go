package rest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/go-on/lib.v3/internal/fat"
	. "gopkg.in/metakeule/pgsql.v5"
)

func (r *CRUD) Create(db DB, json_ []byte, validateOnly bool, singleField string) (id string, err error) {
	m := map[string]interface{}{}
	err = json.Unmarshal(json_, &m)
	if err != nil {
		return
	}

	_, fields := r.queryParams("C")
	fatstruc := r.newObject()

	if singleField != "" {
		exists := false

		for _, fff := range fields {
			if fff == singleField {
				exists = true
				break
			}
		}

		if !exists {
			err = fieldNotAllowed{} // fmt.Errorf("field not allowed for method")
			return
		}

		fields = []string{singleField}
	}

	errs := r.setMapToStruct(m, fields, fatstruc)
	if len(errs) > 0 {
		//err = NewValidationError(errs)
		err = &validationError{errs}
		// fmt.Println("set map failed")
		return
	}

	row := NewRow(db, r.table)
	fs := reflect.ValueOf(fatstruc).Elem()

	for _, fld := range fields {

		if fs.FieldByName(fld).IsNil() {
			row.Set(r.field(fld), nil)
			continue
		}

		ff := fs.FieldByName(fld).Interface().(*fat.Field)

		v := ff.Get()

		switch v.(type) {
		case []fat.Type:
			vl := ff.String()
			vl = strings.Replace(vl, "[", "{", -1)
			vl = strings.Replace(vl, "]", "}", -1)
			err = row.Set(r.field(fld), vl)
		case map[string]fat.Type:
			err = row.Set(r.field(fld), ff.String())
		default:
			err = row.Set(r.field(fld), v)
		}
		if err != nil {
			fmt.Printf("err setting %s: %s\n", fld, err.Error())
			return
		}
	}
	err = row.Validate()
	if err != nil {
		err = &validationError{map[string]error{"": err}}
		// fmt.Printf("error in validation: %s\n", err)
		return
	}

	if validateOnly {
		return
	}
	err = row.Insert()
	id = row.GetString(r.primaryKey)
	return
}

func (r *CRUD) addWhereId(query []interface{}, id string) (q []interface{}, err error) {
	var w *WhereStruct
	if r.pKeyIsString {
		w = Where(
			Equals(
				r.primaryKey,
				r.primaryKey.MustValue(id),
			),
		)
	} else {
		var i int
		i, err = strconv.Atoi(id)
		if err != nil {
			return
		}
		w = Where(
			Equals(
				r.primaryKey,
				r.primaryKey.MustValue(i),
			),
		)
	}
	q = append(query, w)
	return
}

func (r *CRUD) _Read(db DB, id string) (rr *Row, fields []string, err error) {
	row := NewRow(db, r.table)
	var query []interface{}
	query, fields = r.queryParams("R")
	query, err = r.addWhereId(query, id)
	if err != nil {
		fmt.Printf("addWhereId error: %s\n", err.Error())
		return
	}
	rr, err = row.Any(query...)
	return
}

func (r *CRUD) Read(db DB, id string) (m map[string]interface{}, err error) {
	row := NewRow(db, r.table)
	query, fields := r.queryParams("R")
	query, err = r.addWhereId(query, id)
	if err != nil {
		fmt.Printf("addWhereId error: %s\n", err.Error())
		return
	}
	var rr *Rows
	rr, err = row.Find(query...)
	if err != nil {
		fmt.Printf("error in sql find query: %s\n", err)
		return
	}
	if !rr.Next() {
		return nil, notFound{}
	}

	ro, e := rr.ScanRow()
	if e != nil {
		err = e
		fmt.Printf("error in scanrow: %s\n", err)
		return
	}
	fatstruc := r.newObject()
	m = map[string]interface{}{}
	err = r.scanToStructAndMap(ro, fields, fatstruc, m)
	if err != nil {
		fmt.Printf("error in scanmap: %s\n", err)
	}
	return
}

func (r *CRUD) Update(db DB, id string, json_ []byte, validateOnly bool, singleField string) (err error) {
	m := map[string]interface{}{}
	err = json.Unmarshal(json_, &m)
	if err != nil {
		return
	}

	row := NewRow(db, r.table)
	query, fields := r.queryParams("U")
	query, err = r.addWhereId(query, id)
	if err != nil {
		fmt.Printf("addWhereId error: %s\n", err.Error())
		return
	}
	// add the primary key to get it back!
	query = append(query, r.primaryKey)

	var rr *Row
	rr, err = row.Any(query...)
	if err != nil || rr == nil {
		return notFound{}
		// fmt.Printf("select error: %s\n", err.Error())
		// return
	}
	fatstruc := r.newObject()
	err = r.scanToStruct(rr, fields, fatstruc)
	if err != nil {
		fmt.Printf("scan error: %s\n", err)
		return err
	}

	if singleField != "" {
		exists := false

		for _, fff := range fields {
			if fff == singleField {
				exists = true
				break
			}
		}

		if !exists {
			err = fieldNotAllowed{} // fmt.Errorf("field not allowed for method")
			return
		}

		fields = []string{singleField}
	}

	errs := r.setMapToStruct(m, fields, fatstruc)
	if len(errs) > 0 {
		err = &validationError{errs}
		fmt.Printf("set error: %s\n", err)
		// return err
		return
	}
	fs := reflect.ValueOf(fatstruc).Elem()
	for _, fld := range fields {
		if fs.FieldByName(fld).IsNil() {
			row.Set(r.field(fld), nil)
			continue
		}
		ff := fs.FieldByName(fld).Interface().(*fat.Field)
		v := ff.Get()
		var err error
		switch v.(type) {
		case []fat.Type:
			vl := ff.String()
			vl = strings.Replace(vl, "[", "{", -1)
			vl = strings.Replace(vl, "]", "}", -1)
			err = row.Set(r.field(fld), vl)
		case map[string]fat.Type:
			err = row.Set(r.field(fld), ff.String())
		default:
			err = row.Set(r.field(fld), ff.Get())
		}
		if err != nil {
			return err
		}
	}

	err = row.Validate()
	if err != nil {
		return &validationError{map[string]error{"": err}}
	}

	if validateOnly {
		return
	}
	return row.Update()
}

func (r *CRUD) Delete(db DB, id string) error {
	row := NewRow(db, r.table)
	err := row.SetId(id)
	if err != nil {
		return &validationError{map[string]error{r.primaryKey.Name: err}}
	}
	err = row.Delete()
	if err != nil {
		return ErrServer
	}
	return nil
}

// , rangeReq.Desc, sortBy, rangeReq.Start
// total, objs, err := r.List(db, limit, rangeReq.Desc, sortBy, rangeReq.Start)

func (r *CRUD) List(db DB, limit int, direction Direction, sortBy *Field, offset int) (total int, ms []map[string]interface{}, err error) {
	ms = []map[string]interface{}{}
	row := NewRow(db, r.table)

	countRow := db.QueryRow(Select(As(Sql(`count("`+r.primaryKey.Name+`")`), "no", IntType), r.table).Sql().String())

	err = countRow.Scan(&total)

	if err != nil {
		return
	}

	query, fields := r.queryParams("L")
	query = append(query, Limit(limit), OrderBy(sortBy, direction), Offset(offset))

	/*
		if len(queryParams) > 0 {
			query = append(query, queryParams...)
		}
	*/

	var rr *Rows
	rr, err = row.Find(query...)
	if err != nil {
		return
	}

	for rr.Next() {
		ro, e := rr.ScanRow()
		if e != nil {
			err = e
			return
		}
		fatstruc := r.newObject()
		m := map[string]interface{}{}
		err = r.scanToStructAndMap(ro, fields, fatstruc, m)
		if err != nil {
			return
		}
		ms = append(ms, m)
	}
	return
}
