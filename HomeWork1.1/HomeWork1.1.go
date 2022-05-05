/*
 * Task     :   1. Написать программу «Анкета». Последовательно задаются вопросы (имя, фамилия, возраст, рост, вес). В результате вся информация выводится в одну строчку.
 *                  -   а) используя склеивание;
 *                  -   б) используя форматированный вывод;
 *                  -   в) *используя вывод со знаком $.
 */

package main

import "fmt"

func main() {
	var firstName, lastName string
	var age, growth, weight int

	fmt.Println("Введите имя ?")
	fmt.Scan(&firstName)
	fmt.Println("Введите фамилию ?")
	fmt.Scan(&lastName)
	fmt.Println("Введите возраст ?")
	fmt.Scan(&age)
	fmt.Println("Введите рост?")
	fmt.Scan(&growth)
	fmt.Println("Введите вес ?")
	fmt.Scan(&weight)

	fmt.Printf("\n Вас зовут : %s %s\n Ваш возраст : %d\n Ваш рост : %d\n Ваш вес : %d\n", firstName, lastName, age, growth, weight)
}
