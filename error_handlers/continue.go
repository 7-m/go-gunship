package error_handlers

import (
	"gunship"
)

type Continue struct {
}

func (f *Continue) HandleError(e error, response interface{}, xchgCtx, ctx map[string]interface{}, defaultErrorHandler gunship.ErrorHandler) {
	// do nothing
}
