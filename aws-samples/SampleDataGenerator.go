package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	var data string
	//generate some sample data
	for index := 1; index < 500; index++ {
		data += "This is a sample data \n"
	}
	numberoffiles := rand.Int() % 500
	for filecount := 1; filecount < numberoffiles; filecount++ {
		filename := os.Args[1] + "/SampleFile" + strconv.Itoa(filecount)
		fmt.Printf("Trying to create file %v\n", filename)
		f, err := os.Create(os.Args[1] + "/SampleFile" + strconv.Itoa(filecount) + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		f.WriteString(data)
		f.Close()
		fmt.Printf("Writing file number %v", filecount)

	}

}
