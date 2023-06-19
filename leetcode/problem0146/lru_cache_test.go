package problem0146

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, constructor func(capacity int) LRUCache) {
	// 测试用例
	testCases := []struct {
		operations []string
		params     [][]int
		want       []int
	}{
		{
			operations: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			params:     [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			want:       []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4},
		},
		{
			operations: []string{"LRUCache", "put", "get", "put", "get", "get"},
			params:     [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}},
			want:       []int{0, 0, 1, 0, -1, 2},
		},
		{
			operations: []string{"LRUCache", "put", "put", "put", "put", "get", "get"},
			params:     [][]int{{2}, {2, 1}, {1, 1}, {2, 3}, {4, 1}, {1}, {2}},
			want:       []int{0, 0, 0, 0, 0, -1, 3},
		},
		{
			operations: []string{"LRUCache", "put", "put", "get", "get", "put", "get", "get", "get"},
			params:     [][]int{{2}, {2, 1}, {3, 2}, {3}, {2}, {4, 3}, {2}, {3}, {4}},
			want:       []int{0, 0, 0, 2, 1, 0, 1, -1, 3},
		},
	}

	for caseIndex, testCase := range testCases {
		obj := constructor(testCase.params[0][0])
		got := make([]int, len(testCase.want))
		for index, operation := range testCase.operations {
			switch operation {
			case "put":
				obj.Put(testCase.params[index][0], testCase.params[index][1])
				got[index] = 0
			case "get":
				got[index] = obj.Get(testCase.params[index][0])
			default:
				got[index] = 0
			}
		}
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLRUCache(t *testing.T) {
	testFramework(t, Constructor)
}
