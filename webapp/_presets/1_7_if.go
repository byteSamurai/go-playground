package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	i := rand.Float32()*4 - 2

	if i < 0 {
		fmt.Println("I'm small")
	} else {
		fmt.Println("I'm huge")
	}

}
