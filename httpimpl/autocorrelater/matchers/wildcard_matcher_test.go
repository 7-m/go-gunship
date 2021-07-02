package matchers

import (
	template2 "gunship/httpimpl/execution"
	"testing"
)

func TestSplitMatcher_Matcher(t *testing.T) {
	// TODO add test cases with trailing slashes
	mathcer := NewWildcardMatcher("/api/v1/{cid}")
	// cat=purr&cat=mew&dog=woof
	rawRequest := template2.RawRequestBuilder().
		SetPath("/api/v1/123").
		SetBaseUrl("https://www.some.com").
		SetQuery(map[string][]string{"cat" :{"purr", "mew"}, "dog" :{"woof"}}).
		Build()
	if !mathcer.Match(&template2.HttpRawExchange{
		Request: rawRequest,
	}) {
		t.Fail()
	}

	ctx := map[string]map[string]string{"cid": {"123": "cid0"}, "cat": {"purr": "variable1"}}

	mathcer.ProcessRequest(rawRequest, ctx)

	if rawRequest.Path != "/api/v1/{cid0}" {
		t.Fail()
	}

	if rawRequest.Query["cat"][0] != "variable1" {
		t.Fail()
	}
}
