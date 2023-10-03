package main

import (
	"fmt"
	"sync"
)

type challenge interface {
	show() string
}

type data struct {
	value string
}

func (d data) show() string {
	return d.value
}

func routine(s string, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s, i)
}

func main() {
	var bisa challenge = data{"bisa"}
	var coba challenge = data{"coba"}
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go routine(bisa.show(), i, &wg)
		go routine(coba.show(), i, &wg)
	}
	wg.Wait()

}
