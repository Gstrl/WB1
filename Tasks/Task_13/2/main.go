package main

import "fmt"

// xor обмен
func main() {
	a := 10
	b := 5

	fmt.Println(a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)
}
