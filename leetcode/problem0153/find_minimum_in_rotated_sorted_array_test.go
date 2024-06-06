package problem0153

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) int) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.nums)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestFindMin_binary(t *testing.T) {
	testFramework(t, findMin_binary)
}

func TestFindMin_traversal(t *testing.T) {
	testFramework(t, findMin_traversal)
}
