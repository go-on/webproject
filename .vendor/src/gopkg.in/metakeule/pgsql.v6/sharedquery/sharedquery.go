/*
The package shared_query allows different queries from different goroutines
to be executed in one single union statement. the endresult is distributed
to the original querier, so that they only get what they asked for.

There are several preconditions that have to be met in order for this to work:
- only queries that affect a single table should be used
- only queries that affect the same table will end up in the same union
- there is a time window in which queries are collected and after that they are
  executed. there is also a timeout for each comined query. the time a querier
  waits for his results would be approximately

    (collectingTime + executingTime) * numberQuries / numberGoroutines

- the joins would happen in the querier

The most useful scenario is to combine details - information queries that refer
to a single id. So if we have 30 concurrent users that need to see their profile
instead of 30 queries for a single id there would be 1 query for 30 ids.

The least useful scenario is for a fulltext search with different returning
columns and sorting orders.

an interesting way would be to add an addional colum that has the id of the
query so that it would be easy to filter out the queried results

*/
package sharedquery

import (
	"database/sql"
	"fmt"
	"gopkg.in/metakeule/pgsql.v6"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

type (
	QueryManager struct {
		CollectingTime           time.Duration
		Timeout                  time.Duration
		NumParallelQueries       int
		NumParallelQueryRequests int
		Store                    *queryStore
		QueryChannel             chan *Query
		Db                       *sql.DB
	}

	Row struct {
		counter    int
		Id         string
		Errors     []error
		RowScanner *RowScanner
	}

	queryStore struct {
		*sync.RWMutex
		Queries map[*pgsql.Table][]*Query
	}

	Query struct {
		Id          string // will be filled by runtime caller combined with unix timestamp
		SelectQuery *pgsql.SelectQuery
		Table       *pgsql.Table
		Response    chan *Row
	}

	RowScanner struct {
		*sync.RWMutex
		Fields []*pgsql.Field
		Rows   []*pgsql.Row
		Result []map[string]interface{}
	}
)

func New(db *sql.DB, numParallelQueryRequests int, numParallelQueries int, collectingTime time.Duration, timeout time.Duration) (ø *QueryManager) {
	ø = &QueryManager{
		Db:                 db,
		CollectingTime:     collectingTime,
		Timeout:            timeout,
		NumParallelQueries: numParallelQueries,
		Store:              &queryStore{&sync.RWMutex{}, map[*pgsql.Table][]*Query{}},
		QueryChannel:       make(chan *Query, numParallelQueryRequests),
	}
	return
}

func NewRowScanner(fields ...*pgsql.Field) (ø *RowScanner) {
	ø = &RowScanner{
		RWMutex: &sync.RWMutex{},
		Fields:  fields,
		Rows:    []*pgsql.Row{},
		Result:  []map[string]interface{}{},
	}
	return
}

func NewRow(scanner *RowScanner, id string) (ø *Row) {
	return &Row{
		Id:         id,
		RowScanner: scanner,
	}
}

func (ø *Row) Next() (res *pgsql.Row, next bool) {
	ø.RowScanner.RLock()
	defer ø.RowScanner.RUnlock()
	if ø.counter >= len(ø.RowScanner.Rows) {
		// fmt.Println("no more results")
		res = nil
		next = false
		return
	}
	r := ø.RowScanner.Rows[ø.counter]
	ø.counter++

	m := r.AsStrings()

	// fmt.Printf("try %s for %s\n", r["sharedqueryid"], id)

	if m["sharedqueryid"] == ø.Id {
		// fmt.Println("ok")
		res = r
		next = true
	} else {
		// fmt.Println("NO WAY")
		res = nil
		next = true
	}
	return
}

func (ø *Row) Next2() (res map[string]interface{}, next bool) {
	ø.RowScanner.RLock()
	defer ø.RowScanner.RUnlock()
	if ø.counter >= len(ø.RowScanner.Result) {
		// fmt.Println("no more results")
		res = nil
		next = false
		return
	}
	r := ø.RowScanner.Result[ø.counter]
	ø.counter++

	var id interface{}
	id = ø.Id

	// fmt.Printf("try %s for %s\n", r["sharedqueryid"], id)

	if r["sharedqueryid"] == id {
		// fmt.Println("ok")
		res = r
		next = true
	} else {
		// fmt.Println("NO WAY")
		res = nil
		next = true
	}
	return
}

func (ø *Row) HasErrors() bool {
	return len(ø.Errors) > 0
}

func (ø *RowScanner) Scan(row *sql.Rows) error {
	//fmt.Println("scanning...")
	vals := []interface{}{}
	for _, fi := range ø.Fields {
		// we need a default value for each db type
		// default_ := fi.Type.Default()
		// but we need to pass a pointer to it
		d := fi.Type.Default()
		//vals = append(vals, reflect.ValueOf(d).Elem().Addr().Interface())
		vals = append(vals, &d)
	}
	idString := ""
	vals = append(vals, &idString)
	err := row.Scan(vals...)
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	for i, fi := range ø.Fields {
		m[fi.Name] = reflect.ValueOf(vals[i]).Elem().Interface()
	}
	m["sharedqueryid"] = reflect.ValueOf(vals[len(vals)-1]).Elem().Interface()
	// fmt.Println(m)
	ø.Result = append(ø.Result, m)
	return nil
}

func (ø *RowScanner) ScanAll(db *sql.DB, row *sql.Rows) (errs []error) {
	errs = []error{}
	fi := []interface{}{}
	for _, f := range ø.Fields {
		fi = append(fi, f)
	}
	as := pgsql.As(pgsql.Sql(`''`), "sharedqueryid", pgsql.TextType)
	fi = append(fi, as)
	for row.Next() {
		r := pgsql.NewRow(db, ø.Fields[0].Table)
		err := r.Scan(row, fi...)
		if err == nil {
			ø.Rows = append(ø.Rows, r)
		} else {
			errs = append(errs, err)
		}
	}
	return
	//func (ø *Row) Scan(row *sql.Rows, fields ...interface{}) (err error) {
}

func (ø *RowScanner) ScanAll2(row *sql.Rows) (errs []error) {
	errs = []error{}
	for row.Next() {
		e := ø.Scan(row)
		if e != nil {
			errs = append(errs, e)
		}
	}
	return
}

func (ø *QueryManager) CheckQuery(query *pgsql.SelectQuery) error {
	if len(query.Fields) < 1 {
		return fmt.Errorf("no fields in query")
	}
	if len(query.Joins) > 0 {
		return fmt.Errorf("joins are currently not supported for shared queries")
	}
	if len(query.FieldsWithAlias) > 0 {
		return fmt.Errorf("fields with alias are currently not supported for shared queries")
	}
	var table = query.Fields[0].Table
	for _, field := range query.Fields {
		if field.Table != table {
			return fmt.Errorf("fields from different tables are currently not supported. Found %s and %s", table, field.Table)
		}
	}

	return nil
}

func (ø *QueryManager) Run() {
	for {
		select {
		case q := <-ø.QueryChannel:
			ø.Store.Lock()
			tf, ok := ø.Store.Queries[q.Table]
			if !ok {
				tf = []*Query{}
			}
			ø.Store.Queries[q.Table] = append(tf, q)
			ø.Store.Unlock()
		case <-time.After(ø.CollectingTime):
			ø.ExecAllQueries()
		}
	}
}

func (ø *QueryManager) ExecAllQueries() {
	ø.Store.Lock()
	for t, _ := range ø.Store.Queries {
		ø.ExecQueriesForTable(t)
	}
	ø.Store.Queries = map[*pgsql.Table][]*Query{}
	ø.Store.Unlock()
}

func (ø *QueryManager) ExecQueriesForTable(table *pgsql.Table) {
	queries := ø.Store.Queries[table]
	s, scanner := ø.Union(queries...)
	rows, err := ø.Db.Query(s.Sql().String())
	if err != nil {
		fmt.Println("errors ", err)
		for _, q := range queries {
			q.Response <- &Row{Errors: []error{err}}
		}
	}
	errs := scanner.ScanAll(ø.Db, rows)
	if len(errs) > 0 {
		fmt.Println("errors: ", errs)
		if err == nil {
			for _, q := range queries {
				q.Response <- &Row{Errors: errs}
			}
		}
		return
	}

	for _, q := range queries {
		rw := NewRow(scanner, q.Id)
		fmt.Println("sending row ", rw)
		q.Response <- rw
	}
}

func (ø *QueryManager) Query(query *pgsql.SelectQuery) (result *Row, err error) {
	if err = ø.CheckQuery(query); err != nil {
		return nil, err
	}

	_, file, line, _ := runtime.Caller(1)
	q := &Query{}
	fs := strings.Split(file, "/")
	q.Id = fmt.Sprintf("%s-%v-%v", fs[len(fs)-1], line, time.Now().UnixNano())
	q.SelectQuery = query
	resp := make(chan *Row, 1)
	q.Response = resp
	q.Table = query.Fields[0].Table
	ø.QueryChannel <- q
	result = <-resp
	return
}

func (ø *QueryManager) Union(queries ...*Query) (query pgsql.Sqler, scanner *RowScanner) {
	fieldMap := map[*pgsql.Field]map[string]int{}
	for _, q := range queries {
		for pos, f := range q.SelectQuery.Fields {
			fi, ok := fieldMap[f]
			if !ok {
				fi = map[string]int{}
			}
			fi[q.Id] = pos
			fieldMap[f] = fi
		}
	}

	orderFields := []*pgsql.Field{}

	for k, _ := range fieldMap {
		orderFields = append(orderFields, k)
	}

	hasField := func(qu *Query, fi *pgsql.Field) bool {
		for _, f := range qu.SelectQuery.Fields {
			if f == fi {
				return true
			}
		}
		return false
	}

	union := []string{}

	// TODO: implement a "Scanner" for the query, i.e. a map[string]interface for a row
	// into which the result is scanned. the scanner must respect the order of the columns queried
	// and the additional alias field "sharedqueryid"

	for _, q := range queries {
		options := []interface{}{orderFields[0].Table}
		for _, of := range orderFields {
			if hasField(q, of) {
				options = append(options, of)
			} else {
				options = append(options, pgsql.As(pgsql.Sqlf(`''`), of.Name, of.Type))
			}
		}
		options = append(options, pgsql.As(pgsql.Sqlf(`'%s'::text`, q.Id), "sharedqueryid", pgsql.TextType))
		options = append(options, q.SelectQuery.Distinct)
		options = append(options, q.SelectQuery.Limit)
		options = append(options, q.SelectQuery.Where)
		options = append(options, q.SelectQuery.OrderBy)
		union = append(union, pgsql.Select(options...).Sql().String())
	}
	scanner = NewRowScanner(orderFields...)
	query = pgsql.Sql("(\n" + strings.Join(union, "\n) UNION (\n") + "\n)")
	fmt.Println(query.Sql().String())
	return
}
