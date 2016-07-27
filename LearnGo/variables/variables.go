package main

import "fmt"

func main() {

	var a string = "hello"
	b := 10
	var c = 12.5
	var d = false
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	// Print the type of the varibales
	fmt.Printf("Variable is of type %T \n", a)
	fmt.Printf("Variable is of type %T \n", b)
	fmt.Printf("Variable is of type %T \n", c)
	fmt.Printf("Variable is of type %T \n", d)

}
