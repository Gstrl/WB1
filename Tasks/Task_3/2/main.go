package main

import (
	"fmt"
	"sync"
)

//Эта программа создает горутины для каждого числа из массива numbers,
//которые конкурентно вычисляют его квадрат и отправляют результат в канал resultChannel.
//Затем основная горутина суммирует квадраты чисел из канала, получая итоговую сумму.
//Конструкция sync.WaitGroup используется для ожидания завершения всех горутин перед закрытием канала и вычислением окончательной суммы.

func calculateSquare(num int, resultChannel chan int) {
	square := num * num
	resultChannel <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	resultChannel := make(chan int)
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			calculateSquare(n, resultChannel)
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	sum := 0
	for square := range resultChannel {
		sum += square
	}

	fmt.Println(sum)
}
