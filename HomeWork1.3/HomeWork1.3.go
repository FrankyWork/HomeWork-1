/*
 * Task     :   а) Написать программу, которая подсчитывает расстояние между точками с координатами x1, y1 и x2,y2
 *                 по формуле r=Math.Sqrt(Math.Pow(x2-x1,2)+Math.Pow(y2-y1,2).
 *                 Вывести результат, используя спецификатор формата .2f (с двумя знаками после запятой);
 *
 *              б) *Выполните предыдущее задание, оформив вычисления расстояния между точками в виде метода;
 */

package main

import (
	"fmt"
	"math"
)

func method(x1, x2, y1, y2 float64) float64 {
	cord := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return cord
}
func main() {
	var sX1, sX2, sY1, sY2 float64

	fmt.Println("Введите координаты первой точки по оси X :")
	fmt.Scan(&sX1)
	fmt.Println("Введите координаты второй точки по оси X :")
	fmt.Scan(&sX2)
	fmt.Println("Введите координаты первой точки по оси Y :")
	fmt.Scan(&sY1)
	fmt.Println("Введите координаты второй точки по оси Y :")
	fmt.Scan(&sY2)

	result := method(sX1, sX2, sY1, sY2)
	if result <= 0 {
		fmt.Println("Ошибка в данных")
		return
	}
	fmt.Println("Расстояния между координатами:", result)

}
