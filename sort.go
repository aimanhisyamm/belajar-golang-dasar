package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age // function to determined which value you want to sorting
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Arya", 26},
		{"Aiman", 24},
		{"Akbar", 29},
		{"Yaquut", 27},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
