package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}

	{
		go say("world")
		say("hello")
	}
	// One possible output:
	// hello
	// world
	// world
	// hello
	// hello
	// world
	// hello
	// world
	// hello
	// world
}
