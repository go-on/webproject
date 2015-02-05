package wsi

import (
	"database/sql"
	"net/http"
	"net/url"
	"strconv"
)

type Encoder func(http.ResponseWriter) (StreamEncoder, error)

type Query struct {
	encFn        Encoder
	mapperFn     RessourceFunc
	fn           QueryFunc
	errorHandler func(*http.Request, error)
}

type QueryOptions struct {
	OrderBy []string
	Limit   int
	Offset  int
	Filter  map[string]string
}

func (wq Query) SetEncoder(e Encoder) Query {
	wq.encFn = e
	return wq
}

func (wq Query) SetErrorCallback(fn func(*http.Request, error)) Query {
	wq.errorHandler = fn
	return wq
}

func (wq Query) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	scanner, err := QueryByRequest(w, r, wq.fn)
	// if we got an error here, the status code has already be written
	if err != nil {
		if wq.errorHandler != nil {
			wq.errorHandler(r, err)
		}
		return
	}

	// we could not construct the scanner properly. fail early.
	err = scanner.Error()
	if err != nil {
		if wq.errorHandler != nil {
			wq.errorHandler(r, err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var enc StreamEncoder
	enc, err = wq.encFn(w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if wq.errorHandler != nil {
			wq.errorHandler(r, err)
		}
		return
	}

	defer enc.Finish()

	// colNum := scanner.ColNum()

	for scanner.Next() {
		mapper := wq.mapperFn()

		err = ScanToMapper(scanner, mapper)
		/*
			cols := make([]interface{}, colNum)
			for i := 0; i < colNum; i++ {
				cols[i] = mapper.Map(scanner.Column(i))
			}
		*/
		// m := map[string]interface{}{}
		// mapper.MapColumns(m)
		// err = scanner.Scan(cols...)

		// we already wrote something to the body, so handle errors gracefully
		if err != nil && wq.errorHandler != nil {
			wq.errorHandler(r, err)
			return
		}

		// we already wrote something to the body, so handle errors gracefully
		err = enc.Encode(mapper)
		if err != nil && wq.errorHandler != nil {
			wq.errorHandler(r, err)
			return
		}
	}
}

// QueryByRequest returns a Scanner with the help of the given search function and parametrized by the given request.
// It does so by using the url query values for the keys "offset", "limit" and "sort", for further information see ScanQueryValues
// If any error happens before scanning, a http.StatusInternalServerError will be written to the ResponseWriter
// and the first call to Next() fails. The error than can retrieved via the Error method of the scanner
func QueryByRequest(w http.ResponseWriter, r *http.Request, fn QueryFunc) (scanner Scanner, err error) {
	options := ScanQueryValues(r.URL.Query())
	options.Filter = map[string]string{}

	for key, _ := range r.URL.Query() {
		options.Filter[key] = r.URL.Query().Get(key)
	}
	/*
		scanner, err = fn(w, r, options)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			scanner = errScanner{err}
		}
		return
	*/
	return fn(options, w, r)
}

// ScanQueryValues scans the query values "offset", "limit" and "sort" out of the given url.Values.
//   offset - if set - must be convertible to an int, defaults to 0 (=no skipping)
//   limit - if set - must be convertible to an int, defaults to maxLimit
//   sort must be in the form "+col" or "-col" where "+col" results in ascending sort of the col and -col results in descending sorting.
//        Multiple query values for sort resulting in mutliple sorts in the order of the values
func ScanQueryValues(values url.Values) (options QueryOptions) {
	options.Offset, _ = strconv.Atoi(values.Get("offset"))
	options.Limit, _ = strconv.Atoi(values.Get("limit"))

	if options.Limit < 0 {
		options.Limit = 0
	}

	// offset must be >= 0
	if options.Offset < 0 {
		options.Offset = 0
	}

	options.OrderBy = convertSorts(values["sort"])
	return
}

// DBQuery returns a Scanner for the given query that allowes iteration over the returned rows from
// the underlying sql query. The Scanner takes of closing the rows
func DBQuery(db *sql.DB, query string, values ...interface{}) (sc Scanner, err error) {
	var (
		rows *sql.Rows
		cols []string
	)

	rows, err = db.Query(query, values...)
	if err != nil {
		return
	}

	cols, err = rows.Columns()
	if err != nil {
		return
	}

	columns := make(map[string]int, len(cols))

	for i, col := range cols {
		columns[col] = i
	}

	sc = &dbScanner{columns: columns, Rows: rows}
	return
}

// convertSorts "+name" or "name" to "name ASC" and "-name" to "name DESC"
func convertSorts(in []string) (out []string) {
	out = make([]string, len(in))

	for i, s := range in {
		switch s[0] {
		case '-':
			out[i] = s[1:] + " DESC"
		default:
			out[i] = s + " ASC"
		}
	}
	return
}
