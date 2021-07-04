package preconcerns

import (
	"fmt"
	"gunship"
	execution2 "gunship/httpimpl/execution"
)

type requestLogger struct {
}

func NewHttpRequestLogger() *requestLogger {
	return &requestLogger{}
}

func (r *requestLogger) ProcessRequest(req gunship.CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	fmt.Println(req.(*execution2.HttpCompiledRequest).Path)
}

func (r *requestLogger) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
