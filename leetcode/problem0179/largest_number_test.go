package problem0179

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) string) {
	testCases := []struct {
		nums []int
		want string
	}{
		{
			nums: []int{10, 2},
			want: "210",
		},
		{
			nums: []int{3, 30, 34, 5, 9},
			want: "9534330",
		},
		{
			nums: []int{432, 43243},
			want: "43243432",
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.nums)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestLargestNumber(t *testing.T) {
	testFramework(t, largestNumber)
}
