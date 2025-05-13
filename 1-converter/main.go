package main

import "fmt"

func main() {
	fmt.Println("__Калькулятор конвертации валюты__")

	amount, fromCurrency, toCurrency := getUserInput()
	result := calculate(amount, fromCurrency, toCurrency)

	fmt.Printf("Результат: %.2f %s\n", result, toCurrency)
}

func getUserInput() (float64, string, string) {
	var from string
	var to string
	var amount float64

	fmt.Println("__Ввод исходной валюты__")
	for {
		fmt.Scan(&from)
		if from == "USD" || from == "EUR" || from == "RUB" {
			break
		}
		fmt.Println("Допустимые значения: USD, EUR, RUB")
	}

	fmt.Println("__Ввод суммы__")
	for {
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Сумма должна быть положительным числом")

			var trash string
			fmt.Scanln(&trash)
			continue
		}
		break
	}

	fmt.Println("__Ввод целевой валюты__")
	for {
		fmt.Scan(&to)
		if to == from {
			fmt.Println("Целевая валюта не должна совпадать с исходной")
		} else if to == "USD" || to == "EUR" || to == "RUB" {
			break
		} else {
			fmt.Println("Допустимые значения: USD, EUR, RUB")
		}
	}
	return amount, from, to
}

func calculate(amount float64, from string, to string) float64 {
	const usdToEur = 0.92
	const usdToRub = 90.0

	var amountInUSD float64

	switch from {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdToEur
	case "RUB":
		amountInUSD = amount / usdToRub
	}

	switch to {
	case "USD":
		return amountInUSD
	case "EUR":
		return amountInUSD * usdToEur
	case "RUB":
		return amountInUSD * usdToRub
	default:
		return 0
	}
}
