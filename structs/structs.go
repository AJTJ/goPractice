package main

import "fmt"

// Vertex is a STRUCT
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 11
	fmt.Println("v", v)

	// using pointers to structs
	d := Vertex{1, 2}
	e := &d
	e.X = 1e9
	fmt.Println("d", d)

	//Struct Literals

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, v2, v3, p)
}
