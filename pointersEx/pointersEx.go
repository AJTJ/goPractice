package main

import "fmt"

func main() {
	// ALL ABOUT POINTERS
	i, j := 42, 2701
	fmt.Println("original", i, j)

	p := &i              // points to i
	fmt.Println("p", *p) // read i through the pointer
	*p = 21              // set i through the pointer
	fmt.Println("i", i)  // see new value of i

	p = &j              // point to j
	*p = *p / 37        // divide j through the pointer
	fmt.Println("j", j) // see new value of j

	/*
		Editing a struct with dot notation
	*/
}
