package iteration

import "strings"

func Repeat(character string, repetition int) string {

	return strings.Repeat(character, repetition)

	// var repeated string
	// for i := 0; i < repetition; i++ {
	// 	repeated += character
	// }
	// return repeated
}
