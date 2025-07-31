package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль нельзя")
	}
	return a / b, nil
}

func mustDivide(a, b int) int {
	if b == 0 {
		panic("Паника")
	}
	return a / b
}

func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановились после паники:", r)
		}
	}()
	return mustDivide(a, b)
}

func main() {
	fmt.Println(divide(10, 0))
	fmt.Println(safeDivide(10, 0))
	fmt.Println(mustDivide(10, 1))
}
