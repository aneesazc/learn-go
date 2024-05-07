package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	// fmt.Println(arr[5]) // panic: runtime error: index out of range

	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	intSlice := []int{4, 5, 6} // slice
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	intSlice2 := append(intSlice, intSlice...)
	fmt.Println(intSlice2)

	newSlice := make([]int, 3, 5) // much faster iteration than without setting capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(newSlice), cap(newSlice))

	myMap := make(map[string]int)
	myMap["John"] = 1
	myMap["Sarah"] = 2
	// var value, ok = myMap["three"] // value = 0, ok = false
	value, ok := myMap["Sarah"] // value = 2, ok = true
	fmt.Println(value, ok)

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
