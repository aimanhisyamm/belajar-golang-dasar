package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var data *ring.Ring = ring.New(7)

	for i := 0; i <= data.Len(); i++ {
		data.Value = "Value " + string(i+96)
		data = data.Next()
	}
	fmt.Println(*data)

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
