package main

import (
	"fmt"
)

func main() {
	const usdToEur float64 = 0.86
	const usdToRub float64 = 81.2
	eur := 1.0
	rub := usdToRub * eur / usdToEur
	fmt.Print(rub)
}
