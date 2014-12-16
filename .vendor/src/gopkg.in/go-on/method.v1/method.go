package method

type Method string

const (
	POST    Method = "POST"
	GET     Method = "GET"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	OPTIONS Method = "OPTIONS"
	HEAD    Method = "HEAD"
	TRACE   Method = "TRACE"
)

func (m Method) String() string {
	return string(m)
}

// Is checks if the given string is the method
func (m Method) Is(meth string) bool {
	return string(m) == meth
}

// IsKnown checks if method is a known http method
func (m Method) IsKnown() bool {
	switch m {
	case GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD, TRACE:
		return true
	default:
		return false
	}
}

// IsREST checks if method one of the REST methods
func (m Method) IsREST() bool {
	switch m {
	case GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD:
		return true
	default:
		return false
	}
}

// MayHaveIfMatch checks if method is allowed to have an IfMatch header
func (m Method) MayHaveIfMatch() bool {
	switch m {
	case GET, PUT, DELETE, PATCH:
		return true
	default:
		return false
	}
}

// MayHaveEtag checks if method is allowed to have an Etag header
func (m Method) MayHaveEtag() bool {
	switch m {
	case GET, HEAD:
		return true
	default:
		return false
	}
}

/*
The following definitions are based on
http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html

and for PATCH
http://tools.ietf.org/html/rfc5789
*/

// a method is considered safe, if it
// has no requested sideeffect by the user agent
func (m Method) IsSafe() bool {
	return m == GET || m == HEAD || m == OPTIONS || m == TRACE
}

// a method is idempotent, if (aside from error or expiration issues)
// the side-effects of N > 0 identical requests is the same as for a single request
func (m Method) IsIdempotent() bool {
	return m != POST && m != PATCH
}

func (m Method) IsResponseCacheable() bool {
	// if it meets the requirements for HTTP caching described in section 13
	return m == GET || m == HEAD
}

// the the method return an empty message body
func (m Method) EmptyBody() bool {
	return m == HEAD || m == OPTIONS
}

/*
if the new field values indicate that the cached entity differs from the current entity (as would be indicated by a change in Content-Length, Content-MD5, ETag or Last-Modified), then the cache MUST treat the cache entry as stale.
*/

/*
The semantics of the GET method change to a "conditional GET" if the request message includes an If-Modified-Since, If-Unmodified-Since, If-Match, If-None-Match, or If-Range header field. A conditional GET method requests that the entity be transferred only under the circumstances described by the conditional header field(s). The conditional GET method is intended to reduce unnecessary network usage by allowing cached entities to be refreshed without requiring multiple requests or transferring data already held by the client.


The semantics of the GET method change to a "partial GET" if the request message includes a Range header field. A partial GET requests that only part of the entity be transferred, as described in section 14.35. The partial GET method is intended to reduce unnecessary network usage by allowing partially-retrieved entities to be completed without transferring data already held by the client.
*/

/*
The metainformation contained in the HTTP headers in response to a HEAD request SHOULD be identical to the information sent in response to a GET request
*/
