package problem0210

import (
	"reflect"
	"slices"
	"testing"
)

func test(t *testing.T, fn func(numCourses int, prerequisites [][]int) []int) {
	testCases := []struct {
		numCourses    int
		prerequisites [][]int
		want          [][]int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          [][]int{{0, 1}},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			want:          [][]int{{0}},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.numCourses, testCase.prerequisites)

		if !slices.ContainsFunc(testCase.want, func(item []int) bool {
			return reflect.DeepEqual(item, got)
		}) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestFindOrder_bfs(t *testing.T) {
	test(t, findOrder_bfs)
}

func TestFindOrder_dfs(t *testing.T) {
	test(t, findOrder_dfs)
}
