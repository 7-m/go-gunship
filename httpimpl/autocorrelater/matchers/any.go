package matchers

import (
	"gunship/correlators"
)

type AnyMatcher struct {

}

func (this AnyMatcher) Match(exchange correlators.RawExchange) bool {
	return true
}
