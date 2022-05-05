/*
 * Task     :   2. Ввести вес и рост человека. Рассчитать и вывести индекс массы тела (ИМТ) по формуле I=m/(h*h); где m — масса тела в килограммах, h — рост в метрах
 */

package main

import (
	"fmt"
)

func IMT(iGrowth, iWeight int) int {
	massa := iGrowth * iGrowth / iWeight
	return massa
}

func main() {
	var Growth, Weight int
	fmt.Println("Введите рост человека ( рост в метрах ) :")
	fmt.Scan(&Growth)
	fmt.Println("Введите вес человека (масса тела в килограммах) :")
	fmt.Scan(&Weight)
	result := IMT(Growth, Weight)
	fmt.Println("Индекс массы тела (ИМТ:)", result)
	if result < 0 {
		fmt.Println("Вы ввели некорректное значение. Рост и вес должны быть > 0")
	}
}
