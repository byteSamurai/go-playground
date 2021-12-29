package main

import "fmt"

func main(){
	k := 5
	p := &k
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
}
