package main

import "fmt"

const languageGerman = "German"

const HelloPrefixEnglisch = "Hello, "
const HelloPrefixGerman = "Hallo, "
const HelloSuffix = "!"

func Hello(name string, language string) string {
	//return "Hello, world!"

	if name == "" {
		name = "World"
	}

	if language == languageGerman {
		return HelloPrefixGerman + name + HelloSuffix
	}
	return HelloPrefixEnglisch + name + HelloSuffix
}

func main() {
	fmt.Println(Hello("world", ""))
}
