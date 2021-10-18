package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var fakta Person
	fakta.Name = "Fakta"

	SayHello(fakta)

	cat := Animal{
		Name: "Ew",
	}

	SayHello(cat)
}
