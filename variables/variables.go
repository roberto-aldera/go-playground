package main

import (
	"fmt"
	"math/cmplx"
)

// var statement can be at package level
var c, python, java bool
var j, k int = 1, 2

// Nice factoring into blocks for readability
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// var statement can also be at function level
	var i int

	// can omit type if initialiser is present
	var h = 5
	fmt.Println(i, 10)
	fmt.Println(i, c, python, java)
	fmt.Println(j, k, h)

	// inside function, can use short assignment statement instead of var if type is implicit
	m := 3
	fmt.Println(m)

	// Print types and values
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Types are set to their zero value if not initialised (0, false, or "")
	var a int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", a, f, b, s)

	// Type conversions
	var q int = 42
	var w float64 = float64(q)
	u := float64(q) // a simpler way to do type conversion
	fmt.Println(q, w, u)

	v := 42.9 // change this to see how inferred type changes
	fmt.Printf("v is of type %T\n", v)

	// Constants - note: can't use := syntax for these
	const Num = 1.234
	fmt.Println(Num)

	// Numeric constants for high precision things
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
}
