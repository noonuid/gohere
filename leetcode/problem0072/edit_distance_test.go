package problem0072

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(word1 string, word2 string) int) {
	// 测试用例。
	testCases := []struct {
		word1 string
		word2 string
		want  int
	}{
		{
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			word1: "intention",
			word2: "execution",
			want:  5,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.word1, testCase.word2)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMinDistance(t *testing.T) {
	testFramework(t, minDistance)
}
