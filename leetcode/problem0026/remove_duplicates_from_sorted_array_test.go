package problem0026

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(nums []int) int) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.nums)
		if !reflect.DeepEqual(got, len(testCase.want)) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, len(testCase.want))
		}
		if !reflect.DeepEqual(testCase.nums[:got], testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, testCase.nums[:got], testCase.want)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	test(t, removeDuplicates)
}
