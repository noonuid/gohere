package observer

import "testing"

func TestObserver(t *testing.T) {
	var subject *Subject = NewSubject()
	var observerOne = &ObserverA{State: "StateA"}
	var observerTwo = &ObserverA{State: "StateA"}

	subject.Attach(observerOne)
	subject.Attach(observerTwo)

	subject.UpdateCtx("StateB")

	if !(observerOne.State == "StateB" && observerTwo.State == "StateB") {
		t.Fatal("StateB")
	}

	subject.Detach(observerTwo)

	subject.UpdateCtx("StateC")

	if !(observerOne.State == "StateC" && observerTwo.State == "StateB") {
		t.Fatal("StateC")
	}
}
