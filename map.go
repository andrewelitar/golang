package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Andrew",
		"address": "Asahan",
	}

	fmt.Print(person)
	fmt.Print(person["name"])
	fmt.Print(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Andrew"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
