package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var n []int
	for _, value := range iterable {
		if predicate(value) {
			n = append(n, value)
		}
	}
	return n
}

func predicateFunc(v int) bool {
	return v%2 == 0
}

func main() {
	src := readInput2()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	res := filter(predicateFunc, src)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
// после ввода нужно нажимать Enter, а затем Ctrl + D
func readInput2() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
