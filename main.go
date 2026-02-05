package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight, userWeight float64 = 1.93, 100
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userWeight)

	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(IMT)
}
