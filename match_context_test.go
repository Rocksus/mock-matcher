package matcher

import (
	"context"
	"testing"
)

func Test_contextMatcher_Matches(t *testing.T) {
	matcher := NewContextMatcher()
	data := context.Background()

	if res := matcher.Matches(data); !res {
		t.Errorf("failure, expected true")
	}

	if res := matcher.Matches(""); res {
		t.Errorf("failure, expected false")
	}
}
