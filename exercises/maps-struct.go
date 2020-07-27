package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	statePopulations["Georgia"] = 10310371 // Adding a new entry.
	delete(statePopulations, "Georgia")    // Deleting

	_, isGeorgia := statePopulations["Georgia"]

	fmt.Println(statePopulations["Georgia"])
	fmt.Printf("isGeorgia: %v\n", isGeorgia)

	// struct
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48.1
	b.CanFly = false
	fmt.Println(b)

	// tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}
