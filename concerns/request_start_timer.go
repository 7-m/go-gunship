package concerns

import (
	"gunship"
	"time"
)

type requestTimer struct {
}

func NewRequestTimer() *requestTimer {
	return &requestTimer{}
}

func (r *requestTimer) ProcessRequest(req gunship.CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	xchngCtx["time"] = time.Now()
}

func (r *requestTimer) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
