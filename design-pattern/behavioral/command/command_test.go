package command

import (
	"reflect"
	"testing"
)

func TestCommand(t *testing.T) {
	var receiver *Receiver = &Receiver{}
	var cmdA Command = &CommandA{receiver: receiver}
	var cmdB Command = &CommandB{receiver: receiver}
	var invoker *Invoker = &Invoker{CmdList: []Command{cmdA, cmdB}}
	result := invoker.Invoke()
	if !reflect.DeepEqual(result, []string{"Receiver.MethodA", "Receiver.MethodB"}) {
		t.Fail()
	}
}
