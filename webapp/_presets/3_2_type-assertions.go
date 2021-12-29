package main

import "fmt"

func main() {
	//var j interface{} = 33.3
	//var j interface{} = "Servus"
	var j interface{} = 5

	switch v := j.(type) {
	case int:
		fmt.Printf("2 * %v = %v\n", v, v*2)
	case string:
		fmt.Printf("%q ist %v Bytes lang\n", v, len(v))
	default:
		fmt.Printf("Unbekannter Typ %T!\n", v)
	}
	
}