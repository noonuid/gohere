package sequentiallist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var list = New(20)

	want := []interface{}{}
	got := []interface{}{}
	for i := 0; i < list.length; i++ {
		got = append(got, (*list.data)[i])
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestIsEmpty(t *testing.T) {
	var list = New(20)

	want := true
	got := list.IsEmpty()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	list.Insert(0, 1)

	want = false
	got = list.IsEmpty()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestIsFull(t *testing.T) {
	var list = New(2)

	want := false
	got := list.IsFull()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	list.Insert(0, 1)
	list.Insert(1, 2)

	want = true
	got = list.IsFull()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestInsert(t *testing.T) {
	var list = New(20)

	// {1, 2, 3}
	list.Insert(0, 1)
	list.Insert(1, 3)
	list.Insert(1, 2)

	want := []interface{}{1, 2, 3}
	got := []interface{}{}
	for i := 0; i < list.length; i++ {
		got = append(got, (*list.data)[i])
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGet(t *testing.T) {
	var list = New(20)

	list.Insert(0, 1)
	list.Insert(1, 3)
	list.Insert(1, 2)

	want := 2
	got, _ := list.Get(1)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGetAll(t *testing.T) {
	var list = New(20)

	list.Insert(0, 1)
	list.Insert(1, 3)
	list.Insert(1, 2)

	want := []interface{}{1, 2, 3}
	got := list.GetAll()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	var list = New(20)

	list.Insert(0, 1)
	list.Insert(1, 3)
	list.Insert(1, 2)

	list.Remove(2)

	want := []interface{}{1, 2}
	got := list.GetAll()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSequentialList(t *testing.T) {
	var list = New(20)

	t.Run("NewSequentialList()", func(t *testing.T) {
		want := []interface{}{}
		got := []interface{}{}
		for i := 0; i < list.length; i++ {
			got = append(got, (*list.data)[i])
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("IsEmpty()", func(t *testing.T) {
		want := true
		got := list.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.Insert(0, 1)
	list.Insert(1, 3)
	list.Insert(1, 2)

	t.Run("Insert()", func(t *testing.T) {
		want := []interface{}{1, 2, 3}
		got := []interface{}{}
		for i := 0; i < list.length; i++ {
			got = append(got, (*list.data)[i])
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("IsEmpty()", func(t *testing.T) {
		want := false
		got := list.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("IsFull()", func(t *testing.T) {
		want := false
		got := list.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Get()", func(t *testing.T) {
		want := 2
		got, _ := list.Get(1)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("GetAll()", func(t *testing.T) {
		want := []interface{}{1, 2, 3}
		got := list.GetAll()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.Remove(2)

	t.Run("Remove()", func(t *testing.T) {
		want := []interface{}{1, 2}
		got := list.GetAll()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
