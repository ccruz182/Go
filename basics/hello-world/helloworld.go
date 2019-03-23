package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	x := 34 // Declaration and asigning a value
	fmt.Println(x)
	x = 50 // Asigning value to a declared variable
	fmt.Println(x)
}

func foo() {
	fmt.Println("I'm in foo")
}
