package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023, 10, 3, 16, 21, 10, 10, time.UTC)
	fmt.Println(utc)

	local := time.Date(2023, 10, 3, 16, 21, 10, 10, time.Local)
	fmt.Println(local)

	// format time.RFC3339
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1999-11-29")
	fmt.Println(parse)

	layout2 := "02-01-2006"
	parse2, _ := time.Parse(layout2, "29-11-1999")
	fmt.Println(parse2)
}
