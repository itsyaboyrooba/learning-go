package main

// Tootsi esimene kalkulaator :OOOOO

import "fmt"

var (
	esimeneNumber, teineNumber float32
	tehe                       string
)

func main() {

	fmt.Printf("Sisesta esimene number: ")
	fmt.Scan(&esimeneNumber)
	fmt.Printf("Sisesta teine number: ")
	fmt.Scan(&teineNumber)
	fmt.Printf("mis tehe: ")
	fmt.Scan(&tehe)

	switch tehe {
	case "+":
		fmt.Println(esimeneNumber + teineNumber)
	case "-":
		fmt.Println(esimeneNumber - teineNumber)
	case "*":
		fmt.Println(esimeneNumber * teineNumber)
	case "/":
		fmt.Println(esimeneNumber / teineNumber)
	default:
		fmt.Println("Pole korrektne tehe!")
	}
}
