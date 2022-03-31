package linkedlist

import (
	"reflect"
	"testing"
)

func TestSingly(t *testing.T) {
	list := NewSinglyLinkedList()
	list.AddAtBeg(1)
	list.AddAtBeg(2)
	list.AddAtBeg(3)

	t.Run("Test AddAtBegin()", func(t *testing.T) {
		want := []interface{}{3, 2, 1}
		got := []interface{}{}
		current := list.head
		got = append(got, current.data)
		for current.next != nil {
			current = current.next
			got = append(got, current.data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.AddAtEnd(4)

	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []interface{}{3, 2, 1, 4}
		got := []interface{}{}
		current := list.head
		got = append(got, current.data)
		for current.next != nil {
			current = current.next
			got = append(got, current.data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DeleteAtBegin()", func(t *testing.T) {
		want := interface{}(3)
		got := list.DelAtBeg()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DeleteAtEnd()", func(t *testing.T) {
		want := interface{}(4)
		got := list.DelAtEnd()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Count()", func(t *testing.T) {
		want := 2
		got := list.Count()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list2 := SinglyLinkedList{}
	list2.AddAtBeg(1)
	list2.AddAtBeg(2)
	list2.AddAtBeg(3)
	list2.AddAtBeg(4)
	list2.AddAtBeg(5)
	list2.AddAtBeg(6)

	t.Run("Test Reverse()", func(t *testing.T) {
		want := []interface{}{1, 2, 3, 4, 5, 6}
		got := []interface{}{}
		list2.Reverse()
		current := list2.head
		got = append(got, current.data)
		for current.next != nil {
			current = current.next
			got = append(got, current.data)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
