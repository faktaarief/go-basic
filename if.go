package main

import "fmt"

func main() {
	var name = "Fakta"

	if name == "Fakta" {
		fmt.Println("Hello Fakta")
	} else if name == "Arief" {
		fmt.Println("Hello Arief")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// var length = len(name)
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
