package template_method

import "testing"

func TestConcreteA(t *testing.T) {
	var c *ConcreteA = NewConcreteA()
	result := c.TemplateMethod()
	if result != "ConcreteA.MethodA ConcreteA.MethodB" {
		t.Fail()
	}
}

func TestConcreteB(t *testing.T) {
	var c *ConcreteB = NewConcreteB()
	result := c.TemplateMethod()
	if result != "ConcreteB.MethodA ConcreteB.MethodB" {
		t.Fail()
	}
}
