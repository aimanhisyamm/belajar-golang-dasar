package main

import "fmt"

func main() {
	var nilai32 int32 = 400
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(string(nilai8))

	var (
		name         = "Aiman Hisyam"
		e       byte = name[0]
		eString      = string(e)
	)
	fmt.Println(name)
	fmt.Println(eString)
}
