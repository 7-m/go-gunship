package httpimpl

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gunship"
	"gunship/correlators"
	"gunship/execution"
	matchers2 "gunship/httpimpl/autocorrelater/matchers"
	post_actions2 "gunship/httpimpl/autocorrelater/post_actions"
	pre_actions2 "gunship/httpimpl/autocorrelater/pre_actions"
	template2 "gunship/httpimpl/execution"
	testutils2 "gunship/httpimpl/testutils"
	"gunship/utils"
	"testing"
)

func Test_correlate(t *testing.T) {

	mockerServer, _ := testutils2.CreateServer1()
	defer mockerServer.Close()

	var exchanges  = []correlators.RawExchange{}
	for _, e := range testutils2.GetExhchanges1(mockerServer.URL) {
		exchanges = append(exchanges, e)
	}

	// match the auth api and extract auth token from response
	authMatcher := matchers2.NewWildcardMatcher("/api/v1/login")
	authXtrct := post_actions2.NewJsonCorrelator(map[string]string{"authorization": "AUTH"})
	authTemplate := correlators.NewBuilder().
		SetMatcher(authMatcher).
		AddResponseProcessor(authXtrct).
		Build()

	// create person and extract its id
	createPersonMatcher := matchers2.NewWildcardMatcher("/api/v1/person")
	personIdXtrct := post_actions2.NewJsonCorrelator(map[string]string{"personId": "PERSON"})
	createPersonTemplate :=  correlators.NewBuilder().
		SetMatcher(createPersonMatcher).
		AddResponseProcessor(personIdXtrct).
		Build()

	// any matcher template to replace all person-id and auth occurences
	anymatcherTemplate :=  correlators.NewBuilder().
		SetMatcher(matchers2.AnyMatcher{}).
		AddRequestProcessor(&pre_actions2.anyTemplater{}).
		Build()

	compiledReqs := gunship.Correlate(exchanges,
		[]*correlators.Template{authTemplate, createPersonTemplate, anymatcherTemplate},
		map[string]map[string]string{},
		template2.FromExchange)

	json := template2.CompiledRequestsToJson(compiledReqs)
	fmt.Println(string(json))


	var bug bytes.Buffer
	encoder := gob.NewEncoder(&bug)
	err := encoder.Encode(compiledReqs)
	utils.Panic(err, "error writing file" )
	var cr []execution.CompiledRequest
	decoder := gob.NewDecoder(&bug)
	err = decoder.Decode(&cr)
	utils.Panic(err, "couldnt decode")


}

