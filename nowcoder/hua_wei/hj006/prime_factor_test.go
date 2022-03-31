package hj006

import (
	"testing"
)

var testCases = []struct {
	n int
}{
	{
		n: 180,
	},
	{
		n: 2000000014,
	},
}

func TestPrimeFactor(t *testing.T) {
	for _, testCase := range testCases {
		PrimeFactor(testCase.n)
	}
}

func ExamplePrimeFactor() {
	for _, testCase := range testCases {
		PrimeFactor(testCase.n)
	}
	// Output:
	// 2 2 3 3 5
	// 2 1000000007
}
