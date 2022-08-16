package variable

import "fmt"

func Example_variable() {
	var c, python, java bool

	{
		var i int
		fmt.Println(i, c, python, java)
	}
	// OUtput:
	// 0 false false false
}

func Example_initialize_variable() {
	var i, j int = 1, 2

	{
		var c, python, java = true, false, "no!"
		fmt.Println(i, j, c, python, java)
	}
	// OUtput:
	// 1 2 true false no!
}

func Example_short_variable_declaration() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	// OUtput:
	// 1 2 3 true false no!
}

const Pi = 3.14

func Example_constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	// Output:
	// Hello 世界
	// Happy 3.14 Day
	// Go rules? true
}

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func Example_numeric_constants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// Output:
	// 21
	// 0.2
	// 1.2676506002282295e+29
}
