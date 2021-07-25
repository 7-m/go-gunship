package executor

import (
	"fmt"
	"gunship"
	"gunship/httpimpl/execution"
	"testing"
)

func Test_defaultHanlder(t *testing.T) {

	defer func() {
		r := recover().(error)
		if r.Error() != "NotFound" {
			t.Fail()
		}
	}()
	mock := &mock{}
	defaultHandler := &defaultErrorHanlder{}
	req := &execution.HttpCompiledRequest{}
	Execute([]gunship.CompiledRequest{req}, mock, nil, nil, defaultHandler, func(m map[string]interface{}) {
		// no-op
	})

}

func Test_requestHandler(t *testing.T) {
	defer func() {
		r := recover().(error)
		if r.Error() != "something went wrong: NotFound" {
			t.Fail()
		}
	}()
	mock := &mock{}
	defaultHandler := &defaultErrorHanlder{}
	req := &execution.HttpCompiledRequest{
		ErrorHandler: mock,
	}
	Execute([]gunship.CompiledRequest{req}, mock, nil, nil, defaultHandler, func(m map[string]interface{}) {
		// no-op
	})
}

// default error handler
type defaultErrorHanlder struct{}

func (d defaultErrorHanlder) HandleError(e error, response interface{}, req gunship.CompiledRequest, xchgCtx, ctx map[string]interface{}, defaultErrorHandler gunship.ErrorHandler) {
	panic(e)
}

type mock struct{}

func (m *mock) HandleError(e error, response interface{}, req gunship.CompiledRequest, xchgCtx, ctx map[string]interface{}, defaultErrorHandler gunship.ErrorHandler) {
	panic(fmt.Errorf("something went wrong: " + e.Error()))
}

func (m *mock) Exchange(request gunship.CompiledRequest) (interface{}, error) {
	return nil, fmt.Errorf("NotFound")
}
