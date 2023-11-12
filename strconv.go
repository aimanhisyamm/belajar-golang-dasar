package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(2000000, 10)
	fmt.Println(value)
	valueBinary := strconv.FormatInt(2000000, 2)
	fmt.Println(valueBinary)
	valueOcta := strconv.FormatInt(2000000, 8)
	fmt.Println(valueOcta)
	valueHexa := strconv.FormatInt(2000000, 16)
	fmt.Println(valueHexa)

	valueInt, err := strconv.Atoi("3700000")
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err.Error())
	}

	valueIntBinary, err := strconv.Atoi("00011100")
	if err == nil {
		fmt.Println(valueIntBinary)
	} else {
		fmt.Println(err.Error())
	}
}
