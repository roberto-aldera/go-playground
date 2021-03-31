package main

import "fmt"

func main() {
	fmt.Println("*** Pointers ***")
	run_pointer_examples()

	fmt.Println("\n*** Arrays and slices ***")
	run_slice_examples()

	fmt.Println("\n*** Running slice exercise ***")
	run_slice_exercise()

	fmt.Println("\n*** Running maps ***")
	run_maps_examples()

	fmt.Println("\n*** Running maps exercise ***")
	run_maps_exercise()

	fmt.Println("\n*** Running functions ***")
	run_functions_examples()

	fmt.Println("\n*** Running function fibonacci exercise ***")
	run_fibonacci_exercise()
}
