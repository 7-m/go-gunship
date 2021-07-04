package execution

import (
	"gunship"
)

// HttpRawExchange represent a pair of raw HTTP request and response
type HttpRawExchange struct {
	Request  *HttpRawRequest
	Response *HttpRawResponse
}

func (h *HttpRawExchange) RawRequest() gunship.RawRequest {
	return h.Request
}

func (h *HttpRawExchange) RawResponse() gunship.RawResponse {
	return h.Response
}
