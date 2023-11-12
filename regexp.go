package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-zA-Z])([a-zA-Z])n")

	fmt.Println(regex.MatchString("aIIn"))
	fmt.Println(regex.MatchString("aima"))
	fmt.Println(regex.MatchString("ailn"))
	fmt.Println(regex.MatchString("aman"))

	regex = regexp.MustCompile("e([a-z])o")
	search := regex.FindAllString("eko edo eka ema ode eto elo", 2)
	fmt.Println(search)

	search = regex.FindAllString("eko edo eka ema ode eto elo", -1) // find all with -1
	fmt.Println(search)
}
