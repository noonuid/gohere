package problem0621

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	tasks []byte
	n     int
	want  int
}{
	{
		tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
		n:     2,
		want:  8,
	},
	{
		tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
		n:     0,
		want:  6,
	},
	{
		tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
		n:     2,
		want:  16,
	},
}

func TestLeastInterval(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		var got = leastInterval(testCase.tasks, testCase.n)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
