package facade

import "fmt"

type SubSystemA struct{}

func (s *SubSystemA) MethodA() string {
	return "SubSystemA.MethodA"
}

type SubSystemB struct{}

func (s *SubSystemB) MethodB() string {
	return "SubSystemB.MethodB"
}

type SubSystemC struct{}

func (s *SubSystemC) MethodC() string {
	return "SubSystemC.MethodC"
}

type SubSystemD struct{}

func (s *SubSystemD) MethodD() string {
	return "SubSystemD.MethodD"
}

type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func NewFacade() *Facade {
	return &Facade{
		a: &SubSystemA{}, b: &SubSystemB{}, c: &SubSystemC{}, d: &SubSystemD{},
	}
}

func (f *Facade) MethodOne() string {
	return fmt.Sprintf("%s %s", f.a.MethodA(), f.b.MethodB())
}

func (f *Facade) MethodTwo() string {
	return fmt.Sprintf("%s %s", f.c.MethodC(), f.d.MethodD())
}
