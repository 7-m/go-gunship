package postconcerns

import (
	"log"
	"net/http"
	"time"
)

type postconcern struct {
}

func NewResponseLogger() *postconcern {
	return &postconcern{}
}

func (p *postconcern) ProcessResponse(resp interface{}, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	r := resp.(*http.Response)
	sessionCtx["logger"].(*log.Logger).Printf("%v, %v, %v, %v, %v\n",
		r.Request.Method,
		r.Request.URL,
		r.StatusCode,
		xchngCtx["attempts"],
		time.Now().Sub(xchngCtx["time"].(time.Time)).Milliseconds())
}

func (p postconcern) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
