package wrapstesting

import (
	"net/http"
	. "launchpad.net/gocheck"
)

type rangeSuite struct{}

var _ = Suite(&rangeSuite{})

func (rs *rangeSuite) TestParseRangeRequest(c *C) {
	rq, _ := http.NewRequest("GET", "/", nil)
	// Range: name ..; order=desc,max=10;
	rq.Header.Set("Range", "name ..;items=1-499,order=desc,max=10;")

	rangeRequest, err := ParseRangeRequest(rq, "name")

	c.Assert(err, Equals, nil)
	c.Assert(rangeRequest.Max, Equals, 10)
	c.Assert(rangeRequest.SortBy, Equals, "name")
	c.Assert(rangeRequest.Desc, Equals, true)
	c.Assert(rangeRequest.Start, Equals, 1)
	c.Assert(rangeRequest.End, Equals, 499)
}
