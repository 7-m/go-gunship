# Go Gunship

## Brief
Gunship is library for load testing with auto-correlation/parameterization. Gunship was created with 3 goals in mind :
- To automatically parametrize/correlated request parameters from any source (Har supported for now)
- To allow for easy coding of the above correlators
- To allow for easy integration of cross-cutting concerns(like logging, metrics collection etc) during load testing

## Is Gunship for me ?
Gunship was designed to make performance testing easy, especially if you have a large # of API's, and many combinations of sequences in which executed hence making it a huge effort task to hand wire test plans. Gunship simplifies things by ingesting recorded requests and responses (such as Har), auto-correlating dynamic data and giving you a test plan which can be saved to disk and/or executed. So, if you have tons of API's and have come to the conclusion that hand wiring test plans is not a good idea, and it's better to record http request and responses, then Gunship is the tool for you.

## Getting Started
### HTTP
- Create `correlator.Template` to associate dyamic data
- Run 'http.Correlate' to create a `[]CompiledRequest`
- Save to disk
- Configure an executor
Refer the example in `httpimpl/correlate_test.go`.

## Concepts
### `RawRequest` and `RawResponse`
- `RawRequest` and `RawResponse` are used to represent requests and response right from the source. For example from a HAR file or cURL's. `RawRequest` and `RawResponse` are just interfaces, so a concrete implementation is required depending on the use case. For example, currently gunship has an implementation for HTTP. This can be extended to other protocols, for example, gRPC.
- `RawExchange` represent a pair of `RawRequest` and `RawResponse`.
- `RawRequest` and `RawResponse` both define an API to set and access a slice of `execution. RequestProcessor`'s and `execution.ResponseProcessor`'s respectively. These are used during 
  execution of  `CompiledRequest`s for the purpose of substituting templates with literals or extraction of literals etc.
  
The implementors of the 3 interfaces are expected to be mutable. They may be mutated by the correlation processors discussed below.


### Correlation
The correlation module resides in the `correlator`  package. They aid in templatizing exchanges, for example, replacing a resource id in the URL with a template and adding related behaviour to replace the template with a literal during execution of a load test. At the center of correlation is the `correlater.Template` which is based on 3 core interfaces :
- **Matcher** : Is a predicate which used to determine whether to process an exchange. The below request and response processors are only applied if they match the predicate 
- **RequestProcessor** : Contains a single method `ProcessRequest` which mutates a `RawRequest`, for example by parameterizing a literal (ex. replacing an authorization header value with a variable). Along with mutating the `RawRequest`, a `RequestProcessor` may also add an `execution.RequestProcessor` which, keeping the example in mind, replaces the variable with a literal during the execution of a load test.
- **ResponseProcessor** : Contains a single method `ProcessResponse` which can be used to extract data from a response, for example, extracting the id of a newly created person from a JSON response. Similar to `RequestProcessor`, `reposnseProcessors` may also add `execution.ResponseProcessor` for a `rawrequest` which, keeping the example in mind, extracts and stores a value from a JSON response for use by subsequent requests.
  
The implementors of the interfaces need to be immutable. The above interfaces have methods which operate on `RawRequest`, `RawResponse` and `RawExchange` and hence may or may not mutate them.

### Execution
- After all `RawExchange`'s have been processed using `correlator.Correlate`, the result is a `[]CompiledRequest` ready for execution. `CompiledRequest`s contain parametrized requests along with pre/post-processors for replacing/extracting data. An `executor` will always operate on a copy of the `CompiledRequest` and never mutate the original copy. Instead it creates a new copy and then applies the `RequestProcessor`. This is to prevent race conditions and the need for synchronization at the cost of garbage. Executors also accept slices of `RequestProcessors` and `ResponseProcessors` to aid in implementing cross-cutting concerns like logging, metric collection or to alter execution for ex. add timers, debug context by peeking into the session context etc.
- `Exchanger` interface is used to 


###

### Working
We begin by defining correlation.Templates which specifies using the Matcher interface, which exchanges do need to be pre or post processed. The preprocessing is done using a list of RequestProcessors wherein each RequestProcessor in the list is applied to the request sequentially. It has the signature `ProcessRequest(req RawRequest, ctx map[string]map[string]string)` where `req` specifies the RawRequest to process and `ctx` holds the list of templates extracted from the previously processed requests.