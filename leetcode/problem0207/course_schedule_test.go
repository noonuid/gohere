package problem0207

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(numCourses int, prerequisites [][]int) bool) {
	// 测试用例。
	testCases := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.numCourses, testCase.prerequisites)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCanFinish_bfs(t *testing.T) {
	testFramework(t, canFinish_bfs)
}

func TestCanFinish_dfs(t *testing.T) {
	testFramework(t, canFinish_dfs)
}
