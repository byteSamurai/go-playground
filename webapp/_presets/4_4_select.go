package main

import (
	"fmt"
	"time"
)

func bomb(tick, boom <-chan time.Time) {
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {

	t := time.Tick(100 * time.Millisecond)
	b := time.After(500 * time.Millisecond)

	bomb(t,b)
}

