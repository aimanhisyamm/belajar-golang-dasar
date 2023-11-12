package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Mamen")
	fmt.Println(result)
	fmt.Println(getGoodBye("Mamen"))

	sayHelloWithFilter("Aimen", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
