package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var fakta Customer
	fakta.Name = "Fakta"
	fakta.Address = "Bogor"
	fakta.Age = 21

	fakta.sayHi("Arief")

	// fmt.Println(fakta)

	// arief := Customer{
	// 	Name:    "Arief",
	// 	Address: "Jakarta",
	// 	Age:     21,
	// }

	// fmt.Println(arief)

	// farhanto := Customer{"Farhanto", "Solo", 21}

	// fmt.Println(farhanto)
}
