package main

import "fmt"

func main() {
	var name = "Hisyam"

	switch name {
	case "Aiman":
		fmt.Println("Hello Aiman")
		fmt.Println("Hello Aiman")
	case "Hisyam":
		fmt.Println("Hello Hisyam")
		fmt.Println("Hello Hisyam")
	default:
		fmt.Println("Kenalan dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
