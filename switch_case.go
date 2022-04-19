package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	var lang string
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}

//func main() {
//	var code string
//	fmt.Scan(&code)
//
//	// определите полное название языка по его коду
//	// и запишите его в переменную `lang`
//	// ...
//	lang := "Unknown"
//	if code == "ru" || code == "rus" {
//		lang = "Russian"
//	} else if code == "fr" {
//		lang = "French"
//	} else if code == "en" {
//		lang = "English"
//	}
//
//	fmt.Println(lang)
//}

//package main
//
//import "fmt"
//
//func main() {
//	list := map[string]string{"en": "English", "fr": "French", "ru": "Russian", "rus": "Russian"}
//
//	var code string
//	fmt.Scanln(&code)
//
//	lang, ok := list[code]
//	if ok {
//		fmt.Println(lang)
//	} else {
//		fmt.Println("Unknown")
//	}
//}
