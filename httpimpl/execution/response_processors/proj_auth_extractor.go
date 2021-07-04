package response_processors

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"gunship/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

type extractor struct {
}

func (e extractor) ProcessResponse(r interface{}, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	resp := r.(*http.Response)
	template := sessionCtx["template"].(map[string]string)
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Panic(err, "error reading json from body")
	}
	url := gjson.Get(string(json), "0.url").String()
	tokens := strings.Split(url, "/")
	template["AUTH0"] = tokens[len(tokens)-1]
	template["PROJECTID0"] = tokens[len(tokens)-2]
}

func (e extractor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		processName string
	}{
		processName: "custom-processor",
	})
}

func NewProjectauthExtractor() *extractor {
	return &extractor{}
}
