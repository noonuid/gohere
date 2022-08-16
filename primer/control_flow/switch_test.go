package control_flow

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Example_switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Output:
	// Go runs on OS X.
}

func Test_switch_evaluation_order(t *testing.T) {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func Test_with_no_condition(t *testing.T) {
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
