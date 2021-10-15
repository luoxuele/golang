package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("main start ...")
	fmt.Println(3.1 * 4.2)
	fmt.Println(decimal.NewFromFloat(3.1).Mul(decimal.NewFromFloat(4.2)))
}
