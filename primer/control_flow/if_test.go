package control_flow

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Example_if() {
	fmt.Println(sqrt(2), sqrt(-4))

	// Output:
	// 1.4142135623730951 2i
}

func Example_if_with_a_short_statement() {
	var pow = func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}

	{
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	}

	// Output:
	// 9 20
}

func Example_if_and_else() {
	var pow = func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("%g >= %g\n", v, lim)
		}
		// 这里开始就不能使用 v 了
		return lim
	}

	{
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	}

	// Output:
	// 27 >= 20
	// 9 20
}
