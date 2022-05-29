package main

import "fmt"

const languageGerman = "German"
const languageFrench = "French"

const HelloPrefixEnglisch = "Hello, "
const HelloPrefixGerman = "Hallo, "
const HelloPrefixFrench = "Bonjour, "
const HelloSuffix = "!"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case languageFrench:
		prefix = HelloPrefixFrench
	case languageGerman:
		prefix = HelloPrefixGerman
	default:
		prefix = HelloPrefixEnglisch
	}
	return
}

func Hello(name string, language string) string {
	//return "Hello, world!"

	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name + HelloSuffix
}

func main() {
	fmt.Println(Hello("world", ""))
}
