package problem0399

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(equations [][]string, values []float64, queries [][]string) []float64) {
	// 测试用例。
	testCases := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			want:      []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			want:      []float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			want:      []float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.equations, testCase.values, testCase.queries)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCalcEquation(t *testing.T) {
	testFramework(t, calcEquation)
}
