package simple_factory

type Fruit interface {
	GetName() string
}

func NewFruit(name string) Fruit {
	if name == "Apple" {
		return &apple{name: name}
	} else if name == "Banana" {
		return &banana{name: name}
	} else if name == "Pear" {
		return &pear{name: name}
	}
	return nil
}

type apple struct {
	Fruit
	name string
}

func (a *apple) GetName() string {
	return a.name
}

type banana struct {
	Fruit
	name string
}

func (b *banana) GetName() string {
	return b.name
}

type pear struct {
	Fruit
	name string
}

func (p *pear) GetName() string {
	return p.name
}
