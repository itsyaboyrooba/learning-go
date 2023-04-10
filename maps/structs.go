package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	episodes   []string // This is an array.
	companions []string
}

func main() {

	{

		// Example #1 (Struct)

		aDoctor := Doctor{
			number:    3,
			actorName: "Jon Pertwee",
			// If you make a slice out of an array in a struct, you have to define what type it is. []string{}
			episodes: []string{
				"Doctor Who?",
				"Who is this?",
			},
			companions: []string{
				"Liz Shaw",
				"Jo Grant",
				"Sarah Jane Smith",
			},
		}
		// Get the second item in the companions slice from struct 'aDoctor'.
		fmt.Printf("Jon Pertwee second companion: %v\n\n", aDoctor.companions[1])

	}

	{

		// Example #2 (Placeholder)

		aDoctor := Doctor{
			// You don't have to list the field names.
			3,
			"Jon Pertwee",
			// If you don't want to use a field ('episodes'), you have to use a placeholder.
			[]string{},
			[]string{
				"Liz Shaw",
				"Jo Grant",
				"Sarah Jane Smith",
			},
		}
		fmt.Printf("Jon Pertwee episodes: %v\n\n", aDoctor.episodes) // Returns empty, because we used a placeholder.

	}

	{

		// Example #3 (Order)

		aDoctor := Doctor{
			// You can use a different order if you define the fields.
			episodes:   []string{},
			companions: []string{},
			number:     3,
			actorName:  "Jon Pertwee",
		}

		fmt.Printf("Jon Pertwee number: %v\n\n", aDoctor.number)

	}

	{

		// Example #4 (Anonymous struct)

		// You can't use this anywhere else because this is anonymous,
		// it doesn't have a independent name that you could refer to.
		aDoctor := struct{ name string }{name: "John Pertwee"}
		anotherDoctor := aDoctor
		anotherDoctor.name = "Tom Baker"
		fmt.Printf("Anonymous struct: %v\n", aDoctor)
		fmt.Printf("Sliced and modifyed anonymous struct: %v\n\n", anotherDoctor)

	}

	{

		// Example #5 (Pointer anonymous struct) // some compile magic

		aDoctor := struct{ name string }{name: "John Pertwee"}

		// If we want to use the same underlying data, we can use &.
		anotherDoctor := &aDoctor
		anotherDoctor.name = "Tom Baker"
		fmt.Printf("Anonymous struct: %v\n", aDoctor)
		fmt.Printf("Sliced, modifyed and addressed anonymous struct: %v\n\n", anotherDoctor)

	}

}
