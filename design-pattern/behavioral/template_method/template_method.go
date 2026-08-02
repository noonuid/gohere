package template_method

type Abstract interface {
	MethodA() string
	MethodB() string
}

type templete struct {
	a Abstract
}

func (t *templete) TemplateMethod() string {
	resultA := t.a.MethodA()
	resultB := t.a.MethodB()
	return resultA + " " + resultB
}

type ConcreteA struct {
	templete
}

func NewConcreteA() *ConcreteA {
	c := &ConcreteA{}
	c.a = c
	return c
}

func (c *ConcreteA) MethodA() string {
	return "ConcreteA.MethodA"
}

func (c *ConcreteA) MethodB() string {
	return "ConcreteA.MethodB"
}

type ConcreteB struct {
	templete
}

func NewConcreteB() *ConcreteB {
	c := &ConcreteB{}
	c.a = c
	return c
}

func (c *ConcreteB) MethodA() string {
	return "ConcreteB.MethodA"
}

func (c *ConcreteB) MethodB() string {
	return "ConcreteB.MethodB"
}
