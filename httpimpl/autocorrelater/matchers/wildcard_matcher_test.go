package matchers

import (
	template2 "gunship/httpimpl/execution"
	"net/http"
	"testing"
)

func TestSplitMatcher_Matcher(t *testing.T) {
	// TODO add test cases with trailing slashes
	mathcer := NewWildcardMatcher("/api/v1/{cid}")

	req, _ := http.NewRequest("", "https://www.some.com/api/v1/123?cat=purr&cat=mew&dog=woof", nil)
	rawRequest := template2.NewRawRequest(req)
	if !mathcer.Match(rawRequest) {
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
