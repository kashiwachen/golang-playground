package main

import "fmt"

func main() {
	var (
		a int = 1
		x, y int = 2, 3
		z bool = true
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("x, y: %v, %v\n", x, y)
	fmt.Printf("%v", z)
}
