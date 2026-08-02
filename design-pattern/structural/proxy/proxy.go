package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (real *RealSubject) Do() string {
	return "Do"
}

type Proxy struct {
	real *RealSubject
}

func NewProxy() *Proxy {
	return &Proxy{real: &RealSubject{}}
}

func (p *Proxy) Do() string {
	var result string

	// 调用真实对象方法之前的工作，如判断权限，检查缓存等。
	result += "PreDo "

	result += p.real.Do()

	// 调用真实对象方法之后的工作，如结果处理，缓存结果等。
	result += " PostDo"

	return result
}
