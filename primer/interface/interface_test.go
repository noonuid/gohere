package interface_

import (
	"fmt"
	"math"
)

// 接口类型 是由一组方法签名定义的集合。
// 接口类型的变量可以保存任何实现了这些方法的值。

type Abser interface {
	Abs() float64
}

func Example_interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v

	fmt.Println(a.Abs())
	// Output:
	// 5
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func Example_interfaces_are_implimented_implicitly() {
	var i I = T{"hello"}
	i.M()
	// Output:
	// hello
}
