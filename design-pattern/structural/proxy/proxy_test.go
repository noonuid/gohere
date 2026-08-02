package proxy

import "testing"

func TestProxy(t *testing.T) {
	var subject Subject = NewProxy()

	result := subject.Do()
	if result != "PreDo Do PostDo" {
		t.Fail()
	}
}
