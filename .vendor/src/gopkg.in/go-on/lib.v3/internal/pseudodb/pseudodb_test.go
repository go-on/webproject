package pseudodb

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"gopkg.in/go-on/router.v2"
)

type person struct {
	Id                  string
	FirstName, LastName string
	Age                 int
}

func (p *person) SetUUID(id string) {
	p.Id = id
}

func (p *person) UUID() string {
	return p.Id
}

type company struct {
	Id   string
	Name string
}

func (c *company) SetUUID(id string) {
	c.Id = id
}

func (c *company) UUID() string {
	return c.Id
}

func TestTransformType(t *testing.T) {
	got := transformType(reflect.TypeOf(&person{}))
	expected := "person"
	if got != expected {
		t.Errorf("got: %#v, expected: %#v", got, expected)
	}
}

func TestGet(t *testing.T) {
	peter := &person{"", "Peter", "Pan", 42}
	app := NewApp(nil, &person{})

	app.Data["peter"] = peter

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/person/peter", nil)

	rt := router.New()
	app.RegisterRoutes(rt)
	rt.Mount("/", nil)
	// router.Mount("/v2", rt)

	rt.ServeHTTP(rec, req)

	peterJson, _ := json.Marshal(peter)

	if rec.Body.String() != string(peterJson) {
		t.Errorf("wrong  get serialization: expecting %s got %s", peterJson, rec.Body.String())
	}

}

func TestIndex(t *testing.T) {
	peter := &person{"", "Peter", "Pan", 42}
	petra := &person{"", "Petra", "Kelly", 102}
	google := &company{"", "google"}
	app := NewApp(nil, &person{}, &company{})

	app.Data["peter"] = peter
	app.Data["petra"] = petra
	app.Data["google"] = google

	// rt := router.New()

	// app.Mount(rt, "/api")

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/person/", nil)

	rt := router.New()
	app.RegisterRoutes(rt)
	rt.Mount("/", nil)
	// router.Mount("/v1", rt)

	rt.ServeHTTP(rec, req)

	persons := []person{}
	// println(rec.Body.String())
	err := json.Unmarshal(rec.Body.Bytes(), &persons)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(persons) != 2 {
		t.Errorf("should be 2 persons, but are: %d", len(persons))
	}

	pete := persons[0]

	if pete.FirstName != "Peter" {
		t.Errorf("wrong  persons['peter'].FirstName: expecting %s got %s", "Peter", pete.FirstName)
	}

	petr := persons[1]

	if petr.FirstName != "Petra" {
		t.Errorf("wrong  persons['petra'].FirstName: expecting %s got %s", "Petra", petr.FirstName)
	}

}

func TestPost(t *testing.T) {
	peter := &person{"", "Peter", "Pan", 42}
	app := NewApp(nil, &person{})

	// rt := router.New()

	// app.Mount(rt, "/api")

	peterJson, _ := json.Marshal(peter)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/person/", bytes.NewReader(peterJson))

	rt := router.New()
	app.RegisterRoutes(rt)
	rt.Mount("/", nil)
	// router.Mount("/v3", rt)

	rt.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("status should be %d, but is %d", http.StatusCreated, rec.Code)
	}

	uuid := rec.Body.String()

	pete, has := app.Data[uuid]

	if !has {
		t.Errorf("not saved: %#v", uuid)
	}

	if pete.(*person).FirstName != peter.FirstName {
		t.Errorf("wrong firstname: expected %#v, got: %#v", peter.FirstName, pete.(*person).FirstName)
	}

	if pete.(*person).Id != uuid {
		t.Errorf("id not set")
	}

}

func TestDelete(t *testing.T) {
	peter := &person{"peter", "Peter", "Pan", 42}
	app := NewApp(nil, &person{})

	app.Data["peter"] = peter

	// rt := router.New()

	// app.Mount(rt, "/api")

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/person/peter", nil)

	// router.Mount("/v4", rt)
	rt := router.New()
	app.RegisterRoutes(rt)
	rt.Mount("/", nil)

	rt.ServeHTTP(rec, req)

	if len(app.Data) > 0 {
		t.Errorf("not deleted")
	}

}

func TestPatch(t *testing.T) {
	peter := &person{"", "Peter", "Pan", 42}
	app := NewApp(nil, &person{})

	app.Data["peter"] = peter

	// rt := router.New()

	// app.Mount(rt, "/api")

	pete := &person{"", "Peter", "Pan", 43}

	peterJson, _ := json.Marshal(pete)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/person/peter", bytes.NewReader(peterJson))

	// router.Mount("/v5", rt)
	rt := router.New()
	app.RegisterRoutes(rt)
	rt.Mount("/", nil)

	rt.ServeHTTP(rec, req)

	if app.Data["peter"].(*person).Age != 43 {
		t.Errorf(" not updated")
	}

}

func TestFile(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "pseudodbtest-")
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	file.Close()

	peter := &person{"", "Peter", "Pan", 42}
	app := NewApp(NewFileStore(file.Name()), &person{})

	// rt := router.New()

	// app.Mount(rt, "/api")

	peterJson, _ := json.Marshal(peter)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/person/", bytes.NewReader(peterJson))

	// router.Mount("/v6", rt)
	rt := router.New()
	app.RegisterRoutes(rt)
	rt.Mount("/", nil)
	rt.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("status should be %d, but is %d", http.StatusCreated, rec.Code)
	}

	// fmt.Printf("%s\n", d)

	var app2 = NewApp(NewFileStore(file.Name()), &person{})
	err = app2.Load()

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	os.Remove(file.Name())

	if len(app2.Data) != 1 {
		t.Errorf("reloaded app should have 1 datum, but has: %d", len(app2.Data))
	}
}
