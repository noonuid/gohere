package problem0039

import (
	"reflect"
	"sort"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, int) [][]int) {
	// 测试用例。
	testCases := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.candidates, testCase.target)

		// 将答案按照升序排序。
		for i := 0; i < len(got); i++ {
			sort.Ints(got[i])
		}
		sort.Slice(got, func(i, j int) bool {
			intsI, intsJ := got[i], got[j]
			m, lenI, lenJ := 0, len(intsI), len(intsJ)
			for m < lenI && m < lenJ {
				if intsI[m] > intsJ[m] {
					return false
				}
				m++
			}
			return true
		})

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestCombinationSum(t *testing.T) {
	testFramework(t, combinationSum)
}
