package stack

import (
	"reflect"
	"testing"
)

func TestStackArray(t *testing.T) {
	stack := NewStackArray()

	start, end := 0, 10

	for i := start; i < end; i++ {
		stack.Push(i)

		t.Run("Test Length()", func(t *testing.T) {
			want := i + 1
			got := stack.Length()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}

	t.Run("Test IsEmpty()", func(t *testing.T) {
		want := false
		got := stack.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	for i := end - 1; i >= start; i-- {
		t.Run("Test Pop()", func(t *testing.T) {
			want := i
			got := stack.Pop()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}

	t.Run("Test IsEmpty()", func(t *testing.T) {
		want := true
		got := stack.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
