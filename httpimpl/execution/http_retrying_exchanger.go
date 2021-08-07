package execution

import (
	"gunship"
	"math/rand"
	"net/http"
	"time"
)

type httpExchanger struct {
	client *http.Client
}

func NewHttpExchanger(client *http.Client) *httpExchanger {
	return &httpExchanger{client: client}
}

func (h *httpExchanger) Exchange(request gunship.CompiledRequest,
	xchngCtx map[string]interface{},
	sessionCtx map[string]interface{}) (interface{}, error) {

	var do *http.Response
	var err error

	attempt :=1
	for ;  attempt <= 10 ; attempt++{
		r := request.(*HttpCompiledRequest)
		do, err = h.client.Do(r.ToHttpRequest())
		if err == nil && do.StatusCode != 500{
			xchngCtx["attempts"] = attempt
			return do, err
		}
		// timeout to prevent overloading
		time.Sleep(time.Duration(rand.Float32() * 10) * time.Second)
	}
	xchngCtx["attempts"] = attempt
	return do, err
}
