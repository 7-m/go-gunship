package executor

import (
	"gunship"
	"log"
	"sync"
)

type Exchanger interface {
	Exchange(request gunship.CompiledRequest) (interface{}, error)
}

func Execute(compiledRequests []gunship.CompiledRequest, exchanger Exchanger,
	preConcerns []gunship.ExecutionRequestProcessor, postConcerns []gunship.ExecutionResponseProcessor,
	defaultHandler gunship.ErrorHandler) {
	sessionCtx := map[string]interface{}{}

	// todo add ctx init to method
	sessionCtx["template"] = map[string]string{}

	for _, e := range compiledRequests {

		reqCpy := e.MakeCopy()
		// perform request processing and prerequest actions
		xchng := map[string]interface{}{}
		reqCpy.ProcessRequest(xchng, sessionCtx)
		for _, b := range preConcerns {
			b.ProcessRequest(reqCpy, xchng, sessionCtx)
		}

		response, err := exchanger.Exchange(reqCpy)

		if err != nil {
			reqCpy.HandleError(err, response, xchng, sessionCtx, defaultHandler)
		}

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
	parallelism int, defaultHandler gunship.ErrorHandler) {

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
			Execute(compiledRequests, getHttpExchanger(), preConcerns, postConcerns, defaultHandler)

		}()

	}
	group.Wait()

}
