package main

import "fmt"

func main() {
	var name string

	name = "Fakta Arief"
	fmt.Println(name)

	name = "Fakta Farhanto"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Fakta"
		lastName  = "Arief"
		address   string
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	address = "Bogor"
	fmt.Println(address)
}
