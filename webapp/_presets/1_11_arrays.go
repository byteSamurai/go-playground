package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)

	b := [2]string{"Hello", "World"}

	fmt.Println(b)

}
