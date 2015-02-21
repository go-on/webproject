package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"gopkg.in/go-on/builtin.v1/db"
	"math"
	"math/rand"
	"strings"
)

// pseudo random string
// based on https://github.com/SDA/passgen/blob/master/lib/passgen.js
func randomString(length int, alphabet string) string {
	alphabetLength := len(alphabet)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		rnd := math.Floor(rand.Float64() * float64(alphabetLength))
		result[i] = alphabet[int(rnd)]
	}

	return string(result)
}

var escapeStr = "$__pg-esc_" + randomString(5, "_-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") + "_pg-esc__$"

func Escape(value interface{}) (out string) {
	out = strings.Replace(fmt.Sprintf("%v", value), escapeStr, "", -1)
	out = escapeStr + out + escapeStr
	return
}

type SQL string

func (s SQL) String() string { return string(s) }

type Value struct {
	v interface{}
}

func (v *Value) String() string { return Escape(v.v) }

func Val(v interface{}) *Value {
	return &Value{v}
}

type Field string

func (f Field) String() string { return `"` + string(f) + `"` }

type Table string

func (t Table) String() string { return `"` + string(t) + `"` }

type Schema string

func (s Schema) String() string { return `"` + string(s) + `"` }

type DB struct{ db.DB }

var ErrInvalidValue = errors.New("invalid value type, must be pg.SQL, *pg.Value, pg.Field, pg.Table or pg.Schema")

func checkVals(vals ...interface{}) error {
	for _, v := range vals {
		switch v.(type) {
		case *Value, Field, Table, Schema, SQL:
		default:
			return ErrInvalidValue
		}
	}
	return nil
}

// Exec runs an exec for the given db via fmt.Sprintf, without the generation of a prepare statement
// vals must be one of SQL, *Value, Field, Table or Schema
func (d *DB) Exec(query string, vals ...interface{}) (sql.Result, error) {
	if err := checkVals(vals...); err != nil {
		return nil, err
	}
	return d.DB.Exec(fmt.Sprintf(query, vals...))
}

// Query runs a query for the given db via fmt.Sprintf, without the generation of a prepare statement
// vals must be one of SQL, *Value, Field, Table or Schema
func (d *DB) Query(query string, vals ...interface{}) (*sql.Rows, error) {
	if err := checkVals(vals...); err != nil {
		return nil, err
	}
	return d.DB.Query(fmt.Sprintf(query, vals...))
}

// QueryRow runs a query for the given db via fmt.Sprintf, without the generation of a prepare statement
// vals must be one of SQL, *Value, Field, Table or Schema
func (d *DB) QueryRow(query string, vals ...interface{}) *sql.Row {
	if err := checkVals(vals...); err != nil {
		return nil
	}
	return d.DB.QueryRow(fmt.Sprintf(query, vals...))
}

func NewDB(d db.DB) db.DB {
	return &DB{d}
}
