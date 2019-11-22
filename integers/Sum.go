package main

import "fmt"

// Add adds two numbers together and returns the result.
func Add(x, y int) (sum int) {
	return x + y
}

func main() {
	fmt.Println(Add(1, 2))
}
