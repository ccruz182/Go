package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

/*
* Basic example of routines

func main() {
	var msg string = "Hello"
	wg.Add(1)

	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "Bye"
	wg.Wait()
}
*/
