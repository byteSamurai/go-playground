package main

import (
	"math"
	"fmt"
)

type I interface {
	M()
}
type F float64

func (f F) M() {
	fmt.Println(f)
}
func main() {
	var i I
	i = F(math.Pi)
	i.M()
}
