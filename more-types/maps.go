package main

import "fmt"

type Coordinate struct {
	Lat, Long float64
}

var m1 map[string]Coordinate

func run_maps_examples() {
	m1 = make(map[string]Coordinate) // this initialises the map, ready for use
	m1["Bell Labs"] = Coordinate{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])

	// Map literals are a bit like struct literals, but you still need the keys
	var m2 = map[string]Coordinate{
		"Bell Labs": Coordinate{
			40.68433, -74.39967,
		},
		"Google": Coordinate{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// We can omit the type here
	var m3 = map[string]Coordinate{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m3)

	// Mutating maps like so:
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v, ok := m4["Answer"] // check if element is in map with second return value
	fmt.Println("The value:", v, "Present?", ok)
}
