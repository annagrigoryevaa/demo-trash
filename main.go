package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор ИМТ")
	for {
		userKg, userHeight, err := getUserInput()
		if err != nil {
			fmt.Println(err)
		} else {
			IMT := calculateIMT(userKg, userHeight)
			outputResult(IMT)
			getTextResultBeauty(IMT)

		}
		if !RepeatCalculation() {
			break
		}
	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
}

func getTextRTesultCourse(IMT float64) {
	if IMT < 16 {
		fmt.Println("у вас сильный дефицит массы тела")
	} else if IMT >= 16 && IMT < 18.5 {
		fmt.Println("У вас дефицит массы тела")
	} else if IMT >= 18.5 && IMT < 25 {
		fmt.Println("У вас нормальный вес")
	} else if IMT >= 25 && IMT < 30 {
		fmt.Println("У вас избыток веса")
	} else if IMT >= 30 && IMT < 35 {
		fmt.Println("У вас первая степень ожирения")
	} else if IMT >= 35 && IMT < 40 {
		fmt.Println("У вас вторая степень ожирения")
	} else {
		fmt.Println("У вас третья степень ожирения")
	}
}

func getTextResultBeauty(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("у вас сильный дефицит массы тела")
	case IMT >= 16 && IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT >= 18.5 && IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT >= 25 && IMT < 30:
		fmt.Println("У вас избыток веса")
	case IMT >= 30 && IMT < 35:
		fmt.Println("У вас первая степень ожирения")
	case IMT >= 35 && IMT < 40:
		fmt.Println("У вас вторая степень ожирения")
	case IMT > 40:
		fmt.Println("У вас третья степень ожирения")
	}
}

func calculateIMT(userKg, userHeight float64) float64 {
	//if userKg <= 0 || userHeight <= 0 {
	//	return 0, errors.New("Не указан вес или рост")
	//	}
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64, error) {
	var userHeight, userKg float64

	fmt.Print("Ввевдите рост в см: ")
	fmt.Scan(&userHeight)
	if userHeight <= 0 {
		return 0, 0, errors.New("Укажите рост верно")

	}
	fmt.Print("Введите вес: ")
	fmt.Scan(&userKg)
	if userKg <= 0 {
		return 0, 0, errors.New("Укажите вес верно")
	}
	return userKg, userHeight, nil

}

func RepeatCalculation() bool {
	var answer string
	for i := 0; ; i++ {
		fmt.Println("Хотите продолжить?")
		fmt.Scanln(&answer)
		if answer == "да" {
			return true
		} else if answer == "нет" {
			return false
		} else {
			fmt.Println("Пожалуйста, введите 'да' или 'нет'.")
		}
	}
}
