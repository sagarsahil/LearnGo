package main

import "fmt"

func main() {

	var avg = average(1, 2, 3, 4, 5, 6, 7, 7, 4, 89)
	fmt.Println("The average of the numbers is ", avg)
	//Here is another way of passing the parameters
	listofitems := []float32{5, 3, 3, 4, 2, 5, 2}
	var avg2 = average(listofitems...)
	fmt.Println("The average of the numbers is ", avg2)
}

func average(list ...float32) float32 {

	var total float32

	for _, num := range list {
		total += num
	}
	return total / float32(len(list))
}
