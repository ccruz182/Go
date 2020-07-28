package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	theNumber := 2038074743
	wg.Add(2)
	quarter, middle := getLimits(theNumber)

	var d1, d2 []int

	go func() {
		d1 = getDividers(theNumber, 2, quarter)
		wg.Done()
	}()

	go func() {
		d2 = getDividers(theNumber, quarter, middle+1)
		wg.Done()
	}()

	wg.Wait()

	dividers := append(d1, d2...)
	if len(dividers) == 0 {
		fmt.Printf("El número %v es primo\n\n", theNumber)
	} else {
		fmt.Printf("El número %v no es primo. Sus divisores son: %v\n\n", theNumber, dividers)
	}

}

func getDividers(number, begin, limit int) []int {
	var dividers []int = []int{}

	for i := begin; i < limit; i++ {
		if number%i == 0 {
			dividers = append(dividers, i)
		}
	}

	return dividers
}

func getLimits(number int) (int, int) {
	return number / 4, number / 2
}
