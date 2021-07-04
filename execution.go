package gunship

type CompiledRequest interface {
	RequestProcessors() []ExecutionRequestProcessor
	ResponseProcessor() []ExecutionResponseProcessor
	MakeCopy() CompiledRequest
	HandleError(e error, resp interface{}, xchgCtx, ctx map[string]interface{}, defaultErrorHandler ErrorHandler)
	ProcessRequest(xchgCtx, ctx map[string]interface{})
	ProcessResponse(response interface{}, xchgCtx, ctx map[string]interface{})
}

// ExecutionRequestProcessor performs any tranformations/logging etc before a request is sent.
// Implementors of this interface shouldn't change state when processing a request
type ExecutionRequestProcessor interface {
	ProcessRequest(req CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{})
	MarshalJSON() ([]byte, error)
}

// ExecutionResponseProcessor  performs any tranformations/logging etc after recieving a response
// Implementors of this interface shouldn't change state when processing a response
type ExecutionResponseProcessor interface {
	ProcessResponse(resp interface{}, xchngCtx map[string]interface{}, sessionCtx map[string]interface{})
	MarshalJSON() ([]byte, error)
}
