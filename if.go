package main

import "fmt"

func main() {
	var name = "Hisyam"

	if name == "Aiman" {
		fmt.Println("Hello Aiman")
	} else if name == "Hisyam" {
		fmt.Println("Hello Hisyam")
	} else if name == "Mamang" {
		fmt.Println("Hello Mamang")
	} else {
		fmt.Println("Hi, boleh kenalan gak?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
