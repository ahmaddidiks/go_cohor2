package main

import (
	"fmt"
	"sync"
)

func firstPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Go routine 3")
}

func secondtPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Go routine 4")

}

func main() {
	fmt.Println("Main go routine")

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Go routine 1")
	}(&wg)

	go firstPrint(&wg)
	go secondtPrint(&wg)

	wg.Wait()
}
