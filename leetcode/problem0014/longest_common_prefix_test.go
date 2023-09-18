package problem0014

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]string) string) {
	// 测试用例。
	testCases := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{""},
			want: "",
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.strs)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestLongestCommonPrefix_sort(t *testing.T) {
	testFramework(t, longestCommonPrefix_sort)
}

func TestLongestCommonPrefix_column_scan(t *testing.T) {
	testFramework(t, longestCommonPrefix_column_scan)
}
