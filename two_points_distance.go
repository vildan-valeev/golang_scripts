package main

import (
	"fmt"
	"math"
)

func main() {
	// объявите переменные x1, y1, x2, y2 типа float64
	var x1, y1, x2, y2 float64

	// считываем числа из os.Stdin
	// гарантируется, что значения корректные
	fmt.Scan(&x1, &y1, &x2, &y2)

	// рассчитайте d по формуле эвклидова расстояния
	// используйте math.Pow(x, 2) для возведения в квардрат
	// используйте math.Sqrt(x) для извлечения корня
	//fmt.Println(math.Hypot(x1-x2, y1-y2))
	d := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))

	// выводим результат в os.Stdout
	fmt.Println(d)
}
