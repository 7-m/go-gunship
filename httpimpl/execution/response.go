package execution

import (
	"gunship/execution"
	"gunship/utils"
	"io/ioutil"
	"net/http"
)

// HttpRawResponse represents an HTTP response as its from the source for example
// from a HAR file or cURL.
type HttpRawResponse struct {
	Headers map[string][]string
	Body    string
	After   []execution.ResponseProcessor
}

func (h *HttpRawResponse) AddResponseProcessor(processor execution.ResponseProcessor) {
	h.After = append(h.After, processor)
}

func (h *HttpRawResponse) ResponseProcessors() []execution.ResponseProcessor {
	return h.After
}

func (h *HttpRawResponse) RawResponse() {
	panic("marker method, not to be used")
}

func NewRawResponse(headers map[string][]string, body string) *HttpRawResponse {
	return &HttpRawResponse{Headers: headers, Body: body, After: []execution.ResponseProcessor{}}

}


func RawResponseFromHttp(response *http.Response) *HttpRawResponse {
	all, err := ioutil.ReadAll(response.Body)
	utils.Panic(err, "error reading response body")
	return &HttpRawResponse{
		Headers: response.Header,
		Body:    string(all),
	}

}

type rawResponseBuilder struct {
	Headers map[string][]string
	Body    string
	After   []execution.ResponseProcessor
}

func RawResponseBuilder() *rawResponseBuilder {
	return &rawResponseBuilder{
		Headers: map[string][]string{},
		Body:    "",
		After:   []execution.ResponseProcessor{},
	}
}
func (this *rawResponseBuilder) SetBody(body string) *rawResponseBuilder {
	this.Body = body
	return this
}
func (this *rawResponseBuilder) AddHeader(key, value string) *rawResponseBuilder {
	if _, ok := this.Headers[key]; !ok {
		this.Headers[key] = []string{}
	}
	this.Headers[key] = append(this.Headers[key], value)
	return this
}
func (this *rawResponseBuilder) Build() *HttpRawResponse {
	return &HttpRawResponse{
		Headers: this.Headers,
		Body:    this.Body,
		After:   this.After,
	}
}