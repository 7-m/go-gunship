package utils

import (
	"gunship"
	execution2 "gunship/httpimpl/execution"
)

func CastToCompiledRequest(uncasted []*execution2.HttpCompiledRequest)[]gunship.CompiledRequest {
	compiledRequests := []gunship.CompiledRequest{}
	for _, e := range uncasted {
		compiledRequests = append(compiledRequests, e)
	}
	return compiledRequests
}

func CastListToRawExchanges(uncasted []*execution2.HttpRawExchange)[]gunship.RawExchange {
	compiledRequests := []gunship.RawExchange{}
	for _, e := range uncasted {
		compiledRequests = append(compiledRequests, e)
	}
	return compiledRequests
}