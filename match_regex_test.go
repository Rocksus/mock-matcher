package matcher

import (
	"testing"
)

func Test_regexMatcher_Matches(t *testing.T) {
	pattern := "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
	matcher := NewRegexMatcher(pattern)

	if res := matcher.Matches("51e5fe10-5325-4a32-bce8-7ebe9708c453"); !res {
		t.Errorf("failure, expected true")
	}

	if res := matcher.Matches("abcdefgh-ijkl-mnop-qrst-uvwxyzabcdef"); res {
		t.Errorf("failure, expected false")
	}
}
