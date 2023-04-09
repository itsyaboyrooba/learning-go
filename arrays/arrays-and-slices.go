package main

import "fmt"

func main() {

	/*
		Array declaration styles:
				1. a := [3]int{1, 2, 3}
				2. a := [...]int{1, 2, 3}
		 		3. var a [3]int

		Access via zero-based index:
				a := [3]int {1, 3, 5} // a [1] == 3

		len() function returns length of array.

	*/

	{

		// Example #1

		// Create an empty array:
		grades_empty := [3]int{}
		fmt.Printf("\nEmpty grades: %v\n", grades_empty)

		// Create an array of 3 objects:
		grades_definedSize := [3]int{97, 85, 82}
		fmt.Printf("Fixed array grades: %v\n", grades_definedSize)

		// Create an array of undefined objects:
		grades_undefinedSize := [...]int{97, 85, 82}
		fmt.Printf("Dynamic array grades: %v\n", grades_undefinedSize)

	}

	{

		// Example #2

		// Put data into an empty array:
		var students [3]string
		fmt.Printf("\nEmpty array of students: %v\n", students)

		students[0] = "Tooc" // Array starts from 0.
		students[1] = "Tolmatsov"
		students[2] = "Obama"
		fmt.Printf("Array of students: %v\n", students)
		fmt.Printf("Student #2: %v\n", students[1])

		// Get the length of the array:
		fmt.Printf("Students in school: %v\n", len(students)) // len() is length()

	}

}
