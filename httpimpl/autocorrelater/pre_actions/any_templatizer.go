package pre_actions

import (
	"gunship/correlators"
	template2 "gunship/httpimpl/execution"
	request_processors2 "gunship/httpimpl/execution/request_processors"
	"strings"
)

type anyTemplater struct {
}

func NewAnyTemplater() *anyTemplater {
	return &anyTemplater{}
}


func (this *anyTemplater) ProcessRequest(request correlators.RawRequest, ctx map[string]map[string]string) {
	req := request.(*template2.HttpRawRequest)
	for _, literalsMap := range ctx {
		for literal, tmpltName := range literalsMap {
			tmplt := "{" + tmpltName + "}"
			req.Path = strings.ReplaceAll(req.Path, literal, tmplt)
			req.Body = strings.ReplaceAll(req.Body, literal, tmplt)

			for _, vals := range req.Query {
				for i, _ := range vals{
					vals[i] = strings.ReplaceAll(vals[i], literal, tmplt)
				}

			}
			for _, vals := range req.Headers {
				for i, _ := range vals{
					vals[i] = strings.ReplaceAll(vals[i], literal, tmplt)
				}

			}
		}
	}
	// Todo, add extra check to add template compilers to only those request
	// which have been templatized
	req.RequestProcessors_ = append(req.RequestProcessors_, request_processors2.NewTemplateCompiler())
}
