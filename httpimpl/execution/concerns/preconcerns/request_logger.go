package preconcerns

import (
	"fmt"
	"gunship/execution"
	execution2 "gunship/httpimpl/execution"
)

type requestLogger struct {

}

func NewHttpRequestLogger() *requestLogger {
	return &requestLogger{}
}


func (r *requestLogger) ProcessRequest(req execution.CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	fmt.Println(req.(*execution2.HttpCompiledRequest).Path);
}

func (r *requestLogger) MarshalJSON() ([]byte, error) {
	panic("implement me")
}



