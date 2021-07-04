package request_processors

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"gunship"
	template2 "gunship/httpimpl/execution"
	"strings"
)

func init() {
	gob.Register(&templateCompiler{})
}

type templateCompiler struct {
}

func (this *templateCompiler) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string
	}{
		Type: "templateCompiler",
	},
	)
}

func NewTemplateCompiler() *templateCompiler {
	return &templateCompiler{}
}

func (t *templateCompiler) ProcessRequest(request gunship.CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	r := request.(*template2.HttpCompiledRequest)
	Ctx := sessionCtx["template"].(map[string]string)
	// replace in path, header and body
	var err error
	r.Path, err = replace(r.Path, Ctx)
	if err != nil {
		panic("error replacing templates")
	}
	// Todo handle multiple headers, Note :-
	// this wont be appicable as multiple values
	// will already be combined into a single csv
	// string, verify it
	for _, vals := range r.Headers {

		for i, _ := range vals {
			vals[i], err = replace(vals[i], Ctx)
		}
		if err != nil {
			panic("error replacing templates")
		}

	}
	// Todo replace template in body

}

// replaces '{xyz}' with Ctx['xyz']
func replace(s string, vars map[string]string) (string, error) {
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
				if s[i] == '}' {
					// replace everything in between
					variable := s[start:i]
					replacement, ok := vars[variable]
					if !ok {
						panic("no variable found for replacement : " + variable)
					}
					sb.WriteString(replacement)
					continue NextChar
				}
			}
			// we come here if we run of characters
			// while looking for }
			return "", fmt.Errorf("Couldn't find closing '}'")
		} else {
			sb.WriteByte(s[i])
		}

	}
	sb.WriteString(s[i:])
	return sb.String(), nil
}
