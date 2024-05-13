package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// Этот вариант докручу чуть позже, но думаю задумка понятна)
func setBit(num int64, bitPosition int, bitValue int) int64 {

	const size = unsafe.Sizeof(int64(0))
	var byteSlise = make([]string, 0, 8)
	var s string

	for i := 0; i < 8; i++ {
		s = strconv.FormatInt(int64(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(size-1) - uintptr(i)))), 2)
		s = strings.Repeat("0", 8-len(s)) + s
		byteSlise = append(byteSlise, s)
	}
	fmt.Println(byteSlise)
	return num
}

func main() {
	var num int64 = -2 // Пример исходного числа
	bitPosition := 2   // Устанавливаемый бит
	bitValue := 1      // Значение, 1 или 0

	num = setBit(num, bitPosition, bitValue)

	fmt.Printf("Результат после установки бита: %d\n", num)
}
