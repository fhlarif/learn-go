package main

import (
	"fmt"
)

const malayHelloPrefix = "Assalam!"

func Hello(name string) string {
	if name == "" {
		name = " Apa Khabar!"
	} else {
		name = " Apa Khabar, " + name
	}
	return malayHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Fathul"))
}
