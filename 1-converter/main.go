package main

import "fmt"

func main() {
	const UsdToEur float64 = 82.12
	const UsdToRub float64 = 0.88

	EurToRub := UsdToEur / UsdToRub

	fmt.Println(EurToRub)
}

func getUserInput() {
	var dataUser string
	fmt.Print("Enter your data: ")
	fmt.Scan(&dataUser)
}

func calculate(amount float64, fromCurrency, toCurrency string) float64 {
	return amount
}
