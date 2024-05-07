package main

// strings, bytes and runes

import "fmt"

func main() {
	myString := []rune("résumé") // rune is an alias for int32, represents a Unicode code point
	// fmt.Println(myString)
	fmt.Printf("%v, %T\n", myString[0], myString[0]) // utf-8 code
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Println(len(myString)) // 8 because of utf-8 encoding without using rune

	myRune := 'a'
	fmt.Printf("%v, %T\n", myRune, myRune) // int32

	myByte := byte('a')
	fmt.Printf("%v, %T\n", myByte, myByte) // uint8
}
