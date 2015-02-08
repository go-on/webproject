package person

import (
	"database/sql"
)

type FakeDB struct {
	LastQuery       string
	LastQueryParams []interface{}
}

func (ø *FakeDB) LastInsertId() (id int64, err error) {
	id = 1
	return
}

func (ø *FakeDB) RowsAffected() (id int64, err error) {
	id = 1
	return
}

func (ø *FakeDB) Exec(query string, args ...interface{}) (s sql.Result, err error) {
	ø.LastQuery = query
	ø.LastQueryParams = args
	s = ø
	return
}

func (ø *FakeDB) Query(query string, args ...interface{}) (s *sql.Rows, err error) {
	ø.LastQuery = query
	ø.LastQueryParams = args
	s = &sql.Rows{}
	return
}

func (ø *FakeDB) QueryRow(query string, args ...interface{}) (s *sql.Row) {
	ø.LastQuery = query
	ø.LastQueryParams = args
	s = &sql.Row{}
	return
}

func (ø *FakeDB) Prepare(query string) (s *sql.Stmt, err error) {
	ø.LastQuery = query
	ø.LastQueryParams = []interface{}{}
	s = &sql.Stmt{}
	return
}
