package gunship

type ExchangeCompilerFunc func(exchange RawExchange, template *Template) CompiledRequest

// Correlate accepts set of exchanges and correlator.Templates and applies the templates to the
// exchanges. It then compiles the exchange to a compiled request. If the exchange matches multiple
// templates then, the processors of all the matching templates are applied but the last exchange
// compiled using exchangeCompiler is considered
func Correlate(exchanges []RawExchange,
	templates []*Template, ctx map[string]map[string]string,
	exchangeCompiler ExchangeCompilerFunc) []CompiledRequest {
	compiledRequests := []CompiledRequest{}

	for _, exchange := range exchanges {
		var compiledRequest CompiledRequest
		for _, tmplt := range templates {
			if tmplt.Matches(exchange) {
				for _, pre := range tmplt.RequestProcessors() {
					pre.ProcessRequest(exchange.RawRequest(), ctx)
				}
				for _, post := range tmplt.ResponseProcessors() {
					post.ProcessResponse(exchange.RawResponse(), ctx)
				}
				compiledRequest = exchangeCompiler(exchange, tmplt)
			}
		}
		compiledRequests = append(compiledRequests, compiledRequest)

	}
	return compiledRequests

}
