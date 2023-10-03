package main

import (
	"fmt"
	"sync"
)

type challenge_interface interface {
	show() string
}

type data struct {
	value string
}

func (d data) show() string {
	return d.value
}

func routine(s string, i int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done() // first defer will be the latest execution
	defer m.Unlock()

	fmt.Println(s, i)
}

func main() {
	var bisa challenge_interface = data{"bisa"}
	var coba challenge_interface = data{"coba"}

	var mutex sync.Mutex
	var wg sync.WaitGroup

	/**
	* Ini sama aja kaya wg.Wait() di setiap iterate
	 */
	for i := 0; i < 4; i++ {
		wg.Add(2)
		mutex.Lock()
		go routine(bisa.show(), i, &wg, &mutex)
		mutex.Lock()
		go routine(coba.show(), i, &wg, &mutex)
	}
	wg.Wait()

}
