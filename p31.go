package main

import (
	"flag"
	"fmt"
)

const (
	S = 200
)

var (
	amount = flag.Int("amount", 0, "total amount")
)

func chargeCount(amount int, c []int) int {
	switch {
	case amount == 0:
		return 1
	case amount < 0:
		return 0
	case len(c) == 0:
		return 0
	}
	return chargeCount(amount, c[1:len(c)]) + chargeCount(amount-c[0], c)
}

func main() {
	flag.Parse()

	if *amount < 0 {
		return
	}

	currency := []int{1, 2, 5, 10, 20, 50, 100, 200}
	fmt.Println(currency)

	fmt.Println(chargeCount(*amount, currency))
}
