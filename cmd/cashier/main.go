package main

import (
	"flag"
	"fmt"
	"github.com/tooooots/cashier"
)

var change = flag.Int("change", 67, "Amount of money to return")

func main() {
	flag.Parse()
	fmt.Printf("change requested: %d cents\n", *change)

	var result = cashier.Cashier(*change)
	fmt.Printf("change to return: %v \n", result)
}
