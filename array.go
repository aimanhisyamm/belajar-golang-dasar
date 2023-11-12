package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhamad"
	names[1] = "Aiman"
	names[2] = "Hisyam"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		85,
		95,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))

}
