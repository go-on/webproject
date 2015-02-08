package pgsqlfat

import (
	"testing"

	"gopkg.in/go-on/lib.v3/internal/fat"
)

// . "gopkg.in/metakeule/pgsql.v6"

/*
ty := reflect.TypeOf(østruct).Elem()
	return "*" + ty.PkgPath() + "." + ty.Name()
*/

type coop struct {
	Id        *fat.Field `type:"string uuid"        db:"id UUIDGEN PKEY" rest:" R DL"`
	Name      *fat.Field `type:"string varchar(66)" db:"name"            rest:"CRU L"`
	FoundedAt *fat.Field `type:"time date"          db:"founded_at NULL" rest:"CRU L"`
}

var COOP = fat.Proto(&coop{}).(*coop)
var registry = NewRegistries()

func init() {
	registry.MustRegisterTable("co_op", COOP)
	// MustRegisterTable("co_op", COOP)
}

func TestRegistry(t *testing.T) {
	f := registry.FieldOf(COOP.Name)

	if f == nil {
		t.Error("field should not be nil")
	}

	ta := registry.TableOf(COOP)

	if ta == nil {
		t.Error("table should not be nil")
	}
}
