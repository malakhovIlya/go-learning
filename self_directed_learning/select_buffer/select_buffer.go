package main

import (
	"fmt"
	"time"
)

func main() {

	fast := make(chan string)
	slow := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(200 * time.Millisecond)
			fast <- "быстро"
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second)
			slow <- "медленно"
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg := <-fast:
			fmt.Println("Получено из fast:", msg)
		case msg := <-slow:
			fmt.Println("Получено из slow:", msg)
		default:
			fmt.Println("Ничего не готово, ждем...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
