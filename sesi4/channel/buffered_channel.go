package main

import "fmt"

func main() {
	c := make(chan bool, 5)

	c <- true

	isiChannelC := <- c

	fmt.Println("seharusnya print", isiChannelC)

	data := make(chan string, 2)
	data <- "Hello"
	// data <- "World"

	// firstWord, secondWord := <-data, <-data
	firstWord := <-data


	// fmt.Println(firstWord, secondWord)
	fmt.Println(firstWord)



}