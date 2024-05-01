package main

import "fmt"

func add(x, y int, z float32) int {
	ans := int(float32(x) + float32(y) + z)
	return ans
}

func main() {
	fmt.Println(add(1, 2, 3.0))
}
