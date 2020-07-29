package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg = sync.WaitGroup{}
var wg2 = sync.WaitGroup{}
var primesCh = make(chan int, 1000)
var doneCh = make(chan struct{})

func main() {
	theNumber := 2038074743
	//theNumber := 110
	wg.Add(2)
	wg2.Add(1)
	quarter, middle := getLimits(theNumber)

	go func(ch <-chan int) {
		var dividers []int
		for {
			if i, ok := <-ch; ok {
				dividers = append(dividers, i)
			} else {
				sort.Ints(dividers)
				if len(dividers) == 0 {
					fmt.Printf("El número %v es primo.\n", theNumber)
				} else {
					fmt.Printf("El número %v no es primo. Sus divisores son: %v.\n", theNumber, dividers)
				}

				wg2.Done()
				break
			}
		}
		//wg.Done()
	}(primesCh)

	go func() {
		getDividers(theNumber, 2, quarter)
		wg.Done()
	}()

	go func() {
		getDividers(theNumber, quarter, middle+1)
		wg.Done()
	}()

	wg.Wait()
	close(primesCh)
	wg2.Wait()
	//doneCh <- struct{}{}
}

func getDividers(number, begin, limit int) {
	for i := begin; i < limit; i++ {
		if number%i == 0 {
			primesCh <- i
		}
	}
}

func getLimits(number int) (int, int) {
	return number / 4, number / 2
}
