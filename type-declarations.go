package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAiman NoKTP = "123617823671"
	var marriedStatus Married = false

	fmt.Println(noKtpAiman)
	fmt.Println(marriedStatus)
}
