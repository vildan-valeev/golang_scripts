package main

//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	numbers := []int{1, 2, 3, 4, 5}
//
//	// Shuffle numbers, swapping corresponding entries in letters at the same time.
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(numbers), func(i, j int) {
//		numbers[i], numbers[j] = numbers[j], numbers[i]
//
//	})
//
//	fmt.Println(numbers)
//
//}

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func shuffle(nums []int) []int {
	// Функция `Seed()` инициализирует генератор случайных чисел
	// здесь мы используем константу `42`, чтобы программу
	// можно было проверить тестами.
	// В реальных задачах не используйте константы!
	// Используйте, например, время в наносекундах:
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(42)
	// перетасуйте `nums` с помощью `rand.Shuffle()``
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]

	})
	return nums
}

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
// после ввода нужно нажимать Enter, а затем Ctrl + D
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
