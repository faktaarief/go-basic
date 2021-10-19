package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	// go run os.go fakta arief
	fmt.Println(args[1]) // fakta
	fmt.Println(args[2]) // arief

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error: ", err)
	}
}
