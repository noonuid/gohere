package problem0063

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(obstacleGrid [][]int) int) {
	testCases := []struct {
		obstacleGrid [][]int
		want         int
	}{
		{
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want:         2,
		},
		{
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			want:         1,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.obstacleGrid)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestUniquePathsWithObstacles_optimization(t *testing.T) {
	test(t, uniquePathsWithObstacles_optimization)
}

func TestUniquePathsWithObstacles(t *testing.T) {
	test(t, uniquePathsWithObstacles)
}
