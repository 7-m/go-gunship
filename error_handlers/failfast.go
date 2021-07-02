package error_handlers

import (
	"gunship"
)

type FailFastErrorHandler struct {

}

func (f *FailFastErrorHandler) HandleError(e error, xchgCtx, ctx map[string]interface{}, defaultErrorHandler gunship.ErrorHandler) {
	panic(e)
}



