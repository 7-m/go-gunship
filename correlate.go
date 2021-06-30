package gunship

import (
	"gunship/correlators"
	"gunship/execution"
)

type ExchangeCompilerFunc func(exchange correlators.RawExchange) execution.CompiledRequest
// Correlate accepts set of exchanges and correlator.Templates and applies the templates to the
// exchanges. It then tranforms the exchanges to compiled requests
func Correlate(exchanges []correlators.RawExchange, templates []*correlators.Template,ctx map[string]map[string]string, exchangeCompiler ExchangeCompilerFunc) []execution.CompiledRequest {
	compiledRequests := []execution.CompiledRequest{}

	for _, exchange := range exchanges{
		for _, tmplt :=range templates {
			if tmplt.Matches(exchange){
				for _, pre := range tmplt.RequestProcessors(){
					pre.ProcessRequest(exchange.RawRequest(), ctx)
				}
				for _, post := range tmplt.ResponseProcessors() {
					post.ProcessResponse(exchange.RawResponse(), ctx)
				}
			}
		}
		compiledRequests = append(compiledRequests, exchangeCompiler(exchange))
	}
	return compiledRequests
	
}
