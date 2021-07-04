package matchers

import (
	"gunship"
)

type AnyMatcher struct {
}

func (this AnyMatcher) Match(exchange gunship.RawExchange) bool {
	return true
}
