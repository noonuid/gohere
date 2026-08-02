package abstract_factory

import "testing"

func TestFruit(t *testing.T) {
	var factory AbstractFactory

	factory = new(ChinaFactory)
	var ca AbstractApple = factory.NewApple("ChinaApple")
	_, ok := ca.(*chinaApple)
	if !ok || ca.GetAppleName() != "ChinaApple" {
		t.Fatal("ChinaApple")
	}
	var cb AbstractBanana = factory.NewBanana("ChinaBanana")
	_, ok = cb.(*chinaBanana)
	if !ok || cb.GetBananaName() != "ChinaBanana" {
		t.Fatal("ChinaBanana")
	}
	var cp AbstractPear = factory.NewPear("ChinaPear")
	_, ok = cp.(*chinaPear)
	if !ok || cp.GetPearName() != "ChinaPear" {
		t.Fatal("ChinaPear")
	}

	factory = new(JapanFactory)
	var ja AbstractApple = factory.NewApple("JapanApple")
	_, ok = ja.(*japanApple)
	if !ok || ja.GetAppleName() != "JapanApple" {
		t.Fatal("JapanApple")
	}
	var jb AbstractBanana = factory.NewBanana("JapanBanana")
	_, ok = jb.(*japanBanana)
	if !ok || jb.GetBananaName() != "JapanBanana" {
		t.Fatal("JapanBanana")
	}
	var jp AbstractPear = factory.NewPear("JapanPear")
	_, ok = jp.(*japanPear)
	if !ok || jp.GetPearName() != "JapanPear" {
		t.Fatal("JapanPear")
	}

	factory = new(AmericanFactory)
	var aa AbstractApple = factory.NewApple("AmericanApple")
	_, ok = aa.(*americanApple)
	if !ok || aa.GetAppleName() != "AmericanApple" {
		t.Fatal("AmericanApple")
	}
	var ab AbstractBanana = factory.NewBanana("AmericanBanana")
	_, ok = ab.(*americanBanana)
	if !ok || ab.GetBananaName() != "AmericanBanana" {
		t.Fatal("AmericanBanana")
	}
	var ap AbstractPear = factory.NewPear("AmericanPear")
	_, ok = ap.(*americanPear)
	if !ok || ap.GetPearName() != "AmericanPear" {
		t.Fatal("AmericanPear")
	}
}
