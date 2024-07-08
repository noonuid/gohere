package problem0045

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
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
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

func TestJump(t *testing.T) {
	test(t, jump)
}

func TestJump_enum(t *testing.T) {
	test(t, jump_enum)
}
