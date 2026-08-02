package command

type Receiver struct{}

func (r *Receiver) MethodA() string {
	return "Receiver.MethodA"
}

func (r *Receiver) MethodB() string {
	return "Receiver.MethodB"
}

type Command interface {
	Execute() string
}

type CommandA struct {
	receiver *Receiver
}

func (c *CommandA) Execute() string {
	return c.receiver.MethodA()
}

type CommandB struct {
	receiver *Receiver
}

func (c *CommandB) Execute() string {
	return c.receiver.MethodB()
}

type Invoker struct {
	CmdList []Command
}

func (i *Invoker) Invoke() []string {
	result := []string{}
	for _, cmd := range i.CmdList {
		result = append(result, cmd.Execute())
	}
	return result
}
