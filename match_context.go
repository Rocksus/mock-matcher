package matcher

import "context"

// NewContextMatcher will return a matcher that matches context type
func NewContextMatcher() *contextMatcher {
	return &contextMatcher{}
}

type contextMatcher struct {
	matchValue bool
}

// Matches implements the gomock Matches interface function
func (m *contextMatcher) Matches(x interface{}) bool {
	_, ok := x.(context.Context)
	return ok
}

// String implements the gomock String interface function
func (m *contextMatcher) String() string {
	return "Matches context type regardless of values"
}
