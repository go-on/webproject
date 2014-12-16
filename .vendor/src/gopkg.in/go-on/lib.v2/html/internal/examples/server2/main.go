package main

import (
	"fmt"

	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/html/internal/element/compiler"
	. "gopkg.in/go-on/lib.v2/html/internal/htmlfat"
	"gopkg.in/go-on/lib.v2/types"
	"gopkg.in/go-on/router.v2"
	. "gopkg.in/go-on/router.v2/internal/routerfat"
	"gopkg.in/go-on/router.v2/route"
	"gopkg.in/go-on/wrap.v2"
	"gopkg.in/go-on/wrap-contrib-testing.v2/wrapstesting"
	// . "gopkg.in/go-on/wrap-contrib.v2/wraps"
	"net/http"
	"sync"

	"gopkg.in/go-on/lib.v2/internal/fat"
)

type Person struct {
	http.ResponseWriter
	Id        *fat.Field `type:"int" details:"id"`
	FirstName *fat.Field `type:"string"`
	LastName  *fat.Field `type:"string"`
	Vita      *fat.Field `type:"string html"`
}

var (
	PERSON  = fat.Proto(&Person{}).(*Person)
	id      = 0
	persons = map[string]*Person{}
	_       = NewPerson("P<ete>r", "Pan", Elements("A ", I("short"), " vita").String())
	_       = NewPerson("Paul", "S<imo>n", Elements("A ", I("long"), " vita").String())
	mutex   = &sync.Mutex{}
)

func init() {
	Register(PERSON)
}

func NewPerson(firstname, lastname, vita string) (p *Person) {
	id++
	p = fat.New(PERSON, &Person{}).(*Person)
	p.FirstName.Set(firstname)
	p.LastName.Set(lastname)
	p.Vita.Set(vita)
	p.Id.Set(id)
	persons[fmt.Sprintf("%d", id)] = p
	return
}

func FindPerson(wr http.ResponseWriter, req *http.Request) http.ResponseWriter {
	id := req.FormValue(":id")

	// prevent race condition for persons[]
	mutex.Lock()
	defer mutex.Unlock()
	pp, found := persons[id]

	if found {
		// make a copy of the person
		p := fat.New(PERSON, &Person{}).(*Person)
		p.FirstName.Set(pp.FirstName.String())
		p.LastName.Set(pp.LastName.String())
		p.Vita.Set(pp.Vita.String())
		// and wrap it to the responsewriter
		p.ResponseWriter = wr
		return p
	} else {
		wr.WriteHeader(404)
		fmt.Fprint(wr, "404 not found")
		return nil
	}
}

func (p *Person) NameView(wr http.ResponseWriter, req *http.Request) {
	H1(p.FirstName.String(), " ", p.LastName.String()).WriteTo(wr)
}

func ListView(wr http.ResponseWriter, req *http.Request) {
	for _, pers := range persons {
		LI(
			AHref(
				MustUrl(personDetails, pers, "details"),
				pers.FirstName.String(), " ", pers.LastName.String(),
			),
		).WriteTo(wr)
	}
}

func (p *Person) VitaView(wr http.ResponseWriter, req *http.Request) {
	DIV(
		types.Class("vita"),
		H2("Vita"),
		P(
			types.HTMLString(p.Vita.String()),
		),
	).WriteTo(wr)
}

var (
	yellow        = types.Style{"background-color", "yellow"}
	green         = types.Style{"background-color", "green"}
	personDetails *route.Route
	personList    *route.Route
)

func listLink(rw http.ResponseWriter, req *http.Request) {
	DIV(AHref(personList.MustURL(), "back to list")).WriteTo(rw)
}

func main() {
	personRouter := router.New()
	personDetails = personRouter.GET("/:id",
		wrap.New(
			wrapstesting.Context(FindPerson),
			wrap.Handler(
				Elements(
					listLink,
					DIV(yellow, wrapstesting.HandlerMethod((*Person).NameView)),
					DIV(green, wrapstesting.HandlerMethod((*Person).VitaView)),
				),
			),
		),
	)

	personList = personRouter.GET("/", UL(ListView))

	personRouter.Mount("/person", http.DefaultServeMux)

	handler := compiler.DocHandler(HTML5(HTML(BODY(personRouter))))

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		panic(err.Error())
	}
}
