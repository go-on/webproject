package rest

import (
	"math"
	"net/http"
	"strconv"
)

func setJsonContentType(wr http.ResponseWriter) {
	wr.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// stolen from https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/ITZV08gAugI
// return rounded version of x with prec precision.
func roundViaFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

func toInt(x float64) int {
	return int(math.Floor(roundViaFloat(x, 0)))
}

func toInt64(x float64) int64 {
	return int64(math.Floor(roundViaFloat(x, 0)))
}
