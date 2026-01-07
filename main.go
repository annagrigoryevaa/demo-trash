package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор ИМТ")
	//userKg, userHeight: = getUserInput()
	IMT := calculateIMT(getUserInput())
	outputresult(IMT)

}

func outputresult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Printf(result)
}

func calculateIMT(userKg, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Ввевдите рост в см: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
