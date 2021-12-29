package main

import	"fmt"

func sayHello() {
	fmt.Println("Whats up Doc?")
}

func main() {
	//"Whats up Doc?"
	sayHello()

	//no output, different thread
	go sayHello()
}
