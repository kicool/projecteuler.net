package main

import (
	"fmt"
	"time"
)

func main() {
	c := 0
	t := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	for t.Year() < 2001 {
		if t.Weekday() == time.Sunday {
			c++
			fmt.Println(c, t)
		}
		t = t.AddDate(0, 1, 0)
	}
	fmt.Println("Total:", c)
}
