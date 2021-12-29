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
	var c = make(chan int, 100)
	go countdownFrom(23, c)
	for {
		number, ok := <-c
		if ok == false {
			break
		}

		fmt.Println(number)
	}
}
