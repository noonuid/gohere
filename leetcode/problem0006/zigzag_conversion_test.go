package problem0006

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(s string, numRows int) string) {
	testCases := []struct {
		s       string
		numRows int
		want    string
	}{
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			s:       "A",
			numRows: 1,
			want:    "A",
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.s, testCase.numRows)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestConvert(t *testing.T) {
	test(t, convert)
}

func TestConvert_simulation_optimized_space(t *testing.T) {
	test(t, convert_simulation_optimized_space)
}

func TestConvert_simulation(t *testing.T) {
	test(t, convert_simulation)
}
