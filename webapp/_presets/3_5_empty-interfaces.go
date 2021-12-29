package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main(){
	describe("Servus!")
	describe(144)
}
