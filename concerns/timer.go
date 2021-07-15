package concerns

import (
	"gunship"
	"math/rand"
	"time"
)

type  Delay struct {
	constant int
	random int
}

func NewDelay(constant int, random int) *Delay {
	if constant < 0 || random < 0 {
		panic("Delay can't be negative")
	}
	return &Delay{constant: constant, random: random}
}

func (t *Delay) ProcessRequest(req gunship.CompiledRequest, xchngCtx map[string]interface{}, sessionCtx map[string]interface{}) {
	time.Sleep(time.Millisecond * time.Duration(t.constant) + time.Millisecond * time.Duration(rand.Intn(t.random)))
}

func (t *Delay) MarshalJSON() ([]byte, error) {
	panic("implement me")
}


