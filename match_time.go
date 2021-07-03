package matcher

import (
	"time"
)

// NewTimeMatcher will return a matcher that matches time type
func NewTimeMatcher() *timeMatcher {
	return &timeMatcher{}
}

type timeMatcher struct {
}

// Matches implements the gomock Matches interface function
func (m *timeMatcher) Matches(x interface{}) bool {
	_, ok := x.(time.Time)
	return ok
}

// String implements the gomock String interface function
func (m *timeMatcher) String() string {
	return "Matches time value"
}
