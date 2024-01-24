package problem0412

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(n int) []string) {
	// 测试用例。
	testCases := []struct {
		n    int
		want []string
	}{
		{
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.n)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestFizzBuzz_string_concatenation(t *testing.T) {
	testFramework(t, fizzBuzz_string_concatenation)
}

func TestFizzBuzz(t *testing.T) {
	testFramework(t, fizzBuzz)
}
