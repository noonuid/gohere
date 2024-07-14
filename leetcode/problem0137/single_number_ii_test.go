package problem0137

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
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		{
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
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

func TestKSingleNumber_bitwise(t *testing.T) {
	test(t, singleNumber_bitwise)
}

func TestKSingleNumber_sort(t *testing.T) {
	test(t, singleNumber_sort)
}
