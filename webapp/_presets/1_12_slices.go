package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	s[0] = 0
	fmt.Println(s)

	slice := make([]int, 5)

	for i := 0; i < 5; i++ {
		slice[i] = i
	}

	fmt.Println(slice)

	//slice[5]=5  <-- Try me ;)

}
