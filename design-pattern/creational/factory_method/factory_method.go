package factory_method

// 抽象的产品。

type Fruit interface {
	GetName() string
}

// 具体的产品。

type apple struct {
	name string
}

func (a *apple) GetName() string {
	return a.name
}

type banana struct {
	name string
}

func (b *banana) GetName() string {
	return b.name
}

type pear struct {
	name string
}

func (p *pear) GetName() string {
	return p.name
}

// 抽象的工厂。

type FruitFactory interface {
	NewFruit(name string) Fruit
}

// 具体的工厂。

type AppleFactory struct{}

func (f *AppleFactory) NewFruit(name string) Fruit {
	return &apple{name: name}
}

type BananaFactory struct{}

func (f *BananaFactory) NewFruit(name string) Fruit {
	return &banana{name: name}
}

type PearFactory struct{}

func (f *PearFactory) NewFruit(name string) Fruit {
	return &pear{name: name}
}
