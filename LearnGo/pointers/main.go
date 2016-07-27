package main

import "fmt"

func main() {
	var number int = 10

	var pointtonum *int = &number

	fmt.Println(number)      // Print the number
	fmt.Println(pointtonum)  // Print the address
	fmt.Println(*pointtonum) // Print the number pointed by address

}
