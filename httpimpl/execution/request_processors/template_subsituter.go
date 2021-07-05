package request_processors

import (
	"encoding/gob"
	"encoding/json"
	"gunship"
	template2 "gunship/httpimpl/execution"
	"strings"
	"unicode"
)

func init() {
	gob.Register(&TemplateCompiler{})
}

type TemplateCompiler struct {
}

func (this *TemplateCompiler) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string
	}{
		Type: "TemplateCompiler",
	},
	)
}

func NewTemplateCompiler() *TemplateCompiler {
	return &TemplateCompiler{}
}

func (t *TemplateCompiler) ProcessRequest(request gunship.CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	r := request.(*template2.HttpCompiledRequest)
	Ctx := sessionCtx["template"].(map[string]string)
	// replace in path, header and body
	r.BaseUrl = replace(r.BaseUrl, Ctx)
	r.Path = replace(r.Path, Ctx)
	// Todo handle multiple headers, Note :-
	// this wont be appicable as multiple values
	// will already be combined into a single csv
	// string, verify it
	for _, vals := range r.Headers {

		for i := range vals {
			vals[i] = replace(vals[i], Ctx)
		}

	}
	for _, vals := range r.Query {

		for i := range vals {
			vals[i] = replace(vals[i], Ctx)
		}

	}
	r.Body = replace(r.Body, Ctx)

}

// replaces '{xyz}' with Ctx['xyz']
func replace(s string, vars map[string]string) string {
	sb := strings.Builder{}
	lim := len(s)
	i := 0
	// 3 bcoz minimal string is {a}
NextChar:
	for ; i < lim-2; i++ {

		if s[i] == '{' {
			// search for the ending }
			i++
			start := i
			for ; i < lim; i++ {
				// if the char is not a }, letter or number, then re read char
				if !(s[i] == '}' || unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))) {
					sb.WriteString(s[start-1 : i])
					i--
					continue NextChar
				}
				if s[i] == '}' {
					// replace everything in between
					variable := s[start:i]
					replacement, ok := vars[variable]
					if !ok {
						panic("no literal found for template {" + variable + "}")
					}
					sb.WriteString(replacement)
					continue NextChar
				}
			}
			// we come here if we run of characters
			// while looking for }
			sb.WriteString(s[start-1 : lim])
		} else {
			sb.WriteByte(s[i])
		}

	}
	if i < lim {
		sb.WriteString(s[i:])
	}
	return sb.String()
}
