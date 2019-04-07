package main

import (
	"fmt"
)

func main() {
	var x [5]int // Array, but GoDoc states that the use of slices is preferable.
	fmt.Println(x)
	x[0] = 182
	fmt.Println(x)

	/*
		SLICES: Values of the same type
	*/
	y := []int{0, 2, 4, 6, 8}
	fmt.Println(y)

	for _, v := range y {
		fmt.Println(v)
	}

	fmt.Println(y[:1])

	q := append(y[:2], y[3:]...)
	fmt.Println(q)

	row1 := []int{0, 2, 4, 6}
	row2 := []int{1, 3, 5, 7}

	table := [][]int{row1, row2}
	fmt.Println(table)

	/* MAPS */
	theMap := map[string]int{
		"Estrella": 14,
		"Kari":     0,
	}
	fmt.Println(theMap)
	name, age := theMap["Kari"]
	fmt.Println(name, age)
}
