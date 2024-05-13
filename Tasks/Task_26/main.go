package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Unique("abCdefAaf"))
}

func Unique(s string) bool {
	r := []rune(strings.ToUpper(s))
	m := make(map[rune]struct{})

	for _, v := range r {
		_, ok := m[v]
		if ok {
			return false
		} else {
			m[v] = struct{}{}
		}
	}
	return true
}
