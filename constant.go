package main

import "fmt"

func main() {
	const firstName string = "Aiman"
	const lastName = "Hisyam"
	const value = 1000

	// Error
	// firstName = "Gak bisa diubah"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		friendFirstName string = "Arya Kunta"
		friendLastName         = "Aji"
		friendValue            = 100
	)

	fmt.Println(friendFirstName)
	fmt.Println(friendLastName)
	fmt.Println(friendValue)
}
