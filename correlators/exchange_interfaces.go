package correlators

import "gunship/execution"

type RawRequest interface {
	AddRequestProcessor(processor execution.RequestProcessor)
	RequestProcessors() []execution.RequestProcessor
}
type RawResponse interface {
	AddResponseProcessor(processor execution.ResponseProcessor)
	ResponseProcessors() []execution.ResponseProcessor
}

type RawExchange interface {
	RawRequest() RawRequest
	RawResponse() RawResponse
}

type RequestProcessor interface {
	ProcessRequest(req RawRequest, ctx map[string]map[string]string)
}
type ResponseProcessor interface {
	ProcessResponse(resp RawResponse, ctx map[string]map[string]string)

}
type Matcher interface {
	Match(exchange RawExchange) bool
}



