package main

import "fmt"

// Vertex is a STRUCT
type Vertex struct {
	X int
	Y int
}

func main() {
	// ALL ABOUT POINTERS
	i, j := 42, 2701
	fmt.Println("original", i, j)

	p := &i         // points to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see new value of j

	/*
		Editing a struct with dot notation
	*/
	v := Vertex{1, 2}
	v.X = 11
	fmt.Println(v)
}
