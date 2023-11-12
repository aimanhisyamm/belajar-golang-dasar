package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Warudo")
}

func sayHelloAiman(index int) {
	fmt.Println("Hello Aiman ke -", index)
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

func getHello(name string) string {
	if name == "" {
		return "Hello Cuy"
	} else {
		return "Hello " + name
	}
}

func getFullName() (string, string, string) {
	return "Muhamad", "Aiman", "Hisyam"
}

func getFullNameNew() (firstName, middleName, lastName string) {
	firstName = "Arya"
	middleName = "Kunta"
	lastName = "Aji"

	return
}

func main() {
	sayHello()
	for i := 1; i <= 3; i++ {
		sayHelloAiman(i)
	}
	firstName := "Aiman"
	sayHelloTo(firstName, "Hisyam")

	result := getHello("Mamen")
	fmt.Println(result)
	fmt.Println(getHello(""))

	_, middleName, lastName := getFullName()
	// fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	a, b, c := getFullNameNew()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
