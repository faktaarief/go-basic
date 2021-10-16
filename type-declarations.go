package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpFakta NoKTP = "1217374999"
	var marriedStatus Married = true
	fmt.Println(noKtpFakta)
	fmt.Println(marriedStatus)
}
