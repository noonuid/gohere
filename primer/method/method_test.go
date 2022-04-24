package method

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.

// A method is a function with a special receiver argument.

// The receiver appears in its own argument list between the func keyword and the method name.

// In this example, the Abs method has a receiver of type Vertex named v.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Example_method() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// Output:
	// 5
}

// Remember: a method is just a function with a receiver argument.

// Here's Abs written as a regular function with no change in functionality.
func Example_method_is_just_a_function() {
	var Abs = func(v Vertex) float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	{
		v := Vertex{3, 4}
		fmt.Println(Abs(v))
	}
	// Output:
	// 5
}

// You can declare a method on non-struct types, too.

// In this example we see a numeric type MyFloat with an Abs method.

// You can only declare a method with a receiver whose type is defined in the
// same package as the method. You cannot declare a method with a receiver whose
// type is defined in another package (which includes the built-in types such as int).

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Example_method_on_non_struct_types() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	// Output:
	// 1.4142135623730951
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Example_pointer_receivers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	// Output:
	// 50
}

// Here we see the Abs and Scale methods rewritten as functions.
func Example_pointers_and_functions() {
	var Abs = func(v Vertex) float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	var Scale = func(v *Vertex, f float64) {
		v.X = v.X * f
		v.Y = v.Y * f
	}

	{
		v := Vertex{3, 4}
		Scale(&v, 10)
		fmt.Println(Abs(v))
	}
	// Output:
	// 50
}
