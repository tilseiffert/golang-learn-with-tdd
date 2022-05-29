package main

import "fmt"

const HelloPrefixEnglisch = "Hello, "
const HelloSuffix = "!"

func Hello(name string) string {
	//return "Hello, world!"
	return HelloPrefixEnglisch + name + HelloSuffix
}

func main() {
	fmt.Println(Hello("world"))
}
