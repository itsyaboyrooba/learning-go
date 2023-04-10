package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100` // Make the field required and max length of 100.
	Origin string
}

type Bird struct {
	Animal   // Embed the 'Animal' struct into the 'Bird' struct, bird is an animal afterall...
	SpeedKPH float32
	CanFly   bool
}

func main() {

	{

		// Example #1 (Embedding)

		b := Bird{}
		b.Name = "Emu"
		b.Origin = "Australia"
		b.SpeedKPH = 48
		b.CanFly = false
		fmt.Printf("Animal name: %v\n", b.Name)

	}

	{

		// Example #2 (Declaring)

		b := Bird{
			// Inside 'Bird' struct, embedded 'Animal struct, declare 'Animal' variables.
			Animal:   Animal{Name: "Chicken", Origin: "Estonia"},
			SpeedKPH: 10,
			CanFly:   false,
		}
		fmt.Printf("Animal name: %v\n", b.Name)
	}

	{

		// Example #3 (Tags)

		t := reflect.TypeOf(Animal{})     // Reflect the struct
		field, _ := t.FieldByName("Name") // Get the "Name" field in the reflected struct 't'.
		fmt.Println(field.Tag)            // Print the tag of the field

	}

}
