package observer

type Subject struct {
	observers []Observer
	ctx       string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Detach(o Observer) {
	for i := 0; i < len(s.observers); i++ {
		if s.observers[i] == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateCtx(ctx string) {
	s.ctx = ctx
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type ObserverA struct {
	State string
}

func (o *ObserverA) Update(s *Subject) {
	o.State = s.ctx
}
