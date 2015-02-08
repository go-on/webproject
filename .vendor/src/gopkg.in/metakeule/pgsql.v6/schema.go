package pgsql

import (
	"database/sql"
	"fmt"
	"gopkg.in/go-on/builtin.v1/db"
	"strings"
)

type Schema struct {
	Name     string
	Database *Database
	Tables   []*Table
}

func NewSchema(name string, options ...interface{}) *Schema {
	s := &Schema{
		Name:   name,
		Tables: []*Table{},
	}
	for _, option := range options {
		switch v := option.(type) {
		case *Database:
			s.Database = v
		case *Table:
			s.AddTable(v)
		}
	}
	return s
}

func (ø *Schema) AddTable(tables ...*Table) {
	for _, f := range tables {
		ø.Tables = append(ø.Tables, f)
		f.Schema = ø
	}
}

func (ø *Schema) Sql() SqlType {
	return Sql(fmt.Sprintf("\"%s\"", ø.Name))
}

type SchemaDB struct {
	name string
	db   db.DBComplete
}

func NewSchemaDB(d db.DBComplete, schemaname string) *SchemaDB {
	return &SchemaDB{schemaname, d}
}

// Transaction receives a function that gets a transaction and returns an error
// it starts a transaction, sets the search path to the schemaname
func (s *SchemaDB) Transaction(fn func(db.DB) error) error {
	tx, err := s.begin()
	if err != nil {
		return err
	}

	err = fn(tx)

	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (s *SchemaDB) begin() (*sql.Tx, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	if strings.Index(s.name, "$") != -1 {
		return nil, fmt.Errorf("$ in schema name not allowed: %#v", s.name)
	}

	_, err = tx.Exec(fmt.Sprintf("SET search_path = $schemaname$%s$schemaname$", s.name))

	if err != nil {
		return nil, err
	}

	return tx, nil
}
