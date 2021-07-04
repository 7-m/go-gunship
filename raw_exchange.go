package gunship

type RawRequest interface {
	AddRequestProcessor(processor ExecutionRequestProcessor)
	RequestProcessors() []ExecutionRequestProcessor
}
type RawResponse interface {
	AddResponseProcessor(processor ExecutionResponseProcessor)
	ResponseProcessors() []ExecutionResponseProcessor
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
