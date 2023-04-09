package main

import "fmt"

func main() {

	/*
			Slice creation styles:
		 		1. Slice existing array or slice.
				2. Literal style.
				3. Via make() function.
						a := make([]int, 10) // Create slice with capacity and length == 10
						a := make([]int, 10, 100) // Create slice with length == 10 and capacity == 100

			len() function returns length of slice.
			cap() function returns length of underlying array.
			append() function to add elements to the slice.
	*/

	{
		// Example #1 (Arrays of arrays)
		var identityMatrix [3][3]int // Create 3 arrays, which hold 3 values, like: [1,0,0][0,1,0][0,0,1]
		identityMatrix[0] = [3]int{1, 0, 0}
		identityMatrix[1] = [3]int{0, 1, 0}
		identityMatrix[2] = [3]int{0, 0, 1}

		fmt.Printf("\nValues in array of array: %v\n", identityMatrix)
		fmt.Printf("Length of array: %v\n", len(identityMatrix))
		fmt.Printf("Capacity of array: %v\n\n", cap(identityMatrix))
	}

	{
		// Example #2 (Modify an array)

		a := [...]int{1, 2, 3}

		// Copy a array values to b.
		b := a
		// Modify second value in b array.
		b[1] = 5
		fmt.Printf("%v\n", a)   // [1 2 3]
		fmt.Printf("%v\n\n", b) // [1 5 3]
	}

	{
		// Example #3 (Addressing)

		a := [...]int{1, 2, 3}

		// Address/point a array values to b, so they are linked.
		b := &a
		// If we change the 5th value in b, it also changes in a.
		b[1] = 5
		fmt.Printf("%v\n", a)   // [1 5 3]
		fmt.Printf("%v\n\n", b) // &[1 5 3]

	}

	{

		// Example #4 (Slices)

		// Unlike arrays, slices don't have to have a fixed length.
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b := a[:]   // slice of all elements
		c := a[3:]  // slice from 4th element to end
		d := a[6:]  // slice first 6 elements
		e := a[3:6] // slice the 4th, 5th and 6th elements

		a[5] = 42 // If you change the value in the array you sliced from, it changes all the slices.
		// Slicing works like pointing/addressing.

		fmt.Printf("%v\n", a)
		fmt.Printf("%v\n", b)
		fmt.Printf("%v\n", c)
		fmt.Printf("%v\n", d)
		fmt.Printf("%v\n\n", e)

	}

	{

		// Example #5 (Make)

		// Slice of length 3, capacity of 100.
		a := make([]int, 3, 100)

		a = append(a, 4) // Appends the array, works just like: a[3] = 4
		a = append(a, 5, 6, 7, 8)

		fmt.Printf("%v\n", a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n\n", cap(a))

		// The output is [0 0 0 4 5 6 7 8], if you create an empty array [0, 0, 0] and append it,
		// it just starts to append/extend the array.

	}

	{

		// Example #6 (Array manipulation)

		// We can't use the 'a' array twice, if we want to use it twice.
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 3, 4, 5}

		c := a[:len(a)-1]            // Creates a slice and removes the last number from 'a' array. [1, 2, 3, 4]
		d := append(b[:2], b[3:]...) // Removes the 3rd number from the 'b' array. Have to use the spread function "..." so append is happy. :)

		fmt.Printf("Removed the last number from 'a' array: %v\n", c)
		fmt.Printf("Removed the 3rd number from 'b' array: %v\n", d)

	}

}
