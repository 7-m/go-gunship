package executor

import (
	"gunship"
	"log"
	"sync"
)

type Exchanger interface {
	Exchange(request gunship.CompiledRequest) (interface{}, error)
}

type SessionCtxCustomizer func(map[string]interface{})

func Execute(compiledRequests []gunship.CompiledRequest, exchanger Exchanger,
	preConcerns []gunship.ExecutionRequestProcessor, postConcerns []gunship.ExecutionResponseProcessor,
	defaultHandler gunship.ErrorHandler, sessionCtxCustomizer SessionCtxCustomizer) {
	sessionCtx := map[string]interface{}{}

	// todo add ctx init to method
	sessionCtx["template"] = map[string]string{}
	sessionCtxCustomizer(sessionCtx)

	for _, e := range compiledRequests {

		reqCpy := e.MakeCopy()
		// perform request processing and prerequest actions
		xchng := map[string]interface{}{}
		reqCpy.ProcessRequest(xchng, sessionCtx)
		for _, b := range preConcerns {
			b.ProcessRequest(reqCpy, xchng, sessionCtx)
		}

		response, err := exchanger.Exchange(reqCpy)

		reqCpy.HandleError(err, response, reqCpy, xchng, sessionCtx, defaultHandler)

		// perform post processing and postresponse actions
		for _, a := range postConcerns {
			a.ProcessResponse(response, xchng, sessionCtx)
		}
		reqCpy.ProcessResponse(response, xchng, sessionCtx)

	}

}

type ExchangerFactory func() Exchanger

func ExecuteParallel(compiledRequests []gunship.CompiledRequest, getHttpExchanger ExchangerFactory,
	preConcerns []gunship.ExecutionRequestProcessor, postConcerns []gunship.ExecutionResponseProcessor,
	parallelism int, defaultHandler gunship.ErrorHandler, customizer SessionCtxCustomizer) {

	group := sync.WaitGroup{}
	for i := 0; i < parallelism; i++ {
		group.Add(1)
		go func() {
			defer func() {
				group.Done()
				if err := recover(); err != nil {
					log.Printf("user stopped after error %v", err)
				}
			}()
			log.Printf("--started-----")
			Execute(compiledRequests, getHttpExchanger(), preConcerns, postConcerns, defaultHandler, customizer)

		}()

	}
	group.Wait()

}
