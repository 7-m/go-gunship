package execution

type CompiledRequest interface {
	RequestProcessors() []RequestProcessor
	ResponseProcessorProcessors() []ResponseProcessor
	MakeCopy() CompiledRequest
	ProcessRequest(xchgCtx, ctx map[string]interface{})
	ProcessResponse(response interface{}, xchgCtx , ctx map[string]interface{})
}

// RequestProcessor performs any tranformations/logging etc before a request is sent.
// Implementors of this interface shouldn't change state when processing a request
type RequestProcessor interface {
	ProcessRequest(req CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{})
	MarshalJSON() ([]byte, error)
}

// ResponseProcessor  performs any tranformations/logging etc after recieving a response
// Implementors of this interface shouldn't change state when processing a response
type ResponseProcessor interface {
	ProcessResponse(resp interface{}, xchngCtx map[string]interface{}, sessionCtx map[string]interface{})
	MarshalJSON() ([]byte, error)
}
