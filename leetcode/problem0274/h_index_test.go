package problem0274

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(citations []int) int) {
	testCases := []struct {
		citations []int
		want      int
	}{
		{
			citations: []int{3, 0, 6, 1, 5},
			want:      3,
		},
		{
			citations: []int{1, 3, 1},
			want:      1,
		},
		{
			citations: []int{4, 4, 0, 0},
			want:      2,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.citations)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestHIndex_counter(t *testing.T) {
	test(t, hIndex_counter)
}

func TestHIndex_sort(t *testing.T) {
	test(t, hIndex_sort)
}
