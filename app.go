package main

import (
	"strings"
)

//"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"

func ReverseWords(str string) string {
	words_array := strings.Split(str, " ")
	var result []string
	for i := len(words_array) - 1; i >= 0; i-- {
		result = append(result, words_array[i])
	}
	return strings.Join(result, " ") // reverse those words
}

func main() {
	ReverseWords("hello world!")
}
