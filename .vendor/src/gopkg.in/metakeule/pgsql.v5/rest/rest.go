package rest

import (
	"gopkg.in/go-on/router.v2/route"
	. "gopkg.in/metakeule/pgsql.v5"
	"gopkg.in/metakeule/pgsql.v5/pgsqlfat"

	"gopkg.in/go-on/router.v2"
)

type rest struct {
	db     DB
	Router *router.Router
	*pgsqlfat.Registry
}

func NewREST(db DB, reg *pgsqlfat.Registry, rt *router.Router) *rest {
	return &rest{db, rt, reg}
}

type action int

const (
	_             = iota
	CREATE action = 1 << iota
	READ
	UPDATE
	DELETE
	LIST
)

var ALL = CREATE | READ | UPDATE | DELETE | LIST

func (r *rest) Mount(proto interface{}, mountPoint string, actions action, options *options) (routes map[action]*route.Route) {
	mounter := NewCRUD(r.Registry, proto).Mount(r.db, r.Router, mountPoint, options)

	routes = map[action]*route.Route{}

	if has(actions, CREATE) {
		routes[CREATE] = mounter.CreateRoute()
	}

	if has(actions, READ) {
		routes[READ] = mounter.ReadRoute()
	}

	if has(actions, UPDATE) {
		routes[UPDATE] = mounter.UpdateRoute()
	}

	if has(actions, DELETE) {
		routes[DELETE] = mounter.DeleteRoute()
	}

	if has(actions, LIST) {
		routes[LIST] = mounter.ListRoute()
	}
	return
}

func has(what action, has action) bool {
	return what&has != 0
}
