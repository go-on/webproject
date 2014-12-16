package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	. "gopkg.in/metakeule/pgsql.v5"
	"os"
	// "reflect"
)

func Connect(str string) *sql.DB {
	cs, ſ := pq.ParseURL(str)
	if ſ != nil {
		panic(ſ)
	}

	db, ſ := sql.Open("postgres", cs)
	if ſ != nil {
		panic(ſ)
	}
	//db.SetMaxIdleConns(0)
	// db.Debug = true
	return db
}

func init() {
	Db = Connect(os.Getenv("PG_URL"))

}

var Db DB
var TABLE = NewTable("koelnart-artist")
var Areas = []string{"painting", "object", "foto", "modern", "contemporary"}
var Id = TABLE.NewField("id", UuidType, PrimaryKey, UuidGenerate).SetQueryField("Id")
var FirstName = TABLE.NewField("firstname", VarChar(255)).SetQueryField("FirstName")
var LastName = TABLE.NewField("lastname", VarChar(255)).SetQueryField("LastName")
var Vita = TABLE.NewField("vita", HtmlType, NullAllowed).SetQueryField("Vita")
var GalleryArtist = TABLE.NewField("galleryartist", BoolType).SetQueryField("GalleryArtist")
var Area = TABLE.NewField("area", StringsType).SetQueryField("Area")

type Artist struct {
	Id            string   `db.select:",all,details,"`
	FirstName     string   `db.select:",all,details,"                db.insert:",galleryartist,other,"`
	LastName      string   `db.select:",all,details,galleryartists," db.insert:",galleryartist,other,"`
	GalleryArtist bool     `db.select:",all,details,"                db.insert:",galleryartist,"`
	Vita          string   `db.select:",all,details,"                db.insert:",galleryartist,"        db.update:",set.birthday,"`
	Area          []string `db.select:",all,details,galleryartists," db.insert:",galleryartist,"        db.update:",set.birthday,"`
}

func All(db DB, all interface{}, options ...interface{}) error {
	return NewRow(db, TABLE).SelectByStructs(all, ",all,", options...)
}

func main() {
	all := make([]Artist, 3)
	//all := []Artist{}
	err := All(Db, all, Where(Equals(GalleryArtist, true)))
	if err != nil {
		fmt.Println(err)
	}
	for _, a := range all {
		fmt.Printf("%v\n", a)
	}

	art := &Artist{}
	err = NewRow(Db, TABLE).SelectByStruct(art, ",all,", Where(Equals(FirstName, "Nadja")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n%v\n", art)
}
