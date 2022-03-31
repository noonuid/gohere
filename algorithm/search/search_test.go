package search

import (
	"testing"
)

type searchTest struct {
	input    []int
	key      int
	expected int
	name     string
}

// Note that these are immutable therefore they are shared among all the search tests.
// If your algorithm is mutating these then it is advisable to create separate test cases.
var searchTests = []searchTest{
	//Sanity
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, "Sanity"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, "Empty"},
}

func TestSequentialSearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue := SequentialSearch(test.input, test.key)

		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'",
				test.name, test.input, test.key, test.expected, actualValue)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue := BinarySearch(test.input, test.key)

		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'",
				test.name, test.input, test.key, test.expected, actualValue)
		}
	}
}

func TestInterpolationSearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue := InterpolationSearch(test.input, test.key)

		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'",
				test.name, test.input, test.key, test.expected, actualValue)
		}
	}
}
