package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_POTATO_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var potatoChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkPotatoPrices(websites[i], potatoChannel)
	}
	sendMessage(chickenChannel, potatoChannel)
}

func checkPotatoPrices(website string, potatoChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var potatoPrice = rand.Float32() * 10
		if potatoPrice <= MAX_POTATO_PRICE {
			potatoChannel <- website
			break
		}
	}

}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, potatoChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Food deal on chicken at %v.", website)
	case website := <-potatoChannel:
		fmt.Printf("\nEmail Sent: Food deal on potatoes at %v.", website)
	}
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int, 5)
// 	go process(c)
// 	for i := range c {
// 		fmt.Println("Received:", i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func process(c chan int) {
// 	defer close(c)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	fmt.Println("Exiting process")
// }
