package problem0378

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(matrix [][]int, k int) int) {
	// 测试用例。
	testCases := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{
			matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
			k:      8,
			want:   13,
		},
		{
			matrix: [][]int{{-5}},
			k:      1,
			want:   -5,
		},
		{
			matrix: [][]int{{-5, -4}, {-5, -4}},
			k:      2,
			want:   -5,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.matrix, testCase.k)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.k)
		}
	}
}

func TestKthSmallest(t *testing.T) {
	testFramework(t, kthSmallest)
}
