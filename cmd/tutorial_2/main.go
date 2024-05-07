package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe()
	res, remainder, err := intDivision(10, 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Result: %d, Remainder: %d\n", res, remainder)
}

func printMe() {
	fmt.Println("Hello, World!")
}

func intDivision(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	res := a / b
	remainder := a % b
	return res, remainder, nil
}
