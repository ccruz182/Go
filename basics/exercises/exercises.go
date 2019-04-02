package main

import "fmt"

var x int

type dogAge int
type y dogAge

var kariAge dogAge
var myVar y

func main() {
	x := 42
	y := "James Bond"
	z := true

	// Single print statement
	fmt.Println(x, y, z)

	// Multiple print statement
	fmt.Println("X := ", x)
	fmt.Println("Y := ", y)
	fmt.Println("Z := ", z)

	// E3
	fmt.Println("Exercise #3")
	s := fmt.Sprintf("x = %v --- y = %v --- z = %v", x, y, z)
	fmt.Println(s)

	// E4
	fmt.Println("Exercise #4")
	fmt.Println(kariAge)
	fmt.Printf("%T\n", kariAge)
	kariAge = 1
	fmt.Println(kariAge)

	// E5
	fmt.Println("Exercise #4")
	myVar = 5
	fmt.Println(myVar)

	bd := 1997

	for bd <= 2019 {
		fmt.Println(bd)
		bd++
	}

	fmt.Println("---------")

	bd = 1997
	for {
		if bd <= 2019 {
			fmt.Println(bd)
			bd++
		} else {
			break
		}
	}

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

}
