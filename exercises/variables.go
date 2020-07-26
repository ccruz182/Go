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
}
