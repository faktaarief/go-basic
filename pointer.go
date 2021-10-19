package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Bogor", "Jawa Barat", "Indonesia"}
	// address2 := address1 // Pass by Value
	var address2 *Address = &address1 // Pointer // Pass by Reference

	address2.City = "Solo"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = new(Address)
	address3.City = "Jakarta"
	fmt.Println(address3)

	var alamat = Address{
		City:     "Bogor",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
