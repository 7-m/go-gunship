package httpimpl

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gunship"
	matchers2 "gunship/httpimpl/autocorrelater/matchers"
	post_actions2 "gunship/httpimpl/autocorrelater/post_actions"
	pre_actions2 "gunship/httpimpl/autocorrelater/pre_actions"
	template2 "gunship/httpimpl/execution"
	"gunship/httpimpl/execution/request_processors"
	"gunship/httpimpl/execution/response_processors"
	testutils2 "gunship/httpimpl/testutils"
	"gunship/utils"
	"testing"
)

func Test_correlate(t *testing.T) {

	mockerServer, _ := testutils2.CreateServer1()
	defer mockerServer.Close()

	var exchanges = []gunship.RawExchange{}
	for _, e := range testutils2.GetExhchanges1() {
		exchanges = append(exchanges, e)
	}

	// match the auth api and extract auth token from response
	authMatcher := matchers2.NewWildcardMatcher("/api/v1/login")
	authXtrct := post_actions2.NewJsonCorrelator(map[string]string{"authorization": "AUTH"})
	authTemplate := gunship.NewBuilder().
		SetMatcher(authMatcher).
		AddResponseProcessor(authXtrct).
		Build()

	// create person and extract its id
	createPersonMatcher := matchers2.NewWildcardMatcher("/api/v1/person")
	personIdXtrct := post_actions2.NewJsonCorrelator(map[string]string{"personId": "PERSON"})
	createPersonTemplate := gunship.NewBuilder().
		SetMatcher(createPersonMatcher).
		AddResponseProcessor(personIdXtrct).
		Build()

	// any matcher template to replace all person-id and auth occurences
	anymatcherTemplate := gunship.NewBuilder().
		SetMatcher(matchers2.AnyMatcher{}).
		AddRequestProcessor(pre_actions2.NewAnyTemplater()).
		Build()

	// we manually add the base urls variable to be BASE0
	correlatorCtx := map[string]map[string]string{"BASE": {"http://example.com": "BASE0"}}
	compiledReqs := gunship.Correlate(exchanges,
		[]*gunship.Template{authTemplate, createPersonTemplate, anymatcherTemplate},
		correlatorCtx,
		template2.FromExchange)

	// assert that the write execution processors were created
	login := compiledReqs[0]
	_ = login.RequestProcessors()[0].(*request_processors.TemplateCompiler)
	_ = login.ResponseProcessor()[0].(*response_processors.JsonExtractor)

	createPerson := compiledReqs[1]
	_ = createPerson.RequestProcessors()[0].(*request_processors.TemplateCompiler)
	_ = createPerson.ResponseProcessor()[0].(*response_processors.JsonExtractor)

	getPerson := compiledReqs[2]
	_ = getPerson.RequestProcessors()[0].(*request_processors.TemplateCompiler)

	updatePerson := compiledReqs[3]
	_ = updatePerson.RequestProcessors()[0].(*request_processors.TemplateCompiler)

	json := template2.CompiledRequestsToJson(compiledReqs)
	fmt.Println(string(json))

	// assert that the requests can be serialized
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(compiledReqs)
	utils.Panic(err, "error writing file")
	var cr []gunship.CompiledRequest
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&cr)
	utils.Panic(err, "couldnt decode")

}
