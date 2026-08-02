package abstract_factory

// 抽象的产品。

type AbstractApple interface {
	GetAppleName() string
}

type AbstractBanana interface {
	GetBananaName() string
}

type AbstractPear interface {
	GetPearName() string
}

// 抽象的工厂。

type AbstractFactory interface {
	NewApple(name string) AbstractApple
	NewBanana(name string) AbstractBanana
	NewPear(name string) AbstractPear
}

// 具体的产品。

type chinaApple struct {
	name string
}

func (a *chinaApple) GetAppleName() string {
	return a.name
}

type chinaBanana struct {
	name string
}

func (b *chinaBanana) GetBananaName() string {
	return b.name
}

type chinaPear struct {
	name string
}

func (p *chinaPear) GetPearName() string {
	return p.name
}

type japanApple struct {
	name string
}

func (a *japanApple) GetAppleName() string {
	return a.name
}

type japanBanana struct {
	name string
}

func (b *japanBanana) GetBananaName() string {
	return b.name
}

type japanPear struct {
	name string
}

func (p *japanPear) GetPearName() string {
	return p.name
}

type americanApple struct {
	name string
}

func (a *americanApple) GetAppleName() string {
	return a.name
}

type americanBanana struct {
	name string
}

func (b *americanBanana) GetBananaName() string {
	return b.name
}

type americanPear struct {
	name string
}

func (p *americanPear) GetPearName() string {
	return p.name
}

// 具体的工厂。

type ChinaFactory struct{}

func (f *ChinaFactory) NewApple(name string) AbstractApple {
	return &chinaApple{name: name}
}

func (f *ChinaFactory) NewBanana(name string) AbstractBanana {
	return &chinaBanana{name: name}
}

func (f *ChinaFactory) NewPear(name string) AbstractPear {
	return &chinaPear{name: name}
}

type JapanFactory struct{}

func (f *JapanFactory) NewApple(name string) AbstractApple {
	return &japanApple{name: name}
}

func (f *JapanFactory) NewBanana(name string) AbstractBanana {
	return &japanBanana{name: name}
}

func (f *JapanFactory) NewPear(name string) AbstractPear {
	return &japanPear{name: name}
}

type AmericanFactory struct{}

func (f *AmericanFactory) NewApple(name string) AbstractApple {
	return &americanApple{name: name}
}

func (f *AmericanFactory) NewBanana(name string) AbstractBanana {
	return &americanBanana{name: name}
}

func (f *AmericanFactory) NewPear(name string) AbstractPear {
	return &americanPear{name: name}
}
