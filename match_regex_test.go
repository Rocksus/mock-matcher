package matcher

import (
	"testing"

	"github.com/google/uuid"
)

func Test_regexMatcher_Matches(t *testing.T) {
	pattern := "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
	matcher := NewRegexMatcher(pattern)

	// generate UUID v4
	data, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	if res := matcher.Matches(data.String()); !res {
		t.Errorf("failure, expected true")
	}

	if res := matcher.Matches("a123av"); res {
		t.Errorf("failure, expected false")
	}
}
