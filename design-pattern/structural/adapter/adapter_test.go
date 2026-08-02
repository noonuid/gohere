package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var adaptee *Adaptee = &Adaptee{}
	var adapter Target = NewAdapter(adaptee)
	result := adapter.Request()
	if result != "Adaptee" {
		t.Fail()
	}
}
