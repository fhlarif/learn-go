package main

import (
	"fmt"
)

const (
	kelantan            = "Kelantenese"
	japan               = "Japanese"
	malayHelloPrefix    = "Assalam, "
	kelantanHelloPrefix = "Guano, "
	japaneseHelloPrefix = "Konichiwa, "
	defaultPrefix       = "Assalam!"
)

func Hello(name string, lang string) string {

	if name == "" {
		return defaultPrefix
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case kelantan:
		prefix = kelantanHelloPrefix
	case japan:
		prefix = japaneseHelloPrefix
	default:
		prefix = malayHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Fathul", ""))
}
