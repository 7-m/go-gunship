package matchers

import (
	"gunship"
	template2 "gunship/httpimpl/execution"
	request_processors2 "gunship/httpimpl/execution/request_processors"
	"gunship/utils"
	"strings"
)

// wildcardMatcher matcher allows matching against urls using wildcards in the
// pattern ex. /api/person/{personid} will match /api/person/per1
// /api/per2 etc.

type wildcardMatcher struct {
	urlTokenizer *utils.UrlTokenizer
}


func NewWildcardMatcher(pattern string) *wildcardMatcher {
	return &wildcardMatcher{utils.NewUrlTokenzier(pattern)}
}

func (this *wildcardMatcher) Match(exchange gunship.RawExchange) bool {
	req:= exchange.(*template2.HttpRawExchange).Request
	// match with just url
	reqAtoms := strings.Split(req.Path[1:], "/")
	if len(reqAtoms) != len(this.urlTokenizer.PathAtoms()){
		return false
	}
	for i, atom := range this.urlTokenizer.PathAtoms() {
			if utils.IsTemplate(atom) {
				continue
			}
			if atom == reqAtoms[i] {
				continue
			}else {
				return false
			}
		}


	return true
}

// ctx[templatename][acutalValue] = nTHTemplate
func (this *wildcardMatcher) ProcessRequest(req *template2.HttpRawRequest, ctx map[string]map[string]string) {
	// assumption that we have already
	// matched it using matcher()
	reqAtoms := strings.Split(req.Path[1:], "/")
	result := ""
	for i, atom := range this.urlTokenizer.PathAtoms() {
		if utils.IsTemplate(atom) {
			templateName := reqAtoms[i]
			result += "/" + "{" + ctx[atom[1: len(atom)-1]][templateName] + "}"
		} else {
			result += "/" + atom
		}
	}

	// process queries
	for key, values := range req.Query {
		for idx, value := range values {
			// todo non existent key check
			values[idx] = ctx[key][value]
		}
	}
	// process header
	for key, values := range req.Headers {
		for idx, value := range values {
			values[idx] = ctx[key][value]
		}
	}

	req.Path = result
	// TODO: add correlation for body too.

	// TODO: use globally shared after/before actions
	// will save alot of need object creation as they
	// are stateless
	req.AddRequestProcessor(request_processors2.NewTemplateCompiler())

}
