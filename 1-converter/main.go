package main

import "fmt"

func main() {
	const UsdToEur float64 = 82.12
	const UsdToRub float64 = 0.88

	EurToRub := UsdToEur / UsdToRub

	fmt.Println(EurToRub)
}
