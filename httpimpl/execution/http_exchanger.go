package execution

import (
	"gunship"
	"net/http"
)

type httpExchanger struct {
	client *http.Client
}

func NewHttpExchanger(client *http.Client) *httpExchanger {
	return &httpExchanger{client: client}
}

func (h *httpExchanger) Exchange(request gunship.CompiledRequest) (interface{}, error) {
	r := request.(*HttpCompiledRequest)
	return h.client.Do(r.ToHttpRequest())
}
