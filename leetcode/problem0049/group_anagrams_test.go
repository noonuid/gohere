package problem0049

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]string) [][]string) {
	// 测试用例。
	testCases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for caseIndex, testCase := range testCases {
		// 将答案按照升序排列。
		sortAnswer := func(answer [][]string) {
			for i := 0; i < len(answer); i++ {
				sort.Strings(answer[i])
			}
			sort.Slice(answer, func(i, j int) bool {
				strI, strJ := strings.Join(answer[i], ""), strings.Join(answer[j], "")
				return strI < strJ
			})
		}
		// 被测方法的返回值。
		got := testFunc(testCase.strs)
		sortAnswer(got)
		sortAnswer(testCase.want)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestGroupAnagrams_map(t *testing.T) {
	testFramework(t, groupAnagrams_map)
}

func TestGroupAnagrams_brute(t *testing.T) {
	testFramework(t, groupAnagrams_brute)
}
