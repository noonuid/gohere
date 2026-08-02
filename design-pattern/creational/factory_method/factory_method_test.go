package factory_method

import "testing"

func TestFruit(t *testing.T) {
	var factory FruitFactory

	factory = new(AppleFactory)
	var a Fruit = factory.NewFruit("Apple")
	_, ok := a.(*apple)
	if !ok || a.GetName() != "Apple" {
		t.Fatal("Apple")
	}

	factory = new(BananaFactory)
	var b Fruit = factory.NewFruit("Banana")
	_, ok = b.(*banana)
	if !ok || b.GetName() != "Banana" {
		t.Fatal("Banana")
	}

	factory = new(PearFactory)
	var p Fruit = factory.NewFruit("Pear")
	_, ok = p.(*pear)
	if !ok || p.GetName() != "Pear" {
		t.Fatal("Pear")
	}
}
