package main

import (
	"fmt"
	"os"
)

func main() {

	args, seperator := "", ""
	var totalargs int
	for _, arg := range os.Args[1:] {
		args += seperator + arg
		seperator = " "
		totalargs++

	}
	fmt.Printf("Total Arguments passed are %v\n", totalargs)
	fmt.Println(args)
}
