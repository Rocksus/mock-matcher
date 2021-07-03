package matcher

import "regexp"

// NewRegexMatcher will return a matcher that matches string value based on the regex
func NewRegexMatcher(pattern string) *regexMatcher {
	return &regexMatcher{
		regex: regexp.MustCompile(pattern),
	}
}

type regexMatcher struct {
	regex *regexp.Regexp
}

// Matches implements the gomock Matches interface function
func (m *regexMatcher) Matches(x interface{}) bool {
	data, ok := x.(string)
	if !ok {
		return false
	}

	return m.regex.MatchString(data)
}

// String implements the gomock String interface function
func (m *regexMatcher) String() string {
	return "Matches string based on regex pattern"
}
