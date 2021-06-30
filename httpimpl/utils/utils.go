package utils

import (
	"gunship/correlators"
	"gunship/execution"
	execution2 "gunship/httpimpl/execution"
)

func CastToCompiledRequest(uncasted []*execution2.HttpCompiledRequest)[]execution.CompiledRequest{
	compiledRequests := []execution.CompiledRequest{}
	for _, e := range uncasted {
		compiledRequests = append(compiledRequests, e)
	}
	return compiledRequests
}

func CastListToRawExchanges(uncasted []*execution2.HttpRawExchange)[]correlators.RawExchange{
	compiledRequests := []correlators.RawExchange{}
	for _, e := range uncasted {
		compiledRequests = append(compiledRequests, e)
	}
	return compiledRequests
}