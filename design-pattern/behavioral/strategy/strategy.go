package strategy

type Strategy interface {
	Method() string
}

type StrategyA struct{}

func (s *StrategyA) Method() string {
	return "StrategyA"
}

type StrategyB struct{}

func (s *StrategyB) Method() string {
	return "StrategyB"
}

type Context struct {
	s Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.s = s
}

func (c *Context) Solve() string {
	return c.s.Method()
}
