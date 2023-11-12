package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Round:", math.Round(1.0))
	fmt.Println("Round:", math.Round(1.7))
	fmt.Println("Round:", math.Round(1.01))

	fmt.Println("Floor:", math.Floor(2))
	fmt.Println("Floor:", math.Floor(1.99))

	fmt.Println("Ceil:", math.Ceil(1.00))
	fmt.Println("Ceil:", math.Ceil(1.001))
	fmt.Println("Ceil:", math.Ceil(1.2))

	fmt.Println("Max:", math.Max(1.2002, 1.22))
	fmt.Println("Min:", math.Min(1.20000001, 1.2001))

}
