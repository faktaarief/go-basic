package main

import (
	"go-basic/helper"
	"fmt"
)

func main() {
	helper.SayHello("Fakta")

	// helper.sayGoodBye("Fakta") // error
	// fmt.Println(helper.version) // error

	fmt.Println(helper.Application)
}
