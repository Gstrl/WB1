package main

import "fmt"

// Интереса ради
func main() {
	sum := 0
	arr := []int{2, 4, 6, 8, 10}
	for _, v := range arr {
		sum += v * v
	}
	fmt.Println(sum)
}
