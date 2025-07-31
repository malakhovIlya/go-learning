package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Voice string
}

type Cat struct {
	Voice string
}

func (cat Cat) Speak() {
	fmt.Println(cat.Voice)
}

func (dog Dog) Speak() {
	fmt.Println(dog.Voice)
}

func MakeAnimalSpeak(animal Animal) {
	animal.Speak()
}

func main() {
	MakeAnimalSpeak(Cat{"mrrrr"})
	MakeAnimalSpeak(Dog{"gafff"})
}
