package strategy

import "testing"

func TestStrategyA(t *testing.T) {
	ctx := &Context{}

	ctx.SetStrategy(&StrategyA{})
	result := ctx.Solve()
	if result != "StrategyA" {
		t.Fail()
	}
}

func TestStrategyB(t *testing.T) {
	ctx := &Context{}

	ctx.SetStrategy(&StrategyB{})
	result := ctx.Solve()
	if result != "StrategyB" {
		t.Fail()
	}
}
