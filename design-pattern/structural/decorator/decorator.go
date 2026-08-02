package decorator

type Component interface {
	Method() string
}

type ComponentA struct{}

func (c *ComponentA) Method() string {
	return "ComponentA"
}

type DecoratorA struct {
	Component
}

func NewDecoratorA(c Component) *DecoratorA {
	return &DecoratorA{Component: c}
}

func (d *DecoratorA) Method() string {
	return d.Component.Method() + " DecoratorA"
}

type DecoratorB struct {
	Component
}

func NewDecoratorB(c Component) *DecoratorB {
	return &DecoratorB{Component: c}
}

func (d *DecoratorB) Method() string {
	return d.Component.Method() + " DecoratorB"
}
