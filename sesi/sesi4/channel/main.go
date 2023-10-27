package main

import "fmt"

func sum(x []int, c chan int) {
	sum := 0
	for _, y := range x {
		sum += y
	}
	c <- sum // send sum to c (channel)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c (channel)

	fmt.Println(x, y, x+y)
}
