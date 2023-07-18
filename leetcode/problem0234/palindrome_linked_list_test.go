package problem0234

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/nowcoder/structure"
)

func testFramework(t *testing.T, testFn func(*ListNode) bool) {
	// 测试用例
	testCases := []struct {
		head *ListNode
		want bool
	}{
		{
			head: structure.Ints2List([]int{1, 2, 2, 1}),
			want: true,
		},
		{
			head: structure.Ints2List([]int{1, 2}),
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFn(testCase.head)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestIsPalindrome_reverse(t *testing.T) {
	testFramework(t, isPalindrome_reverse)
}

func TestIsPalindrome_copy(t *testing.T) {
	testFramework(t, isPalindrome_copy)
}
