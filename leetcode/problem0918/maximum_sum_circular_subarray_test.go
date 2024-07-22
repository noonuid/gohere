package problem0918

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(nums []int) int) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, -2, 3, -2},
			want: 3,
		},
		{
			nums: []int{5, -3, 5},
			want: 10,
		},
		{
			nums: []int{3, -2, 2, -3},
			want: 3,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.nums)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMaxSubarraySumCircular(t *testing.T) {
	test(t, maxSubarraySumCircular)
}

func TestMaxSubarraySumCircular_dynamic_programming(t *testing.T) {
	test(t, maxSubarraySumCircular_dynamic_programming)
}

func TestMaxSubarraySumCircular_enum(t *testing.T) {
	test(t, maxSubarraySumCircular_enum)
}
