package problem0200

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(grid [][]byte) int) {
	// 测试用例。
	testCases := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
			want: 1,
		},
		{
			grid: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
			want: 3,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.grid)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestNumIslands_bfs(t *testing.T) {
	testFramework(t, numIslands_bfs)
}

func TestNumIslands_dfs(t *testing.T) {
	testFramework(t, numIslands_dfs)
}
