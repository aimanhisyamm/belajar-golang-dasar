package main

import "fmt"

func main() {
	name := "Aiman"
	counter := 0

	increment := func() {
		name := "Mamen"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}
	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
