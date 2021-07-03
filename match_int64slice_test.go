package matcher

import "testing"

func Test_int64SliceMatcher_Matches(t *testing.T) {
	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	matcher := NewInt64SliceMatcher(data)

	// simulate random map assignment
	mapData := make(map[int64]int64)
	for _, v := range data {
		mapData[v] = v
	}
	newData := make([]int64, 0)
	for _, v := range mapData {
		newData = append(newData, v)
	}

	if res := matcher.Matches(newData); !res {
		t.Errorf("failure, expected true")
	}

	if res := matcher.Matches([]int64{1, 2, 9, 8, 7, 6, 5, 4, 3}); res {
		t.Errorf("failure, expected false")
	}
}
