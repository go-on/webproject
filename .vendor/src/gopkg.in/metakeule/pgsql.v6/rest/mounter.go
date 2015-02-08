package rest

import (
	"encoding/json"
	"fmt"
	"gopkg.in/go-on/builtin.v1/db"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"

	"gopkg.in/go-on/router.v2/route"

	. "gopkg.in/metakeule/pgsql.v6"
	"gopkg.in/go-on/lib.v3/internal/fat"
	"gopkg.in/go-on/router.v2"
	"gopkg.in/go-on/wrap-contrib-testing.v2/wrapstesting"
)

type Mounter struct {
	CRUD        *CRUD
	db          db.DB
	mountPoint  string
	placeholder string
	Router      *router.Router
	options     *options
}

type options struct {
	maxLimit int
	sortKeys []string
}

func MaxLimit(m int) (o *options) {
	o = &options{}
	o.SetMaxLimit(m)
	return
}

func SortFields(sortFields ...*fat.Field) (o *options) {
	o = &options{}
	o.SetSortFields(sortFields...)
	return
}

func (o *options) SetMaxLimit(m int) *options {
	o.maxLimit = m
	return o
}

func (o *options) SetSortFields(sortFields ...*fat.Field) *options {
	o.sortKeys = []string{}
	for _, sF := range sortFields {
		o.sortKeys = append(o.sortKeys, sF.Name())
	}
	return o
}

// default maxLimit is 100, may be changed with SetListOptions
func Mount(r *CRUD, d db.DB, rt *router.Router, mountPoint string, options *options) *Mounter {
	if options == nil {
		options = MaxLimit(100)
	}
	return &Mounter{
		CRUD:        r,
		db:          d,
		mountPoint:  mountPoint,
		placeholder: mountPoint + "_id",
		Router:      rt,
		options:     options,
	}
}

func (m *Mounter) itemPath() string {
	return path.Join(m.mountPoint, ":"+m.placeholder)
}

//func (m *Mounter) getId(vars *router.Vars) string {
func (m *Mounter) getId(rq *http.Request) string {
	return router.GetRouteParam(rq, m.placeholder)
	//return rq.FormValue(":" + m.placeholder)
	// return vars.Get(m.placeholder)
}

func (m *Mounter) CreateRoute() *route.Route {
	return m.Router.POST("/"+m.mountPoint, wrapstesting.HandlerMethod(m.serveCreate))
}

func (m *Mounter) ListRoute() *route.Route {
	return m.Router.GET("/"+m.mountPoint, wrapstesting.HandlerMethod(m.serveList))
}

func (m *Mounter) UpdateRoute() *route.Route {
	return m.Router.PATCH("/"+m.itemPath(), wrapstesting.HandlerMethod(m.serveUpdate))
}

func (m *Mounter) DeleteRoute() *route.Route {
	return m.Router.DELETE("/"+m.itemPath(), wrapstesting.HandlerMethod(m.serveDelete))
}

func (m *Mounter) ReadRoute() *route.Route {
	return m.Router.GET("/"+m.itemPath(), wrapstesting.HandlerMethod(m.serveRead))
}

func (m *Mounter) serveCreate(wr http.ResponseWriter, rq *http.Request) {
	json_, err := ioutil.ReadAll(rq.Body)

	if err != nil {
		wr.WriteHeader(400)
		return
	}

	var id string

	validation := strings.TrimSpace(rq.Header.Get("X-Validation"))
	onlyValidation := false

	switch validation {
	case "*":
		_, err = m.CRUD.Create(m.db, json_, true, "")
		onlyValidation = true
	case "":
		id, err = m.CRUD.Create(m.db, json_, false, "")
	default:
		onlyValidation = true
		_, err = m.CRUD.Create(m.db, json_, true, validation)
		// err = m.validateField(validation, string(json_), wr)
	}
	//id, err = m.CRUD.Create(m.db, json_)

	// TODO: respect validationError and handle it differently
	if err != nil {
		serveError(err, wr, rq)
		return
	}

	if onlyValidation {
		serveSuccess(wr, rq)
		return
	}

	route := m.Router.Route("/" + m.itemPath())
	if route != nil {
		wr.Header().Set("Location", route.MustURL(m.placeholder, id))
	}
	okCreated(id).ServeHTTP(wr, rq)
}

