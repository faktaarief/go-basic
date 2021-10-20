package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fakta Arief", "Fakta"))
	fmt.Println(strings.Contains("Fakta Arief", "Farhan"))

	fmt.Println(strings.Split("Fakta Arief", " "))

	fmt.Println(strings.ToLower("Fakta Arief"))
	fmt.Println(strings.ToUpper("Fakta Arief"))
	fmt.Println(strings.ToTitle("fakta arief"))

	fmt.Println(strings.Trim(" Fakta Arief  ", " "))
	fmt.Println(strings.ReplaceAll("Fakta Fakta", "Fakta", "Arief"))
}
