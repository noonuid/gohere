package hj077

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	trains []int
	want   []string
}{
	{
		trains: []int{1, 2, 3},
		want: []string{
			"1 2 3",
			"1 3 2",
			"2 1 3",
			"2 3 1",
			"3 2 1",
		},
	},
	// {
	// 	trains: []int{1, 2, 3, 4},
	// 	want:   []string{},
	// },
	// {
	// 	trains: []int{6, 1, 5, 3, 2, 7, 4},
	// 	want:   []string{},
	// },
}

func TestTrainsPullsInStaion(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var sequences = [][]int{}
		trainsPullsInStaion(&sequences, []int{}, []int{}, testCase.trains)
		var got = sortResult(sequences)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
