package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767 // + 1 will give compile time error
	intNum = intNum + 1      // runtime error: constant 32768 overflows int16
	fmt.Println(intNum)

	var floatNum float64 = 3.14 // more precise than float32
	fmt.Println(floatNum)

	// cant add int16 and int32 or other different types

	var myString string = "Hello, World!" + " I am a string"
	fmt.Println(myString)

	fmt.Println(len(myString)) // length of string -> not the number of characters but the number of bytes
	// go uses utf-8 encoding, so some characters take more than 1 byte
	fmt.Println(utf8.RuneCountInString(myString)) // number of characters

	var myRune rune = 'a' // alias for int32
	fmt.Println(myRune)   // prints the ascii value of 'a'

	// default values for numerics is 0, for strings is empty string, for bool is false

	myVar := "text" // short hand declaration
	fmt.Println(myVar)

	const myConst float64 = 3.14 // constant
	fmt.Println(myConst)
}
