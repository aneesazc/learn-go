package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{} // wait group to wait for all goroutines to finish
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // increment the wait group counter
		go dbCall(i)
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Printf("\nTotal time taken: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
}

func dbCall(i int) {
	// simulate db call delay
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done() // decrement the wait group counter
}
