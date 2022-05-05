/*
 * Task     :   1. Написать программу «Анкета». Последовательно задаются вопросы (имя, фамилия, возраст, рост, вес). В результате вся информация выводится в одну строчку.
 *                  -   а) используя склеивание;
 *                  -   б) используя форматированный вывод;
 *                  -   в) *используя вывод со знаком $.
 */

//Main - Анкета
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//Функция ошибки
func eError(x string) string {
	if len(x) <= 2 {
		fmt.Println("***Ошибка***")
		log.Fatal()
	}
	return x
}

//Основная функция
func main() {

	//Переменная для использования интерфейса io.Reader
	reader := bufio.NewReader(os.Stdin)

	//Ввод имени
	fmt.Println("Введите ваше имя :")
	inputFirstName, _ := reader.ReadString('\n')
	eError(inputFirstName)
	//Ввод фамилии
	fmt.Println("Введите фамилию :")
	inputLastName, _ := reader.ReadString('\n')
	eError(inputLastName)
	//Ввод возраста
	fmt.Println("Введите возраст :")
	inputAge, _ := reader.ReadString('\n')
	eError(inputAge)
	//Ввод роста
	fmt.Println("Введите рост :")
	inputGrowth, _ := reader.ReadString('\n')
	eError(inputGrowth)
	//Ввод веса
	fmt.Println("Введите вес :")
	inputWeight, _ := reader.ReadString('\n')
	eError(inputWeight)

	//Вывод введенных параметров
	fmt.Printf("\n Вас зовут : \n %v %v Ваш возраст : \n %v\n Ваш рост : \n %v\n Ваш вес : \n %v\n",
		strings.ToUpper(inputFirstName), strings.ToUpper(inputLastName), inputAge, inputGrowth, inputWeight)
}
