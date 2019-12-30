package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

// Sqrt replicating square root
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		x := z - (z*z-x)/(2*z)
		if d := math.Abs(z - x); d < delta {
			fmt.Println("here")
			z = x
			return z
		}
		z = x
		fmt.Println(i, z)
	}
	return z * z
}

func main() {

	fmt.Println("final", Sqrt(120))
}
