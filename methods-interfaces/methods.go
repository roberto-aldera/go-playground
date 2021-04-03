package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// A method is just a function with a receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Same thing as above, but now as a regular function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can also declare a method on non-struct types too
type MyFloat float64

// Just make sure the type is defined within the same package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Using a pointer receiver for the method allows us to change the underlying instance rather than a copy
// In practice, this is more commonly used than value receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	ScaleFunc(&v, 10) // here we MUST dereference
	fmt.Println(AbsFunc(v))

	// But if we had been using a pointer receiver, Go would do that for you with (&p)
	v = Vertex{3, 4}
	v.Scale(2) // no need to specify here, interpreted as (&v).Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3) // no need to specify here
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	// Same thing applies for value receivers, when passed a pointer Go does the coversion (*p)
	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex{3, 4}
	fmt.Println(p.Abs())     // here there's no need, it's interpreted as (*p).Abs()
	fmt.Println(AbsFunc(*p)) // here we have to specify

	// Generally all methods on a given type will use either value or pointer receivers, but not a mixture of both
	// Pointers are more efficient if the underlying type is large, or great for modifying the underlying data
	v1 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v1, v1.Abs())
	v1.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v1, v1.Abs())
}
