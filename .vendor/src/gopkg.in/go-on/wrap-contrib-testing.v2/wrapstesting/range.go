package wrapstesting

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Next-Range
// Prev-Range

/*
Range bytes=0-499
bytes=-500
bytes=9500-
*/

func writeRangeStatusCode(w http.ResponseWriter, start, end, total int) {
	if total > end-start+1 {
		w.WriteHeader(206)
	}
}

func WriteContentRange(w http.ResponseWriter, start, end, total int) {
	w.Header().Set("Accept-Ranges", "items")
	if end < 1 || total < 1 {
		w.Header().Set("Content-Range", "*")
		w.WriteHeader(200)
		return
	}

	if start == 0 && end == 0 {
		w.Header().Set("Content-Range", fmt.Sprintf("items %d-%d/%d", start, end, total))
		writeRangeStatusCode(w, start, end, total)
		return
	}
	// Content-Range: bytes 21010-47021/47022
	if end-start < total {
		w.Header().Set("Content-Range", fmt.Sprintf("items %d-%d/%d", start, end, total))
		writeRangeStatusCode(w, start, end, total)
		return
	}
}

// Accept: application/vnd.heroku+json; version=3
// Authorization: $TUTORIAL_KEY

type RangeRequest struct {
	AcceptRanges []string
	Max          int    // max number of results
	SortBy       string // key for sorting
	Desc         bool   // if false, the sort order is ascending (default), otherwise descending
	Start        int
	End          int
}

func ParseRangeRequest(rq *http.Request, acceptedKeys ...string) (*RangeRequest, error) {
	r := rq.Header.Get("Range")
	if r == "" {
		return nil, nil
	}

	rr := &RangeRequest{}
	rr.Max = -1
	rr.Start = -1
	rr.End = -1
	rr.AcceptRanges = acceptedKeys

	rangeArr := strings.Split(r, ";")

	if len(rangeArr) > 1 {
		var options = map[string]string{}
		var optStr = rangeArr[1]

		for _, option := range strings.Split(optStr, ",") {
			pos := strings.Index(option, "=")

			if pos == -1 {
				continue
			}

			options[strings.TrimSpace(strings.ToLower(option[0:pos]))] = strings.TrimSpace(strings.ToLower(option[pos+1:]))
		}

		for k, v := range options {

			switch k {
			case "order":
				switch v {
				case "desc":
					rr.Desc = true
				case "asc":
					rr.Desc = false
				default:
					return rr, fmt.Errorf("invalid value for order: %s. supported are asc and desc", v)

				}
			case "max":
				m, err := strconv.Atoi(v)
				if err != nil {
					return rr, fmt.Errorf("max must be an integer")
				}
				rr.Max = m
			case "items":
				pos := strings.Index(v, "-")
				if pos > 0 {
					start := v[:pos]
					end := v[pos+1:]

					s, err := strconv.Atoi(start)
					if err != nil {
						return rr, fmt.Errorf("invalid value for range items; should be specified like this: items=2-5")
					}

					e, err := strconv.Atoi(end)
					if err != nil {
						return rr, fmt.Errorf("invalid value for range items; should be specified like this: items=2-5")
					}

					rr.Start = s
					rr.End = e
				}
			default:
				return rr, fmt.Errorf("unknown range option: %s", k)
			}

		}
	}

	pos := strings.Index(rangeArr[0], " ")
	sortBy := rangeArr[0][:pos]
	var found bool
	for _, sortKey := range rr.AcceptRanges {
		if sortKey == sortBy {
			found = true
			break
		}
	}
	if !found {
		return rr, fmt.Errorf("invalid sort key %s, allowed are %s", sortBy, strings.Join(rr.AcceptRanges, ","))
	}
	rr.SortBy = sortBy
	return rr, nil
}
