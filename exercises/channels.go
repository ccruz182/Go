package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(3)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 20
		close(ch)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		for k := 0; k < 5; k++ {
			//	ch <- k
		}
		wg.Done()
	}(ch)

	wg.Wait()
}
