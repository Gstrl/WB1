package main

import (
	"fmt"
	"time"
)

// Глобальная переменная для остановки горутины
var stopFlag bool

func goroutine() {
	defer fmt.Println("Горутина завершена")
	for {
		fmt.Println("Горутина работает")
		time.Sleep(1 * time.Second)
		if stopFlag {
			break
		}
	}
}

func main() {
	// Остановка горутины через флаг
	go goroutine()
	time.Sleep(5 * time.Second)
	stopFlag = true
	time.Sleep(1 * time.Second)
}
