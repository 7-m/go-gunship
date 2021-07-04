package response_processors

import (
	"encoding/gob"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func init() {
	gob.Register(&jsonExtractor{})
}

type jsonExtractor struct {
	Extractors map[string]string // paths -> variables
}

func NewJsonBodyExtractor(extractors map[string]string) *jsonExtractor {
	return &jsonExtractor{Extractors: extractors}
}

func (j *jsonExtractor) ProcessResponse(r interface{}, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	response := r.(*http.Response)
	ctx := sessionCtx["template"].(map[string]string)
	jsonBlob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("failed to read json from response body")
	}

	for path, variable := range j.Extractors {
		match := gjson.Get(string(jsonBlob), path)
		if match.Exists() {
			ctx[variable] = match.String()
		}
	}
}

func (j *jsonExtractor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       string
		Extractors map[string]string
	}{
		Type:       "jsonExtractor",
		Extractors: j.Extractors,
	})
}
