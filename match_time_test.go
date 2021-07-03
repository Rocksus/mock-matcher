package matcher

import (
	"testing"
	"time"
)

func Test_timeMatcher_Matches(t *testing.T) {
	matcher := NewTimeMatcher()
	data := time.Now()

	if res := matcher.Matches(data); !res {
		t.Errorf("failure, expected true")
	}

	if res := matcher.Matches(""); res {
		t.Errorf("failure, expected false")
	}
}
