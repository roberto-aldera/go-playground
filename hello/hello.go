package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")

	rand.Seed(1) // this is the default seed if none is set
	fmt.Println("Random number:", rand.Intn(100))

	fmt.Printf("String formatting says the sqrt(5) is %g .\n", math.Sqrt(5))

	fmt.Println(math.Pi)

	fmt.Println(add(3, 4))
	fmt.Println(another_add(3, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(53))
}

// First function
func add(x int, y int) int {
	return x + y
}

// No need to specify types if they are the same for consecutive input parameters
func another_add(x, y int) int {
	return x + y
}

// We can return multiple variables
func swap(x, y string) (string, string) {
	return y, x
}

// An example of a naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
