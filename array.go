package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Fakta"
	names[1] = "Arief"
	names[2] = "Farhanto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	// Dynamic index length
	var values2 = [...]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values2)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}
