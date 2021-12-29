package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addVariant(x, y int) int {
	return x + y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addVariant(1, 2))
	fmt.Println(split(5))
	fmt.Println(swap("Ave", "Gaius"))
}
