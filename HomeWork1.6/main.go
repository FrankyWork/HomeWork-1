/*
 * Person   :   Dolgosheev Alexander
 * Date     :   27.11.2020
 *
 * Task     :   6. *Создать класс с методами, которые могут пригодиться в вашей учебе (Print, Pause).
 */
package main

func person(firstName, lastName string) {
	print(firstName + " " + lastName)
}

const firstName, lastName string = "Dany", "Dolgosheev"

func main() {
	person(firstName, lastName)
}
