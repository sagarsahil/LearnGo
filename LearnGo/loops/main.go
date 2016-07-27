package main

import "fmt"

func main() {

	//We have only for loops in go.
	//There are three ways

	//Traditional
	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}

	fmt.Println("====================")

	//replacement of while
	index := 1
	for index < 10 {
		fmt.Println(index)
		index++
	}

	fmt.Println("====================")

	// Using the break statement .. infinite loop.
	index = 1
	for {

		fmt.Println(index)
		index++
		if index == 10 {
			break

		}

	}

}
