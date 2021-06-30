package pre_actions

import (
	"gunship/correlators"
	template2 "gunship/httpimpl/execution"
	"gunship/utils"
	"strconv"
)

// Extrract and store the literal at the position specified by a literal
type urlExtractor struct {
	tokenziner *utils.UrlTokenizer

}

func NewUrlExtractor(tokeizer string) *urlExtractor {
	return &urlExtractor{utils.NewUrlTokenzier(tokeizer)}
}

func (u urlExtractor) ProcessRequest(r correlators.RawRequest, ctx map[string]map[string]string) {
	// req has the url values
	// /api/endpoint/xxxx-yyyy-zzzz
	// /api/endpoint/{value}

	// tokenize the request

	req := r.(*template2.HttpRawRequest)
	reqtokenizer := utils.NewUrlTokenzier(req.Path)
	// process Path section

	for i, token := range u.tokenziner.PathAtoms() {
		// api endpoint xxxx-yyyy-zzz
		// api endpoint {value}
		if utils.IsTemplate(token) {
			token = token[1 : len(token)-1] // remove braces
			utils.EnsureValue(ctx, token, map[string]string{})
			literalMap := ctx[token]
			literalMap[reqtokenizer.PathAtoms()[i]] = token +  strconv.Itoa(len(literalMap))
		}
	}

}

