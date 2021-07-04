package post_actions

import (
	"github.com/tidwall/gjson"
	"gunship"
	template2 "gunship/httpimpl/execution"
	response_processors2 "gunship/httpimpl/execution/response_processors"
	"gunship/utils"
	"strconv"
)

type jsonCorrelator struct {
	extractors map[string]string // paths -> variables
	count      int
}

func NewJsonCorrelator(extractors map[string]string) *jsonCorrelator {
	return &jsonCorrelator{extractors: extractors}
}

// ctx : variable -> value->variableN
func (this *jsonCorrelator) ProcessResponse(response gunship.RawResponse, ctx map[string]map[string]string) {
	resp := response.(*template2.HttpRawResponse)
	json := resp.Body
	templateVars := map[string]string{}
	for path, variable := range this.extractors {
		match := gjson.Get(json, path)
		if !match.Exists() {
			panic("variable not found in json body : " + path)
		}
		utils.EnsureValue(ctx, variable, map[string]string{})
		templateVar := variable + strconv.Itoa(this.count)
		ctx[variable][match.String()] = templateVar
		this.count++

		templateVars[path] = templateVar
	}
	resp.AddResponseProcessor(response_processors2.NewJsonBodyExtractor(templateVars))

}
