/*
 *
 * Task     :   4. Написать программу обмена значениями двух переменных.
 *                  -   а) с использованием третьей переменной;
 *                  -   б) *без использования третьей переменной.
 */

package main

import (
	"fmt"
)

var varA int
var varB int
var varC int

func main() {

	fmt.Println("Программа обмена значениями двух переменных \n \t Введите переменную A :")
	_, _ = fmt.Scan(&varA)
	fmt.Println("\tВведите переменную B : ")
	_, _ = fmt.Scan(&varB)

	fmt.Printf("Переменные \n varA = %v\n varB = %v\n", varA, varB)

	varC = varB
	varB = varA
	varA = varC

	fmt.Printf("Поменяли переменные ( через 3ю ) \n varA = %v\n varB = %v\n", varA, varB)

	varA = varA + varB
	varB = varA - varB
	varA -= varB

	fmt.Printf("Поменяли переменные ( без использования третьей переменной ) \n varA = %v\n varB = %v\n", varA, varB)
}
