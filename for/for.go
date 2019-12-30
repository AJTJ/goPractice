package main

import "fmt"

func main() {
	sum := 0
	// base for loop
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)

	tum := 1
	// without init and post statements
	// basically a "while" loop
	for tum < 1000 {
		tum += tum
	}
	fmt.Println("tum", tum)

	//an infinite loop
	// for {

	// }
}
