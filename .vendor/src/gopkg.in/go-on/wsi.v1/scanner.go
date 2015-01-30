package wsi

import (
	"database/sql"
	"errors"
)

// Scanner is a more comfortable scanner that works similar to sql.Rows but does not have to be closed.
type Scanner interface {

	// Next should return false if there are no rows left or if any error happened before
	Next() bool

	// Scan allows scanning by column name instead of column position
	// If an error did happen, every successing call to Scan should return that error without doing any scanning
	Scan(map[string]interface{}) error

	// Error should return the first error that did happen
	Error() error
}

// NewTestScanner returns a new faking Scanner using the given function to fake the scanning.
// The function is called each time the scanners Scan method is called and the target map is passed to fn.
// If fn returns an error, each further call of Scan will return this error and fn is no longer called.
// If fn returns a stop or an error the Next method of the scanner will return false.
func NewTestScanner(fn func(targets map[string]interface{}) (stop bool, err error)) *TestScanner {
	return &TestScanner{
		fn: fn,
	}
}

type TestScanner struct {
	fn   func(targets map[string]interface{}) (stop bool, err error)
	err  error
	stop bool
}

// Error returns the first error that did happen
func (f *TestScanner) Error() error { return f.err }

// Next returns false if there are no rows left or if any error happened before
func (f *TestScanner) Next() bool {
	// fmt.Println("next called")
	if f.err != nil {
		return false
	}
	return !f.stop
}

// Scan allows scanning by column name instead of column position
func (f *TestScanner) Scan(targets map[string]interface{}) error {
	// fmt.Println("scan called")
	if f.err != nil {
		return f.err
	}
	f.stop, f.err = f.fn(targets)
	return f.err
}

// errScanner is a Scanner that does nothing but returning errors
type errScanner struct{ err error }

func (e errScanner) Error() error                      { return e.err }
func (e errScanner) Next() bool                        { return false }
func (e errScanner) Scan(map[string]interface{}) error { return e.err }

// dbScanner wraps sql.Rows to autoclose the rows and allow scanning by a map of fields to targets
// so that the call is independant from the position of the returned columns
type dbScanner struct {
	columns map[string]int
	*sql.Rows
	err    error
	closed bool
}

// Error returns the first error that did happen
func (sc *dbScanner) Error() error {
	return sc.err
}

// Next returns false if there are no rows left or if any error happened before
func (sc *dbScanner) Next() bool {
	if sc.err != nil {
		if !sc.closed {
			sc.Rows.Close()
			sc.closed = true
		}
		return false
	}
	return sc.Rows.Next()
}

// Scan allows scanning by column name instead of column position
// If an error did happen, every successing call to Scan will return that error without doing any scanning
func (sc *dbScanner) Scan(targets map[string]interface{}) error {
	if sc.err != nil {
		return sc.err
	}
	vals := make([]interface{}, len(sc.columns))

	for colName, val := range targets {
		i, ok := sc.columns[colName]
		if !ok {
			return errors.New("unknown column " + colName)
		}
		vals[i] = val
	}

	sc.err = sc.Rows.Scan(vals...)
	if sc.err != nil && !sc.closed {
		sc.Rows.Close()
		sc.closed = true
	}
	return sc.err
}
