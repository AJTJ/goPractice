package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		//adds the "i" because a negative number is an impossible number.
		return sqrt(-x) + "i"
	}
	//Sprint formats and returns a string
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
