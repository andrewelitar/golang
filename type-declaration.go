package main

import "fmt"

func main() {
	type NOKTP string
	type Married bool

	var noKTPAndrew NOKTP = "123456789101112"
	var marriedStatus Married = false

	fmt.Println(noKTPAndrew)
	fmt.Println(marriedStatus)
}
