package adapter

type Target interface {
	Request() string
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee"
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(a *Adaptee) *Adapter {
	return &Adapter{adaptee: a}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
