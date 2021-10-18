package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApllication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApllication(0)
}
