package problem0131

import (
	"reflect"
	"sort"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string) [][]string) {
	// 测试用例。
	testCases := []struct {
		s    string
		want [][]string
	}{
		{
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			s:    "a",
			want: [][]string{{"a"}},
		},
	}

	// 将答案按照升序排序。
	sortAnswer := func(answer [][]string) {
		sort.Slice(answer, func(i, j int) bool {
			answerI, answerJ := answer[i], answer[j]
			lenI, lenJ := len(answerI), len(answerJ)
			for m, n := 0, 0; m < lenI && n < lenJ; m, n = m+1, n+1 {
				if answerI[m] > answerJ[n] {
					return false
				}
			}
			return lenI <= lenJ
		})
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.s)
		sortAnswer(got)
		sortAnswer(testCase.want)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestPartition(t *testing.T) {
	testFramework(t, partition)
}

func TestPartition_backtrack_dp(t *testing.T) {
	testFramework(t, partition_backtrack_dp)
}
