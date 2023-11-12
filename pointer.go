package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	address3 := &address1

	address2.City = "Bandung" // if address2 not pointer, address1 will not change

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // all variable with same reference will change
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(*address4)

	var alamat Address = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
