package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	/*
		switch written with an expression to define os
		switch always exits upon 1st successful case
	*/
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows
		fmt.Printf("%s. \n", os)
	}

	/*
		DEFER
		Is evaluated immediately, but not executed until the surrounding functions return.
		Last-In-First-Out
	*/
	defer fmt.Println("The deferred statement")
	defer fmt.Println("The second deferred")

	/*
		Switch without condition is switch true
		Good for long if-then-else chains
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
