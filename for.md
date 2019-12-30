CURRENT STATUS
https://tour.golang.org/flowcontrol/1

# Loops in Go
- Go only has the `for` loop
```
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

### For loops are made up of three parts
1. the init statement: executed before the first iteration
2. the condition expression: evaluated before every iteration
3. the post statement: executed at the end of every iteration
