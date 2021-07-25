package errorhandlers

import (
	"fmt"
	"gunship"
	"net/http"
)

type FailFastErrorHandler struct {
}

func (f *FailFastErrorHandler) HandleError(e error, response interface{}, request gunship.CompiledRequest, xchgCtx, ctx map[string]interface{}, defaultErrorHandler gunship.ErrorHandler) {
	resp := response.(*http.Response)
	if e != nil {
		panic(e)
	}
	if resp.StatusCode == 500 {
		panic(fmt.Sprintf("Request\n %v Response %v", request, response))
	}
}


