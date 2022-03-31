package hj001

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	str  string
	want int
}{
	{
		str:  "hello nowcoder",
		want: 8,
	},
}

func TestLengthOfTheLastWord(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		got := lengthOfTheLastWord(testCase.str)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
