package main

import (
	"fmt"
	"time"
)

func main() {

	firstGuess := 42

	switch firstGuess{
	case 42:
		fmt.Println("Correct!")
	default:
		fmt.Println("That's wrong!")
	}

	currentTime := time.Now()
	switch {
	case currentTime.Hour() < 12:
		fmt.Println("Good morning!")
	case currentTime.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
