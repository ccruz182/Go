package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)

	d, e := divide(5.2, 5.0)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Division:", d)

	g := greeter{
		greeting: "Hola",
		name:     "Cesar",
	}
	g.greet()
}

func sum(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}

	return &result
}

func sum2(values ...int) (result int) {
	for _, v := range values {
		result += v
	}

	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
