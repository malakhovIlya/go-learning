package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("Завершение")
	done := make(chan bool)
	go countTo(100, done)
	fmt.Println("Работаем")
	<-done
}

func countTo(n int, done chan<- bool) {
	for i := 1; i <= n; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(i)
	}
	done <- true
}
