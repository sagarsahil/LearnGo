package support

import "fmt"

//PackageLevel is an exposed variable outside the package
var PackageLevel = "I am an expored variable"

//PrintString prints the data as an exported function
func PrintString() {
	fmt.Println("I am an exported method")
	packscope(10)
}

func packscope(param int) {
	fmt.Printf("I can be called only from package and can have args %v \n", param)
}
