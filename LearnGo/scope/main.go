package main

import "fmt"
import "github.com/sagarsahil/LearnGo/scope/support"

func main() {
	fmt.Println(support.PackageLevel)
	support.PrintString()
	support.FromOtherFile()
}
