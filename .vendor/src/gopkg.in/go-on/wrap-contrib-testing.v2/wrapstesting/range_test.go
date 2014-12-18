package wrapstesting

import (
	"net/http"
	"testing"
	// . "launchpad.net/gocheck"
)

func TestParseRangeRequest(t *testing.T) {
	rq, _ := http.NewRequest("GET", "/", nil)
	// Range: name ..; order=desc,max=10;
	rq.Header.Set("Range", "name ..;items=1-499,order=desc,max=10;")

	rangeRequest, err := ParseRangeRequest(rq, "name")

	if err != nil {
		t.Error(err)
	}

	if rangeRequest.Max != 10 {
		t.Errorf("rangeRequest.Max = %d // expected 10", rangeRequest.Max)
	}

	if rangeRequest.SortBy != "name" {
		t.Errorf("rangeRequest.SortBy = %#v // expected \"name\"", rangeRequest.SortBy)
	}

	if !rangeRequest.Desc {
		t.Errorf("rangeRequest.Desc = false // expected true")
	}

	if rangeRequest.Start != 1 {
		t.Errorf("rangeRequest.Start = %d // expected 1", rangeRequest.Start)
	}

	if rangeRequest.End != 499 {
		t.Errorf("rangeRequest.End = %d // expected 499", rangeRequest.End)
	}
}
