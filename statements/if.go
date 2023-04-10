package main

import (
	"fmt"
)

func main() {

	{

		// #1 Simple example

		// Non-interactive number guessing game
		number := 50
		guess := 30
		if guess < number {
			fmt.Printf("Too low! %v\n", guess)
		}
		if guess > number {
			fmt.Printf("Too high! %v\n", guess)
		}
		if guess == number {
			fmt.Printf("You got it! %v\n", guess)
		}

	}

	{

		// #2 Booleans

		// Print out booleans
		number := 50
		guess := 30
		fmt.Println(number <= guess, number >= guess, number != guess)

	}

}
