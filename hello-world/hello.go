package main

import "fmt"

import "rsc.io/quote"

func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(quote.Go())
}
