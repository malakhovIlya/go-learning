package main

import "fmt"

func main() {
	var name string = "Alex"
	age := 30

	if age > 18 {
		fmt.Println(name, "is an adult.")
	} else {
		fmt.Println(name, "is a minor.")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}
}
