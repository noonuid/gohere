package problem0070

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	n    int
	want int
}{
	{
		n:    2,
		want: 2,
	},
	{
		n:    3,
		want: 3,
	},
	{
		n:    4,
		want: 5,
	},
}

func TestClimbStairs(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		var got = climbStairs(testCase.n)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestClimbStairsScrolly(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		var got = climbStairsScroll(testCase.n)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
