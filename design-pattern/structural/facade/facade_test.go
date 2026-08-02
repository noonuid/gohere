package facade

import "testing"

func TestMethodOne(t *testing.T) {
	f := NewFacade()
	one := f.MethodOne()
	if one != "SubSystemA.MethodA SubSystemB.MethodB" {
		t.Fatal()
	}
}

func TestMethodTwo(t *testing.T) {
	f := NewFacade()
	two := f.MethodTwo()
	if two != "SubSystemC.MethodC SubSystemD.MethodD" {
		t.Fatal()
	}
}
