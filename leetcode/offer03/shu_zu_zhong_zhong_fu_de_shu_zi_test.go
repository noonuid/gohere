package offer03

import (
	"testing"

	"golang.org/x/exp/slices"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 3, 1, 0, 2, 5, 3},
			want: []int{2, 3},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums)
		if !slices.Contains(testCase.want, got) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFindRepeatNumber_hash(t *testing.T) {
	testFramework(t, findRepeatNumber_hash)
}

func TestFindRepeatNumber_brute(t *testing.T) {
	testFramework(t, findRepeatNumber_brute)
}
