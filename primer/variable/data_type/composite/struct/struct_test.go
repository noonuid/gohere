package struct_

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Example_struct() {
	fmt.Println(Vertex{1, 2})

	// Output:
	// {1 2}
}

func Example_struct_fields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Output:
	// 4
}

// 结构体字段可以通过结构体指针来访问。
// 如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，
// 所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
func Example_pointers_to_struct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// Output:
	// {1000000000 2}
}

func Example_struct_literals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	{
		fmt.Println(v1, p, v2, v3)
	}
	// Output:
	// {1 2} &{1 2} {1 0} {0 0}
}
