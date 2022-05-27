/*
 *
 * Task     :   5.
 *                  -   а) Написать программу, которая выводит на экран ваше имя, фамилию и город проживания.
 *                  -   б) Сделать задание, только вывод организуйте в центре экрана
 *                  -   в) *Сделать задание б с использованием собственных методов (например, Print(string ms, int x,int y)
 */
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
	firstName, _ := reader.ReadString('\n')
	eError(firstName)
	//Ввод фамилии
	fmt.Println("Введите вашу фамилию :")
	lastName, _ := reader.ReadString('\n')
	eError(lastName)
	//Ввод города
	fmt.Println("Введите ваш город :")
	yourCity, _ := reader.ReadString('\n')
	eError(yourCity)

	//Вывод введенных параметров
	fmt.Printf("\t *Вас зовут : %v%v\t*Ваш город :%v", strings.ToUpper(firstName), strings.ToUpper(lastName), strings.ToUpper(yourCity))
}
