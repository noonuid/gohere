package problem0763

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string) []int) {
	// 测试用例。
	testCases := []struct {
		s    string
		want []int
	}{
		{
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			s:    "eccbbbbdec",
			want: []int{10},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.s)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestPartitionLabels(t *testing.T) {
	testFramework(t, partitionLabels)
}
