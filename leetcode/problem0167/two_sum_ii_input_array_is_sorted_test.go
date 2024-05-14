package problem0167

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(numbers []int, target int) []int) {
	testCases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.numbers, testCase.target)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	testFramework(t, twoSum)
}

func TestTwoSum_binary(t *testing.T) {
	testFramework(t, twoSum_binary)
}
