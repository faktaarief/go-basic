package main

import (
	"fmt"
	"go-basic/database"
)

// Blank Identifier
// import _ "go-basic/database"

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
