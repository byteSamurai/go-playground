package main

import "fmt"


type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main(){

	var v I

	var t *T
	v = t
	describe(v)
	v.M()



	//nil interface
	var i I
	describe(i)

	//i.M() //runtime error
}