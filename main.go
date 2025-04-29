package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	//var userHeight, userKg = 1.8,100
	fmt.Println("__Калькулятор индекса массы тела__")
	userKg, userHeight := getUserInput()

	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)

}

func outputResult(imt float64) {
	fmt.Printf("Ваш идекс массы тела: %v", imt)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	userHeight := 1.8
	userKg := 100.0

	fmt.Println("Введите свой рост в см:")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес в кг:")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
