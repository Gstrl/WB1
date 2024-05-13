package main

import (
	"fmt"
	"time"
)

// В этом случае канал читается, используя for-range.
// Когда канал будет закрыт и из него будет прочитано последнее сообщение, произойдёт выход из цикла.
func goroutine(ch chan string) {
	defer fmt.Println("Горутина завершена")
	for v := range ch {
		time.Sleep(time.Second)
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan string)
	go goroutine(ch)
	go func(ch chan string) {
		ch <- "Горутина работает"
	}(ch)
	time.Sleep(2 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)
}
