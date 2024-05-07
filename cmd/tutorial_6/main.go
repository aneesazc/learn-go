package main

// pointers

import "fmt"

func sliceExample(arg []int) {
	// modify the slice
	arg[0] = 100
}

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println(a, b)
	fmt.Printf("The value a points to is %d\n", *b)
	fmt.Printf("The address of a is %v\n", &a)
	fmt.Printf("The value of b is %v\n", b)

	myMap := make(map[string]int)
	myMap["adam"] = 1
	myMap["john"] = 2
	myMap["joe"] = 3
	mapCopy := myMap
	mapCopy["dee"] = 4
	fmt.Println(myMap)
	fmt.Println(mapCopy)

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	sliceExample(mySlice)
	fmt.Println(mySlice)
}

// Create a new map
// newMap := make(map[string]int)

// Copy each element from the original map
// for key, value := range myMap {
//     newMap[key] = value
// }

// Now, modifications to newMap will not affect myMap
// newMap["dee"] = 4

// fmt.Println(myMap)    // Will not have the key "dee"
// fmt.Println(newMap)   // Will have the key "dee"
