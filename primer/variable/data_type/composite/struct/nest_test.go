package struct_

import "fmt"

func Example_struct_nest() {
	type Person struct {
		Name string
		Age  int
	}

	type Student struct {
		Person Person
		Score  float64
	}

	student := Student{
		Person{
			"Alice",
			20,
		},
		100,
	}
	fmt.Println(student)

	// Output: {{Alice 20} 100}
}

func Example_anonymous_field() {
	type Person struct {
		Name string
		Age  int
	}

	type Student struct {
		Person
		Score float64
	}

	student := Student{
		Person: Person{
			Name: "Alice",
			Age:  20,
		},
		Score: 100,
	}
	fmt.Println(student)
	fmt.Println(student.Person.Name, student.Person.Age, student.Score)
	fmt.Println(student.Name, student.Age, student.Score)

	// Output:
	// {{Alice 20} 100}
	// Alice 20 100
	// Alice 20 100
}
