package main

import (
	"fmt"
	"strings"
)

func run_slice_examples() {
	// Arrays are fixed, not resizeable
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices are resizeable, and a lot more commonly used in practice
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices refer to underlying arrays, and changing them will change those arrays (and any other slices)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sl := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sl, sl[0].i, sl[0].b)

	// Default limits
	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:2]
	fmt.Println(s1)

	s1 = s1[1:]
	fmt.Println(s1)

	// Using length and capacity of slices
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	s2 = s2[:0] // Slice the slice to give it zero length.
	printSlice(s2)

	s2 = s2[:4] // Extend its length. Capacity won't change. Not quite sure why it changes for the next example but not this one.
	printSlice(s2)

	s2 = s2[2:] // Drop its first two values. Capacity will fall by 2, because there are two fewer elements in the array
	printSlice(s2)

	// Nil slice - has no underlying array
	var nil_slice []int
	fmt.Println(nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("nil!")
	}

	// Use make - this is how you create dynamically-sized arrays
	a2 := make([]int, 5) // len(a2)=5, cap(a2)=5
	printSlice2("a", a2)

	b2 := make([]int, 0, 5) // len(b2)=0, cap(b2)=5
	printSlice2("b", b2)

	c2 := b2[:2]
	printSlice2("c", c2)

	d2 := c2[2:5]
	printSlice2("d", d2)

	// Create a tic-tac-toe board. Slices of slices.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to slices will grow them
	var s3 []int
	printSlice(s3)

	s3 = append(s3, 0) // append works on nil slices.
	printSlice(s3)

	s3 = append(s3, 1) // The slice grows as needed.
	printSlice(s3)

	s3 = append(s3, 2, 3, 4) // We can add more than one element at a time.
	printSlice(s3)

	// Using range to iterate over a slice - returns index and copy of element at index
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// This is how you'd use either the index or value only
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
