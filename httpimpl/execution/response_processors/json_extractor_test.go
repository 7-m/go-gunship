package response_processors

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestJsonExtractor_after(t *testing.T) {
	ctx := map[string]string{}
	json := `
			{
				"user" : {
					"profile" :{
						"id" : 123
					}
				}
			}
			`
	response := &http.Response{Body: ioutil.NopCloser(strings.NewReader(json))}
	jsonextractor := getTestCase(ctx)

	jsonextractor.ProcessResponse(response, nil, map[string]interface{}{"template": ctx})

	if _, ok := ctx["id"]; !ok {
		t.Fail()
	}
}

func getTestCase(ctx map[string]string) *jsonExtractor {

	extractor := map[string]string{"user.profile.id": "id"}
	return NewJsonBodyExtractor(extractor)
}