func (m *Mounter) serveUpdate(wr http.ResponseWriter, rq *http.Request) {
	id := m.getId(rq)
	body, err := ioutil.ReadAll(rq.Body)

	if id == "" || err != nil {
		wr.WriteHeader(400)
		return
	}

	if err == nil {
		// TODO: do the same for create
		// TODO: translate validation errors somehow
		validation := strings.TrimSpace(rq.Header.Get("X-Validation"))

		switch validation {
		case "*":
			err = m.CRUD.Update(m.db, id, body, true, "")
		case "":
			err = m.CRUD.Update(m.db, id, body, false, "")
		default:
			err = m.CRUD.Update(m.db, id, body, true, validation)
			// err = m.validateField(validation, string(body), wr)
		}
	}

	// TODO: check if err is a validationError and handle it differently
	if err != nil {
		serveError(err, wr, rq)
		return
	}
	serveSuccess(wr, rq)
}

func (m *Mounter) serveDelete(wr http.ResponseWriter, rq *http.Request) {
	id := m.getId(rq)

	if id == "" {
		wr.WriteHeader(400)
		return
	}

	err := m.CRUD.Delete(m.db, id)

	if err != nil {
		serveError(err, wr, rq)
		return
	}

	serveSuccess(wr, rq)
}

func (m *Mounter) serveRead(wr http.ResponseWriter, rq *http.Request) {

	// func (m *Mounter) serveRead(wr http.ResponseWriter, rq *http.Request) {
	//	var vars = &router.Vars{}

	//	wrapstesting.MustUnWrap(wr, &vars)

	//	fmt.Println("serving read")
	id := m.getId(rq)

	if id == "" {
		//	fmt.Println("id empty")
		wr.WriteHeader(400)
		return
	}

	// TODO Read should return not found if item could not be found
	obj, err := m.CRUD.Read(m.db, id)

	var j []byte
	if err == nil {
		j, err = json.MarshalIndent(obj, "", "  ")
	}

	if err != nil {
		serveError(err, wr, rq)
		return
	}

	setJsonContentType(wr)

	//	if rq.Method == "HEAD" {
	wr.Header().Set("Content-Length", strconv.Itoa(len(j)))
	//	}
	wr.Write(j)
}

func (m *Mounter) serveList(wr http.ResponseWriter, rq *http.Request) {

	rangeReq, err := wrapstesting.ParseRangeRequest(rq, m.options.sortKeys...)

	if err != nil {
		wr.WriteHeader(400)
		fmt.Fprintln(wr, err)
		return
	}

	var total int
	var objs []map[string]interface{}
	var start int
	var reqLimit = m.options.maxLimit
	var sortBy string

	dir := ASC

	if err == nil {
		if rangeReq != nil {
			if rangeReq.End > -1 {
				requestedLimit := rangeReq.End - rangeReq.Start + 1
				if requestedLimit < m.options.maxLimit {
					reqLimit = requestedLimit
				}
			}

			if rangeReq.Max > -1 && rangeReq.Max < reqLimit {
				reqLimit = rangeReq.Max
			}

			if rangeReq.SortBy != "" {
				sortBy = rangeReq.SortBy
			}

			if rangeReq.Desc {
				dir = DESC
			}

			if rangeReq.Start > -1 {
				start = rangeReq.Start
			}
		}

		f := m.CRUD.primaryKey
		if sortBy != "" {
			f = m.CRUD.Field(m.CRUD.typeString(), sortBy)
		}
		total, objs, err = m.CRUD.List(m.db, reqLimit, dir, f, start)
	}

	var j []byte

	if err == nil {
		j, err = json.MarshalIndent(objs, "", "  ")
	}

	if err != nil {
		ErrServer.ServeHTTP(wr, rq)
		return
	}

	setJsonContentType(wr)

	// 416 Requested Range Not Satisfiable
	ifrange := rq.Header.Get("If-Range")

	if rangeReq != nil && len(objs) == 0 && ifrange == "" {
		wr.WriteHeader(416)
		wr.Write([]byte(`[]`))
		return
	}

	wrapstesting.WriteContentRange(wr, start, start+len(objs)-1, total)
	// if rq.Method == "HEAD" {
	wr.Header().Set("Content-Length", strconv.Itoa(len(j)))
	// }
	wr.Write(j)
}
