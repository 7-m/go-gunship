package execution

import (
	"gunship/execution"
	"gunship/utils"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HttpRawRequest represents an HTTP request right from the source. For example,
// from HAR file or from cURL etc. These are then tranformed in place
// using their repective processors.
type HttpRawRequest struct {
	Method  string
	BaseUrl string
	Path    string
	Query   map[string][]string
	Headers map[string][]string
	Body    string
	// requestProcessors and responseProcessors action to perfrom requestProcessors
	// and responseProcessors execution of the actual http request
	RequestProcessors_ []execution.RequestProcessor
}

func (h *HttpRawRequest) AddRequestProcessor(processor execution.RequestProcessor) {
	h.RequestProcessors_ = append(h.RequestProcessors_, processor)

}

func (h HttpRawRequest) RequestProcessors() []execution.RequestProcessor {
	return h.RequestProcessors_
}


func (h HttpRawRequest) RawRequest() {
	panic("marker method do not use")
}

func NewRawRequest(r *http.Request) *HttpRawRequest {
	all, err := ioutil.ReadAll(r.Body)
	utils.Panic(err, "error reading request body")
	return &HttpRawRequest{
		BaseUrl:            r.URL.Scheme + "://" + r.URL.Host,
		Path:               r.URL.Path,
		Query:              r.URL.Query(),
		Headers:            r.Header,
		RequestProcessors_: []execution.RequestProcessor{},
		Body:               string(all),
	}
}

// ************* RawRequest Builder *************

type rawRequestBuilder struct {
	Method  string
	BaseUrl string
	Path    string
	Query   map[string][]string
	Headers map[string][]string
	Body    string
	// requestProcessors and responseProcessors action to perfrom requestProcessors
	// and responseProcessors execution of the actual http request
	RequestProcessors []execution.RequestProcessor
}


func RawRequestBuilder() *rawRequestBuilder {
	return &rawRequestBuilder{
		Method:            "GET",
		BaseUrl:           "",
		Path:              "",
		Query:             map[string][]string{},
		Headers:           map[string][]string{},
		Body:              "",
		RequestProcessors: []execution.RequestProcessor{},
	}
}
func (r *rawRequestBuilder) SetMethod(method string) *rawRequestBuilder {
	r.Method = method
	return r
}
func (r *rawRequestBuilder) SetBaseUrl(baseUrl string) *rawRequestBuilder {
	r.BaseUrl = baseUrl
	return r
}
func (r *rawRequestBuilder) SetFromUrl(url *url.URL) *rawRequestBuilder {
	r.BaseUrl = url.Scheme + "://" + url.Host
	r.Path = url.Path
	r.Query = url.Query()
	return r
}
func (r *rawRequestBuilder) SetPath(path string) *rawRequestBuilder {
	r.Path = path
	return r
}
func (r *rawRequestBuilder) SetQuery(query map[string][]string) *rawRequestBuilder {
	r.Query = query
	return r
}
// Deprecated use AddHeader()
func (r *rawRequestBuilder) SetHeaders(headers map[string][]string) *rawRequestBuilder {
	r.Headers = headers
	return r
}
func (r *rawRequestBuilder) SetBody(body string) *rawRequestBuilder {
	r.Body = body
	return r
}
func (r *rawRequestBuilder) AddRequestProcessor(processor execution.RequestProcessor) *rawRequestBuilder {
	r.RequestProcessors = append(r.RequestProcessors, processor)
	return r
}

func (r *rawRequestBuilder) Build() *HttpRawRequest {
	if r.BaseUrl == "" {
		panic("base url cannot be empty in raw request")
	}
	if r.Method == "" {
		panic("Method cannot be empty in raw request")
	}

	return &HttpRawRequest{
		Method:             r.Method ,
		BaseUrl:            r.BaseUrl,
		Path:               r.Path,
		Query:              r.Query,
		Headers:            r.Headers,
		Body:               r.Body,
		RequestProcessors_: r.RequestProcessors,
	}
}

func (r *rawRequestBuilder) AddHeader(key string, value string) *rawRequestBuilder {
	if _, ok := r.Headers[key]; !ok {
		r.Headers[key] = []string{}
	}
	r.Headers[key] = append(r.Headers[key], value)
	return r
}
