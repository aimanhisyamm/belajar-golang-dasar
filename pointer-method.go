package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { // if don't use pointer method, stuct will be pass by value, and struct variable will not change
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	aiman := Man{"Aiman"}
	aiman.Married()
	fmt.Println(aiman.Name)
}
