package main

import "fmt"

//func Solution(word string) string {
//	r := []rune(word)
//	fmt.Println(r)
//	var result []rune
//	for i := len(r) - 1; i >= 0; i-- {
//		result = append(result, r[i])
//	}
//	fmt.Printf("Result: %s\n", string(result))
//	return string(result)
//}

func Solution(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	return result
}
func main() {
	fmt.Printf(Solution("world"))
	fmt.Printf(Solution("word"))
}
