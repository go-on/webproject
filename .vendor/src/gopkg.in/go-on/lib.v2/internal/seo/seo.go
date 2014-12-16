package seo

import (
	h "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/types"
)

func DCSubject(content string, other ...interface{}) HTMLer {
	e := h.META(h.Attrs_("content", content, "name", "DC.subject"))
	if len(other) > 0 {
		e.Add(other...)
	}
	return e
}

func DCTitle(content string, other ...interface{}) HTMLer {
	e := h.META(h.Attrs_("content", content, "name", "DC.title"))
	if len(other) > 0 {
		e.Add(other...)
	}
	return e
}

func DCCreator(content string, other ...interface{}) HTMLer {
	e := h.META(h.Attrs_("content", content, "name", "DC.creator"))
	if len(other) > 0 {
		e.Add(other...)
	}
	return e
}

func Keywords(content string, other ...interface{}) HTMLer {
	e := h.META(h.Attrs_("content", content, "name", "keywords"))
	if len(other) > 0 {
		e.Add(other...)
	}
	return e
}

func Description(content string, other ...interface{}) HTMLer {
	e := h.META(h.Attrs_("content", content, "name", "description"))
	if len(other) > 0 {
		e.Add(other...)
	}
	return e
}

func Canonical(url string, other ...interface{}) HTMLer {
	e := h.LINK(h.Attrs_("rel", "canonical", "href", url))
	if len(other) > 0 {
		e.Add(other...)
	}
	return e
}
