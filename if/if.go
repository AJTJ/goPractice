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

func pow(x, n, lim float64) float64 {
	fmt.Println("this:", math.Pow(x, n))
	if v := math.Pow(x, n); v < lim {
		fmt.Println("here lies v", v)

		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
