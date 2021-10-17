package main

import "fmt"

func main() {
	name := "Fakta"
	counter := 0

	increment := func() {
		name := "Arief"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(name)
	fmt.Println(counter)
}
