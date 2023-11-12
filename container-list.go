package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Muhamad")
	data.PushBack("Aiman")
	data.PushBack("Hisyam")
	data.PushBack("Mamen")

	// fmt.Println(data.Front().Next().Prev().Value)
	// fmt.Println(data.Back().Prev().Prev().Value)

	// iteration from front
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// reverse iteration from back
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
