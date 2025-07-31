package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayHi() {
	fmt.Println("Привет, меня зовут,", u.Name)
}

func (u *User) GrowOld() {
	u.Age++
}
