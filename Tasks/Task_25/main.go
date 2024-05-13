package main

import (
	"context"
	"fmt"
	"time"
)

type Sleep struct {
	time context.Context
}

func (s *Sleep) Run() {
	select {
	case <-s.time.Done():
		return
	}
}

func NewSleep(time time.Duration) Sleep {
	ctx, _ := context.WithTimeout(context.Background(), time)
	return Sleep{ctx}
}

func main() {
	s := NewSleep(time.Second * 5)
	s.Run()
	fmt.Println("Проснулись")
}
