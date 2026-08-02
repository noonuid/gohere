package problem0216

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func testFramework(t *testing.T, testFunc func(k int, n int) [][]int) {
	// 测试用例。
	testCases := []struct {
		k    int
		n    int
		want [][]int
	}{
		{
			k:    3,
			n:    7,
			want: [][]int{{1, 2, 4}},
		},
		{
			k:    3,
			n:    9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			k:    4,
			n:    1,
			want: [][]int{},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.k, testCase.n)

		// 对返回值与期望值排序，方便后面判断二者是否相等。
		slices.SortFunc(got, func(a, b []int) int {
			for i := 0; i < len(a) && i < len(b); i++ {
				if c := cmp.Compare(a[i], b[i]); c != 0 {
					return c
				}
			}
			return cmp.Compare(len(a), len(b))
		})
		slices.SortFunc(testCase.want, func(a, b []int) int {
			for i := 0; i < len(a) && i < len(b); i++ {
				if c := cmp.Compare(a[i], b[i]); c != 0 {
					return c
				}
			}
			return cmp.Compare(len(a), len(b))
		})

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestCombinationSum3(t *testing.T) {
	testFramework(t, combinationSum3)
}

func TestCombinationSum3_binary(t *testing.T) {
	testFramework(t, combinationSum3_binary)
}
