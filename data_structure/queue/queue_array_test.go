package queue

import (
	"reflect"
	"testing"
)

func TestQueueArray(t *testing.T) {
	queue := NewQueueArray()

	start, end := 0, 10

	for i := start; i < end; i++ {
		queue.EnQueue(i)

		t.Run("Test Length()", func(t *testing.T) {
			want := i + 1
			got := queue.Length()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}

	t.Run("Test IsEmpty()", func(t *testing.T) {
		want := false
		got := queue.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	for i := start; i < end; i++ {
		t.Run("Test DeQueue()", func(t *testing.T) {
			want := i
			got := queue.DeQueue()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}

	t.Run("Test DeQueue() when queue is empty", func(t *testing.T) {
		var want interface{} = nil
		got := queue.DeQueue()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test IsEmpty()", func(t *testing.T) {
		want := true
		got := queue.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
