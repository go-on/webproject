package pg

import (
	"gopkg.in/metakeule/dbwrap.v2"
	"testing"
)

func TestEscape(t *testing.T) {

	tests := []struct {
		input    interface{}
		expected string
	}{
		{
			"hi" + escapeStr + "ho",
			escapeStr + "hiho" + escapeStr,
		},
		{
			escapeStr + "hi" + escapeStr + "ho",
			escapeStr + "hiho" + escapeStr,
		},
		{
			escapeStr + "$hi$" + escapeStr + "hu$hi$",
			escapeStr + "$hi$hu$hi$" + escapeStr,
		},
		{
			"$hi$hu$hi$",
			escapeStr + "$hi$hu$hi$" + escapeStr,
		},
		{
			"hi",
			escapeStr + "hi" + escapeStr,
		},
		{
			0.45423,
			escapeStr + "0.45423" + escapeStr,
		},
		{
			434,
			escapeStr + "434" + escapeStr,
		},
	}

	for _, test := range tests {
		if got, want := Escape(test.input), test.expected; got != want {
			t.Errorf("Escape(%v) = %#v; want %#v", test.input, got, want)
		}
	}

}

func TestExec(t *testing.T) {
	fake, d := dbwrap.NewFake()

	db := NewDB(d)

	tests := []struct {
		query  string
		vals   []interface{}
		result string
		err    error
	}{
		{
			"select %s from %s",
			[]interface{}{Field("a"), Table("b")},
			`select "a" from "b"`,
			nil,
		},
		{
			"select %s from %s",
			[]interface{}{"a", Table("b")},
			``,
			ErrInvalidValue,
		},
		{
			"select %s from %s where id = %s",
			[]interface{}{Field("a"), Table("b"), Val(3)},
			`select "a" from "b" where id = ` + escapeStr + "3" + escapeStr,
			nil,
		},
		{
			"select %[2]s.%[3]s, %[2]s.%[4]s from %[1]s.%[2]s where %[3]s = %[5]s",
			[]interface{}{Schema("testdb"), Table("person"), Field("id"), Field("name"), Val(3)},
			`select "person"."id", "person"."name" from "testdb"."person" where "id" = ` + escapeStr + "3" + escapeStr,
			nil,
		},
	}

	for _, test := range tests {
		_, err := db.Exec(test.query, test.vals...)
		if got, want := err, test.err; got != want {
			t.Errorf("db.Exec(%#v,%#v); err = %#v; want %#v", test.query, test.vals, got, want)
		}

		if err == nil {
			q, _ := fake.LastQuery()
			if got, want := q, test.result; got != want {
				t.Errorf("db.Exec(%#v,%#v) = %#v; want %#v", test.query, test.vals, got, want)
			}
		}
	}
}

func TestQuery(t *testing.T) {
	fake, d := dbwrap.NewFake()

	db := NewDB(d)

	tests := []struct {
		query  string
		vals   []interface{}
		result string
		err    error
	}{
		{
			"select %s from %s",
			[]interface{}{Field("a"), Table("b")},
			`select "a" from "b"`,
			nil,
		},
		{
			"select %s from %s",
			[]interface{}{"a", Table("b")},
			``,
			ErrInvalidValue,
		},
		{
			"select %s from %s where id = %s",
			[]interface{}{Field("a"), Table("b"), Val(3)},
			`select "a" from "b" where id = ` + escapeStr + "3" + escapeStr,
			nil,
		},
		{
			"select %[2]s.%[3]s, %[2]s.%[4]s from %[1]s.%[2]s where %[3]s = %[5]s",
			[]interface{}{Table("testdb"), Table("person"), Field("id"), Field("name"), Val(3)},
			`select "person"."id", "person"."name" from "testdb"."person" where "id" = ` + escapeStr + "3" + escapeStr,
			nil,
		},
	}

	for _, test := range tests {
		_, err := db.Query(test.query, test.vals...)
		if got, want := err, test.err; got != want {
			t.Errorf("db.Query(%#v,%#v); err = %#v; want %#v", test.query, test.vals, got, want)
		}

		if err == nil {
			q, _ := fake.LastQuery()
			if got, want := q, test.result; got != want {
				t.Errorf("db.Query(%#v,%#v) = %#v; want %#v", test.query, test.vals, got, want)
			}
		}
	}
}

func TestQueryRow(t *testing.T) {
	fake, d := dbwrap.NewFake()

	db := NewDB(d)

	tests := []struct {
		query  string
		vals   []interface{}
		result string
		err    error
	}{
		{
			"select %s from %s",
			[]interface{}{Field("a"), Table("b")},
			`select "a" from "b"`,
			nil,
		},
		{
			"select %s from %s",
			[]interface{}{"a", Table("b")},
			``,
			ErrInvalidValue,
		},
		{
			"select %s from %s where id = %s",
			[]interface{}{Field("a"), Table("b"), Val(3)},
			`select "a" from "b" where id = ` + escapeStr + "3" + escapeStr,
			nil,
		},
		{
			"select %[2]s.%[3]s, %[2]s.%[4]s from %[1]s.%[2]s where %[3]s = %[5]s",
			[]interface{}{Table("testdb"), Table("person"), Field("id"), Field("name"), Val(3)},
			`select "person"."id", "person"."name" from "testdb"."person" where "id" = ` + escapeStr + "3" + escapeStr,
			nil,
		},
	}

	for _, test := range tests {
		row := db.QueryRow(test.query, test.vals...)
		if test.err != nil && row != nil {
			t.Errorf("db.QueryRow(%#v,%#v); row not nil (should be)", test.query, test.vals)
		}

		if test.err == nil {
			q, _ := fake.LastQuery()
			if got, want := q, test.result; got != want {
				t.Errorf("db.QueryRow(%#v,%#v) = %#v; want %#v", test.query, test.vals, got, want)
			}
		}
	}
}
