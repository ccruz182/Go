package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 42
	x := 30
	fmt.Println(i, x)

	// String conversion
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

	// Complex numbers
	var myComplex complex64 = 1 + 2i
	var otherComplex complex64 = 2 - 8i
	complexMulti := myComplex * otherComplex
	fmt.Printf("Complex operation: %v %v\n", real(complexMulti), imag(complexMulti))

	// Text - rune
	var r rune = 'a'
	fmt.Printf("%v %T\n", r, r)

}
