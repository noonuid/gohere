package problem0050

import (
	"math"
	"testing"
)

func test(t *testing.T, fn func(x float64, n int) float64) {
	testCases := []struct {
		x    float64
		n    int
		want float64
	}{
		{
			x:    2.00000,
			n:    10,
			want: 1024.00000,
		},
		{
			x:    2.10000,
			n:    3,
			want: 9.26100,
		},
		{
			x:    2.00000,
			n:    -2,
			want: 0.25000,
		},
	}

	// 允许的误差范围（epsilon）。
	const epsilon = 1e-5

	for caseIndex, testCase := range testCases {
		got := fn(testCase.x, testCase.n)

		if math.Abs(testCase.want-got) > epsilon {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestMyPow(t *testing.T) {
	test(t, myPow)
}

func TestMyPow_iterate(t *testing.T) {
	test(t, myPow_iterate)
}
