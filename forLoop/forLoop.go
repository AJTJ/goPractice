package main

import (
	"fmt"
	"math"
)

/*
	IF STATEMENT
		but with a "pre" statement
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
		/* internal variables (v) available within IF and ELSE */
	} else {
		/* printing within a format */
		fmt.Printf("%g >= %g\n", v, lim)
	}
	/* internal variable (v) NOT available here */
	return lim
}

func sqrt(x float64) string {
	/*
		IF STATMENT
		if statements are like for loops. Parentheses aren't required around the expression, but braces ARE required.
	*/
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println("hello",
		/* both pow calls return before fmt.Println begins */
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	sum := 0
	/*
		FOR LOOP
		init statement
		condition expression
		post statement

		will iterate until the condition expression evaluates to false
	*/
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	/*
		FOR LOOP (ALSO WHILE LOOP)
		init and post statements are optional
		This is basically a WHILE loop
	*/

	total := 1

	for total < 1000 {
		total += total
	}

	fmt.Println(total)

	/*
		INFINITE LOOP

		for {
			...because there is no loop condition in here
		}
	*/

	fmt.Println("square root", sqrt(2233), sqrt(-4))

}
