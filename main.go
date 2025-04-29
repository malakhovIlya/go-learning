package main

import (
	"fmt"
	"math"
)

func main() {
	//var userHeight, userKg = 1.8,100
	const IMTPower = 2
	userHeight := 1.8
	userKg := 100.0
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Println(IMT)
}
