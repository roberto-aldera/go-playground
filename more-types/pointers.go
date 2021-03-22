package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2 = &Vertex{1, 2} // has type *Vertex
)

func run_pointer_examples() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Create an instance of type Vertex (a struct)
	fmt.Println(Vertex{1, 2})

	// Access its members
	v := Vertex{1, 2}
	fmt.Println(v.X)
	v.X = 4
	fmt.Println(v)

	// Access it with a pointer
	p1 := &v
	fmt.Println(*p1)
	p1.X = 10 // we can access and set member fields directly, no need to use (*p1).X
	fmt.Println(*p1)

	fmt.Println(v1, p2, v2, v3)
}
