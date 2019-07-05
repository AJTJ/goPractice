package main

import (
	"fmt"
	"math/cmplx"
)

//variable declarations
var i, j int = 1, 2

//variables declared without an explicit initial value are given their zero value
var f float64
var bzz bool
var s string

var (
	// ToBe comment here
	ToBe = false
	// MaxInt comment here
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(2, 4))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	// examples of initializers
	// := short assignment, var not necessary
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("%v %v %q\n", f, bzz, s)
}
