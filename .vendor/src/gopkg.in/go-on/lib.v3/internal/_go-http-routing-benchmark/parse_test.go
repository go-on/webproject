// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package benchmark

import (
	"net/http"
	"testing"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []route{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var (
	parseGocraftWeb http.Handler
	parseGorillaMux http.Handler
	parseHttpRouter http.Handler
	parseMartini    http.Handler
	parsePat        http.Handler
	parseTigerTonic http.Handler
	parseTraffic    http.Handler
	parseGoon       http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))

	parseGocraftWeb = loadGocraftWeb(parseAPI)
	parseGorillaMux = loadGorillaMux(parseAPI)
	parseHttpRouter = loadHttpRouter(parseAPI)
	parseMartini = loadMartini(parseAPI)
	parsePat = loadPat(parseAPI)
	parseTigerTonic = loadTigerTonic(parseAPI)
	parseTraffic = loadTraffic(parseAPI)
	parseGoon = loadGoon(parseAPI)
}

// Static
func BenchmarkGocraftWeb_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGocraftWeb, req)
}
func BenchmarkGorillaMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkMartini_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPat_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkTigerTonic_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkGoon_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGoon, req)
}

// One Param
func BenchmarkGocraftWeb_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGocraftWeb, req)
}
func BenchmarkGorillaMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkMartini_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPat_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkTigerTonic_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkGoon_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGoon, req)
}

// Two Params
func BenchmarkGocraftWeb_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGocraftWeb, req)
}
func BenchmarkGorillaMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkMartini_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPat_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkTigerTonic_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkGoon_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGoon, req)
}

// All Routes
func BenchmarkGocraftWeb_ParseAll(b *testing.B) {
	benchRoutes(b, parseGocraftWeb, parseAPI)
}
func BenchmarkGorillaMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseGorillaMux, parseAPI)
}
func BenchmarkHttpRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpRouter, parseAPI)
}
func BenchmarkMartini_ParseAll(b *testing.B) {
	benchRoutes(b, parseMartini, parseAPI)
}
func BenchmarkPat_ParseAll(b *testing.B) {
	benchRoutes(b, parsePat, parseAPI)
}
func BenchmarkTigerTonic_ParseAll(b *testing.B) {
	benchRoutes(b, parseTigerTonic, parseAPI)
}
func BenchmarkTraffic_ParseAll(b *testing.B) {
	benchRoutes(b, parseTraffic, parseAPI)
}
func BenchmarkGoon_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoon, parseAPI)
}
