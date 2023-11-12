package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument: ")
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hostname:", hostname)
	}

	gopath := os.Getenv("GOPATH")
	goroot := os.Getenv("GOROOT")

	fmt.Println(gopath)
	fmt.Println(goroot)

}
