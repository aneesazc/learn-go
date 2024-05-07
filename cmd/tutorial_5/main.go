package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	kWh     uint8
	gallons uint8
}

func (g gasEngine) String() string {
	return fmt.Sprintf("This car has %d mpg and %d gallons.", g.mpg, g.gallons)
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.kWh * 3
}

func (e gasEngine) setMPG(mpg uint8) {
	e.mpg = mpg
}

type emgine interface {
	milesLeft() uint8
}

func canMakeIt(e emgine, miles uint8) {
	if e.milesLeft() >= miles {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

func main() {
	myEngine := gasEngine{mpg: 30, gallons: 10}
	// myEngine2 := struct {
	// 	mpg     uint8
	// 	gallons uint8
	// }{mpg: 30, gallons: 10}
	fmt.Println(myEngine)
	// fmt.Println(myEngine.owner.name)
	// myEngine.drive()
	myEngine.setMPG(40) // this will not change the value of mpg because the receiver is a value receiver
	// you have to use a pointer receiver to change the value of mpg or directly change the value of mpg
	fmt.Println(myEngine)
	myEngine.mpg = 50
	fmt.Println(myEngine) // this will change the value of mpg
	myElectricEngine := electricEngine{kWh: 10, gallons: 10}
	fmt.Println(myElectricEngine)
	canMakeIt(myEngine, 60)
	canMakeIt(myElectricEngine, 30)
}
