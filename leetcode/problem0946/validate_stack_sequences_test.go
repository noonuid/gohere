package problem0946

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	pushed []int
	popped []int
	want   bool
}{
	{
		pushed: []int{1, 2, 3, 4, 5},
		popped: []int{4, 5, 3, 2, 1},
		want:   true,
	},
	{
		pushed: []int{1, 2, 3, 4, 5},
		popped: []int{4, 3, 5, 1, 2},
		want:   false,
	},
}

func TestChorus(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		var got = validateStackSequences(testCase.pushed, testCase.popped)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
