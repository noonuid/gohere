package problem0406

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([][]int) [][]int) {
	// 测试用例
	testCases := []struct {
		people [][]int
		want   [][]int
	}{
		{
			people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			want:   [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
		{
			people: [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}},
			want:   [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.people)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestReconstructQueue(t *testing.T) {
	testFramework(t, reconstructQueue)
}

func TestReconstructQueue_brute_force(t *testing.T) {
	testFramework(t, reconstructQueue_brute_force)
}
