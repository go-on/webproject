package wsi

import (
	"gopkg.in/go-on/builtin.v1/sqlnull"

	"database/sql"
)

// Scanner is a more comfortable scanner that works similar to sql.Rows but does not have to be closed.
type Scanner interface {

	// Next should return false if there are no rows left or if any error happened before
	Next() bool

	// Scan allows scanning by column name instead of column position
	// If an error did happen, every successing call to Scan should return that error without doing any scanning
	Scan(vals ...interface{}) error

	// ColNum returns the number of columns
	// ColNum() int

	// Column returns the column name for the given position
	// Column(pos int) string

	Columns() []string

	// Error should return the first error that did happen
	Error() error
}

// ScanToMapper scans the values from a scanner to a mapper
func ScanToMapper(sc Scanner, m interface{}) error {
	// colNum := sc.ColNum()
	// vals := make([]interface{}, colNum)
	/*
		for i := 0; i < colNum; i++ {
			vals[i] = m.Map(sc.Column(i))
		}
	*/
	ptrs, err := ColumnPtrs(m, sc.Columns())
	if err != nil {
		return err
	}
	return sc.Scan(ptrs...)
}

// NewTestScanner returns a new faking Scanner using the given function to fake the scanning.
// The function is called each time the scanners Scan method is called and the target map is passed to fn.
// If fn returns an error, each further call of Scan will return this error and fn is no longer called.
// If fn returns a stop or an error the Next method of the scanner will return false.
func NewTestScanner(cols []string, fn func(targets map[string]interface{}) (stop bool, err error)) *TestScanner {
	return &TestScanner{
		fn:   fn,
		cols: cols,
	}
}

type TestScanner struct {
	fn   func(targets map[string]interface{}) (stop bool, err error)
	err  error
	stop bool
	cols []string
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

/*
func (f *TestScanner) ColNum() int {
	return len(f.cols)
}
*/

func (f *TestScanner) Columns() []string {
	return f.cols
}

/*
func (f *TestScanner) Column(pos int) string {
	return f.cols[pos]
}
*/

// Scan allows scanning by column name instead of column position
func (f *TestScanner) Scan(vals ...interface{}) error {

	cols := f.Columns()
	m := map[string]interface{}{}

	for i, val := range vals {
		m[cols[i]] = val
	}

	// fmt.Println("scan called")

	if f.err != nil {
		return f.err
	}
	f.stop, f.err = f.fn(m)

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

/*
func (sc *dbScanner) ColNum() int {
	cols, _ := sc.Rows.Columns()
	return len(cols)
}
*/

func (sc *dbScanner) Columns() []string {
	cols, _ := sc.Rows.Columns()
	return cols
}

/*
func (sc *dbScanner) Column(pos int) string {
	cols, _ := sc.Rows.Columns()
	return cols[pos]
}
*/

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
func (sc *dbScanner) Scan(vals ...interface{}) error {
	if sc.err != nil {
		return sc.err
	}
	/*
		vals := make([]interface{}, len(sc.columns))

		for colName, val := range targets {
			i, ok := sc.columns[colName]
			if !ok {
				return errors.New("unknown column " + colName)
			}
			vals[i] = val
		}
	*/
	wr := sqlnull.Wrap(sc.Rows)
	sc.err = wr.Scan(vals...)
	if sc.err != nil && !sc.closed {
		sc.Rows.Close()
		sc.closed = true
	}
	return sc.err
}
