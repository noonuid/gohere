package simple_factory

import (
	"testing"
)

func TestApple(t *testing.T) {
	fruit := NewFruit("Apple")
	_, ok := fruit.(*apple)
	if !ok || fruit.GetName() != "Apple" {
		t.Fail()
	}
}

func TestBanana(t *testing.T) {
	fruit := NewFruit("Banana")
	_, ok := fruit.(*banana)
	if !ok || fruit.GetName() != "Banana" {
		t.Fail()
	}
}

func TestPear(t *testing.T) {
	fruit := NewFruit("Pear")
	_, ok := fruit.(*pear)
	if !ok || fruit.GetName() != "Pear" {
		t.Fail()
	}
}

func TestNil(t *testing.T) {
	fruit := NewFruit("Dog")
	if fruit != nil {
		t.Fail()
	}
}
