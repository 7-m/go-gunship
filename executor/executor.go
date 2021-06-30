package executor

import (
	"fmt"
	"gunship/execution"
	"sync"
)

type Exchanger interface {
	Exchange(request execution.CompiledRequest)(interface{}, error)
}

func Execute(compiledRequests []execution.CompiledRequest, exchanger Exchanger,
	preConcerns []execution.RequestProcessor, postConcerns []execution.ResponseProcessor)  {
	sessionCtx := map[string]interface{}{}

	// todo add ctx init to method
	sessionCtx["template"] = map[string]string{}

	for _, e := range compiledRequests {

		reqCpy := e.MakeCopy()
		// perform request processing and prerequest actions
		reqCpy.ProcessRequest(nil, sessionCtx)
		for _, b:= range preConcerns {
			b.ProcessRequest(reqCpy, nil, sessionCtx)
		}


		response, err := exchanger.Exchange(reqCpy)
		if err != nil {
			panic(fmt.Sprintf("an error occured while executing request: %v", err))
		}

		// perform post processing and postresponse actions
		for _,a := range postConcerns {
			a.ProcessResponse(response, nil, sessionCtx)
		}
		reqCpy.ProcessResponse(response, nil, sessionCtx)

	}

}


type ExchangerFactory func()Exchanger

func ExecuteParallel(compiledRequests []execution.CompiledRequest, getHttpExchanger ExchangerFactory,
	preConcerns []execution.RequestProcessor, postConcerns []execution.ResponseProcessor,  parallelism int){

	group := sync.WaitGroup{}
	for i := 0; i < parallelism; i++ {
		group.Add(1)
		go func() {
			fmt.Printf("------------------------------------started-----------------------")

			Execute(compiledRequests, getHttpExchanger(), preConcerns, postConcerns)
			group.Done()
		}()

	}
	group.Wait()

}

