package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aiman Hisyam", "Arya"))
	fmt.Println(strings.Contains("Aiman Hisyam", "Aiman"))

	fmt.Println(strings.Split("Muhamad Aiman Hisyam", " "))

	fmt.Println(strings.ToLower("Muhamad Aiman Hisyam"))
	fmt.Println(strings.ToUpper("Muhamad Aiman Hisyam"))
	fmt.Println(strings.ToTitle("muhamad aiman hisyam"))

	fmt.Println(strings.Trim("     muhamad aiman hisyam     ", " ")) // hanya menghilangkan kiri dan kanan yang kosong
	fmt.Println(strings.Trim("a     muhamad aiman hisyam     ", " "))
	fmt.Println(strings.Trim("a     muhamad aiman hisyam     a", " "))

	fmt.Println(strings.ReplaceAll("Aiman Arya Aiman Yaquut", "Aiman", "Mamen"))

}
