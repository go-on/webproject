package main

import (
	. "gopkg.in/metakeule/pgsql.v5"
	"time"
)

type Artist struct {
	Id            string     `db.select: "all,details"`
	FirstName     string     `db.select: "all,details"                db.insert: "galleryartist,other"`
	LastName      string     `db.select: "all,details,galleryartists" db.insert: "galleryartist,other"`
	GalleryArtist bool       `db.select: "all,details"                db.insert: "galleryartist"`
	Age           int        `db.select: "    details"                db.insert: "galleryartist"        db.update: "set.birthday"`
	DateOfBirth   *time.Time `db.select: "all,details,galleryartists" db.insert: "galleryartist"        db.update: "set.birthday"`
}

func All() (all []*Artist, err error) {
	all = [20]*Artist{}
	err = SelectByStructs("all", &all)
	return
}

func GalleryArtists() (ga []*Artist, err error) {
	ga = [20]*Artist{}
	err = SelectByStructs("galleryartists", &ga, Where(Equals(artist.GalleryArtists, true)))
	return
}

func (ø *Artist) forMe() *WhereStruct { return Where(Equals(artist.Id, ø.Id)) }

func (ø *Artist) Details() (err error) { return SelectByStruct(ø, ø.forMe()) }

func (ø *Artist) AddAsGalleryArtist() error {
	ø.GalleryArtist = true
	return InsertByStruct("galleryartist", ø)
}

func (ø *Artist) AddAsOther() error {
	ø.GalleryArtist = false
	return InsertByStruct("other", ø)
}

func AddGalleryArtists(ga ...*Artist) {
	for _, a := range ga {
		a.GalleryArtist = true
	}
	return InsertByStructs("galleryartist", a...)
}

func (ø *Artist) SetBirthday() error { return UpdateByStruct("set.birthday", ø, ø.forMe()) }
