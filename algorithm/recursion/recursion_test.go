package recursion

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	var want = []interface{}{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	var got = []interface{}{}
	for i := 0; i < 11; i++ {
		got = append(got, Factorial(i))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestFactorialIterative(t *testing.T) {
	var want = []interface{}{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	var got = []interface{}{}
	for i := 0; i < 11; i++ {
		got = append(got, FactorialIterative(i))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestFibonacci(t *testing.T) {
	var want = []interface{}{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	var got = []interface{}{}
	for i := 0; i < 11; i++ {
		got = append(got, Fibonacci(i))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestFibonacciMemorization(t *testing.T) {
	var want = []interface{}{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	var got = []interface{}{}
	for i := 0; i < 11; i++ {
		got = append(got, FibonacciMemorization(i))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestFibonacciIterative(t *testing.T) {
	var want = []interface{}{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	var got = []interface{}{}
	for i := 0; i < 11; i++ {
		got = append(got, FibonacciIterative(i))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
