package main

import (
	"fmt"
)

func countdownFrom(n int, c chan int) {
	for {
		c <- n
		if n <= 0 {
			close(c)
			return
		}
		n--
	}
}
func main() {
	c := make(chan int)
	go countdownFrom(23, c)

	for number := range c {
		fmt.Println(number)
	}
}
