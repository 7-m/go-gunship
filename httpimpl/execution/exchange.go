package execution

import (
	"gunship/correlators"
)

// HttpRawExchange represent a pair of raw HTTP request and response
type HttpRawExchange struct {
	Request  *HttpRawRequest
	Response *HttpRawResponse
}

func (h *HttpRawExchange) RawRequest() correlators.RawRequest {
	return h.Request
}

func (h *HttpRawExchange) RawResponse() correlators.RawResponse {
	return h.Response
}




