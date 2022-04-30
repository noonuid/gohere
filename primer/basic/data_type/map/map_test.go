package map_

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func Example_map() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	// Output:
	// {40.68433 -74.39967}
}

func Example_map_literals() {
	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	{
		fmt.Println(m)
	}
	// Output:
	// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}

func Example_mutating_maps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Output:
	// The value: 42
	// The value: 48
	// The value: 0
	// The value: 0 Present? false
}
