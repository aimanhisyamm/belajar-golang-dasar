package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func sayHi1(name string, customer Customer) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (customer Customer) sayHi2(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	var aiman Customer
	aiman.Name = "Aiman"
	aiman.Address = "Jakarta"
	aiman.Age = 23
	fmt.Println(aiman.Name)
	fmt.Println(aiman.Address)
	fmt.Println(aiman.Age)

	arya := Customer{
		Name:    "Arya",
		Address: "Yogyakarta",
		Age:     26,
	}
	fmt.Println(arya)

	yaquut := Customer{"Yaquut", "Bogor", 27}
	fmt.Println(yaquut)

	sayHi1("Mamen", aiman)
	yaquut.sayHi2("Ayat")
	arya.sayHuuu()
}
