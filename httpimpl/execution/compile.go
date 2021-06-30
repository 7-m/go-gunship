package execution

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"gunship/correlators"
	"gunship/execution"
	"io/ioutil"
	"net/http"
)

func init() {
	gob.Register(&HttpCompiledRequest{})
}
// HttpCompiledRequest represents a templated HTTP request with its associated request and response processors.
// The request and response processors, transform/edit the request to prepare it to be convereted
// to a valid http.Request. Always MakeCopy() and then call its request processors so as to
// not modify the original template

type HttpCompiledRequest struct {
	Method              string
	BaseUrl             string
	Path                string
	Query               map[string][]string
	Headers             map[string][]string
	// perform template replacements and other actions here
	RequestProcessors_ []execution.RequestProcessor
	// perform data extractions actions here
	ResponseProcessors_ []execution.ResponseProcessor
	Body                string
	//// perform cross cutting actions just before a request is send
	//// and just after a response is recieved below
	//PreRequestActions   []RequestProcessor
	//PostResponseActions []ResponseProcessor
}

func (this HttpCompiledRequest) RequestProcessors() []execution.RequestProcessor {
	return this.RequestProcessors_
}

func (this HttpCompiledRequest) ResponseProcessorProcessors() []execution.ResponseProcessor {
	return this.ResponseProcessors_
}

func FromExchange(e correlators.RawExchange) execution.CompiledRequest {
	exchange := e.(*HttpRawExchange)
	req := exchange.Request
	resp := exchange.Response

	return &HttpCompiledRequest{
		Method:              req.Method,
		BaseUrl:             req.BaseUrl,
		Path:                req.Path,
		Query:               req.Query,
		Headers:             req.Headers,
		Body:                req.Body,
		RequestProcessors_:  req.RequestProcessors_,
		ResponseProcessors_: resp.After,
	}

}
func (this HttpCompiledRequest) makeQueryString() string {
	if len(this.Query) == 0 {
		return ""
	}
	qs := "?"

	for key, vals := range this.Query {
		for _, val := range vals {
			qs += key + "=" + val + "&"
		}
	}
	return qs

}
func (this HttpCompiledRequest) ToHttpRequest() *http.Request {
	req, err := http.NewRequest(this.Method, this.BaseUrl+this.Path+this.makeQueryString(),
		ioutil.NopCloser(bytes.NewBufferString(this.Body)))
	if err != nil {
		panic("Couldnt create url")
	}

	for key, vals := range this.Headers {
		req.Header[key] = vals
	}
	return req

}

func (this HttpCompiledRequest) MakeCopy() execution.CompiledRequest {
	cpy := &HttpCompiledRequest{
		Body:                this.Body,
		Method:              this.Method,
		BaseUrl:             this.BaseUrl,
		Path:                this.Path,
		Query:               nil,
		Headers:             nil,
		RequestProcessors_:  this.RequestProcessors_,
		ResponseProcessors_: this.ResponseProcessors_,
	}

	query := map[string][]string{}
	for key, vals := range this.Query {
		query[key] = make([]string, len(vals))
		copy(query[key], vals)

	}
	cpy.Query = query

	headers := map[string][]string{}
	for key, vals := range this.Headers {
		headers[key] = make([]string, len(vals))
		copy(headers[key], vals)

	}
	cpy.Headers = headers

	return cpy
}

func (this *HttpCompiledRequest) ProcessRequest(xchgCtx, sessionCtx map[string]interface{}) {
	for _, b := range this.RequestProcessors_ {
		b.ProcessRequest(this, xchgCtx, sessionCtx)
	}
}

func (this *HttpCompiledRequest) ProcessResponse(response interface{}, xchgCtx, sessionCtx map[string]interface{}) {
	for _, b := range this.ResponseProcessors_ {
		b.ProcessResponse(response.(*http.Response), xchgCtx, sessionCtx)
	}
}

func (this *HttpCompiledRequest) ToJson() string {
	str, err := json.Marshal(this)
	if err != nil {
		panic("error  converting to json")
	}
	return string(str)
}

func CompiledRequestsToJson(compiledrequests []execution.CompiledRequest) []byte {
	str, err := json.MarshalIndent(compiledrequests, "", " ")
	if err != nil {
		panic("error  converting to json")
	}
	return str

}
