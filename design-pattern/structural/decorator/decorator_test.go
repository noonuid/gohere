package decorator

import "testing"

func TestDecoratorA(t *testing.T) {
	var c Component = &ComponentA{}

	var decoratorA *DecoratorA = NewDecoratorA(c)
	result := decoratorA.Method()
	if result != "ComponentA DecoratorA" {
		t.Fail()
	}
}

func TestDecoratorB(t *testing.T) {
	var c Component = &ComponentA{}

	var decoratorB *DecoratorB = NewDecoratorB(c)
	result := decoratorB.Method()
	if result != "ComponentA DecoratorB" {
		t.Fail()
	}
}
