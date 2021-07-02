package gunship

// CorrelationTemplate
type Template struct {
	matcher            Matcher // checks if this should be applied ornot
	requestProcessors  []RequestProcessor
	responseProcessors []ResponseProcessor
	errorCallback      ErrorHandler
}

func NewTemplate(matcher Matcher, before []RequestProcessor, after []ResponseProcessor) *Template {
	return &Template{matcher: matcher, requestProcessors: before, responseProcessors: after}
}

func (t *Template) ErrorCallback() ErrorHandler {
	return t.errorCallback
}

func (t *Template) RequestProcessors() []RequestProcessor {
	return t.requestProcessors
}
func (t *Template) ResponseProcessors() []ResponseProcessor {
	return t.responseProcessors
}
func (t *Template) Matcher() Matcher {
	return t.matcher
}

func (t *Template) Matches(exchange RawExchange) bool {
	return t.matcher.Match(exchange)
}

type ErrorHandler interface {
	HandleError(e error, xchgCtx, ctx map[string]interface{}, defaultErrorHandler ErrorHandler)
}

// ******** Template Builder *******
type templateBuilder struct {
	matcher            Matcher
	requestProcessors  []RequestProcessor
	responseProcessors []ResponseProcessor
	errorHandler       ErrorHandler
}

func NewBuilder() *templateBuilder {
	builder := templateBuilder{}
	builder.requestProcessors = []RequestProcessor{}
	builder.responseProcessors = []ResponseProcessor{}
	return &builder
}
func (this *templateBuilder) SetMatcher(matcher Matcher) *templateBuilder {
	this.matcher = matcher
	return this
}
func (this *templateBuilder) AddRequestProcessor(requestProcessor RequestProcessor) *templateBuilder {
	this.requestProcessors = append(this.requestProcessors, requestProcessor)
	return this
}
func (this *templateBuilder) SetErrorHandler(handler ErrorHandler) *templateBuilder {
	this.errorHandler = handler
	return this
}
func (this *templateBuilder) AddResponseProcessor(responseProcessor ResponseProcessor) *templateBuilder {
	this.responseProcessors = append(this.responseProcessors, responseProcessor )
	return this
}
func (this *templateBuilder) Build() *Template {
	// validate fields
	if this.matcher == nil {
		panic("no correlator matcher set")
	}
	if len(this.responseProcessors) == 0 && len(this.requestProcessors) == 0 {
		panic("Both request and response correlators are empty")
	}
	return &Template{matcher: this.matcher, requestProcessors: this.requestProcessors,
		responseProcessors: this.responseProcessors}
}

