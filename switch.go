package main

import "fmt"

func main() {
	name := "Fakta"

	switch name {
	case "Fakta":
		fmt.Println("Hello Fakta!")
		fmt.Println("Good Morning")
	case "Arief":
		fmt.Println("Hello Arief!")
	default:
		fmt.Println("Apa Kabar?")
	}

	// Short Statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	// Swicth Tanpa Kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
