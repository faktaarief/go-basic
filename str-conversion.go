package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String to Boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// String to Number
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(10000000, 10)
	fmt.Println(value)

	// String to Integer
	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)

	// Integer to String
	valueStr := strconv.Itoa(2000000)
	fmt.Println(valueStr)
}
