package matcher

// NewInt64SliceMatcher will return a matcher that matches int64 slice regardless of order
func NewInt64SliceMatcher(value []int64) *int64SliceMatcher {
	return &int64SliceMatcher{
		matching: value,
	}
}

type int64SliceMatcher struct {
	matching []int64
}

// Matches implements the gomock Matches interface function
func (m *int64SliceMatcher) Matches(x interface{}) bool {
	// check whether interface is of type int64 slice
	data, ok := x.([]int64)
	if !ok {
		return false
	}

	if len(m.matching) != len(data) {
		return false
	}

	matcher := make(map[int64]struct{}, len(m.matching))

	for _, v := range data {
		matcher[v] = struct{}{}
	}

	for _, v := range m.matching {
		if _, ok := matcher[v]; !ok {
			return false
		}
	}

	return true
}

// String implements the gomock String interface function
func (m *int64SliceMatcher) String() string {
	return "Matches int64 slice regardless of ordering"
}
