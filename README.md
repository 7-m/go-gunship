# Go Gunship

## Brief
Gunship is a library for load testing with auto-correlation/parameterization. Gunship was created with 3 goals in mind :
- To automatically parameterize/correlate request parameters from any source (Har supported for now)
- To allow for easy coding of the above correlators
- To allow for easy integration of cross-cutting concerns(like logging, metrics collection etc) during load testing

## Is Gunship for me ?
Gunship was designed to make performance testing easy, especially if you have a large # of API's, and many combinations of sequences in which they can be executed hence making it a big effort task to hand wire test plans. Gunship simplifies things by ingesting recorded requests and responses (such as Har), auto-correlating dynamic data and giving you a test plan which can be saved to disk and/or executed. So, if you have tons of API's and have come to the conclusion that hand wiring test plans is not a good idea and that it's better to record http request and responses, then Gunship is the tool for you.

## Getting Started
### HTTP
- Create a `Template` to associate dynamic data
- Run `Correlate` to create a `[]CompiledRequest`
- [Optional] Save `[]CompileRequest` to disk
- Configure and run the `[]CompiledRequest` with an executor

Refer the example in `httpimpl/correlate_test.go`.

#### [Check out the Wiki for more details](https://github.com/7-m/go-gunship/wiki)