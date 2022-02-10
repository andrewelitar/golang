package main

import "fmt"

func main() {
	var name string

	name = "Andrew Eliezer"
	fmt.Println(name)

	name = "Andrew Tarigan"
	fmt.Println(name)

	var friendName string = "Dadang"
	fmt.Println(friendName)

	var age int8 = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Andrew"
		lastName  = "Tarigan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
