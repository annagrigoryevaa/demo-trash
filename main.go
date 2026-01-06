package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight, userKg float64
	fmt.Print("Калькулятор ИМТ\n")
	fmt.Print("Ввевдите рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш индекс массы тела: ", IMT)
}
