package postconcerns

import (
	"fmt"
	"net/http"
)

type postconcern struct {
}

func NewResponseLogger() *postconcern {
	return &postconcern{}
}

func (p *postconcern) ProcessResponse(resp interface{}, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	r := resp.(*http.Response)
	fmt.Println(r.StatusCode)
}

func (p postconcern) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
