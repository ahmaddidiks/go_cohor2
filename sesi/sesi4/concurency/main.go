package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main go routine")

	go func() {
		fmt.Println("Go routine 1")
	}()

	go firstPrint()
	go secondtPrint()

	time.Sleep(time.Millisecond * 2)
}

func firstPrint() {
	fmt.Println("Go routine 3")
}

func secondtPrint() {
	fmt.Println("Go routine 4")
}
