package main

import (
	"fmt"
)

const kelantan = "Kelantenese"
const japan = "Japanese"
const malayHelloPrefix = "Assalam!"
const kelantanHelloPrefix = "Guano "
const japaneseHelloPrefix = "Konichiwa, "

func Hello(name string, lang string) string {
	if lang == kelantan {
		return kelantanHelloPrefix + name + "? Bereh?"
	}
	if lang == japan {
		return japaneseHelloPrefix + name
	}

	if name == "" {
		name = " Apa Khabar!"
	} else {
		name = " Apa Khabar, " + name
	}

	return malayHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Fathul", ""))
}
