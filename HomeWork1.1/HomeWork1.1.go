/*
 * Task     :   1. Написать программу «Анкета». Последовательно задаются вопросы (имя, фамилия, возраст, рост, вес). В результате вся информация выводится в одну строчку.
 *                  -   а) используя склеивание;
 *                  -   б) используя форматированный вывод;
 *                  -   в) *используя вывод со знаком $.
 */

//Анкета
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Основная функция
func main() {

	//Переменная для использования интерфейса io.Reader
	reader := bufio.NewReader(os.Stdin)

	//Ввод имени
	fmt.Println("Введите ваше имя :")
	inputFirstName, _ := reader.ReadString('\n')
	if len(inputFirstName) <= 2 {
		fmt.Println("Вводи свое имя!")
		return
	}
	//Ввод фамилии
	fmt.Println("Введите фамилию :")
	inputLastName, _ := reader.ReadString('\n')
	if len(inputLastName) <= 2 {
		fmt.Println("Вводи свою фамилию!")
		return
	}
	//Ввод возраста
	fmt.Println("Введите возраст :")
	inputAge, _ := reader.ReadString('\n')
	if len(inputAge) <= 2 {
		fmt.Println("Вводи свой возраст!")
		return
	}
	//Ввод роста
	fmt.Println("Введите рост :")
	inputGrowth, _ := reader.ReadString('\n')
	if len(inputGrowth) <= 2 {
		fmt.Println("Вводи свой рост!")
		return
	}
	//Ввод веса
	fmt.Println("Введите вес :")
	inputWeight, _ := reader.ReadString('\n')
	if len(inputWeight) <= 2 {
		fmt.Println("Вводи свой вес!")
		return
	}

	//Вывод введенных параметров
	fmt.Printf("\n Вас зовут : \n %v %v Ваш возраст : \n %v\n Ваш рост : \n %v\n Ваш вес : \n %v\n",
		strings.ToUpper(inputFirstName), strings.ToUpper(inputLastName), inputAge, inputGrowth, inputWeight)
}
