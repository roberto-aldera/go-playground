package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	// if-statement doesn't need () but does need {}
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow1(x, n, lim float64) float64 {
	// can execute a short statement in this if statement, no need to do it beforehand for v
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// A function to manually implement a square-root finder
func get_sqrt(x float64) float64 {
	z := x
	prev_z := 0.0
	for math.Abs(z-prev_z) > 0.01 {
		prev_z = z
		z -= (z*z - x) / (2 * z) // Newton's method
		fmt.Println(z)
	}
	return z
}

func main() {

	// This is evaluated immediately but only executes last
	defer fmt.Println("Finished!")

	// simple for-loop
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// the init and post statements are actually optional (that is cool)
	sum2 := 1
	// for ; sum2 < 1000; { -> this has to be a comment because my IDE gets rid of the semicolons,
	// Go doesn't need them, this is basically a while-loop
	for sum2 < 1000 {
		sum2 += sum2
		fmt.Println(sum2)
	}

	// Infinite loop - commented out for now
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	fmt.Println("Sqrt is:", get_sqrt(100))

	// A switch-case statement - notice the lack of 'break' as would be needed in other languages
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch statement evaluation order example
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch but without a condition - a clean way to write long if...else statements
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// An example of defer statements being pushed onto a stack, LIFO
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
