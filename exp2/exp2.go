package main

import (
	"fmt"
	"math"
)

// Pi const declarations cannot be changed with :=
const Pi = 3.14

// Big shifts the number left by 100 zeroes
// Small removes 99 of those zeroes
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)

	v := 52.6
	//the type of v is inferred based on the righthand side
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
}
