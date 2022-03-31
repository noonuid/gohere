package hj024

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	heights []int
	want    int
}{
	{
		heights: []int{186, 186, 150, 200, 160, 130, 197, 200},
		want:    4,
	},
	{
		heights: []int{1, 2, 30, 8, 9, 20, 9, 8},
		want:    1,
	},
}

func TestChorus(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		var got = chorus(testCase.heights)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
