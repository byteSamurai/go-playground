package main

import "fmt"

func main() {
	sum := 1
	//2^10
	for i := 0; i < 10; i++ {
		sum += sum
	}
	fmt.Println(sum)

}
